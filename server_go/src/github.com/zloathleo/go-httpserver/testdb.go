package main

import (
	"github.com/Cistern/catena"
	"fmt"
)

func insertData(){
	db, err := catena.NewDB("d:/abc1", 500, 20)
	if err != nil {
		fmt.Println(err)
	}

	var rows [5000]catena.Row

	for i:=0; i < len(rows); i++ {
		rows[i] = catena.Row{
			Source: "src",
			Metric: "met_1" ,
			Point:catena.Point{Timestamp:int64(1514736000 + i),Value:99.8 + float64(i)},
		}
	}

	rowss  := rows[:]
	fmt.Println("InsertRows")
	db.InsertRows(rowss )
	fmt.Println("InsertRows ed")
	err = db.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func searchData(start, end int64){
	db, err := catena.OpenDB("d:/abc1", 500, 20)
	if err != nil {
		fmt.Println(err)
	}

	partitionList := db.GetPartitionList()
	//fmt.Println(partitionList)
	i := partitionList.NewIterator()
	for i.Next() {
		val, _ := i.Value()
		val.Hold()
		if val.MaxTimestamp() > start && val.MinTimestamp() < end {
			it,_ := val.NewIterator("src","met_1")
            for  it.Next() == nil {
				p := it.Point()
				fmt.Printf("%d - %f\n",p.Timestamp   ,p.Value)
			}
		}
		val.Release()
	}

	fmt.Println("searchData ed")
	err = db.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//insertData()

	searchData(1514736000,1514736200)
	//ts := int64(0)

	//parallelism := 4
	//runtime.GOMAXPROCS(parallelism)
	//wg := sync.WaitGroup{}
	//wg.Add(parallelism)
	//
	//work := make(chan []catena.Row, parallelism)
	//
	//for i := 0; i < parallelism; i++ {
	//	go func() {
	//		for rows := range work {
	//			err := db.InsertRows(rows)
	//			if err != nil {
	//				wg.Done()
	//				fmt.Println(err)
	//			}
	//		}
	//
	//		wg.Done()
	//	}()
	//}

	//for n := 0; n < 500; n++ {
	//
	//	rows := []catena.Row{}
	//	for i := 0; i < 10000; i++ {
	//		rows = append(rows, catena.Row{
	//			Source: "src",
	//			Metric: "met_" + strconv.Itoa(i),
	//			Point: catena.Point{
	//				Timestamp: ts,
	//				Value:     float64(i),
	//			},
	//		})
	//	}
	//
	//	ts++
	//
	//	work <- rows
	//}

	//close(work)
	//wg.Wait()


}
