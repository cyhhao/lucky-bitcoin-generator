package tools

import (
	"time"
	"fmt"
	"strconv"
)

type TimeTest struct {
	Name      string
	startTime int64
	lastTime  int64
}

func (this *TimeTest) Start(stepName string) {
	this.Name = stepName
	this.startTime = time.Now().UnixNano()

	this.lastTime = this.startTime
	fmt.Println(this.Name + " start at: " + time.Unix(0, this.startTime).Format("2006-01-02 15:04:05"))
}

func (this *TimeTest) Ding(name ...string) {
	now := time.Now().UnixNano()
	var printName = ""
	if len(name) > 0 {
		printName = name[0]
	}
	delta := float64(now-this.lastTime) / 10e8
	fmt.Println(printName + " used: " + strconv.FormatFloat(delta, 'f', 2, 64) + " s")
	this.lastTime = now

}

func (this *TimeTest) End(name ...string) {
	now := time.Now().UnixNano()
	if len(name) > 0 {
		this.Ding(name[0])
	} else {
		this.Ding()
	}
	delta := float64(now-this.startTime) / 10e8
	fmt.Println(this.Name + " total used: " + strconv.FormatFloat(delta, 'f', 4, 64) + " s")
	fmt.Println(this.Name + " end at: " + time.Unix(0, now).Format("2006-01-02 15:04:05"))
}
