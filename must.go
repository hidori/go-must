package must

func Do[T any](fn func() (T, error)) T {
	v, err := fn()
	if err != nil {
		panic(err)
	}

	return v
}

func Do2[T1 any, T2 any](fn func() (T1, T2, error)) (T1, T2) {
	v1, v2, err := fn()
	if err != nil {
		panic(err)
	}

	return v1, v2
}

func Do3[T1 any, T2 any, T3 any](fn func() (T1, T2, T3, error)) (T1, T2, T3) {
	v1, v2, v3, err := fn()
	if err != nil {
		panic(err)
	}

	return v1, v2, v3
}
