package util

func FlatMap[A, B any](source []A, f func(A) []B) []B {
	var result []B
	for _, v := range source {
		result = append(result, f(v)...)
	}
	return result
}

func Map[A, B any](source []A, f func(A) B) []B {
	var result []B
	for _, v := range source {
		result = append(result, f(v))
	}
	return result
}

func Filter[A any](source []A, f func(A) bool) []A {
	var result []A
	for _, v := range source {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}
