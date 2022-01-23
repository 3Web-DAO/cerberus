package payload

type ReqUser struct {
    Name        string `json:"user_name"`
    Pwd         string `json:"password"`
}

type ReqTwitterUser struct {
    UserName    string `json:"user_name"`
    Name        string `json:"name"`
}

type ReqTwitterUserList struct {
    UserName    string `json:"user_name"`
    StartIndex    int `json:"start_index"`
    SearchNum     int `json:"search_num"`
}

type RespTwitterUserListData struct {
    Name        string `json:"name"`
    Id          string `json:"id"`
}

type RespTwitterUserList struct {
    State       string `json:"state"`
    Data        []RespTwitterUserListData `json:"data"`
}

type ReqEmailAdd struct {
    UserName    string `json:"user_name"`
    Email       string `json:"email_address"`
}
