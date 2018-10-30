package dstruct

import "time"

type Alarm struct {
	Module string `json:"module"`
	Err string `json:"err"`
	TimeStamp int64  `json:"timestamp"`
}

type PushDas struct {
	Rows      []*Das `json:"rows"`
	TimeStamp uint64 `json:"timestamp"`
}

type Das struct {
	PointName string  `json:"name,omitempty"`
	Value     float64 `json:"value"`
	TimeStamp uint64  `json:"timestamp"`
}

type HDas struct {
	V float64 `json:"v"`
	T uint64  `json:"t"`
}

type HisReqParam struct {
	From     time.Time
	To       time.Time
	Times 	 []*HisReqParamTime
	Dur      int // 时长
	Interval int // 时间戳间隔[1...60]
	Count    int // Dur/Interval + 1
}

type HisReqParamTime struct {
	From     int64
	To       int64
	Count    int // Dur/Interval + 1
	Interval int // 时间戳间隔[1...60]
}

//历史请求总的上下文
type HisReqContext struct {
	Count     int//历史时间戳总数
	Index 	  int//历史时间戳序号
}
