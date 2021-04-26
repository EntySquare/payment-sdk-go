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
func CalculateString(x string, y string, operator string) (i string, definedErr *MessageError) {
	if x == "" {
		x = "0"
	}
	if y == "" {
		y = "0"
	}
	a := new(big.Float)
	b := new(big.Float)
	xbf, xok := a.SetString(x)
	if !xok {
		definedErr := NewMsgError(4, "error in trans string into big float")
		return "", definedErr
	}
	ybf, yok := b.SetString(y)
	if !yok {
		definedErr := NewMsgError(4, "error in trans string into big float")
		return "", definedErr
	}
	switch operator {
	case "add":
		ai, xErr := strconv.Atoi(x)
		bi, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			definedErr = NewMsgError(4, "error in trans string into int")
		}
		i := strconv.Itoa(ai + bi)
		return i, definedErr
	case "sub":
		ai, xErr := strconv.Atoi(x)
		bi, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			definedErr = NewMsgError(4, "error in trans string into int")
		}
		if ai < bi {
			return "", NewMsgError(0, "wrong order")
		}
		i := strconv.Itoa(ai - bi)
		return i, definedErr
	case "mul":
		ai, xErr := strconv.Atoi(x)
		bi, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			definedErr = NewMsgError(4, "error in trans string into int")
		}
		i := strconv.Itoa(ai * bi)
		return i, definedErr
	case "div":
		ai, xErr := strconv.Atoi(x)
		bi, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			definedErr = NewMsgError(4, "error in trans string into int")
		}
		if bi == 0 {
			return "", NewMsgError(0, "the denominator is zero")
		}
		i := strconv.Itoa(ai / bi)
		return i, definedErr
	case "cmp":
		ai, xErr := strconv.Atoi(x)
		bi, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			definedErr = NewMsgError(4, "error in trans string into int")
		}
		if ai > bi {
			i = "1"
		} else if ai < bi {
			i = "-1"
		} else {
			i = "0"
		}
		return i, definedErr
	case "addBigFU":
		ibf := xbf.Add(xbf, ybf)
		i = ibf.Text('f', 18)
		return i, definedErr
	case "subBigFU":
		if xbf.Cmp(ybf) < 0 {
			definedErr := NewMsgError(0, "wrong order")
			return "", definedErr
		}
		ibf := xbf.Sub(xbf, ybf)
		i = ibf.Text('f', 18)
		return i, definedErr
	case "addBigFH":
		ibf := xbf.Add(xbf, ybf)
		i = ibf.Text('f', 2)
		return i, definedErr
	case "cmpBigF":
		i = strconv.Itoa(xbf.Cmp(ybf))
		return i, definedErr
	case "divBigF":
		ibf := xbf.Quo(xbf, ybf)
		i = ibf.Text('f', 18) //保留6位小数
		return i, definedErr
	case "mulBigF":
		ibf := xbf.Mul(xbf, ybf)
		i = ibf.Text('f', 18) //保留6位小数
		return i, definedErr
	case "subBigFH":
		if xbf.Cmp(ybf) < 0 {
			definedErr := NewMsgError(0, "wrong order")
			return "", definedErr
		}
		ibf := xbf.Sub(xbf, ybf)
		i = ibf.Text('f', 2)
		return i, definedErr
	}

	return i, definedErr
}
func Digit(x string, operator string) (i string, definedErr *MessageError) {
	unit := new(big.Float)
	a := new(big.Float)
	bf, ok := a.SetString(x)
	if !ok {
		definedErr := NewMsgError(4, "error in trans string into bigint")
		return "", definedErr
	}
	switch operator {
	case "div18":
		unit.SetString("1000000000000000000")
		bf.Quo(bf, unit)
		i = bf.Text('f', 18)
		return i, definedErr
	case "div2":
		unit.SetString("100")
		bf.Quo(bf, unit)
		i = bf.Text('f', 2)
		return i, definedErr
	case "mul2":
		unit.SetString("100")
		bf.Mul(bf, unit)
		i = bf.Text('f', 0)
		return i, definedErr
	case "mul18":
		unit.SetString("1000000000000000000")
		bf.Mul(bf, unit)
		i = bf.Text('f', 0)
		return i, definedErr
	}

	return i, definedErr
}
