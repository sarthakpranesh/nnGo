package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	porter "github.com/reiver/go-porterstemmer"
	"github.com/sarthakpranesh/nnGo"
)

// RawData is non processed csv data
type RawData struct {
	Tweets	   []string
	Category   []int64    // two categories: 0 means not hate speech and 1 is hate speech
}

// TokenizedData is tokenized data
type TokenizedData struct {
	TokenizedTweets [][]string
	Category   []int64
}

var (
	stopWords = []string{"i", "me", "my", "myself", "we", "our", "ours", "ourselves",
		"you", "your", "yours", "yourself", "yourselves", "he", "him", "his", "himself",
		"she", "her", "hers", "herself", "it", "its", "itself", "they", "them", "their",
		"theirs", "themselves", "what", "which", "who", "whom", "this", "that", "these",
		"those", "am", "is", "are", "was", "were", "be", "been", "being", "have", "has",
		"had", "having", "do", "does", "did", "doing", "a", "an", "the", "and", "but", "if",
		"or", "because", "as", "until", "while", "of", "at", "by", "for", "with",
		"about", "against", "between", "into", "through", "during", "before", "after",
		"above", "below", "to", "from", "up", "down", "in", "out", "on", "off", "over",
		"under", "again", "further", "then", "once", "here", "there", "when", "where",
		"why", "how", "all", "any", "both", "each", "few", "more", "most", "other",
		"some", "such", "no", "nor", "not", "only", "own", "same", "so", "than", "too",
		"very", "s", "t", "can", "will", "just", "don", "should", "now", "@user"}
	punctuation   = []string{"#", "'", ",", ".", "?", "!"}
)

const (
	testIndex = 814
)

// LoadCsv helps in importing csv data into a RawData type
func LoadCsv(name string) (RawData, error) {
	f, err := os.Open(name)
	if err != nil {
		return RawData{}, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return RawData{}, err
	}

	var Tweets []string
	var TweetCategory []int64
	for _, line := range lines {
		Tweets = append(Tweets, line[0])
		i, err := strconv.ParseInt(line[1], 10, 64)
		if err != nil {
			i = 0
		}
		TweetCategory = append(TweetCategory, i)
	}

	return RawData{
		Tweets: Tweets,
		Category:   TweetCategory,
	}, nil
}

// Tokenize helps convert the string tweet into a list of words
func Tokenize(rawData RawData) TokenizedData {
	tokenizedData := TokenizedData{
		Category: rawData.Category,
	}
	
	for _, tweet := range rawData.Tweets {
		var tokenizedTweet []string
		for _, word := range strings.Split(tweet, " ") {
			if strings.Trim(word, " ") == "" {
				continue
			} else {
				tokenizedTweet = append(tokenizedTweet, word)
			}
		}
		tokenizedData.TokenizedTweets = append(tokenizedData.TokenizedTweets, tokenizedTweet)
	}

	return tokenizedData
}

// contains finds if the string s has any of the substrings present in stopwords
func contains(s string, stopWords []string) bool {
	for _, a := range stopWords {
		if a == s {
			return true
		}
	}
	return false
}

// RemoveStopWords are used to remove any stop words
func RemoveStopWords(tokenizedData TokenizedData) TokenizedData {
	var removedStopWords [][]string
	for _, tweet := range tokenizedData.TokenizedTweets {
		var tweetWithOutStopWords []string
		for _, word := range tweet {
			isThere := contains(word, stopWords)
			if isThere == false {
				tweetWithOutStopWords = append(tweetWithOutStopWords, word)
			}
		}
		removedStopWords = append(removedStopWords, tweetWithOutStopWords)
	}
	tokenizedData.TokenizedTweets = removedStopWords

	return tokenizedData
}

// cleanPunctuation removes all the punctuations from the string
func cleanPunctuation(word string, p []string) string {
	for _, a := range p {
		word = strings.Replace(word, a, "", -1)
	}
	return word
}

// RemovePun removes all the punctuation defined
func RemovePun(tokenizedData TokenizedData) TokenizedData {
	var removedPunctuations [][]string
	for _, tweet := range tokenizedData.TokenizedTweets {
		var tweetWithOutPun []string
		for _, word := range tweet {
			word := cleanPunctuation(word, punctuation)
			tweetWithOutPun = append(tweetWithOutPun, word)
		}
		removedPunctuations = append(removedPunctuations, tweetWithOutPun)
	}
	tokenizedData.TokenizedTweets = removedPunctuations

	return tokenizedData
}

