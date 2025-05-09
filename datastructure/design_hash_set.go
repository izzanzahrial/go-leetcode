package datastructure

// https://leetcode.com/problems/design-hashset/
type MyHashSet struct {
	hash map[int]struct{}
}

func ConstructorHashSet() MyHashSet {
	return MyHashSet{
		hash: make(map[int]struct{}),
	}
}

func (this *MyHashSet) Add(key int) {
	this.hash[key] = struct{}{}
}

func (this *MyHashSet) Remove(key int) {
	delete(this.hash, key)
}

func (this *MyHashSet) Contains(key int) bool {
	_, contains := this.hash[key]
	return contains
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
