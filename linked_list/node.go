package linked_list

type dnode[T any] struct {
	data T
	next *dnode[T]
	prev *dnode[T]
}

type snode[T any] struct {
	data T
	next *snode[T]
}
