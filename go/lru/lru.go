package lru

import (
	"fmt"
	"strings"
)

type LRU struct {
	currentSize int
	capacity    int
	caches      map[interface{}]*Node
	first       *Node
	last        *Node
}

func NewLRU(size int) *LRU {
	return &LRU{
		currentSize: 0,
		capacity:    size,
		caches:      make(map[interface{}]*Node),
	}
}

// 添加新元素
func (l *LRU) Put(key, value interface{}) {
	if node, ok := l.caches[key]; !ok {
		// new element
		// if size > cap, remove last node
		if len(l.caches) >= l.capacity {
			l.removeLast()
		}
		// create new node
		node = NewNode(key, value)
		l.caches[key] = node
		l.currentSize++
	} else {
		node.Value = value
	}
	l.moveToHead(l.caches[key])
}

// 移除最后一个节点
func (l *LRU) removeLast() {
	delete(l.caches, l.last.Key)
	if l.last != nil {
		l.last = l.last.pre
		if l.last == nil {
			l.first = nil
		} else {
			l.last.next = nil
		}
	}
}

// 节点移动至头部
func (l *LRU) moveToHead(node *Node) {
	if l.first == node {
		return
	}
	if node.next != nil {
		node.next.pre = node.pre
	}

	if node.pre != nil {
		node.pre.next = node.next
	}

	if node == l.last {
		l.last = l.last.pre
	}

	if l.first == nil || l.last == nil {
		l.first = node
		l.last = node
		return
	}

	node.next = l.first
	l.first.pre = node
	l.first = node
	l.first.pre = nil
}

// 根据key移除节点
func (l *LRU) Remove(key interface{}) {
	if node, ok := l.caches[key]; ok {
		if node.pre != nil {
			node.pre.next = node.next
		}

		if node.next != nil {
			node.next.pre = node.pre
		}

		if node == l.first {
			l.first = node.next
		}

		if node == l.last {
			l.last = node.pre
		}
	}
	delete(l.caches, key)
}

// 根据key获取元素
func (l *LRU) Get(key interface{}) interface{} {
	if node, ok := l.caches[key]; ok {
		l.moveToHead(node)
		return node.Value
	} else {
		return nil
	}
}

// 清楚所有节点
func (l *LRU) Clear() {
	l.first = nil
	l.last = nil
	l.caches = make(map[interface{}]*Node)
}

func (l *LRU) String() string {
	sb := strings.Builder{}
	for node := l.first; node != nil; node = node.next {
		sb.WriteString(fmt.Sprintf("%v:%v", node.Key, node.Value))
		sb.WriteString(",")
	}
	return strings.TrimRight(sb.String(), ",")
}
