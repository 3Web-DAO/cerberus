package handler

import (
    "encoding/json"
    "github.com/google/uuid"
    "github.com/kataras/iris/v12"
    "time"
    "twitterservice/constant"
    "twitterservice/models"
    "twitterservice/payload"
)

func AddUser(ctx iris.Context) {
    user := ctx.Values().Get(constant.ROUTE_VALUE).(*payload.ReqUser)
    session := models.Engine.NewSession()
    defer session.Close()

    userId := uuid.New().String()[:6]
    affectedRows, err := session.Insert(&models.User{
        UserId:     userId,
        UserName:   user.Name,
        Password:   user.Pwd,
        CreateTime: time.Now(),
    })

    var resp []byte
    resp, _ = json.Marshal(RespSuccess)
    if err != nil {
        resp, _ = json.Marshal(RespSqlFail)
    } else if affectedRows == 0 {
        resp, _ = json.Marshal(RespSqlInsertFail)
    }
    ctx.WriteString(string(resp))
}



