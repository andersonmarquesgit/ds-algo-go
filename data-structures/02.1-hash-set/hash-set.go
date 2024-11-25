package main

type MyHashSet struct {
	bucket [][]int
	size   int
}

// Initialize data structure.
func Constructor() MyHashSet {
	return MyHashSet{
		bucket: make([][]int, 1000),
		size:   1000,
	}
}

// Function hash to generate an index for the key
func (hs *MyHashSet) hash(key int) int {
	return key % hs.size
}

func (hs *MyHashSet) Add(key int) {
	if !hs.Contains(key) {
		hash := hs.hash(key)
		hs.bucket[hash] = append(hs.bucket[hash], key)
	}
}

func (hs *MyHashSet) Remove(key int) {
	hash := hs.hash(key)
	for i, v := range hs.bucket[hash] {
		if v == key {
			hs.bucket[hash] = append(hs.bucket[hash][:i], hs.bucket[hash][i+1:]...)
		}
	}
}

func (hs *MyHashSet) print() {
	for i, v := range hs.bucket {
		if len(v) != 0 {
			println("Key: ", i, "Value: ", v[0])
		}

	}
}

// Returns true if this set contains the specified element
func (hs *MyHashSet) Contains(key int) bool {
	hash := hs.hash(key)
	for _, v := range hs.bucket[hash] {
		if v == key {
			return true
		}
	}
	return false
}

func main() {
	hashSet := Constructor()
	hashSet.Add(1)
	hashSet.Add(2)
	println(hashSet.Contains(1)) // true
	println(hashSet.Contains(3)) // false
	hashSet.Add(2)
	println(hashSet.Contains(2)) // true
	hashSet.Remove(2)
	println(hashSet.Contains(2)) // false
	hashSet.Add(3)
	hashSet.Add(1) // Não deve adicionar pois já existe
	hashSet.print()
}
