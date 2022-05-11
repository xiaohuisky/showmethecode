package channel

import (
	"sync"
	"unsafe"
)

// hchan channel 结构体
// https://github.com/golang/go/blob/master/src/runtime/chan.go#L33
type hchan struct {
	qcount   uint // 元素个数
	dataqsiz uint // 循环队列的长度
	buf      unsafe.Pointer // 缓存区数据指针
	elemsize uint16 // 能够收发的元素大小
	closed   uint32
	elemtype *_type // 能够收发的元素类型
	sendx uint // 发送操作处理到的位置
	recvx uint // 接收操作处理到的位置
	recvq waitq // 存储了当前 channel 由于缓存区不足而阻塞的 goroutine 列表，这些等待队列使用双向链表表示
	sendq waitq // 存储了当前 channel 由于缓存区不足而阻塞的 goroutine 列表，这些等待队列使用双向链表表示

	lock sync.Mutex
}

type waitq struct { // 表示一个在等待列表中的 Goroutine
	first *sudog // 前指针
	last  *sudog // 后指针
}

// _type https://github.com/golang/go/blob/master/src/runtime/type.go#L35
type _type struct {}

// sudog https://github.com/golang/go/blob/master/src/runtime/runtime2.go#L348
type sudog struct {}