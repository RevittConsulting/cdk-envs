package hexadecimal

import "fmt"

func blockToHex(block int) string {
	return fmt.Sprintf("0x%X", block)
}
