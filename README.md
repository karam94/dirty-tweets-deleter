# dirty-tweets-deleter
A small project to try out Go as part of a personal monthly challenge on my [blog](http://www.karam.io/2018/November-2018-Best-thing-since-bread/). This console application allows users to bulk delete all old tweets prior to a specified date with the option to also instead specify a comma separated file of swear words which then deletes only tweets prior to the specified date that contains one of the swear words.

**DISCLAIMER:** Running this application MAY delete all of your tweets. I am not to be held responsible for the result of your actions of using this application.

## Getting started
You need a [Twitter developers account](https://apps.twitter.com/) and a registered application. You can then specify your application's Twitter API consumer key, consumer secret, access token and access secret within a file called "config.json". Use "config.json.template" as a template to work off.

You can download the archive of your Twitter accounts tweets by going to Twitter > Settings > Account > Content > Request your archive. Place the tweet archive CSV file e-mailed to you in the root folder of your application as "tweets.csv".

If you want to delete only tweets that contain specific words, place your own CSV of words in the root folder as "swearwords.csv". You can download a generic one from [here](http://www.bannedwordlist.com/) if you prefer.

## How to use
Providing you have [Go](https://golang.org/) installed on your machine, you can compile the code by running `go build` in the root directory which will give you an .exe file to run. E.g.`./main.go` or `./main.exe`. You can also run the unit tests by running `go test` in the root directory.

## dirty-tweets-deleter-web
As this application was created for learning purposes, I also created a similar web based alternative that can be found [here](https://github.com/karam94/dirty-tweets-deleter-web).

## Credits
- [go-twitter by dghubble](https://github.com/dghubble/go-twitter)
- [Go OAuth1 by dghubble](https://github.com/dghubble/oauth1)
- [jodaTime by vjeantet](https://github.com/vjeantet/jodaTime)
