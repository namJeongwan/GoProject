package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestSquare1(t *testing.T) {
//	assert := assert.New(t)
//	assert.Equal(82, square(9), "square(9) should be 81")
//}
//
//func TestSquare2(t *testing.T) {
//	rst := square(3)
//	if rst != 9 {
//		t.Errorf("Sqaure(3) should be 9 but square(3) returns %d\n", rst)
//	}
//}

func TestFibonacci1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci1(-1), "fibonacci1(-1) should be 0")
	assert.Equal(0, fibonacci1(0), "fibonacci1(0) should be 0")
	assert.Equal(1, fibonacci1(1), "fibonacci1(1) should be 1")
	assert.Equal(2, fibonacci1(3), "fibonacci1(2) should be 2")
	assert.Equal(233, fibonacci1(13), "fibonacci1(13) should be 233")
}

func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci1(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}
