// grpc拦截器基本原理，链式调用注册的拦截函数,适用于统一入口，
// 需要多层处理的部分，可以进行很好的业务解偶
package main

import (
	"context"
	"net/http"
)

func main() {
	ctx := context.Background()
	req := http.Request{}
	// grpc实际拦截器
	interceptor := ChainInterceptor(registerInterceptor()...)
	resp, err := interceptor(ctx, req, nil, handler)
	print(resp, err)
}

// handler grpc实际执行器
func handler(ctx context.Context, req interface{}) (resp interface{}, err error) {
	println("I am grpc handler")
	return http.Response{}, nil
}

type UnaryServerInfo struct {
}

// Handler 处理函数可以自己根据业务而定义，grpc中是处理http请求，所以一个req然后一个resp，ctx存metadata,
// 该函数用于真正处理grpc信号
type Handler func(ctx context.Context, req interface{}) (resp interface{}, err error)

// Interceptor 拦截器，该方法的设计特点：入参要有handler的入参以及handler，返回值要与handler相同
// info这种附带的值可以根据业务需要而定
type Interceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler Handler) (resp interface{}, err error)

// ChainInterceptor 拦截器调用链，生成最终链式拦截器
func ChainInterceptor(interceptorArr ...Interceptor) Interceptor {
	n := len(interceptorArr)

	return func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler Handler) (resp interface{}, err error) {
		chainerFunc := func(interceptor Interceptor, handler Handler) Handler {
			return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
				// 拦截器实际上在这里被装载进去，并被链式调用
				return interceptor(ctx, req, info, handler)
			}
		}
		chanHandler := handler
		// 倒序是为了顺序执行拦截器
		for i := n-1; i >= 0; i-- {
			// 实际上的handler函数在这里被层层传递进去拦截器链中
			chanHandler = chainerFunc(interceptorArr[i], chanHandler)
		}
		return chanHandler(ctx, req)
	}
}

// registerInterceptor 可以注册自己想要注册的拦截器函数
func registerInterceptor() []Interceptor {
	// grpc的默认拦截器
	defaultInterceptor := []Interceptor{func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler Handler) (resp interface{}, err error) {
		println("I am default interceptor")
		return handler(ctx, req)
	}}
	return append(defaultInterceptor,
		interceptor1,
		interceptor2,
		interceptor3)
}

func interceptor1(ctx context.Context, req interface{}, _ *UnaryServerInfo, handler Handler) (resp interface{}, err error) {
	println("interceptor1")
	// 这里可以对传进来的ctx和req进行操作,如注入全局参数，或是对grpc的传参req进行检验等
	return handler(ctx, req)
}

func interceptor2(ctx context.Context, req interface{}, _ *UnaryServerInfo, handler Handler) (resp interface{}, err error) {
	println("interceptor2")
	// 这里可以对传进来的ctx和req进行操作,如注入全局参数，或是对grpc的传参req进行检验等
	return handler(ctx, req)
}

func interceptor3(ctx context.Context, req interface{}, _ *UnaryServerInfo, handler Handler) (resp interface{}, err error) {
	println("interceptor3")
	// 这里可以对传进来的ctx和req进行操作,如注入全局参数，或是对grpc的传参req进行检验等
	return handler(ctx, req)
}
