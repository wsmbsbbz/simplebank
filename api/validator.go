package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/wsmbsbbz/simplebank/util"
)

var supportedCurrencies = map[string]bool{
	"USD": true,
	"EUR": true,
	"CAD": true,
}

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
