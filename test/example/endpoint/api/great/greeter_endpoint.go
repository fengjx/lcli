// Code generated by lc pbgen. DO NOT EDIT.
package great

import (
	"github.com/fengjx/lctest/proto/pbgreet"
	"github.com/fengjx/luchen"
)

// RegisterGreeterGRPCHandler 注册 GRPC 服务处理器
func RegisterGreeterGRPCHandler(gs *luchen.GRPCServer) {
	pbgreet.RegisterGreeterGRPCHandler(gs, GreeterEndpointImpl)
}

// RegisterGreeterHTTPHandler 注册 HTTP 服务处理器
func RegisterGreeterHTTPHandler(hs *luchen.HTTPServer) {
	pbgreet.RegisterGreeterHTTPHandler(hs, GreeterEndpointImpl)
}

// GreeterEndpointImpl 默认的服务实现
var GreeterEndpointImpl = &GreeterEndpoint{
	handler: &GreeterHandlerImpl{},
}

// GreeterHandlerImpl 服务处理器实现
type GreeterHandlerImpl struct {
}

// GreeterEndpoint 服务 Endpoint 定义
type GreeterEndpoint struct {
	handler pbgreet.GreeterHandler
}
