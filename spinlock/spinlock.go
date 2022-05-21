package spinlock

import (
	"runtime"
	"sync"
	"sync/atomic"
)

// =================================================================
// 自旋锁
// =================================================================

type spinLock uint32

func (sl *spinLock) Lock() {
	backoff := 1
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		for i := 0; i < backoff; i++ {
			runtime.Gosched()
		}
		backoff <<= 1
	}
}
func (sl *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}

// NewSpinLock 实例化自旋锁
func NewSpinLock() sync.Locker {
	return new(spinLock)
}
