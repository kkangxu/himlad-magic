package wrapper

import (
	"context"
	"reflect"

	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/server"
	"github.com/kkangxu/himlad-magic/log"
	"github.com/kkangxu/himlad-proto/base"
)

type logWrapper struct {
	client.Client
}

func (o *logWrapper) Publish(ctx context.Context, p client.Message, opts ...client.PublishOption) error {

	log.Infow("start client Publish logWrapper", "topic", p.Topic(), "contentType", p.ContentType())
	if err := o.Client.Publish(ctx, p, opts...); err != nil {
		log.Infow("in client Publish logWrapper", "topic", p.Topic(), "contentType", p.ContentType(), "err", err)
		return err
	}

	log.Infow("end client Publish logWrapper", "topic", p.Topic(), "contentType", p.ContentType())
	return nil
}

func (w *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {

	log.Infow("start client Call logWrapper", "method", req.Method(), "body", req.Body())

	err := w.Client.Call(ctx, req, rsp, opts...)
	if err != nil {
		log.Infow("in client Call logWrapper", "method", req.Method(), "err", err)
		return err
	}

	result, ok := reflectBaseRet(rsp)
	if result == nil || !ok {
		log.Warnw("in client Call LogWrapper, unknown result type", "method", req.Method(), "result", result)
		return nil
	}

	log.Infow("end client Call logWrapper", "method", req.Method(), "rsp", rsp)
	return nil
}

func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}

func LogWrapper(fn server.HandlerFunc) server.HandlerFunc {

	return func(ctx context.Context, req server.Request, rsp interface{}) error {

		log.Infow("start handler LogWrapper", "headers", req.Header(), "method", req.Method(), "body", req.Body())
		defer func() {
			result, ok := reflectBaseRet(rsp)
			if result == nil || !ok {
				log.Warnw("in handler LogWrapper, unknown result type", "method", req.Method(), "result", result)
				return
			}
			log.Infow("end handler LogWrapper", "method", req.Method(), "code", result.Code, "msg", result.Msg)
		}()

		return fn(ctx, req, rsp)
	}
}

func reflectBaseRet(rsp interface{}) (*base.BaseRet, bool) {
	v := reflect.ValueOf(rsp)
	f := v.Elem().FieldByName("Result")
	if f.IsValid() {
		r := f.Interface()
		result, ok := r.(*base.BaseRet)
		return result, ok
	}
	return nil, false
}
