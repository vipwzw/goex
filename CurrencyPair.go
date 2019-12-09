package goex

import "strings"

type Currency string

func NewCurrency(c string) Currency {
	return Currency(strings.ToUpper(c))
}

func (c Currency) String() string {
	return string(c)
}

type CurrencyPair struct {
	Base  Currency
	Quote Currency
}

//base currency const
var (
	UNKNOWN = Currency("UNKNOWN")
	CNY     = Currency("CNY")
	USD     = Currency("USD")
	USDT    = Currency("USDT")
	PAX     = Currency("PAX")
	USDC    = Currency("USDC")
	EUR     = Currency("EUR")
	KRW     = Currency("KRW")
	JPY     = Currency("JPY")

	BTC  = Currency("BTC")
	XBT  = Currency("XBT")
	BCH  = Currency("BCH")
	LTC  = Currency("LTC")
	ETH  = Currency("ETH")
	ETC  = Currency("ETC")
	EOS  = Currency("EOS")
	BTS  = Currency("BTS")
	QTUM = Currency("QTUM")
	SC   = Currency("SC")
	ANS  = Currency("ANS")
	ZEC  = Currency("ZEC")
	DCR  = Currency("DCR")
	XRP  = Currency("XRP")
	NEO  = Currency("NEO")
	BSV  = Currency("BSV")
	OKB  = Currency("OKB")
	HT   = Currency("HT")
	BNB  = Currency("BNB")
	TRX  = Currency("TRX")
	FT   = Currency("FT")

	//currency pair

	BTC_CNY = CurrencyPair{BTC, CNY}
	LTC_CNY = CurrencyPair{LTC, CNY}
	ETH_CNY = CurrencyPair{ETH, CNY}
	ETC_CNY = CurrencyPair{ETC, CNY}
	EOS_CNY = CurrencyPair{EOS, CNY}
	BTS_CNY = CurrencyPair{BTS, CNY}

	BTC_KRW = CurrencyPair{BTC, KRW}
	ETH_KRW = CurrencyPair{ETH, KRW}
	ETC_KRW = CurrencyPair{ETC, KRW}
	LTC_KRW = CurrencyPair{LTC, KRW}
	BCH_KRW = CurrencyPair{BCH, KRW}

	BTC_USD = CurrencyPair{BTC, USD}
	LTC_USD = CurrencyPair{LTC, USD}
	ETH_USD = CurrencyPair{ETH, USD}
	ETC_USD = CurrencyPair{ETC, USD}
	BCH_USD = CurrencyPair{BCH, USD}
	XRP_USD = CurrencyPair{XRP, USD}
	EOS_USD = CurrencyPair{EOS, USD}
	BSV_USD = CurrencyPair{BSV, USD}

	BTC_USDT = CurrencyPair{BTC, USDT}
	LTC_USDT = CurrencyPair{LTC, USDT}
	BCH_USDT = CurrencyPair{BCH, USDT}
	ETC_USDT = CurrencyPair{ETC, USDT}
	ETH_USDT = CurrencyPair{ETH, USDT}
	NEO_USDT = CurrencyPair{NEO, USDT}
	EOS_USDT = CurrencyPair{EOS, USDT}
	XRP_USDT = CurrencyPair{XRP, USDT}
	BSV_USDT = CurrencyPair{BSV, USDT}
	OKB_USDT = CurrencyPair{OKB, USDT}
	HT_USDT  = CurrencyPair{HT, USDT}
	BNB_USDT = CurrencyPair{BNB, USDT}
	PAX_USDT = CurrencyPair{PAX, USDT}
	TRX_USDT = CurrencyPair{TRX, USDT}

	XRP_EUR = CurrencyPair{XRP, EUR}

	BTC_JPY = CurrencyPair{BTC, JPY}
	LTC_JPY = CurrencyPair{LTC, JPY}
	ETH_JPY = CurrencyPair{ETH, JPY}
	ETC_JPY = CurrencyPair{ETC, JPY}
	BCH_JPY = CurrencyPair{BCH, JPY}

	LTC_BTC = CurrencyPair{LTC, BTC}
	ETH_BTC = CurrencyPair{ETH, BTC}
	ETC_BTC = CurrencyPair{ETC, BTC}
	BCH_BTC = CurrencyPair{BCH, BTC}
	DCR_BTC = CurrencyPair{DCR, BTC}
	XRP_BTC = CurrencyPair{XRP, BTC}
	NEO_BTC = CurrencyPair{NEO, BTC}
	EOS_BTC = CurrencyPair{EOS, BTC}
	BSV_BTC = CurrencyPair{BSV, BTC}
	OKB_BTC = CurrencyPair{OKB, BTC}
	HT_BTC  = CurrencyPair{HT, BTC}
	BNB_BTC = CurrencyPair{BNB, BTC}
	TRX_BTC = CurrencyPair{TRX, BTC}

	ETC_ETH = CurrencyPair{ETC, ETH}
	EOS_ETH = CurrencyPair{EOS, ETH}
	ZEC_ETH = CurrencyPair{ZEC, ETH}
	NEO_ETH = CurrencyPair{NEO, ETH}
	LTC_ETH = CurrencyPair{LTC, ETH}

	UNKNOWN_PAIR = CurrencyPair{UNKNOWN, UNKNOWN}
)

func (c CurrencyPair) String() string {
	return c.ToSymbol("_")
}

func (c CurrencyPair) Eq(c2 CurrencyPair) bool {
	return c.String() == c2.String()
}

func NewCurrencyPair(currencyA Currency, currencyB Currency) CurrencyPair {
	return CurrencyPair{currencyA, currencyB}
}

func NewCurrencyPair2(currencyPairSymbol string) CurrencyPair {
	return NewCurrencyPair3(currencyPairSymbol, "_")
}

func NewCurrencyPair3(currencyPairSymbol string, sep string) CurrencyPair {
	currencys := strings.Split(currencyPairSymbol, sep)
	if len(currencys) >= 2 {
		return CurrencyPair{Currency(currencys[0]), Currency(currencys[1])}
	}
	return UNKNOWN_PAIR
}

func (pair CurrencyPair) ToSymbol(joinChar string) string {
	return strings.Join([]string{string(pair.Base), string(pair.Quote)}, joinChar)
}

func (pair CurrencyPair) ToSymbol2(joinChar string) string {
	return strings.Join([]string{string(pair.Quote), string(pair.Base)}, joinChar)
}

func (pair CurrencyPair) AdaptUsdtToUsd() CurrencyPair {
	if pair.Quote == USDT {
		return NewCurrencyPair(pair.Base, USD)
	}
	return pair
}

func (pair CurrencyPair) AdaptUsdToUsdt() CurrencyPair {
	if pair.Quote == USD {
		return NewCurrencyPair(pair.Base, USDT)
	}
	return pair
}

//for to symbol lower , Not practical '==' operation method
func (pair CurrencyPair) ToLower() CurrencyPair {
	return CurrencyPair{Currency(strings.ToLower(string(pair.Base))),
		Currency(strings.ToLower(string(pair.Quote)))}
}

func (pair CurrencyPair) Reverse() CurrencyPair {
	return CurrencyPair{pair.Quote, pair.Base}
}
