package main

type Pair struct {
	key   int
	value int
}

type MyHashMap struct {
	bucket [][]Pair
	size   int
}

// Initialize data structure.
func Constructor() MyHashMap {
	return MyHashMap{
		bucket: make([][]Pair, 1000),
		size:   1000,
	}
}

// Hash function to generate an index for the key.
func (hm *MyHashMap) hash(key int) int {
	return key % hm.size
}

// Put a key-value pair into the HashMap.
func (hm *MyHashMap) Put(key int, value int) {
	hash := hm.hash(key)
	for i := range hm.bucket[hash] {
		if hm.bucket[hash][i].key == key {
			hm.bucket[hash][i].value = value // Update existing value
			return
		}
	}
	// If key does not exist, append new Pair
	hm.bucket[hash] = append(hm.bucket[hash], Pair{key, value})
}

// Get the value corresponding to the key.
func (hm *MyHashMap) Get(key int) int {
	hash := hm.hash(key)
	for _, pair := range hm.bucket[hash] {
		if pair.key == key {
			return pair.value
		}
	}
	return -1 // Key not found
}

// Remove the key-value pair from the HashMap.
func (hm *MyHashMap) Remove(key int) {
	hash := hm.hash(key)
	for i := range hm.bucket[hash] {
		if hm.bucket[hash][i].key == key {
			// Remove the Pair by slicing
			hm.bucket[hash] = append(hm.bucket[hash][:i], hm.bucket[hash][i+1:]...)
			return
		}
	}
}

func main() {
	myHashMap := Constructor()
	myHashMap.Put(1, 1)       // The map is now [[1,1]]
	myHashMap.Put(2, 2)       // The map is now [[1,1], [2,2]]
	println(myHashMap.Get(1)) // return 1, The map is now [[1,1], [2,2]]
	println(myHashMap.Get(3)) // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
	myHashMap.Put(2, 1)       // The map is now [[1,1], [2,1]] (i.e., update the existing value)
	println(myHashMap.Get(2)) // return 1, The map is now [[1,1], [2,1]]
	myHashMap.Remove(2)       // remove the mapping for 2, The map is now [[1,1]]
	println(myHashMap.Get(2)) // return -1 (i.e., not found), The map is now [[1,1]]
}
