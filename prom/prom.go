package prom

import (
	"net"
	"os"
	"time"

	"github.com/kkangxu/himlad-magic/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	ReqGinCounterTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "counter_gin_requests_total",
			Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
		[]string{"code", "method", "path"},
	)

	ReqGinSummaryDur = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "summary_gin_request_duration_seconds",
			Help:       "The HTTP request latencies in seconds.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.05, 0.95: 0.05, 0.99: 0.05},
		},
		[]string{"code", "method", "path"},
	)

	ReqGinHistogramDur = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "histogram_gin_request_duration_seconds",
			Help:    "Histogram of response latency (seconds) of http handlers.",
			Buckets: []float64{0.0005, 0.001, 0.002, 0.005, 0.010, 0.020, 0.050, 0.1, 0.5, 1, 5},
		}, []string{"code", "method", "path"})

	ReqGinSummarySizeByte = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "summary_gin_request_size_bytes",
			Help: "The HTTP request sizes in bytes.",
		}, []string{})

	RspGinSummarySizeByte = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "summary_gin_response_size_bytes",
			Help: "The HTTP response sizes in bytes.",
		}, []string{})
)

func RunProm(service, url string, collectors ...prometheus.Collector) {
	for i := 0; i < len(collectors); i++ {
		prometheus.MustRegister(collectors[i])
	}
	pusher := push.New(url, service).Gatherer(prometheus.DefaultGatherer).Grouping("instance", getIPV4())
	go func() {
		for {
			time.Sleep(time.Second * 15)
			err := pusher.Push()
			if err != nil {
				log.Infow("pusher.Push err", "err", err)
				continue
			}
			log.Info("... pusher.Push ...")
		}
	}()
}

func getIPV4() string {
	var ip string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		ipNet, ok := a.(*net.IPNet)
		if !ok {
			continue
		}

		if !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ip = ipNet.IP.String()
			break
		}
	}
	return ip
}
