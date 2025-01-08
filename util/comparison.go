package util

// IfThenElse returns the given trueValue if the given condition is true or the given falseValue if it is not.
func IfThenElse[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}
