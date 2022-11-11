package main

import (
	"github.com/fvbock/endless"
	"log"
	"middleground/app"
	"middleground/routes"
)

const env = "develop"

//const env = "master"

func main() {
	// 初始化加载配置
	app.Init(env)

	// 初始化路由配置
	r := routes.RouteApiInit()

	port := ":8800"

	if env == "develop" {
		_ = r.Run(port)
	} else {
		if err := endless.ListenAndServe(port, r); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}
}
