package dstruct

import "time"

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
	Times 	 []int64
	Dur      int // 时长
	Interval int // 时间戳间隔[1...60]
	Count    int // Dur/Interval + 1
}

type HisReqContext struct {
	Count     int//历史时间戳总数
	Index 	  int//历史时间戳序号
}
