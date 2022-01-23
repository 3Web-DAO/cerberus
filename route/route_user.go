package route

import "twitterservice/handler"
import "twitterservice/middlewares"
import "twitterservice/payload"

func userRouteInit() {
    userRoute := app.Party("/user")
    userRoute.Post("/add", middlewares.NewPayloadMiddleware(payload.ReqUser{}),handler.AddUser)

}
