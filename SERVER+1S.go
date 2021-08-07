package main

import (
    "log"
	"github.com/robfig/cron/v3"
	"fmt"
    "os"
)

func init() {
    // get log
    // 已 只写入文件|没有时创建|文件尾部追加 的形式打开这个文件
    logFile, err := os.OpenFile(`./+1s.log`, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        panic(err)
    }
    // fileoutput
    log.SetOutput(logFile)
}

func main() {
    log.Println(`给长者+1s`)
    log.Println(`人的一生当然要靠自我奋斗,当然也要考虑历史的进程`)
	log.Println(`我为长者续一秒，气的轮子满地爬`)
}


//定时部分
func main() {
	cron2 := cron.New() //create 1 cron

	//5s for 1
	err:= cron2.AddFunc("*/5 * * * * *", xuming)
	if err!=nil{
	   fmt.Println(err)
	}

	//on/off
	cron2.Start()
	defer cron2.Stop()
	select {
	  //find,just for a for
	}
}

func xuming()  {
	fmt.Println("我为长者续一秒，气的轮子满地爬")
	log.Println(`给长者+1s`)
    log.Println(`人的一生当然要靠自我奋斗,当然也要考虑历史的进程`)
	log.Println(`我为长者续一秒，气的轮子满地爬`)
}
