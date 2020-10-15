package lru_test

import (
	"fmt"
	"golang_learn/go/lru"
	"testing"
)

func TestLRU(t *testing.T) {
	oneLRU := lru.NewLRU(5)
	oneLRU.Put(1,"a")
	oneLRU.Put(2,"b")
	oneLRU.Put(3,"c")
	oneLRU.Put(4,"d")
	oneLRU.Put(5,"e")

	fmt.Println("原始链表为:"+oneLRU.String())

	oneLRU.Get(4)
	fmt.Println("获取key为4的元素之后的链表:"+oneLRU.String())
	oneLRU.Put(6,"f")
	fmt.Println("新添加一个key为6之后的链表:"+oneLRU.String())
	oneLRU.Remove(3)
	fmt.Println("移除key=3的之后的链表:"+oneLRU.String())
}
