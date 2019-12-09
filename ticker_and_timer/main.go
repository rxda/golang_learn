package main

import (
	"math/rand"
	"time"
)

//模拟获取接口数据
func mockGetData(){
	r := rand.Intn(100)
	if r== 1{
		time.Sleep(time.Second * 5)
	}else {
		time.Sleep(time.Second)
	}
}

//启动时立即查询一次当前小时数据， 再定时每小时的11分查询小时数据
//timer定时为当前小时（当前分钟 <10 ）或下小时的11分（当前分钟 >10 ）
//问题： ticker 创建的第0秒不会触发，会少一小时的数据
func timingGetData(timer *time.Timer){
	mockGetData()
	<- timer.C
	ticker := time.NewTicker(time.Hour) //每隔1小时查询一次
	for range ticker.C{
		mockGetData()
	}
}
// mockGetData()这一段在原程序里很长，比较杂乱
func timingGetData1(timer *time.Timer){
	mockGetData()
	<- timer.C
	ticker := time.NewTicker(time.Hour) //每隔1小时查询一次
	for ;; <-ticker.C{
		mockGetData()
	}
}

// 如果mockGetData 获取数据时间变长，可能会错过ticker的触发，
func timingGetData2(timer *time.Timer){
	firstRun := true
	var ticker *time.Ticker
	for{
		mockGetData()
		if firstRun{
			<- timer.C
			firstRun = false
			ticker = time.NewTicker(time.Hour) //每隔1小时查询一次
		}else{
			<-ticker.C
		}
	}
}

//准备改成这样, 或者timingGetData2
func timingGetData3(timer *time.Timer){
	firstRun := true
	ticker := time.NewTicker(time.Hour) //每隔1小时查询一次
	for;; <-ticker.C{
		go mockGetData()
		if firstRun{
			<- timer.C
			firstRun = false
		}
	}
}