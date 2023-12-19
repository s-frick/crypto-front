package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/s-frick/crypto-front/model"
	"github.com/s-frick/crypto-front/view"
	"github.com/s-frick/crypto-front/view/orderbook"
)

type OrderbookHandler struct {
}

func generateData() []model.Limit {
	limits := make([]model.Limit, 0)
	for i := 0; i < 40; i++ {
		if i < 20 {
			asks := make([]model.Order, 1)
			limits = append(limits, model.Limit{
				Asks:  append(asks, model.Order{Type: model.ASK, Size: 1}),
				Bids:  make([]model.Order, 0),
				Price: (1000.0 + (float64(i * 10))),
			})

		} else {
			bids := make([]model.Order, 1)
			limits = append(limits, model.Limit{
				Asks:  make([]model.Order, 0),
				Bids:  append(bids, model.Order{Type: model.BID, Size: 1}),
				Price: (1000.0 + (float64(i * 10))),
			})
		}
	}
	for i, j := 0, len(limits)-1; i < j; i, j = i+1, j-1 {
		limits[i], limits[j] = limits[j], limits[i]
	}
	return limits
}

func (h OrderbookHandler) OrderbookShow(c echo.Context) error {
	limits := generateData()

	return view.Render(c, orderbook.Show(limits))
}