func StemData(tokenizedData TokenizedData) TokenizedData {
	var stemmedData [][]string
	for _, tweet := range tokenizedData.TokenizedTweets {
		var stemmedTweet []string
		for _, word := range tweet {
			word := porter.StemString(word)
			stemmedTweet = append(stemmedTweet, word)
		}
		stemmedData = append(stemmedData, stemmedTweet)
	}
	tokenizedData.TokenizedTweets = stemmedData

	return tokenizedData
}

// CreateDictionaries creates two dictionaries, one for positive (non hate speech) and the other for negative (hate speech)
func CreateDictionaries(tokenizedData TokenizedData) (map[string]float64, map[string]float64) {
	positiveDic := make(map[string]float64)
	negativeDic := make(map[string]float64)
	for i, tweet := range tokenizedData.TokenizedTweets {
		for _, word := range tweet {
			if tokenizedData.Category[i] == 0 {
				val, _ := positiveDic[word]
				val = val + 1
				positiveDic[word] = val
			} else if tokenizedData.Category[i] == 1 {
				val, _ := negativeDic[word]
				val = val + 1
				negativeDic[word] = val
			} else {
				fmt.Println("--> No Proper Class")
			}
		}
	}

	return positiveDic, negativeDic
}

// NormalizeDictionary converts the given map into a probabilistic map
func NormalizeDictionary(dict map[string]float64) map[string]float64 {
	var max float64
	for _, v := range dict {
		if max < v {
			max = v
		}
	}

	for k, v := range dict {
		dict[k] = v / max
	}

	return dict
}

func GenerateNNData(tokenizedData TokenizedData, positiveDic, negativeDic map[string]float64) ([][]float64, [][]float64) {
	var tData, tLabel [][]float64
	for i, tweet := range tokenizedData.TokenizedTweets {
		if tokenizedData.Category[i] == 0 {
			tLabel = append(tLabel, []float64{1, 0})
		} else {
			tLabel = append(tLabel, []float64{0, 1})
		}
		var posi, nega float64
		for _, word := range tweet {
			posi = posi + positiveDic[word]
			nega = nega + negativeDic[word]
		}
		tData = append(tData, []float64{1, posi, nega})
	}

	return tData, tLabel
}

func main() {
	/*
		Training Phase
	*/
	// This loads the data in from a csv file:>
	rawCsvData, err := LoadCsv("tweets.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Tweet:", rawCsvData.Tweets[testIndex], "\tLabel:", rawCsvData.Category[testIndex])

	// Tokenize data
	tokenizedData := Tokenize(rawCsvData)
	fmt.Println("Tweet:", tokenizedData.TokenizedTweets[testIndex], "\tLabel:", tokenizedData.Category[testIndex])

	// Removing stop words
	dataWithNoStopWords := RemoveStopWords(tokenizedData)
	fmt.Println("Tweet:", dataWithNoStopWords.TokenizedTweets[testIndex], "\tLabel:", dataWithNoStopWords.Category[testIndex])

	// Removing punctuations
	dataWithNoPunc := RemovePun(dataWithNoStopWords)
	fmt.Println("Tweet:", dataWithNoPunc.TokenizedTweets[testIndex], "\tLabel:", dataWithNoPunc.Category[testIndex])

	// Stem data
	stemedData := StemData(dataWithNoPunc)
	fmt.Println("Tweet:", stemedData.TokenizedTweets[testIndex], "\tLabel:", stemedData.Category[testIndex])

	// creating dictionary
	positiveDic, negativeDic := CreateDictionaries(stemedData)
	positiveDic = NormalizeDictionary(positiveDic)
	negativeDic = NormalizeDictionary(negativeDic)
	/*
	* positiveDic - All non hate speech tweets labeled with zero in this dictionary
	* negativeDic - All hate speech tweets labeled with one in this dictionary
	 */

	// create trainable data
	TweetData, TweetLabel := GenerateNNData(stemedData, positiveDic, negativeDic)
	fmt.Println("Tweet:", TweetData[testIndex], "\tLabel:", TweetLabel[testIndex])

	trainData := TweetData[:4000]
	trainLabel := TweetLabel[:4000]
	testData := TweetData[4000:]
	testLabel := TweetLabel[4000:]

	// Neural Network Model
	nn := nnGo.NewNN(3, 100, 2, 0.5, "sgd", 20)

	nn.Train(trainData, trainLabel)

	fmt.Println("Actual Encoded Tweet:", testData[21])
	fmt.Println("Actual:", testLabel[21])
	nn.Predict(testData[21])

	fmt.Println("Actual Encoded Tweet:", testData[22])
	fmt.Println("Actual:", testLabel[22])
	nn.Predict(testData[22])
}
