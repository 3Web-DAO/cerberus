package payload

type Tweet struct {
    CreatedAt	string `json:"created_at"`
    Text 		string `json:"text"`
    Id          string `json:"id"`
}
type Tweets struct {
    Data 	[]Tweet `json:"data"`
}

type TwitterUser struct {
    Name 		string `json:"name"`
    Id 			string `json:"id_str"`
}

type TwitterUsers struct {
    Users 		[]TwitterUser `json:"users"`
}
