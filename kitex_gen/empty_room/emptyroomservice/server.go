// Code generated by Kitex v0.8.0. DO NOT EDIT.
package emptyroomservice

import (
	server "github.com/cloudwego/kitex/server"
	empty_room "github.com/west2-online/fzuhelper-server/kitex_gen/empty_room"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler empty_room.EmptyRoomService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}