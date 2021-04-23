package util

import (
	"math/big"
	"strconv"
)

func CalculateInt64(x int64, y int64, operator string) (i int64) {
	switch operator {
	case "add":
		a := big.NewInt(x)
		b := big.NewInt(y)
		z := a.Add(a, b)
		i := z.Int64()
		return i
	case "sub":
		a := big.NewInt(x)
		b := big.NewInt(y)
		z := a.Sub(a, b)
		i := z.Int64()
		return i
	case "mul":
		a := big.NewInt(x)
		b := big.NewInt(y)
		z := a.Mul(a, b)
		i := z.Int64()
		return i
	case "div":
		a := big.NewInt(x)
		b := big.NewInt(y)
		z := a.Div(a, b)
		i := z.Int64()
		return i
	}
	return i
}
func CalculateString(x string, y string, operator string) (i string, err error) {
	switch operator {
	case "add":
		var a, b int
		a, err = strconv.Atoi(x)
		b, err = strconv.Atoi(y)
		i := strconv.Itoa(a + b)
		return i, err
	case "sub":
		var a, b int
		a, err = strconv.Atoi(x)
		b, err = strconv.Atoi(y)
		if a < b {
			return "", NewMsgError(0, "wrong order")
		}
		i := strconv.Itoa(a - b)
		return i, err
	case "mul":
		var a, b int
		a, err = strconv.Atoi(x)
		b, err = strconv.Atoi(y)
		i := strconv.Itoa(a * b)
		return i, err
	case "div":
		var a, b int
		a, err = strconv.Atoi(x)
		b, err = strconv.Atoi(y)
		if b == 0 {
			return "", NewMsgError(0, "the denominator is zero")
		}
		i := strconv.Itoa(a / b)
		return i, err
	case "cmp":
		var a, b int
		a, err = strconv.Atoi(x)
		b, err = strconv.Atoi(y)
		if a > b {
			i = "1"
		} else if a < b {
			i = "-1"
		} else {
			i = "0"
		}
		return i, err
	case "addBig":
		a := new(big.Int)
		b := new(big.Int)
		xbi, xok := a.SetString(x, 10)
		if !xok {
			definedErr := NewMsgError(4, "error in trans string into bigint")
			return "", definedErr
		}
		ybi, yok := b.SetString(y, 10)
		if !yok {
			definedErr := NewMsgError(4, "error in trans string into bigint")
			return "", definedErr
		}
		ibi := xbi.Add(xbi, ybi)
		i = ibi.String()
		return i, err
	case "cmpBig":
		a := new(big.Int)
		b := new(big.Int)
		xbi, xok := a.SetString(x, 10)
		if !xok {
			definedErr := NewMsgError(4, "error in trans string into bigint")
			return "", definedErr
		}
		ybi, yok := b.SetString(y, 10)
		if !yok {
			definedErr := NewMsgError(4, "error in trans string into bigint")
			return "", definedErr
		}
		i = strconv.Itoa(xbi.Cmp(ybi))
		return i, err
	case "subBig":
		a := new(big.Int)
		b := new(big.Int)
		xbi, xok := a.SetString(x, 10)
		if !xok {
			definedErr := NewMsgError(4, "error in trans string into bigint")
			return "", definedErr
		}
		ybi, yok := b.SetString(y, 10)
		if !yok {
			definedErr := NewMsgError(4, "error in trans string into bigint")
			return "", definedErr
		}
		if xbi.Cmp(ybi) < 0 {
			definedErr := NewMsgError(0, "wrong order")
			return "", definedErr
		}
		ibi := xbi.Sub(xbi, ybi)
		i = ibi.String()
		return i, err
	}
	return i, err
}
