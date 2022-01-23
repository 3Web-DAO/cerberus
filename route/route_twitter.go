package route

import (
    "twitterservice/handler"
    "twitterservice/middlewares"
    "twitterservice/payload"
)

func twitterRouteInit() {
    twitterRoute := app.Party("/twitteruser")
    twitterRoute.Post("/add", middlewares.NewPayloadMiddleware(payload.ReqTwitterUser{}), handler.AddTwitterUser)
    twitterRoute.Post("/list", middlewares.NewPayloadMiddleware(payload.ReqTwitterUserList{}), handler.ListTwitterUser)
}
