package main

import (
	"fmt"
	"github.com/kataras/iris"
	"strconv"
)

func main() {
	app := iris.Default()
	app.Get("/hello", func(ctx iris.Context) {
		ctx.WriteString("hello world by iris")
	})
	app.Get("/user/{id:uint64 min(2)}", GetUserInfo)
	app.Get("/orders/{action:string prefix(a_)}", func(ctx iris.Context) {
		ctx.WriteString("hello world by iris")
	})
	err := app.Run(iris.Addr(":8082"))
	fmt.Println(err)
}
func GetUserInfo(ctx iris.Context) {
	id := ctx.Params().GetUint64Default("id", 0)
	ctx.WriteString(strconv.Itoa(int(id)))
}
