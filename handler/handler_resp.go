package handler

type Resp struct {
    State       int32 `json:"state"`
    Msg         string `json:"msg"`
}

var RespSuccess = Resp{0, ""}

var RespSqlFail = Resp{1000, "execute sql failed"}
var RespSqlInsertFail = Resp{1001, "affected rows is zero"}
var RespSqlGetFail = Resp{1001, "query rows is zero"}
var RespSqlFindFail = Resp{1001, "find sql failed"}
var RespTwitterFail = Resp{2000, "request twitter api failed"}

