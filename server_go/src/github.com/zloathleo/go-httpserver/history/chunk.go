package history

func slocksToBytes(blocks []*Block ){

	count:= len(blocks)
	for i:=0; i<count;i++  {
		block := blocks[i]
		blockToBytes(0,block.ID,block.Dass)
	}
}