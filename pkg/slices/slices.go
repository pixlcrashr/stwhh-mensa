package slices

func AddOrSet[T any](
	ts []T,
	predFn func(v T) bool,
	setFn func(v T) T,
	addFn func() T,
) []T {
	for i, t := range ts {
		if !predFn(t) {
			continue
		}

		ts[i] = setFn(t)
		return ts
	}

	return append(ts, addFn())
}
