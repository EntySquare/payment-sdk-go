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
	switch operator {
	case "add":
		a, xErr := strconv.Atoi(x)
		b, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			definedErr = NewMsgError(4, "error in trans string into int")
		}
		i := strconv.Itoa(a + b)
		return i, definedErr
	case "sub":
		a, xErr := strconv.Atoi(x)
		b, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			definedErr = NewMsgError(4, "error in trans string into int")
		}
		if a < b {
			return "", NewMsgError(0, "wrong order")
		}
		i := strconv.Itoa(a - b)
		return i, definedErr
	case "mul":
		a, xErr := strconv.Atoi(x)
		b, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			definedErr = NewMsgError(4, "error in trans string into int")
		}
		i := strconv.Itoa(a * b)
		return i, definedErr
	case "div":
		a, xErr := strconv.Atoi(x)
		b, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			definedErr = NewMsgError(4, "error in trans string into int")
		}
		if b == 0 {
			return "", NewMsgError(0, "the denominator is zero")
		}
		i := strconv.Itoa(a / b)
		return i, definedErr
	case "cmp":
		a, xErr := strconv.Atoi(x)
		b, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			definedErr = NewMsgError(4, "error in trans string into int")
		}
		if a > b {
			i = "1"
		} else if a < b {
			i = "-1"
		} else {
			i = "0"
		}
		return i, definedErr
	case "addBigFU":
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
		ibf := xbf.Add(xbf, ybf)
		i = ibf.Text('f', 18)
		return i, definedErr
	case "subBigFU":
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
		if xbf.Cmp(ybf) < 0 {
			definedErr := NewMsgError(0, "wrong order")
			return "", definedErr
		}
		ibf := xbf.Sub(xbf, ybf)
		i = ibf.Text('f', 18)
		return i, definedErr
	case "addBigFH":
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
		ibf := xbf.Add(xbf, ybf)
		i = ibf.Text('f', 2)
		return i, definedErr
	case "cmpBigF":
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
		i = strconv.Itoa(xbf.Cmp(ybf))
		return i, definedErr
	case "divBigF":
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
		ibf := xbf.Quo(xbf, ybf)
		i = ibf.Text('f', 6) //保留6位小数
		return i, definedErr
	case "subBigFH":
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
