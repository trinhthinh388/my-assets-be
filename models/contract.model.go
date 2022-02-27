package models

type Contract struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Address   string `json:"name" gorm:"notnull;type:varchar(42);index:idx_contract_address"`
	Name      string `json:"name" gorm:"notnull;type:varchar(255)"`
	ABI       string `json:"abi" gorm:"notnull;type:text`
	CreatedAt string `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updatedAt" gorm:"autoUpdateTime"`
}
