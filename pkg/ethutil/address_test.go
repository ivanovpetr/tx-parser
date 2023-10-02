package ethutil

import "testing"

func TestParseAddressFromString(t *testing.T) {
	cases := []struct {
		address        string
		expectedOutput Address
		ErrorExpected  bool
	}{
		{
			"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
			Address{0xa0, 0xb8, 0x69, 0x91, 0xc6, 0x21, 0x8b, 0x36, 0xc1, 0xd1, 0x9d, 0x4a, 0x2e, 0x9e, 0xb0, 0xce, 0x36, 0x6, 0xeb, 0x48},
			false,
		},
		{
			"a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
			ZeroAddress,
			true,
		},
		{
			"0xa0b86991c6218b36c1d19d4a2e9eb0qe3606eb48",
			ZeroAddress,
			true,
		},
		{
			"0xa0b86991c6218b36c1d19d4a2e9eb0qe3606",
			ZeroAddress,
			true,
		},
	}
	for _, tt := range cases {
		t.Run("address parsing "+tt.address, func(t *testing.T) {
			addr, err := ParseAddressFromString(tt.address)
			if tt.ErrorExpected {
				if err == nil {
					t.Error()
				}
			} else {
				if addr != tt.expectedOutput {
					t.Errorf("Expected %#v got %#v", tt.expectedOutput, addr)
				}
			}
		})
	}
}

func TestAddressStringConversion(t *testing.T) {
	cases := []struct {
		address        Address
		expectedOutput string
	}{
		{
			Address{0xa0, 0xb8, 0x69, 0x91, 0xc6, 0x21, 0x8b, 0x36, 0xc1, 0xd1, 0x9d, 0x4a, 0x2e, 0x9e, 0xb0, 0xce, 0x36, 0x6, 0xeb, 0x48},
			"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		},
	}
	for _, tt := range cases {
		t.Run("string conversion", func(t *testing.T) {
			converted := tt.address.String()
			if tt.expectedOutput != converted {
				t.Errorf("Expected %#v got %#v", tt.expectedOutput, converted)
			}
		})
	}
}
