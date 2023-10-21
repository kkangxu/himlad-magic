module github.com/kkangxu/himlad-magic

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/IBM/sarama v1.41.3
	github.com/Shopify/sarama v1.29.1 // indirect
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/asim/go-micro/plugins/broker/kafka/v3 v3.7.0
	github.com/asim/go-micro/plugins/registry/etcd/v3 v3.7.0
	github.com/asim/go-micro/plugins/transport/grpc/v3 v3.7.0
	github.com/asim/go-micro/plugins/wrapper/monitoring/prometheus/v3 v3.7.0
	github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3 v3.7.0
	github.com/asim/go-micro/v3 v3.7.1
	github.com/gin-gonic/gin v1.9.1
	github.com/gomodule/redigo v1.8.9
	github.com/kkangxu/himlad-proto v0.0.0-20231021110838-eaa5089b1cb6
	github.com/opentracing/opentracing-go v1.2.0
	github.com/prometheus/client_golang v1.11.0
	github.com/spf13/viper v1.7.1
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/uptrace/opentelemetry-go-extra/otelgorm v0.2.3
	go.uber.org/zap v1.17.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	gorm.io/driver/mysql v1.5.2
	gorm.io/driver/sqlserver v1.5.2
	gorm.io/gorm v1.25.5
	gorm.io/plugin/dbresolver v1.4.7
)
