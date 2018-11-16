package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/vjeantet/jodaTime"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Tweet struct {
	tweet_id string
	timestamp string
	text string
}

func main(){
	deleteBeforeDate, onlySwearwords := talkToUser()
	setupTwitterClient()
	tweets := readTweetsCsv()
	swearWords := readSwearWordsCsv()
	tweetsToDelete := cleanTweets(tweets, swearWords, deleteBeforeDate, onlySwearwords)
	deleteTweets(tweetsToDelete)
}

func talkToUser() (deleteBeforeDate time.Time, onlySwearwords bool){
	onlySwearwords = true
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Tweet Deleter!")
	fmt.Println("Please enter the date you want to delete tweets before in the format DD/MM/YYYY.")

	scanner.Scan()
	text := scanner.Text()
	deleteBeforeDate, _ = jodaTime.Parse("dd/MM/yyyy", text)
	fullDate := jodaTime.Format("dd/MM/yyyy", deleteBeforeDate)

	fmt.Println("Do you want to delete only tweets with swear words?")
	fmt.Println("Y/N?")
	scanner.Scan()
	text = scanner.Text()

	if strings.ToUpper(scanner.Text()) != "Y" {
		onlySwearwords = false
	}

	fmt.Println("All your tweets before the date " + fullDate + " will be PERMANENTLY DELETED.")
	fmt.Println("Y/N?")
	scanner.Scan()

	if strings.ToUpper(scanner.Text()) != "Y" {
		os.Exit(3)
	}

	return
}

func setupTwitterClient() (client *twitter.Client){
	cfg := loadConfiguration()
	config := oauth1.NewConfig(cfg.ConsumerKey, cfg.ConsumerSecret)
	token := oauth1.NewToken(cfg.AccessToken, cfg.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client = twitter.NewClient(httpClient)

	return
}

type Config struct {
	ConsumerKey string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
	AccessToken string `json:"accessToken"`
	AccessSecret string `json:"accessSecret"`
}

func loadConfiguration() Config {
	var config Config

	configFile, err := os.Open("config.json")
	defer configFile.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func readTweetsCsv() (tweets []Tweet) {
	csvFile, _ := os.Open("tweets.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		tweets = append(tweets, Tweet {
			tweet_id: line[0],
			timestamp: line[3],
			text: line[5],
		})
	}

	return
}

func readSwearWordsCsv() (swearWords [][]string) {
	csvFile, _ := os.Open("swearWords.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	swearWords, error := reader.ReadAll()

	if error != nil {
		log.Fatal(error)
	}

	return
}

func cleanTweets(tweets []Tweet, swearWords [][]string, deleteBeforeTime time.Time, onlySwearwords bool) (tweetsToDelete []Tweet){

	for _, tweet := range tweets {
		if tweet.timestamp != "timestamp" {
			tweetDate, _ := jodaTime.Parse("yyyy-MM-dd HH:mm:ss Z", tweet.timestamp)

			if onlySwearwords {
				for _, swearWord := range swearWords[0] {
					if strings.Contains(strings.ToLower(tweet.text), " "+swearWord+" ") && tweetDate.Before(deleteBeforeTime) {
						if tweetDate.Before(deleteBeforeTime) {
							tweetsToDelete = append(tweetsToDelete, tweet)
							break
						}
					}
				}
			} else {
				if tweetDate.Before(deleteBeforeTime) {
					tweetsToDelete = append(tweetsToDelete, tweet)
				}
			}
		};
	}

	return
}

func deleteTweets(tweetsToDelete []Tweet){
	client := setupTwitterClient()

	for i, tweet := range tweetsToDelete {
		currentIndex := fmt.Sprintf("%.2f",float64(i)/float64(len(tweetsToDelete)) * 100)
		id, _ := strconv.ParseInt(tweet.tweet_id, 0, 64)
		client.Statuses.Destroy(id, nil)
		fmt.Println("Deleting tweets... " + currentIndex + "%")
	}

	fmt.Println("Tweets Deleted.")
}