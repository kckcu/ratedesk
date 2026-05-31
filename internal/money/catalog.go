package money

type CurrencyMeta struct {
	Name  string
	Scale int
}

var Catalog = map[string]CurrencyMeta{
	"USD": {Name: "US Dollar", Scale: 2},
	"EUR": {Name: "Euro", Scale: 2},
	"GBP": {Name: "British Pound", Scale: 2},
	"RUB": {Name: "Russian Ruble", Scale: 2},
}
