package numbers

import (
	"fmt"
	"strconv"
)

func HexToDec(hexStr string) (string, error) {
	val, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return hexStr, err
	}
	return fmt.Sprintf("%d", val), nil
}

func BinToDec(binStr string) (string, error) {
	val, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		return binStr, err
	}
	return fmt.Sprintf("%d", val), nil
}
