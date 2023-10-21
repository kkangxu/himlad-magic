package microservice

import (
	"fmt"
	"sync"

	"github.com/IBM/sarama"
	"github.com/asim/go-micro/plugins/broker/kafka/v3"
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/plugins/transport/grpc/v3"
	"github.com/asim/go-micro/plugins/wrapper/monitoring/prometheus/v3"
	"github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/broker"
	"github.com/asim/go-micro/v3/registry"
	"github.com/kkangxu/himlad-magic/log"
	"github.com/kkangxu/himlad-magic/morm"
	"github.com/kkangxu/himlad-magic/mredis"
	"github.com/kkangxu/himlad-magic/prom"
	"github.com/kkangxu/himlad-magic/tracer"
	"github.com/kkangxu/himlad-magic/wrapper"
	"github.com/spf13/viper"
)

var MicroService micro.Service
var once sync.Once

func Config(configFile string) {
	initConfig(configFile)
}

func New(options ...micro.Option) micro.Service {
	once.Do(func() {
		options = append(options,
			micro.Name(viper.GetString("module.name")),
			micro.Version("latest"),
			micro.Transport(grpc.NewTransport()),
			micro.WrapClient(wrapper.NewLogWrapper),
			micro.WrapHandler(wrapper.LogWrapper),
			micro.WrapHandler(prometheus.NewHandlerWrapper()),
		)
		options = append(options, initEtcd()...)
		options = append(options, initBroker()...)
		options = append(options, initTracer()...)
		MicroService = micro.NewService(options...)
	})

	return MicroService
}

func initProm() {
	if !viper.GetBool("otel.metric.enable") {
		return
	}
	prom.RunProm(viper.GetString("broker.kafka.user"), viper.GetString("broker.kafka.user"))
}

func initBroker() []micro.Option {

	if !viper.GetBool("broker.kafka.enable") {
		return nil
	}
	saramaConfig := sarama.NewConfig()
	saramaConfig.Net.SASL.Enable = viper.GetBool("broker.kafka.enable_sasl")
	saramaConfig.Net.SASL.User = viper.GetString("broker.kafka.user")
	saramaConfig.Net.SASL.Password = viper.GetString("broker.kafka.password")
	saramaConfig.Version = sarama.V0_10_2_0

	return []micro.Option{
		micro.Broker(kafka.NewBroker(
			broker.Addrs(viper.GetString("broker.kafka.addr")),
			kafka.BrokerConfig(saramaConfig),
			kafka.ClusterConfig(saramaConfig),
		)),
	}
}

func initTracer() []micro.Option {
	if !viper.GetBool("otel.tracer.enable") {
		return nil
	}

	t, _, err := tracer.NewTracer(viper.GetString("otel.tracer.service"), viper.GetString("otel.tracer.url"))
	if err != nil {
		panic(err)
	}

	tracer.WrapperTracer()

	return []micro.Option{
		micro.WrapHandler(opentracing.NewHandlerWrapper(t)),
		micro.WrapClient(opentracing.NewClientWrapper(t)),
	}
}

func initEtcd() []micro.Option {
	if !viper.GetBool("registry.etcd.enable") {
		return nil
	}

	return []micro.Option{
		micro.Registry(
			etcd.NewRegistry(
				etcd.Auth(
					viper.GetString("registry.etcd.username"),
					viper.GetString("registry.etcd.password"),
				),
				registry.Addrs(
					fmt.Sprintf("%s:%s", viper.GetString("registry.etcd.host"), viper.GetString("registry.etcd.port")),
				),
			),
		),
	}
}

func initConfig(configFile string) {
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	log.Init(
		viper.GetString("log.path"),
		viper.GetString("log.level"),
		log.SetCompress(true),
		log.SetCaller(true),
	)

	var settings []*morm.DBSetting
	if err := viper.UnmarshalKey("dbs", &settings); err != nil {
		panic(err)
	}

	if len(settings) != 0 {
		if err := morm.InitInstance(settings); err != nil {
			panic(err)
		}
	}

	if viper.GetBool("redis.enable") {
		var conf mredis.RedisSetting
		if err := viper.UnmarshalKey("redis", &conf); err != nil {
			panic(err)
		}
		mredis.Init(&conf)
	}
}
