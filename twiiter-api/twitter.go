package twitterapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"twitterservice/payload"
)


func GetLatestTweetByUserId(userId, sinceId string) ([]payload.Tweet, error) {
	user_id := userId
	url := fmt.Sprintf("https://api.twitter.com/2/users/%s/tweets?tweet.fields=created_at&max_results=10&since_id=%s", user_id, sinceId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("build req failed, url:", url)
		return nil, err
	}
	resp, err := httpRequest(req, nil)
	if err != nil {
		fmt.Println("build req failed, url:", url)
		return nil, err
	}

	respBody := resp.([]byte)
	var t payload.Tweets
	err = json.Unmarshal(respBody, &t)
	if err != nil {
		return nil, err
	}

	//fmt.Println(len(t.Data))
	//fmt.Println(string(respBody))
	/*
	{"data":[{"created_at":"2021-12-26T01:05:48.000Z","id":"1474909274459578380","text":"@NASA @NASAWebb Congratulations, this is major!"},{"c
	reated_at":"2021-12-26T01:04:04.000Z","id":"1474908839979925508","text":"@WorldAndScience Me \uD83D\uDCAF"},{"created_at":"2021-12-25T21:5
	3:31.000Z","id":"1474860884359094275","text":"Floki Santa https://t.co/y3CTq16bGi"},{"created_at":"2021-12-25T20:17:37.000Z","id":"1474836
	752619741188","text":"Merry Christmas \uD83C\uDF84\uD83C\uDF81 ⛄️ ⭐️ \uD83D\uDE03"},{"created_at":"2021-12-25T20:07:59.000Z","id":"147
	4834326571397120","text":"@JonErlichman Or Tesla!"}],"meta":{"oldest_id":"1474834326571397120","newest_id":"1474909274459578380","result_c
	ount":5,"next_token":"7140dibdnow9c7btw3z3pz9qwlflxnl3ul8ouohoy8yeo"}}
	 */
	return t.Data, nil
}

func GetFollowersByUserId(userId string) ([]payload.TwitterUser, error) {
	user_id := userId
	url := fmt.Sprintf("https://api.twitter.com/1.1/followers/list.json?user_id=%s&cursor=-1&&skip_status=true&include_user_entities=false", user_id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("build req failed, url:", url)
		return nil, err
	}
	resp, err := httpRequest(req, nil)
	if err != nil {
		fmt.Println("build req failed, url:", url)
		return nil, err
	}

	respBody := resp.([]byte)
	var t payload.TwitterUsers
	err = json.Unmarshal(respBody, &t)
	if err != nil {
		return nil, err
	}
	return t.Users, nil
}

func GetUserIdByUserName(userName string) (string, error) {
	usernames := "usernames=elonmusk"
	user_fields := "user.fields=description,created_at"
	url := fmt.Sprintf("https://api.twitter.com/2/users/by?%s&%s", usernames, user_fields)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("build req failed, url:", url)
		return "", err
	}
	resp, err := httpRequest(req, nil)
	if err != nil {
		fmt.Println("build req failed, url:", url)
		return "", err
	}

	respBody := resp.([]byte)
	var t payload.Tweets
	err = json.Unmarshal(respBody, &t)
	if err != nil {
		return "", err
	}
	fmt.Println(string(respBody))
	return "", nil

}
func httpRequest(req *http.Request, respType interface{}) (interface{}, error) {
	var bearer_token = `AAAAAAAAAAAAAAAAAAAAAN9KWwEAAAAASIg1Q%2Fc%2FMzehE2tqXERqQLl8LGk%3D8FFz9LBEVyvwh3Y5Tgl1VFeFle7m1dN7IUlyz8oftic8S54gpp`
	bearer := "Bearer " + bearer_token
	req.Header.Add("content-type", "text/plain")
	req.Header.Add("Authorization", bearer)
	req.Header.Add("User-Agent", "v2UserLookupGo")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return respBody, nil
	//switch respType.(type) {
	//case payload.Tweets:
	//	var t payload.Tweet
	//	err = json.Unmarshal(respBody, &t)
	//	if err != nil {
	//		return nil, err
	//	}
	//	return t, nil
	//default:
	//	fmt.Println("unsupported type")
	//}
}

