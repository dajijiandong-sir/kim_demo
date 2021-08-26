package main

import (
	"context"
	"kim_demo/server"
)

/*********************************************
* @Author: 戴建东
* @Date: 2021/8/26 15:43
* @File: main.go
* @Pack: kim_demo
* @Pro: kim_demo
* @Ides: GoLand
* @Desc:
**********************************************/
const version = "v1"

func main() {
	ctx := context.Background()
	opt := &server.ServerStartOptions{}
	server.RunServerStart(ctx, opt, version)
}
