package store

import (
	"errors"
	"fmt"
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/src/store/handler"
)

type Store struct {
	Currency string
	Env      environment.Env
	Handler  handler.Handler
}

func MakeStore(driver string, currency string) (*Store, error) {
	env, err := environment.Make()

	if err != nil {
		return nil, errors.New("the store was unable to load the environment information")
	}

	store := &Store{
		Currency: currency,
		Env: *env,
	}

	if err := store.build(driver); err != nil {
		return nil, err
	}

	return store, nil
}

func (current *Store) build(driver string) error {
	switch driver {
	case "currency-layer":
		(*current).Handler = handler.CurrencyLayer{
			Env: current.Env,
			Currency: current.Currency,
		}
	default:
		return errors.New(fmt.Sprintf("The given driver [%s] is invalid", driver))
	}

	return nil
}
