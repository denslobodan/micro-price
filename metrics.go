package main

import "context"

type MetricService struct {
	next PriceService
}

func NewMetricSevice(next PriceService) PriceService {
	return &MetricService{next: next}
}

func (s *MetricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	//  your metrcis storage. Push to prometheus (gauge, counters)
	return s.next.FetchPrice(ctx, ticker)
}
