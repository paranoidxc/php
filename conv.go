package php

import "strconv"

func String2Int64(str string) (int64, error) {
	num, err := strconv.ParseInt(str, 10, strconv.IntSize)
	if err != nil {
		return num, err
	}

	return num, nil
}

func Int642String(num int64) string {
	return strconv.FormatInt(num, 10)
}

func String2Int(str string) (int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func intToString(num int) string {
	return strconv.Itoa(num)
}

func String2Uint(str string) (uint, error) {
	var uintNum uint
	num, err := strconv.ParseUint(str, 10, strconv.IntSize)
	if err != nil {
		return uintNum, err
	}
	uintNum = uint(num)

	return uintNum, nil
}

func Float642String(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

func String2Float64(str string) (float64, error) {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return num, err
	}

	return num, nil
}
