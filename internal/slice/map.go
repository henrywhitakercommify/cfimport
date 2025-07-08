package slice

func Map[T any, U any](in []T, m func(T) U) []U {
	out := []U{}
	for _, v := range in {
		out = append(out, m(v))
	}
	return out
}
