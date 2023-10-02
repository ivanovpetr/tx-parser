package ethclient

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/ivanovpetr/tx-parser/pkg/ethutil"
)

// HexedInt64 json wrapper for hexed representation of an int64 for example '0x1f47ed'
type HexedInt64 struct {
	Value int64
}

// MarshalJSON implements json.Marshaler interface
func (hi *HexedInt64) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"0x%x\"", hi.Value)), nil
}

// UnmarshalJSON implements json.Unmarshaler interface
func (hi *HexedInt64) UnmarshalJSON(data []byte) error {
	str := string(data[3 : len(data)-1])

	number, err := strconv.ParseInt(str, 16, 64)
	if err != nil {
		return fmt.Errorf("failed to to parse hex string %s", str)
	}
	hi.Value = number
	return nil
}

// Address json wrappper for eth address
type Address struct {
	Value ethutil.Address
}

// MarshalJSON implements json.Marshaler interface
func (a *Address) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", a.Value)), nil
}

// UnmarshalJSON implements json.Unmarshaler interface
func (a *Address) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" {
		a.Value = ethutil.ZeroAddress
		return nil
	}
	trimmedQuotes := string(data[1 : len(data)-1])

	addr, err := ethutil.ParseAddressFromString(trimmedQuotes)
	if err != nil {
		return fmt.Errorf("failed to parse eth address %s : %w", str, err)
	}
	a.Value = addr

	return nil
}

// HexedBigInt json wrapper for hexed representation of a big.Int for example '0x10000000000001f47ed'
type HexedBigInt struct {
	Value *big.Int
}

// MarshalJSON implements json.Marshaler interface
func (hb *HexedBigInt) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"0x%s\"", hb.Value.Text(16))), nil
}

// UnmarshalJSON implements json.Unmarshaler interface
func (hb *HexedBigInt) UnmarshalJSON(data []byte) error {
	str := string(data[3 : len(data)-1])

	num := new(big.Int)
	if _, success := num.SetString(str, 16); !success {
		return fmt.Errorf("failed to parse hex string %s", str)
	}
	hb.Value = num
	return nil
}
