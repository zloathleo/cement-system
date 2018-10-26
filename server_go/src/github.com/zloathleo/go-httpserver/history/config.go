package history

//
///**
//*历史文件存储
//1.文件头
//1)magic number 0x13304F3
//2)info
//   -info-length uint32
//   -info-content version uint32
//3)index块 24 * 6 = 144 个时间块(10分钟) 位置偏移
//	144 * 位置偏移[int64(8Byte)]
//
//2.时间块chunk 10分钟
//   block1 - id[uint64(8Byte)] + 实时数据长度[uint16(2Byte)] + 统计数据长度[uint16(2Byte)] + 实时数据[ 时间偏移uint16(2Byte) + float32(4Byte) * n ] + 数据统计[statistics]
//   block2 - id[uint64(8Byte)] + 实时数据长度[uint16(2Byte)] + 统计数据长度[uint16(2Byte)] + 实时数据[ 时间偏移uint16(2Byte) + float32(4Byte) * n ] + 数据统计[statistics]
//	......
// */
//
//var (
//	fileHeader = []byte{20,12, 8,19}//文件头
//	indexBlock = make([]byte, 144 * 8)//index块
//	todayFilePath string
//)
//
//func init(){
//	//History Data
//	todayFilePath = "d:/20180930.hd"
//	initializeTodayFile()
//}
//
//func check(err error){
//	logger.Println(err)
//}
//
//func isFileExist(f string) bool {
//	_, err := os.Stat(f)
//	return err == nil || os.IsExist(err)
//}
//
//func initializeTodayFile(){
//	_exist := isFileExist(todayFilePath)
//	var file *os.File
//	var err error
//    if _exist{
//		file, err = os.Open(todayFilePath)
//		check(err)
//		defer file.Close()
//
//		//check header
//		header := make([]byte, 4)
//		count, err := file.Read(header)
//		check(err)
//
//		//检查包头
//		if count == 4 && fileHeader[0] == header[0] && fileHeader[1] == header[1] && fileHeader[2] == header[2] && fileHeader[3] == header[3] {
//
//		}
//		logger.Println("count:",count)
//	}else{
//		file, err = os.Create(todayFilePath)
//		check(err)
//		defer file.Close()
//
//		//
//		count, err := file.Write(fileHeader)
//		check(err)
//		logger.Println("count:",count)
//
//		//144 int64
//		count, err = file.Write(indexBlock)
//		check(err)
//		logger.Println("count:",count)
//	}
//}
//
//
//
//func Print(){
//
//}
