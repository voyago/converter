package conversion

import (
	"github.com/voyago/converter/pkg/model"
	"github.com/voyago/converter/pkg/store"
)

type Converter struct {
	store store.Store
}

func Make(store store.Store) Converter {
	return Converter{store: store}
}

func (current Converter) Convert(price model.Price) (model.Price, error) {
	rates, err := current.store.GetExchangeRates()

	if err != nil {
		return model.Price{}, err
	}

	currency, err := rates.Find(current.store.GetSource())

	if err != nil {
		return model.Price{}, err
	}

	return ConvertTo(price, currency)
}

func ConvertTo(price model.Price, currency model.Currency) (model.Price, error) {
	rate, err := currency.ToRatePrice()

	if err != nil {
		return model.Price{}, err
	}

	target := float64(price.Amount) / rate.ToFloat()
	result, _ := model.MakePrice(currency, target)

	return result, nil
}
