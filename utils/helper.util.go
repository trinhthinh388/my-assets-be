package utils

import (
	"math/big"
	"strconv"
	"strings"
)

func UDecToHex(iValue uint64) string {
	return "0x" + strconv.FormatUint(iValue, 16)
}

func HexToInteger(hexaString string) (int64, error) {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return strconv.ParseInt(numberStr, 16, 64)
}

func RemoveLeadingZero(s string) string {
	input := strings.TrimLeft(s, "0x")
	return strings.Join([]string{"0x", strings.TrimLeft(input, "0")}, "")
}

func TopicToAddress(topic string) string {
	input := []byte(strings.TrimSpace(topic))
	return strings.ToLower(string(append(input[:2], input[26:]...)))
}

func GetCoinAmount(value string, decimals int) (result big.Float) {
	base := big.NewInt(10)
	amount := big.Int{}
	amount.SetString(value, 10)

	base.Exp(base, big.NewInt(int64(decimals)), nil)

	amountF := new(big.Float).SetInt(&amount)
	baseF := new(big.Float).SetInt(base)

	result.Quo(amountF, baseF)

	return result
}

func IsValidAddress(address string) bool {
	address = strings.ToLower(address)
	if len(address) < 42 || address[:3] != "0x" {
		return false
	}
	return true
}
