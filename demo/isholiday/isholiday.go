package isholiday

import (
	"fmt"
	"time"
)

type date struct{
	year int
	month int
	day int
}
var chinese_holiday :=[]date{
{2018,5,12},
}
//判断是否为假期，通过遍历策略判断
func isholiday(year int, month int, day int) {
	local, _ := time.LoadLocation("Local")
	theDay := time.Date(year, time.Month(month), day, 0, 0, 0, 0, local)
	theDay.Weekday()
	fmt.Println(theDay)
}
