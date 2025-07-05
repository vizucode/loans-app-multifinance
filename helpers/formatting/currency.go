package formatting

import (
	"fmt"
	"strconv"
	"strings"
)

func Currency2Str(x float64) string {
	return strconv.FormatFloat(x, 'f', 2, 64)
}

func Str2Currency(x string) (float64, error) {
	f, err := strconv.ParseFloat(x, 64)
	return f, err
}

func HumanCurrency(prefix string, currency float64) (resp string) {
	// Format the number with 2 decimal places and comma as thousand separator
	str := fmt.Sprintf("%.2f", currency)
	str = strings.ReplaceAll(str, ".", ",") // replace dot with comma for decimal

	// Split into whole and decimal parts
	parts := strings.Split(str, ",")
	whole := parts[0]
	decimal := parts[1]

	// Add thousand separator to the whole part
	n := len(whole)
	for i := n - 3; i > 0; i -= 3 {
		whole = whole[:i] + "." + whole[i:]
	}

	return prefix + " " + whole + "," + decimal
}
