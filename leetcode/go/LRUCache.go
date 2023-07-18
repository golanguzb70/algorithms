package leetcode

/*https://leetcode.com/problems/lru-cache/*/

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	space      int
	cache      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	cache := map[int]*Node{}
	head, tail := &Node{0, 0, nil, nil}, &Node{0, 0, nil, nil}
	head.next = tail
	tail.prev = head
	return LRUCache{capacity, cache, head, tail}
}

func (lru *LRUCache) AddNode(node *Node) {
	node.next = lru.head.next
	node.prev = lru.head
	node.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) RemoveNode(node *Node) {
	prev, next := node.prev, node.next
	node.prev.next = next
	node.next.prev = prev
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		lru.RemoveNode(node)
		lru.AddNode(node)
		return node.value
	} else {
		return -1
	}
}

func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.cache[key]; ok {
		node.value = value
		lru.RemoveNode(node)
		lru.AddNode(node)
	} else {
		if lru.space == 0 {
			toRemove := lru.tail.prev
			lru.RemoveNode(toRemove)
			delete(lru.cache, toRemove.key)
		} else {
			lru.space--
		}
		newNode := &Node{key, value, nil, nil}
		lru.AddNode(newNode)
		lru.cache[key] = newNode
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
