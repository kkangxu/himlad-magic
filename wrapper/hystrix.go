package wrapper

import (
	"context"
	"errors"
	"reflect"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/server"
	"github.com/kkangxu/himlad-magic/log"
	"github.com/kkangxu/himlad-proto/base"
	"github.com/kkangxu/himlad-proto/code"
)

type hystrixWrapper struct {
	client.Client
}

func (pw *hystrixWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	hystrix.ConfigureCommand(cmdName, hystrix.CommandConfig{})

	log.Infow("start hystrixWrapper", "headers", cmdName)

	defer func() {
		result, ok := reflectBaseRet(rsp)
		if result == nil || !ok {
			log.Warnw("in hystrixWrapper, unknown result type", "method", req.Method(), "result", result)
			return
		}
	}()

	var hystrixErr error
	err := hystrix.Do(
		cmdName,
		func() error {
			return pw.Client.Call(ctx, req, rsp, opts...)
		},
		func(err error) error {
			if err != nil {
				hystrixErr = err
				log.Warnw("in hystrixWrapper err", "method", req.Method(), "err", err)
				return nil
			}
			return nil
		},
	)

	defer func() {
		if hystrixErr != nil {
			v := reflect.ValueOf(rsp)
			f := v.Elem().FieldByName("Result")
			f.Set(reflect.ValueOf(WrapperBase(hystrixErr)))
		}
	}()

	return err
}

func NewHystrixWrapper(c client.Client) client.Client {
	return &hystrixWrapper{c}
}

func HystrixWrapper(fn server.HandlerFunc) server.HandlerFunc {

	return func(ctx context.Context, req server.Request, rsp interface{}) error {

		cmdName := req.Service() + "." + req.Endpoint()
		hystrix.ConfigureCommand(cmdName, hystrix.CommandConfig{})

		var hystrixErr error
		err := hystrix.Do(
			cmdName,
			func() error {
				return fn(ctx, req, rsp)
			},
			func(err error) error {
				if err != nil {
					hystrixErr = err
					log.Warnw("in hystrixWrapper err", "method", req.Method(), "err", err)
					return nil
				}
				return nil
			},
		)

		defer func() {
			if hystrixErr != nil {
				v := reflect.ValueOf(rsp)
				f := v.Elem().FieldByName("Result")
				f.Set(reflect.ValueOf(WrapperBase(hystrixErr)))
			}
		}()

		return err
	}
}

func WrapperBase(err error) *base.BaseRet {
	b := &base.BaseRet{
		Code: code.ErrCode_InternalError,
	}

	if errors.Is(err, hystrix.ErrMaxConcurrency) || errors.Is(err, hystrix.ErrCircuitOpen) || errors.Is(err, hystrix.ErrTimeout) {
		b = &base.BaseRet{
			Code: code.ErrCode_LimitExceeded,
			Msg:  err.Error(),
		}
	}

	return b
}
