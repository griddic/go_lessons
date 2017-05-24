package main

import (
	"fmt"
	"math"
)

var suffixes = []string{"B", "KB", "MB", "GB"}

func main() {
	fmt.Println(humanize(100))
	fmt.Println(humanize(12345))
}


func humanize(n uint64) string {
	fmt.Println(n)
	len := int(math.Floor(math.Log10(float64(n))))
	var devider uint64 = 1
	if len > 2 {
		devider = uint64(math.Pow10(len - 2))
	}
	to_print := n / devider
	point_location := (len - 2) % 3
	suffix_index := len / 3
	to_print_float := float64(to_print) / float64(math.Pow10(point_location))
	fmt.Println(len, devider, to_print, to_print_float,
		point_location, suffixes[suffix_index])

	return fmt.Sprint(to_print_float) + suffixes[suffix_index]
}