//KODE FOR TESTING AV NUMERISKE SYSTEMER.

package sum

import "testing"

var sum_test_int8 = []struct {
	n1       int8
	n2       int8
	expected int8
}{
	{90, 37, 127},
	{126, 1, 127},
	{35, -45, -10},
}

func TestSumInt8(t *testing.T) {
	for _, v := range sum_test_int8 {
		if val := SumInt8(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum (%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

var sum_tests_int32 = []struct {
	n1       int32
	n2       int32
	expected int32
}{
	{1, 2, 3},
	{-2145126647, -2357000, -2147483647},
	{2145126647, 2357000, 2147483647},
}

func TestSumInt32(t *testing.T) {
	for _, v := range sum_tests_int32 {
		if val := SumInt32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

//tester om variablene overskrider grenseverdiene for int64
var sum_tests_int64 = []struct {
	n1       int64
	n2       int64
	expected int64
}{
	{1, 2, 3},
	{2635, 2357, 4992},
	{5, -7, -2},
}

func TestSumInt64(t *testing.T) {
	for _, v := range sum_tests_int64 {
		if val := SumInt64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

//Tester om variablene overskrider grenseverdi for float64
var sum_tests_float64 = []struct {
	n1       float64
	n2       float64
	expected float64
}{
	{1.4, 1.3, 2.7},
	{4.5, 5.6, 10.1},
	{21.73, 3.12, 24.85},
}

func TestSumFloat64(t *testing.T) {
	for _, v := range sum_tests_float64 {
		if val := SumFloat64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%f, %f) returned %f, expected %f", v.n1, v.n2, val, v.expected)
		}
	}
}

var sum_tests_uint32 = []struct {
	n1       uint32
	n2       uint32
	expected uint32
}{
	{1, 2, 3},
	{4, 5, 9},
	{0, 5, 5},
}

func TestSumUint32(t *testing.T) {
	for _, v := range sum_tests_uint32 {
		if val := SumUint32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}
