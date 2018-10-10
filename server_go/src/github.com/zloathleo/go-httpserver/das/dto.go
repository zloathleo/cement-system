package das

type PushDas struct {
	Rows      []*Das `json:"rows"`
	TimeStamp uint64 `json:"timestamp"`
}

type Das struct {
	PointName string  `json:"name"`
	Value     float64 `json:"value"`
	TimeStamp uint64   `json:"timestamp"`
}