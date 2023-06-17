package _interface

type Service interface {
	Init()

	Run()

	Stop()
}

// 服务基础数据
type ServiceData struct {
	ServiceId uint64 // 服务ID
}
