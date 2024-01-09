package main

import (
	"fmt"
	"log"
	"math"
)

type FieldElement struct {
	num   int
	prime int
}

func NewFieldElement(num int, prime int) FieldElement {
	if num >= prime || num < 0 {
		log.Panicf("Num %d not in field range 0 to %d", num, prime-1)
	}
	return FieldElement{num, prime}
}

func (fe *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d(%d)", fe.prime, fe.num)
}

// 有限体の加算
func (fe *FieldElement) Add(other *FieldElement) FieldElement {
	if fe.prime != other.prime {
		log.Panicf("Cannot add two numbers in different Fields")
	}
	num := (fe.num + other.num) % fe.prime
	return FieldElement{num, fe.prime}
}

// 有限体の減算
func (fe *FieldElement) Sub(other *FieldElement) FieldElement {
	if fe.prime != other.prime {
		log.Panicf("Cannot sub two numbers in different Fields")
	}
	num := modLikePython(fe.num-other.num, fe.prime)
	return FieldElement{num, fe.prime}
}

// 有限体の乗算
func (fe *FieldElement) Mul(other *FieldElement) FieldElement {
	if fe.prime != other.prime {
		log.Panicf("Cannot mul two numbers in different Fields")
	}
	num := (fe.num * other.num) % fe.prime
	return FieldElement{num, fe.prime}
}

// 有限体のべき乗
func (fe *FieldElement) Pow(exponent int) FieldElement {
	num := (int(math.Pow(float64(fe.num), float64(exponent)))) % fe.prime
	return FieldElement{num, fe.prime}
}

// Goの場合、モジュロ演算するとマイナスになっちゃうので、Pythonのように正の値になるように調整する
// https://stackoverflow.com/a/43018347
func modLikePython(d int, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}
