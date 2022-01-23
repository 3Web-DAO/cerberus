package main

import (
	"fmt"
	"time"
	"twitterservice/models"
	"twitterservice/notification"
	_ "twitterservice/notification"
	"twitterservice/route"
	"twitterservice/twiiter-api"
)


func tweetTask() {
    const TwitterTimeout = 15 //minute
    const MaxRequests = 450
	session := models.Engine.NewSession()
	defer session.Close()

	for {
		time.Sleep(TwitterTimeout * time.Minute)
		//time.Sleep(TwitterTimeout * time.Second)
		var twitterUsers= make([]models.TwitterUser, 0)
		err := session.Find(&twitterUsers)
		if err != nil {
			fmt.Println(err)
			continue
		}

		userId2emailAddress := make(map[string]string, 0)
		for _, twitterUser := range twitterUsers {
			email := models.Email{UserId: twitterUser.UserId}
		    isExist, err := session.Get(&email)
		    if err != nil || !isExist {
		    	fmt.Println("debug0", isExist)
		    	continue
			}
		    userId2emailAddress[twitterUser.UserId] = email.EmailAddress
		}

		count := 0
		for _, twitterUser := range twitterUsers {
			count++
			if count > MaxRequests {
				fmt.Println("out of service")
				break
			}
			fmt.Println("debug0", twitterUser.UserId, twitterUser.LastTweetId)
			lastTweets, err := twitterapi.GetLatestTweetByUserId(twitterUser.Id, twitterUser.LastTweetId)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if len(lastTweets) == 0 {
				continue
			}
			fmt.Println("debug1", len(lastTweets))
			maxTweetId := "0"
			for _, tweet := range lastTweets {
				content := "time: " + tweet.CreatedAt + ", content:" + tweet.Text
				fmt.Println("debug2", content)
				_ = content
				if tweet.Id > maxTweetId {
					maxTweetId = tweet.Id
				}
				err := notification.SendEmail(userId2emailAddress[twitterUser.UserId], twitterUser.Name, content)
				if err != nil {
					continue
				}
			}
			fmt.Println("debug3", maxTweetId, twitterUser.UserId)
			session.Update(&models.TwitterUser{LastTweetId: maxTweetId}, models.TwitterUser{UserId: twitterUser.UserId})
		}
		//fmt.Println("exit")
		//break
	}
}

func main() {
    //init database
    //models.Init("root:hep2uv?s1TZwV@(127.0.0.1:3306)/blockchaindata?charset=utf8")
	fmt.Println("debug0")
	models.Init("root:hep2uv?s1TZwV@(1.14.148.7:3306)/blockchaindata?charset=utf8")
	fmt.Println("debug1")

	//setup time task
	go tweetTask()

	//init iris
	fmt.Println("debug2")
	route.Init()
	fmt.Println("debug3")

}