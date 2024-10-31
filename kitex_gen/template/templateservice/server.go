/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kitex v0.11.3. DO NOT EDIT.
package templateservice

import (
	server "github.com/cloudwego/kitex/server"

	template "github.com/west2-online/fzuhelper-server/kitex_gen/template"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler template.TemplateService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler template.TemplateService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
