package main

import (
	"log"
	"time"

	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

type fxMetrics struct {
	rate metrics.Gauge
}

// NewMetrics initialises the metrics
func NewMetrics() fxMetrics {
	rate := kitprometheus.NewGaugeFrom(prometheus.GaugeOpts{
		Namespace: "fx",
		Name:      "rate",
		Help:      "fx rate is the exchange rate between the base and quote e.g. USD and GBP.",
		//ConstLabels: prometheus.Labels{},
	}, []string{"base", "quote"})

	return fxMetrics{rate}
}

// collect collects metrics
func (m fxMetrics) collect(appId string) error {
	res, err := request(appId)
	if err != nil {
		return errors.Wrap(err, "requestFixture for data failed")
	}

	for quoteSymbol, quotePrice := range res.Rates {
		// TODO: take base param
		m.rate.With("base", "USD", "quote", quoteSymbol).Set(quotePrice)
	}

	return nil
}

func runCollector(appId string, m fxMetrics, duration time.Duration) chan error {
	errs := make(chan error)
	ticker := time.NewTicker(duration)

	go func() {
		for {
			log.Println("running collection")
			err := m.collect(appId)
			log.Println("finished collection")
			if err != nil {
				errs <- err
			}
			<-ticker.C
		}
	}()

	return errs
}
