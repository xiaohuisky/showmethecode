package algorithm

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash 函数类型，默认为 crc32.ChecksumIEEE
type Hash func(data []byte) uint32

// Map 一致性哈希算法的主数据结构
type Map struct {
	hash     Hash
	replicas int
	keys     []int // Sorted
	hashMap  map[int]string
}

// New 构造函数，允许自定义虚拟节点倍数和 Hash 函数
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}

	return m
}

// Add 允许传入 0 或 多个真实节点的名称
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ { // 创建 replicas 个虚拟节点，通过添加编号的方式区分不同虚拟节点
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key // 添加虚拟节点和真实节点的映射关系
		}
	}
	sort.Ints(m.keys) // 对环的哈希值进行排序
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))                   // 计算 key 的 哈希值
	idx := sort.Search(len(m.keys), func(i int) bool { // 顺时针找到第一个匹配的虚拟节点的下标
		return m.keys[i] >= hash
	})
	return m.hashMap[m.keys[idx%len(m.keys)]] // 通过 hashMap 映射得到真实的节点。（因为 m.keys 是一个环状的结构，所以用取余的方式）
}
