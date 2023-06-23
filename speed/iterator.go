package speed

type Iterator[v any] interface {
	HasNext() bool
	Next() v
}
