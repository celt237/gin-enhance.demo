package model

type ResultGeneric[T any] struct {
	// 错误码 200:成功 其他:失败
	Code int `json:"code"`

	// 数据
	Data T `json:"data"`

	// 错误消息
	Message string `json:"message"`

	// 耗时 单位:毫秒
	Time int64 `json:"time"`

	// 跟踪ID
	TraceId string `json:"traceId"`
}

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Time    int64       `json:"time"`
	TraceId string      `json:"traceId"`
}
