package main

import (
	"fmt"
	"math"
)

var suffixes = []string{"B", "KB", "MB", "GB"}

func main() {
	fmt.Println(humanize(5))
	fmt.Println(humanize(100))
	fmt.Println(humanize(12345))
	fmt.Println(humanize(2100000))
}


func humanize(n uint64) string {
	//fmt.Println(n)
	len := uint(math.Floor(math.Log10(float64(n))))
	if len > 11 {
		len = 11
	}
	var devider uint64 = 1
	var point_location uint = 0
	if len > 2 {
		devider = uint64(math.Pow10(int(len) - 2))
		point_location = (5 - len + 3 * len) % 3
	}
	to_print := n / devider

	suffix_index := len / 3
	to_print_float := float64(to_print) / float64(math.Pow10(int(point_location)))
	//fmt.Println(len, devider, to_print, to_print_float,
	//	point_location, suffixes[suffix_index])

	return fmt.Sprint(to_print_float) + suffixes[suffix_index]
}