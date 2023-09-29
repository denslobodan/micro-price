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

// 1. Найти созданную в goose базу данных
// 2. Зачем открывать в каждом методе соединение с базой
// 3. Что такое формат hcl и как с ним работать
// 4. Вычищение html тегов
