package utils

import (
	erc20 "my-assets-be/constants/ERC20"
	"my-assets-be/geth"
)

func FilterTransferLogs(logs []geth.TxReceiptLog) (transferLogs []geth.TxReceiptLog) {
	for i := 0; i < len(logs); i++ {
		if logs[i].Topics[0] == erc20.TransferEventHash {
			transferLogs = append(transferLogs, logs[i])
		}
	}

	return transferLogs
}
