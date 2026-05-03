package util

const (
	USD = "USD"
	GBP = "GBP"
	THB = "THB"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, GBP, THB:
		return true
	}
	return false
}
