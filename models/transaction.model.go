package models

import (
	"time"
)

type Transaction struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Hash        string     `json:"hash" gorm:"notnull;type:varchar(255);index:idx_hash_chain_index"`
	From        string     `json:"from" gorm:"notnull;type:varchar(42)"`
	To          string     `json:"to" gorm:"notnull;type:varchar(42)"`
	Value       string     `json:"value" gorm:"notnull;type:varchar(255)"`
	Gas         string     `json:"gas" gorm:"notnull;type:varchar(255)"`
	GasPrice    string     `json:"gasPrice" gorm:"notnull;type:varchar(255)"`
	Input       string     `json:"input" gorm:"type:text"`
	BlockNumber int        `json:"blockNumber" gorm:"notnull;index:idx_block_time_asc,sort:asc;index:idx_block_time_desc,sort:desc"`
	Timestamp   *time.Time `json:"timestamp" gorm:"index:idx_block_time_asc,sort:asc;index:idx_block_time_desc,sort:desc"`
	Chain       string     `json:"chain" gorm:"type:varchar(255);index:hash_chain_index"`
	Owner       string     `json:"owner" gorm:"notnull;type:varchar(42)"`
}
