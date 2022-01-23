package route

import (
    "github.com/kataras/iris/v12"
)

var app *iris.Application

func Init() {
    app = iris.New()

    userRouteInit()
    twitterRouteInit()
    emailRouteInit()

    app.Listen(":8080")
}

