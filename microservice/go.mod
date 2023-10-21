module microservice

go 1.16

require (
	github.com/kkangxu/himlad-magic/log v0.0.0-00010101000000-000000000000
	github.com/kkangxu/himlad-magic/morm v0.0.0-00010101000000-000000000000
	github.com/kkangxu/himlad-magic/prom v0.0.0-00010101000000-000000000000
	github.com/kkangxu/himlad-magic/mredis v0.0.0-00010101000000-000000000000
	github.com/kkangxu/himlad-magic/tracer v0.0.0-00010101000000-000000000000
	github.com/kkangxu/himlad-magic/wrapper v0.0.0-00010101000000-000000000000
	github.com/Shopify/sarama v1.29.1
	github.com/asim/go-micro/plugins/broker/kafka/v3 v3.7.0
	github.com/asim/go-micro/plugins/registry/etcd/v3 v3.7.0
	github.com/asim/go-micro/plugins/transport/grpc/v3 v3.7.0
	github.com/asim/go-micro/plugins/wrapper/monitoring/prometheus/v3 v3.7.0
	github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3 v3.7.0
	github.com/asim/go-micro/v3 v3.7.1
	github.com/gomodule/redigo v1.8.8 // indirect
	github.com/spf13/viper v1.7.1
)

replace (
	github.com/kkangxu/himlad-magic/log => ../log
	github.com/kkangxu/himlad-magic/morm => ../morm
	github.com/kkangxu/himlad-magic/prom => ../prom
	github.com/kkangxu/himlad-magic/mredis => ../mredis
	github.com/kkangxu/himlad-magic/tracer => ../tracer
	github.com/kkangxu/himlad-magic/wrapper => ../wrapper
	github.com/kkangxu/himlad-proto => ../../proto
)
