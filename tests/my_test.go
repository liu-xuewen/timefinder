package tests

import (
	"fmt"
	"github.com/liu-xuewen/timefinder"
	"testing"
	"time"
)

func TestMyTimeFinder(t *testing.T) {
	var msg string

	// 初始化timefinder
	segmenter := timefinder.New()

	msg = "每月1号上午提醒我取公积金"
	test(segmenter, msg)

}

func test(segmenter *timefinder.TimeFinder, msg string) {
	start := time.Now()

	// 解析话语词汇
	extract, repeat := segmenter.TimeExtract(msg)
	fmt.Println(msg)
	if len(extract) > 0 {
		fmt.Println(extract[0].Format(timeFormat), repeat)
	}
	fmt.Println(time.Now().Sub(start).Seconds())

}
