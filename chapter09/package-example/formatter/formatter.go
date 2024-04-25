package formatter

import "fmt"

// Format returns a formatted string
func Format(num int) string {
	return fmt.Sprintf("Number: %d", num)
}
