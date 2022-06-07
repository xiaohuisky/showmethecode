package groupcache

// 缓存值的抽象与封装

type ByteView struct {
	b []byte // 存储真实的缓存值
}

// Len 查看所占的内存大小
func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b) // b 是只读的, 返回一个拷贝，防止缓存值被外部程序修改。
	return c
}
