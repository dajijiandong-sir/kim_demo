package server

import "context"
/*********************************************
* @Author: 戴建东
* @Date: 2021/8/26 15:54
* @File: start.go
* @Pack: server
* @Pro: kim_demo
* @Ides: GoLand
* @Desc:
**********************************************/

type ServerStartOptions struct {
	id     string
	listen string
}

func RunServerStart(ctx context.Context, opts *ServerStartOptions, version string) error {
	server := NewServer(opts.id, opts.listen)
	defer server.Shutdown()
	return server.Start()
}