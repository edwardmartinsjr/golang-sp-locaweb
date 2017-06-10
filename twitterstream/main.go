package main

import (
	. "bayesian"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"
)

//./sentiment -track=Trump -maxtweet=1 > sentimentresult.txt
//DefaultThresh -
const DefaultThresh = .95

const (
	consumerKey       string = "xxx"
	consumerSecret    string = "xxx"
	accessToken       string = "xxx-xxx"
	accessTokenSecret string = "xxx"
)

var track *string        // comma-delimited list of tracking keywords for twitter api
var maxtweet *string     // number of tweets to be classified
var clssfier *Classifier // the classifier
var san *Sanitizer       // the sanitizer
var exclList *string     // list of excluded terms
var count [2]int         // the count of all classifications
var highCount [2]int     // the count of all learned classifications
var thresh *float64      // threshold for learning
var printOnly *bool      // suppress classification?
var loadFile *string

func init() {
	// command-line flags
	maxtweet = flag.String("maxtweet", "", "number of tweets to be classified")
	track = flag.String("track", "", "comma-separated list of tracking terms")
	thresh = flag.Float64("thresh", DefaultThresh, "the confidence threshold required to learn new content")
	exclList = flag.String("exclude", "", "comma-separated list of keywords excluded from classification")
	printOnly = flag.Bool("print-only", false, "only print the Tweets, do not classify them")
	loadFile = flag.String("load-file", "", "specify classifier file")
	flag.Parse()

	// load and train the classifier
	if *loadFile != "" {
		// from a file
		c, err := NewClassifierFromFile(*loadFile)
		if err != nil {
			println("error loading:", err.Error)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "classifier is loaded: %v\n", c.WordCount())
	} else {
		// from scratch
		clssfier = NewClassifier(Positive, Negative)
		LearnFile(clssfier, "data/positive.txt", Positive)
		LearnFile(clssfier, "data/negative.txt", Negative)
		fmt.Fprintf(os.Stderr, "classifier is trained: %v\n", clssfier.WordCount())
	}

	// init the sanitizer
	excl := strings.Split(*exclList, ",")
	if *exclList != "" {
		fmt.Fprintf(os.Stderr, "excluding: %v\n", excl)
	}
	fmt.Printf("Track: %s\n", *track)
	stopWords := ReadFile("data/stopwords.txt")
	fmt.Printf("stop words: %v\n", stopWords)
	san = NewSanitizer(
		ToLower,
		NoMentions,
		NoLinks,
		NoNumbers,
		Punctuation,
		NoSmallWords,
		CombineNots,
		Exclusions(excl),
		Exclusions(stopWords),
	)
	// listen for Ctrl-C
	go signalHandler()
}

func signalHandler() {
	for {

		var signalChannel chan os.Signal
		signalChannel = make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt)
		go func() {
			<-signalChannel

			t := time.Now()
			name := t.Format("15-04-05") + ".data"
			println("\nsaving classifier to", name)
			err := clssfier.WriteToFile(name)
			if err != nil {
				println("error", err)
			}
			os.Exit(0)

		}()
	}
}

func main() {

	var twitterList = TwitterTrack(*maxtweet, *track, consumerKey, consumerSecret, accessToken, accessTokenSecret)

	// process the tweets
	for _, item := range twitterList {
		if !*printOnly {
			process(item.Tweet)
		} else {
			fmt.Println(item.Tweet)
		}
	}
}

// Obtain a the tweet, classify it, learn it
// if necessary, calculate the positive tweet
// rate and print information.
func process(document string) {
	fmt.Printf("\n> %v\n\n", document)
	// the sanitized document
	doc := san.GetDocument(document)
	if len(doc) < 1 {
		return
	}

	// classification of this document
	//fmt.Printf("---> %s\n", doc)
	scores, inx, _ := clssfier.ProbScores(doc)
	logScores, logInx, _ := clssfier.LogScores(doc)
	class := clssfier.Classes[inx]
	logClass := clssfier.Classes[logInx]

	// the rate of positive sentiment
	posrate := float64(count[0]) / float64(count[0]+count[1])
	learned := ""

	// if above the threshold, then learn
	// this document
	if scores[inx] > *thresh {
		clssfier.Learn(doc, class)
		learned = "***"
	}

	// print info
	prettyPrintDoc(doc)
	fmt.Printf("%7.5f %v %v\n", scores[inx], class, learned)
	fmt.Printf("%7.2f %v\n", logScores[logInx], logClass)
	if logClass != class {
		// incorrect classification due to underflow
		fmt.Println("CLASSIFICATION ERROR!")
	}
	fmt.Printf("%7.5f (posrate)\n", posrate)
	//fmt.Printf("%5.5f (high-probability posrate)\n", highrate)
}

// pretty print all the classification information
func prettyPrintDoc(doc []string) {
	//fmt.Printf("\n%v\n", doc)
	fmt.Printf("\t")
	for _, word := range doc {
		fmt.Printf("%7s", abbrev(word, 5))
	}
	fmt.Println("")

	freqs := clssfier.WordFrequencies(doc)
	for i := 0; i < 2; i++ {
		fmt.Printf("%6s", clssfier.Classes[i])
		for j := 0; j < len(doc); j++ {
			fmt.Printf("%7.4f", freqs[i][j])
		}
		fmt.Println("")
	}
}

// abbrev will abreavate a word with ".." if it
// is too long. It is used for display purposes.
func abbrev(word string, max int) (result string) {
	result = word
	if max < 5 {
		panic("max must be at least 5")
	}
	if len(word) > max {
		result = word[0:max-2] + ".."
	}
	return
}
