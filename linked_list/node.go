package linked_list

type dnode struct {
	data any
	next *dnode
	prev *dnode
}

type snode struct {
	data any
	next *snode
}
