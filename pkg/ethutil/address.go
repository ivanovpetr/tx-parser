package ethutil

import (
	"encoding/hex"
	"errors"
	"fmt"
)

// Address represents ethereum address
type Address [20]byte

// ZeroAddress represents zero ethereum address 0x0000000000000000000000000000000000000000
var ZeroAddress = Address{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// String implements fmt.Stringer interface for Address
func (a Address) String() string {
	return "0x" + hex.EncodeToString(a[:])
}

// ParseAddressFromString parses ethereum address from string
func ParseAddressFromString(addressString string) (Address, error) {
	if addressString[:2] != "0x" {
		return ZeroAddress, errors.New("Invalid address format")
	}
	addressBytes, err := hex.DecodeString(addressString[2:])
	if err != nil {
		return ZeroAddress, fmt.Errorf("Failed to parse Ethereum address: %w", err)
	}

	if len(addressBytes) != 20 {
		return ZeroAddress, fmt.Errorf("Invalid Ethereum address length: %d bytes", len(addressBytes))
	}

	var address [20]byte
	copy(address[:], addressBytes)

	return address, nil
}
