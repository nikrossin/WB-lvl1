package main

import (
	"fmt"
	"math/big"
)

type BigNum big.Int

func NewBigNum(num int64) *BigNum {
	return (*BigNum)(big.NewInt(num))
}

func (n *BigNum) Add(m *BigNum) *BigNum {
	var newNum big.Int
	return (*BigNum)(newNum.Add((*big.Int)(n), (*big.Int)(m)))
}

func (n *BigNum) Sub(m *BigNum) *BigNum {
	var newNum big.Int
	return (*BigNum)(newNum.Sub((*big.Int)(n), (*big.Int)(m)))
}
func (n *BigNum) Mul(m *BigNum) *BigNum {
	var newNum big.Int
	return (*BigNum)(newNum.Mul((*big.Int)(n), (*big.Int)(m)))
}
func (n *BigNum) Div(m *BigNum) *BigNum {
	var newNum big.Int
	return (*BigNum)(newNum.Div((*big.Int)(n), (*big.Int)(m)))
}

func (n *BigNum) String() string {
	return (*big.Int)(n).String()
}

func main() {
	var a int64 = 18446744073709551
	var b int64 = 16446744073709551

	num1 := NewBigNum(a)
	num2 := NewBigNum(b)

	num := num1.Add(num2)
	fmt.Println(num)
	num = num1.Sub(num2)
	fmt.Println(num)
	num = num1.Mul(num2)
	fmt.Println(num)
	num = num1.Div(num2)
	fmt.Println(num)

}
