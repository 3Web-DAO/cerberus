package twitterapi

import "testing"

func TestGetUserIdByUserName(t *testing.T) {
    userName, err := GetUserIdByUserName("elonmusk")
    if err != nil {
        t.Error(err)
        return
    }
    t.Log(userName)
}

func TestGetLatestTweetByUserIds(t *testing.T) {
    lastTweets, err := GetLatestTweetByUserId("44196397", "1476473842059161602")
    if err != nil {
        t.Error(err)
        return
    }
    t.Log(lastTweets)
}

func TestGetFollowersByUserId(t *testing.T) {
    followers, err := GetFollowersByUserId("44196397")
    if err != nil {
        t.Error(err)
        return
    }
    t.Log(followers)
}
