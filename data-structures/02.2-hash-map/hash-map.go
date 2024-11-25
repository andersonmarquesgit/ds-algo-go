package main

type MyHashMap struct {
	bucket [][]int
	size   int
}

// Initialize data structure.
func Constructor() MyHashMap {
	return MyHashMap{
		bucket: make([][]int, 1000),
		size:   1000,
	}
}

// Function hash to generate an index for the key
func (hm *MyHashMap) hash(key int) int {
	return key % hm.size
}

func (hm *MyHashMap) put(key int, value int) {
	hash := hm.hash(key)
	for i, v := range hm.bucket[hash] {
		if v == key {
			hm.bucket[hash][i+1] = value
			return
		}
	}
	hm.bucket[hash] = append(hm.bucket[hash], key, value)
}

func (hm *MyHashMap) get(key int) int {
	hash := hm.hash(key)
	for i, v := range hm.bucket[hash] {
		if v == key {
			return hm.bucket[hash][i+1]
		}
	}
	return -1
}

func (hm *MyHashMap) remove(key int) {
	hash := hm.hash(key)
	for i, v := range hm.bucket[hash] {
		if v == key {
			hm.bucket[hash] = append(hm.bucket[hash][:i], hm.bucket[hash][i+2:]...)
		}
	}
}

func (hm *MyHashMap) print() {
	for i, v := range hm.bucket {
		if len(v) != 0 {
			println("Key: ", i, "Value: ", v[0])
		}
	}
}

func main() {
	myHashMap := Constructor()
	myHashMap.put(1, 1)       // The map is now [[1,1]]
	myHashMap.put(2, 2)       // The map is now [[1,1], [2,2]]
	println(myHashMap.get(1)) // return 1, The map is now [[1,1], [2,2]]
	println(myHashMap.get(3)) // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
	myHashMap.put(2, 1)       // The map is now [[1,1], [2,1]] (i.e., update the existing value)
	println(myHashMap.get(2)) // return 1, The map is now [[1,1], [2,1]]
	myHashMap.remove(2)       // remove the mapping for 2, The map is now [[1,1]]
	println(myHashMap.get(2)) // return -1 (i.e., not found), The map is now [[1,1]]
}
