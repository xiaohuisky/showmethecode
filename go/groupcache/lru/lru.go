package lru

import "container/list"

// Cache 使用最近最少使用算法，并发访问不安全
type Cache struct {
	maxBytes  int64                         // 允许使用的最大内存
	nbytes    int64                         // 当前已经使用的内存
	ll        *list.List                    // 双向链表
	cache     map[string]*list.Element      // 键是字符串、值是双向链表中对应节点的指针
	OnEvicted func(key string, value Value) // 可选参数，某条记录被移除时的回调函数，在清除条目时执行
}

// entry 双向链表节点的数据类型，在链表中仍保存每个值对应的 key 的好处在于，淘汰队首节点时，需要用 key 从字典中删除对应的映射
type entry struct {
	key   string
	value Value
}

// Value 使用长度来计算字节数，此处使用 interface 允许值是任意类型
type Value interface {
	Len() int // 返回值所占用的内存大小
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Get 从字典中找到对应的双向链表的节点，并将节点移动到队尾
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok { // 如果键对应的节点存在
		c.ll.MoveToFront(ele) // 将对应节点移动到队尾（双向链表作为队列，队首队尾是相对的，这里约定 front 为队尾），并返回查找到的值
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOldest 删除，实际上是缓存淘汰，即移除最近最少访问的节点（队首）
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back() // 取队首节点，从链表中删除
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)                         // 从字典中删除该节点的映射关系
		c.nbytes -= int64(len(kv.key) + kv.value.Len()) // 更新当前使用内存长度
		if c.OnEvicted != nil {                         // 当回调函数不为空时进行调用
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Add 新增/修改
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok { // 如果对应键存在，更新对应节点的值，并将该节点移到队尾
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len() - kv.value.Len()) // 更新内存
		kv.value = value
	} else { // 如果对应键不存在则是新增场景，
		ele := c.ll.PushFront(&entry{key: key, value: value}) // 队尾添加新节点，并在字典中添加 key 和节点的映射关系
		c.cache[key] = ele
		c.nbytes += int64(len(key) + value.Len()) // 更新内存
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytes { // 如果超过设定的最大值，移除最少访问的节点
		c.RemoveOldest()
	}
}

// Len 查询添加了多少条数据
func (c *Cache) Len() int {
	return c.ll.Len()
}
