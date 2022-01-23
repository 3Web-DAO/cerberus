package route

import "twitterservice/handler"
import "twitterservice/middlewares"
import "twitterservice/payload"

func emailRouteInit() {
    userRoute := app.Party("/email")
    userRoute.Post("/add", middlewares.NewPayloadMiddleware(payload.ReqEmailAdd{}),handler.AddEmail)

}
