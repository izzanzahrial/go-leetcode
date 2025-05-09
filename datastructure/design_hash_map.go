package datastructure

// https://leetcode.com/problems/design-hashmap/description/
type MyHashMap struct {
	hash map[int]int
}

func ConstructorHashMap() MyHashMap {
	return MyHashMap{
		hash: make(map[int]int),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	this.hash[key] = value
}

func (this *MyHashMap) Get(key int) int {
	val, exists := this.hash[key]
	if !exists {
		return -1
	}

	return val
}

func (this *MyHashMap) Remove(key int) {
	delete(this.hash, key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
