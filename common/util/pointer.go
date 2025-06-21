package util

func GetPointer[T any](input T) *T {
	return &input
}

func GetNilPointer[T comparable](input T) *T {
	var zeroValue T
	if input == zeroValue {
		return nil
	}

	return &input
}

func GetDerefPointer[T any](input *T) T {
	var zeroValue T
	if input == nil {
		return zeroValue
	}

	return *input
}
