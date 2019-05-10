package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/makisenpai/is105-ica01/sum/sum"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Bruk to argumenter") //tester om det er to parametere
		os.Exit(2)
	}

	if a, err := strconv.Atoi(os.Args[1]); err == nil {
		if b, err := strconv.Atoi(os.Args[2]); err == nil {
			i8a, i32a, u32a, i64a, f64a := int8(a), int32(a), uint32(a), int64(a), float64(a)
			i8b, i32b, u32b, i64b, f64b := int8(b), int32(b), uint32(b), int64(b), float64(b)

			i8 := sum.SumInt8(i8a, i8b)
			i32 := sum.SumInt32(i32a, i32b)
			u32 := sum.SumUint32(u32a, u32b)
			i64 := sum.SumInt64(i64a, i64b)
			f64 := sum.SumFloat64(f64a, f64b)

			fmt.Printf("Int8: %v, int32: %v, Uint32: %v, Int64: %v, Float64: %v", i8, i32, u32, i64, f64)
		}
	}
}
