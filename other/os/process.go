package os

type LockT struct {
	Flag int
}

func start(mutex *LockT) {
	mutex.Flag = 0
}

func lock(mutex *LockT) {
	// 自旋等待
	for mutex.Flag == 1 {
		// 如果在此发生中断，两个线程都将持有锁，都可以进入临界区
	}
	mutex.Flag = 1
}

func unlock(mutex *LockT) {
	mutex.Flag = 0
}
