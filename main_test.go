package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/vjeantet/jodaTime"
	"testing"
)

// WHEN Swearwords Enabled...
func Test_OnlyTweetsWithSwearwordsAreRemoved(t *testing.T){
	var tweets = []Tweet {
		Tweet {
			tweet_id: "1",
			timestamp: "2018-11-15 08:56:54 +0000",
			text: "this is a karam tweet",
		},
		Tweet {
			tweet_id: "2",
			timestamp: "2018-11-15 08:56:54 +0000",
			text: "this is a good tweet",
		},
		Tweet {
			tweet_id: "3",
			timestamp: "2018-11-15 08:56:54 +0000",
			text: "this is a nice tweet",
		},
	}

	var swearWords [][]string
	swearWord := []string{"karam"}
	swearWords = append(swearWords, swearWord)

	deleteBeforeDate, _ := jodaTime.Parse("dd/MM/yyyy", "30/11/2018")

	tweetsToDelete := cleanTweets(tweets, swearWords, deleteBeforeDate, true)
	assert.Equal(t, 1, len(tweetsToDelete), "Tweet containing swear word should be removed.")
	assert.Equal(t, tweetsToDelete[0].tweet_id, "1")
}

func Test_OnlyTweetsContainingSwearwordsAndBeforeDateAreRemoved(t *testing.T){
	var tweets = []Tweet {
		Tweet {
			tweet_id: "1",
			timestamp: "2018-11-15 08:56:54 +0000",
			text: "this is a karam tweet",
		},
		Tweet {
			tweet_id: "4",
			timestamp: "2018-11-15 08:56:54 +0000",
			text: "this is a lovely tweet",
		},
		Tweet {
			tweet_id: "2",
			timestamp: "2018-11-16 08:56:54 +0000",
			text: "this is a good tweet",
		},
		Tweet {
			tweet_id: "3",
			timestamp: "2018-11-16 08:56:54 +0000",
			text: "this is a nice tweet",
		},
	}

	var swearWords [][]string
	swearWord := []string{"karam"}
	swearWords = append(swearWords, swearWord)

	deleteBeforeDate, _ := jodaTime.Parse("dd/MM/yyyy", "16/11/2018")

	tweetsToDelete := cleanTweets(tweets, swearWords, deleteBeforeDate, true)
	assert.Equal(t, 1, len(tweetsToDelete), "One tweet will be removed.")
	assert.Equal(t, tweetsToDelete[0].tweet_id, "1")
}

func Test_NoSwearwordsSoNoTweetsAreRemoved(t *testing.T){
	var tweets = []Tweet {
		Tweet {
			tweet_id: "1",
			timestamp: "2018-11-15 08:56:54 +0000",
			text: "this is a lovely tweet",
		},
		Tweet {
			tweet_id: "2",
			timestamp: "2018-11-16 08:56:54 +0000",
			text: "this is a good tweet",
		},
		Tweet {
			tweet_id: "3",
			timestamp: "2018-11-16 08:56:54 +0000",
			text: "this is a nice tweet",
		},
	}

	var swearWords [][]string
	swearWord := []string{"karam"}
	swearWords = append(swearWords, swearWord)

	deleteBeforeDate, _ := jodaTime.Parse("dd/MM/yyyy", "16/11/2018")

	tweetsToDelete := cleanTweets(tweets, swearWords, deleteBeforeDate, true)
	assert.Equal(t, 0, len(tweetsToDelete), "One tweet will be removed.")
}

// WHEN Swearwords Disabled...
func Test_WhenSwearwordsDisabledAllTweetsRemoved(t *testing.T){
	var tweets = []Tweet {
		Tweet {
			tweet_id: "1",
			timestamp: "2018-11-15 08:56:54 +0000",
			text: "this is a karam tweet",
		},
		Tweet {
			tweet_id: "2",
			timestamp: "2018-11-15 08:56:54 +0000",
			text: "this is a good tweet",
		},
		Tweet {
			tweet_id: "3",
			timestamp: "2018-11-15 08:56:54 +0000",
			text: "this is a nice tweet",
		},
	}

	var swearWords [][]string
	swearWord := []string{"karam"}
	swearWords = append(swearWords, swearWord)

	deleteBeforeDate, _ := jodaTime.Parse("dd/MM/yyyy", "30/11/2018")

	tweetsToDelete := cleanTweets(tweets, swearWords, deleteBeforeDate, false)
	assert.Equal(t, 3, len(tweetsToDelete), "All tweets will be removed.")
}

func Test_WhenSwearwordsDisabledTweetsBeforeDateAreRemoved(t *testing.T){
	var tweets = []Tweet {
		Tweet {
			tweet_id: "1",
			timestamp: "2018-11-15 08:56:54 +0000",
			text: "this is a lovely tweet",
		},
		Tweet {
			tweet_id: "2",
			timestamp: "2018-11-16 08:56:54 +0000",
			text: "this is a good tweet",
		},
		Tweet {
			tweet_id: "3",
			timestamp: "2018-11-16 08:56:54 +0000",
			text: "this is a nice tweet",
		},
	}

	var swearWords [][]string
	swearWord := []string{"karam"}
	swearWords = append(swearWords, swearWord)

	deleteBeforeDate, _ := jodaTime.Parse("dd/MM/yyyy", "16/11/2018")

	tweetsToDelete := cleanTweets(tweets, swearWords, deleteBeforeDate, false)
	assert.Equal(t, 1, len(tweetsToDelete), "One tweet will be removed.")
}