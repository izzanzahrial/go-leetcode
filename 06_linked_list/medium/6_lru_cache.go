package linkedList

// https://leetcode.com/problems/lru-cache/description/
type Node struct {
	key, val   int
	next, prev *Node
}

type LRUCache struct {
	cacheMap   map[int]*Node
	head, tail *Node
	capacity   int
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return LRUCache{
		cacheMap: make(map[int]*Node, 0),
		head:     head,
		tail:     tail,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exist := this.cacheMap[key]; exist {
		this.moveToFront(node)
		return node.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exist := this.cacheMap[key]; exist {
		node.val = value
		this.moveToFront(node)
		return
	}

	if len(this.cacheMap) == this.capacity {
		lru := this.tail.prev
		this.remove(lru)
		delete(this.cacheMap, lru.key)
	}

	node := &Node{key: key, val: value}
	this.cacheMap[key] = node
	this.addToFront(node)
}

func (this *LRUCache) addToFront(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToFront(node *Node) {
	this.remove(node)
	this.addToFront(node)
}
