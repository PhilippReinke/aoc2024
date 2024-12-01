package shared

type Number interface {
	int
}

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
