package util

import (
	"math/rand"
	"time"
)

func Sleep() {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
}

func SleepN(n int) {
	time.Sleep(time.Duration(n) * time.Second)
}

func Loop(f func(i int, n int), n int) {
	for i := 0; i < n; i++ {
		f(i, n)
	}
}
func ILoop(f func(i int)) {
	for i := 0; ; i++ {
		f(i)
	}
}
