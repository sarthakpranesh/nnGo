package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	porter "github.com/reiver/go-porterstemmer"
	"github.com/sarthakpranesh/nnGo/nnGo"
)

// RawData is non processed csv data
type RawData struct {
	tweets []string
	hate   []int64
}

// TokenizedData is tokenized data
type TokenizedData struct {
	tweets [][]string
	hate   []int64
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
	rawData       RawData
	tokenizedData TokenizedData
	positiveDic   = make(map[string]float64)
	negativeDic   = make(map[string]float64)
)

// LoadData helps in importing csv data into a RawData type
func LoadData(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}
	var tweetslist []string
	var hatelist []int64
	for _, line := range lines {
		tweetslist = append(tweetslist, line[2])
		i, err := strconv.ParseInt(line[1], 10, 64)
		if err != nil {
			i = 0
		}
		hatelist = append(hatelist, i)
	}

	rawData = RawData{
		tweets: tweetslist,
		hate:   hatelist,
	}

	return nil
}

func Tokenize() {
	tokenizedData = TokenizedData{
		hate: rawData.hate,
	}
	for _, tweet := range rawData.tweets {
		var tokenizedTweet []string
		for _, word := range strings.Split(tweet, " ") {
			if strings.Trim(word, " ") == "" {
				continue
			} else {
				tokenizedTweet = append(tokenizedTweet, word)
			}
		}
		tokenizedData.tweets = append(tokenizedData.tweets, tokenizedTweet)
	}
}

func contains(s string, stopWords []string) bool {
	for _, a := range stopWords {
		if a == s {
			return true
		}
	}
	return false
}

// RemoveStopWords are used to remove any stop words
func RemoveStopWords() {
	var removedStopWords [][]string
	for _, tweet := range tokenizedData.tweets {
		var tweetWithOutStopWords []string
		for _, word := range tweet {
			isThere := contains(word, stopWords)
			if isThere == false {
				tweetWithOutStopWords = append(tweetWithOutStopWords, word)
			}
		}
		removedStopWords = append(removedStopWords, tweetWithOutStopWords)
	}
	tokenizedData.tweets = removedStopWords
}

func cleanPunctuation(word string, p []string) string {
	for _, a := range p {
		word = strings.Replace(word, a, "", -1)
	}
	return word
}

// RemovePun removes all the punctuation defined
func RemovePun() {
	var removedPunctuations [][]string
	for _, tweet := range tokenizedData.tweets {
		var tweetWithOutPun []string
		for _, word := range tweet {
			word := cleanPunctuation(word, punctuation)
			tweetWithOutPun = append(tweetWithOutPun, word)
		}
		removedPunctuations = append(removedPunctuations, tweetWithOutPun)
	}
	tokenizedData.tweets = removedPunctuations
}

func StemData() {
	var stemmedData [][]string
	for _, tweet := range tokenizedData.tweets {
		var stemmedTweet []string
		for _, word := range tweet {
			word := porter.StemString(word)
			stemmedTweet = append(stemmedTweet, word)
		}
		stemmedData = append(stemmedData, stemmedTweet)
	}
	tokenizedData.tweets = stemmedData
}

func CreateDictionaries() {
	for i, tweet := range tokenizedData.tweets {
		for _, word := range tweet {
			if tokenizedData.hate[i] == 0 {
				val, _ := positiveDic[word]
				val = val + 1
				positiveDic[word] = val
			} else if tokenizedData.hate[i] == 1 {
				val, _ := negativeDic[word]
				val = val + 1
				negativeDic[word] = val
			} else {
				fmt.Println("--> No Proper Class")
			}
		}
	}
}

func NormalizeDictionary(dict map[string]float64) {
	var max float64
	for _, v := range dict {
		if max < v {
			max = v
		}
	}

	for k, v := range dict {
		dict[k] = v / max
	}
}

func GenerateNNData(tData [][]float64, tLabel [][]float64) ([][]float64, [][]float64) {
	for i, tweet := range tokenizedData.tweets {
		if tokenizedData.hate[i] == 0 {
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
	// This loads the data in :>
	LoadData("train.csv")
	// Tokenize data
	Tokenize()
	// Removing stop words
	RemoveStopWords()
	// Removing punctuations
	RemovePun()
	// Stem data
	StemData()

	// creating dictionary
	CreateDictionaries()
	NormalizeDictionary(positiveDic)
	NormalizeDictionary(negativeDic)
	/*
	* positiveDic - All non hate speech tweets labeled with zero in dataset
	* negativeDic - All hate speech tweets labeled with one in dataset
	 */
	nn := nnGo.NewNN(3, 8, 2, 0.6, "sgd", 20)
	var TweetData, TweetLabel [][]float64
	TweetData, TweetLabel = GenerateNNData(TweetData, TweetLabel)

	trainData := TweetData[:30000]
	trainLabel := TweetLabel[:30000]
	nn.Train(trainData, trainLabel)

	/*
		Testing Phase
	*/
	testData := TweetData[30000:]
	testLabel := TweetLabel[30000:]

	fmt.Println("Processed Tweet:", tokenizedData.tweets[31948], "\t Tweet Label:", tokenizedData.hate[31948])
	fmt.Println("Processed Tweet:", TweetData[31948], "\t Tweet Label:", TweetLabel[31948])
	nn.Predict(testData[1948])
	fmt.Println(testLabel[1948])
}
