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

type MyHashSet2 struct {
	set [1000001]bool
}

func Constructor() MyHashSet2 {
	return MyHashSet2{}
}

func (this *MyHashSet2) Add(key int) {
	this.set[key] = true
}

func (this *MyHashSet2) Remove(key int) {
	this.set[key] = false
}

func (this *MyHashSet2) Contains(key int) bool {
	return this.set[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
