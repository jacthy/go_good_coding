package main

func testCache()  {
	cache := newLRUCache(2)
	cache.put(1,1)
	cache.put(2,2)
	println("-")
	println(cache.get(2),"-",cache.get(1))
	cache.put(1,11)
	cache.put(2,21)
	println("--")
	println(cache.get(2),"-",cache.get(1))
	cache.put(3,3)
	println("---")
	println(cache.get(2),"-",cache.get(1))
}

//请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
//实现 LRUCache 类：
//LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
//int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
//void put(int key, int value) 如果关键字 key
//已经存在，则变更其数据值 value ；
//如果不存在，则向缓存中插入该组 key-value 。
//如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
//函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

type LRUCache struct {
	cap   int
	// 存队头&队尾实现O(1)的复杂度
	head  *value
	tail  *value
	store map[int]*value
}

type value struct {
	key  int
	val  int
	prev *value
	next *value
}

func newLRUCache(cap int) *LRUCache {
	return &LRUCache{cap: cap,store: make(map[int]*value)}
}

func (l *LRUCache) put(key, val int) {
	if v, ok := l.store[key]; ok {
		v.val = val
		if v.prev == nil {//本身是队头
			return
		}
		v.prev = v.next
		l.head = v
		return
	}
	if len(l.store) == l.cap {
		delete(l.store, l.tail.key)
		l.tail = l.tail.prev
		l.tail.next = nil
	}
	newKey := &value{key: key}
	if l.head == nil {
		l.head = newKey
	}else {
		l.head.prev = newKey
		l.head = newKey
	}
	if l.tail == nil {
		l.tail = newKey
		l.head.next = newKey
		l.tail.prev = l.head
	}
	l.store[key] = newKey
}

func (l *LRUCache) get(key int) int{
	v,ok :=l.store[key]
	if ok {
		return v.val
	}
	return -1
}
