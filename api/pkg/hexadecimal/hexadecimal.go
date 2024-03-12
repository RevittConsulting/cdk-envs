package hexadecimal

import (
	"fmt"
	"strconv"
	"strings"
)

func BlockToHex(block int) string {
	return fmt.Sprintf("0x%X", block)
}

func HashToUint64(hash string) (uint64, error) {
	cleanHash := strings.TrimPrefix(hash, "0x")
	return strconv.ParseUint(cleanHash, 16, 64)
}
