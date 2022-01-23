package middlewares

import (
    "encoding/json"
    "fmt"
    "github.com/kataras/iris/v12"
    "reflect"
    "twitterservice/constant"
)

type PayloadMiddleware struct {
    pltype      reflect.Type
}

func NewPayloadMiddleware(payload interface{}) iris.Handler {
    pltype := reflect.TypeOf(payload)
    handler := &PayloadMiddleware{pltype: pltype}
    return handler.parse
}

func (pm *PayloadMiddleware) parse(ctx iris.Context) {
    if pm.pltype == nil {
        ctx.Next()
    }

    payload := reflect.New(pm.pltype).Interface()
    reqBody, err := ctx.GetBody()
    if err != nil {
        return
    }

    if response := payloadParse(payload, reqBody); response != nil {
        return
    }
    ctx.Values().Set(constant.ROUTE_VALUE, payload)
    ctx.Next()
}

func payloadParse(data interface{}, body []byte) error {
    if len(body) == 0 {
        return nil
    }
    err := json.Unmarshal(body, &data)
    if err != nil {
        fmt.Println("request body json fail")
        return err
    }
    return nil
}