package utils

import (
	erc20 "my-assets-be/constants/ERC20"
	"my-assets-be/geth"
	"strconv"
	"strings"
)

func FilterTransferLogs(logs []geth.TxReceiptLog) (transferLogs []geth.TxReceiptLog) {
	for i := 0; i < len(logs); i++ {
		if logs[i].Topics[0] == erc20.TransferEventHash {
			transferLogs = append(transferLogs, logs[i])
		}
	}

	return transferLogs
}

func UDecToHex(iValue uint64) string {
	return "0x" + strconv.FormatUint(iValue, 16)
}

func HexToInteger(hexaString string) (int64, error) {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return strconv.ParseInt(numberStr, 16, 64)
}
