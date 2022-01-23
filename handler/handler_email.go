package handler

import (
    "encoding/json"
    "github.com/kataras/iris/v12"
    "twitterservice/constant"
    "twitterservice/models"
    "twitterservice/payload"
)

func AddEmail(ctx iris.Context) {
    reqPara := ctx.Values().Get(constant.ROUTE_VALUE).(*payload.ReqEmailAdd)
    session := models.Engine.NewSession()
    var resp []byte
    defer session.Close()
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

    affectedRows, err := session.Insert(&models.Email{
        UserId:     user.UserId,
        EmailAddress: reqPara.Email,
    })

    if err != nil {
        resp, _ = json.Marshal(RespSqlFail)
    } else if affectedRows == 0 {
        resp, _ = json.Marshal(RespSqlInsertFail)
    }

    resp, _ = json.Marshal(RespSuccess)
}