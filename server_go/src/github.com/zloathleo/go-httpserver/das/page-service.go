package das

func renderPageDashboard() *PushDas {
	var rows []*Das

	addData(&rows,"1_00001")
	addData(&rows,"1_00002")
	addData(&rows,"4_40001")
	addData(&rows,"4_40002")
	addData(&rows,"1_00008")

	addData(&rows,"4_40003")
	addData(&rows,"5_40013")
	addData(&rows,"1_00004")
	addData(&rows,"1_00005")

	addData(&rows,"4_40004")
	addData(&rows,"5_40083")
	addData(&rows,"1_00006")
	addData(&rows,"1_00007")

	return &PushDas{rows, 0}
}

//控制页面数据
func renderPageControl() *PushDas {

	var rows []*Das

	addData(&rows,"1_00001")
	addData(&rows,"1_00002")
	addData(&rows,"4_40001")
	addData(&rows,"4_40002")
	addData(&rows,"1_00008")

	addData(&rows,"4_40003")
	addData(&rows,"5_40013")
	addData(&rows,"1_00004")
	addData(&rows,"1_00005")

	addData(&rows,"4_40004")
	addData(&rows,"5_40083")
	addData(&rows,"1_00006")
	addData(&rows,"1_00007")

	return &PushDas{rows, 0}

}

func addData(rows *[]*Das, key string){
	exist,value := getCurrentData(key)
	if exist {
		*rows = append(*rows, &Das{key,value,0} )
	}
}
