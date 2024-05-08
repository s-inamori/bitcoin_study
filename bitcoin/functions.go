package main

import (
	"math"
	"math/big"
)

// PythonのPowの第三引数の挙動を再現する
// https://stackoverflow.com/a/77542076
func Pow(a, b int, m *int) int {
	var mod *big.Int
	if m != nil {
		mod = big.NewInt(int64(*m))
	}
	result := new(big.Int).Exp(
		big.NewInt(int64(a)),
		big.NewInt(int64(b)),
		mod,
	)
	return int(result.Uint64())
}

// シンプルなPow関数
// https://stackoverflow.com/a/64109952
func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
