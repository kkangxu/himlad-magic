module github.com/kkangxu/himlad-magic/wrapper

go 1.16

require (
	github.com/kkangxu/himlad-magic/log v0.0.0-00010101000000-000000000000
	github.com/kkangxu/himlad-proto v0.0.0-00010101000000-000000000000
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/asim/go-micro/v3 v3.7.1
	github.com/gin-gonic/gin v1.7.7
	github.com/opentracing/opentracing-go v1.2.0
)

replace (
	github.com/kkangxu/himlad-magic/log => ../log
	github.com/kkangxu/himlad-proto => ../../proto
)
