package history

type Das struct {
	Value     float32
	TimeStamp uint64
}

type Block struct {
	ID     uint64
	Dass   []*Das
}
