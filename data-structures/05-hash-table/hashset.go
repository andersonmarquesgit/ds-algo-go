package main

type HashSet struct {
	data map[interface{}]bool
}

func NewHashSet() *HashSet {
	return &HashSet{
		data: make(map[interface{}]bool),
	}
}

func (h *HashSet) Add(item interface{}) {
	h.data[item] = true
}
