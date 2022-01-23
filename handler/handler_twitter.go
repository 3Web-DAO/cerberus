package handler

import (
    "encoding/json"
    "github.com/kataras/iris/v12"
    "twitterservice/constant"
    "twitterservice/models"
    "twitterservice/payload"
    twitterapi "twitterservice/twiiter-api"
)


func AddTwitterUser(ctx iris.Context) {
    reqPara := ctx.Values().Get(constant.ROUTE_VALUE).(*payload.ReqTwitterUser)
    session := models.Engine.NewSession()
    defer session.Close()
    var resp []byte
    defer func() {
        ctx.WriteString(string(resp))
    }()

    twitterUserId, err := twitterapi.GetUserIdByUserName(reqPara.Name)
    if err != nil {
        resp, _ = json.Marshal(RespTwitterFail)
        return
    }

    var user = models.User{UserName: reqPara.UserName}
    {
        isExist, err := session.Get(&user)
        if err != nil {
            resp, _ = json.Marshal(RespSqlFail)
            return
        } else if !isExist {
            resp, _ = json.Marshal(RespSqlGetFail)
            return
        }
    }

    affectedRows, err := session.Insert(&models.TwitterUser{
        UserId:     user.UserId,
        Id:         twitterUserId,
        Name:       reqPara.Name,
    })

    resp, _ = json.Marshal(RespSuccess)
    if err != nil {
        resp, _ = json.Marshal(RespSqlFail)
        return
    } else if affectedRows == 0 {
        resp, _ = json.Marshal(RespSqlInsertFail)
        return
    }
}

func ListTwitterUser(ctx iris.Context) {
    reqPara := ctx.Values().Get(constant.ROUTE_VALUE).(*payload.ReqTwitterUserList)
    session := models.Engine.NewSession()
    defer session.Close()
    var resp []byte
    defer func() {
        ctx.WriteString(string(resp))
    }()

    var user = models.User{UserName: reqPara.UserName}
    {
        isExist, err := session.Get(&user)
        if err != nil {
            resp, _ = json.Marshal(RespSqlFail)
            return
        } else if !isExist {
            resp, _ = json.Marshal(RespSqlGetFail)
            return
        }
    }

    var twitterUsers = make([]models.TwitterUser, 0)
    err := session.Limit(reqPara.SearchNum, reqPara.StartIndex).Find(&twitterUsers,
        models.TwitterUser{UserId: user.UserId})
    if err != nil {
        resp, _ = json.Marshal(RespSqlFindFail)
        return
    }

    var data payload.RespTwitterUserList
    data.Data = make([]payload.RespTwitterUserListData, 0)
    for _, twitterUser := range twitterUsers {
        data.Data = append(data.Data, payload.RespTwitterUserListData{
            Name: twitterUser.Name,
            Id: twitterUser.Id,
        })
    }
    resp, _ = json.Marshal(data)
}
