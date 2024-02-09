package hexadecimal

import "testing"

func Test_HexDec(t *testing.T) {
	scenarios := map[string]struct {
		block       int
		expectedHex string
	}{
		"block 0": {
			block:       0,
			expectedHex: "0x0",
		},
		"block 1": {
			block:       1,
			expectedHex: "0x1",
		},
		"block 5251189": {
			block:       5251189,
			expectedHex: "0x502075",
		},
	}

	for name, scenario := range scenarios {
		t.Run(name, func(t *testing.T) {
			actualHex := blockToHex(scenario.block)
			if actualHex != scenario.expectedHex {
				t.Errorf("Expected %s, got %s", scenario.expectedHex, actualHex)
			}
		})
	}
}
