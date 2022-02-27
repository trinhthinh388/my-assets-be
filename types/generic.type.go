package types

import (
	"math/big"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

type BigInt struct {
	i *big.Int
}

func NewBigInt(bigint *big.Int) *BigInt {
	return &BigInt{i: bigint}
}

func (bi *BigInt) Int() *big.Int {
	return bi.i
}

func (bi *BigInt) MarshalJSON() ([]byte, error) {
	txt, err := bi.i.MarshalText()
	if err != nil {
		return nil, err
	}
	return []byte(txt), err
}

func (bi *BigInt) UnmarshalJSON(data []byte) error {
	input := strings.Trim(string(data), "\"")
	if bi.i == nil {
		bi.i = big.NewInt(0)
	}
	err := bi.i.UnmarshalText([]byte(input))
	if err != nil {
		return err
	}
	return nil
}

func (bi *BigInt) MarshalBSON() ([]byte, error) {
	txt, err := bi.i.MarshalText()
	if err != nil {
		return nil, err
	}
	a, err := bson.Marshal(map[string]string{"i": string(txt)})
	return a, err
}

func (bi *BigInt) UnmarshalBSON(data []byte) error {
	var d bson.D
	err := bson.Unmarshal(data, &d)
	if err != nil {
		return err
	}
	if v, ok := d.Map()["i"]; ok {
		bi.i = big.NewInt(0)
		return bi.i.UnmarshalText([]byte(v.(string)))
	}
	return err
}
