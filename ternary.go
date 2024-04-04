package php

func Ternary[T any](condition bool, trueOut T, falseOut T) T {
	if condition {
		return trueOut
	}

	return falseOut
}
