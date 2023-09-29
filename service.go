package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

var prices = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 999.99,
	"GG":  100_000.0,
}

// PriceFetcher is an interface that can fetch a price.
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceService implements the PriceFfetcher interface.
type priceService struct{}

func (s *priceService) FetchPrice(_ context.Context, ticker string) (float64, error) {
	price, ok := prices[ticker]
	if !ok {
		return 0.0, fmt.Errorf("price for ticker (%s) is not available", ticker)
	}
	return price, nil
}

type logginingService struct {
	priceService
}

func (s logginingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		reqID := ctx.Value("requestID")

		logrus.WithFields(logrus.Fields{
			"requestID": reqID,
			"took":      time.Since(begin),
			"err":       err,
			"price":     price,
			"ticker":    ticker,
		}).Info("FetchPrice")
	}(time.Now())

	return s.priceService.FetchPrice(ctx, ticker)
}
