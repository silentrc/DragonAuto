package utils

import (
	"go.uber.org/zap"
	"runtime/debug"
)

// Recover 出错打印错误堆栈
func Recover() {
	if err := recover(); err != nil {
		zap.L().Sugar().Errorf("Panic: %v\n%s", err, string(debug.Stack()))
	}
}

// Go 带recover模式带go
func Go(f func()) {
	go func() {
		defer Recover()
		f()
	}()
}
