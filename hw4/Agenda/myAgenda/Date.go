package myAgenda

import "fmt"

//包装时间数据
type Date struct {
	M_year   int
	M_month  int
	M_day    int
	M_hour   int
	M_minute int
}

//判断一个Date结构体数据是否合法
var date = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};

func isleap(year int) bool {
	if year%4 == 0 {
		if year%400 != 0 && year%100 == 0 {
			return false
		}
		return true
	}
	return false
}
func IsValid(t_date Date) bool {
	judge := true
	if isleap(t_date.M_year) {
		date[1] = 29
	} else {
		date[1] = 28
	}
	if t_date.M_year > 9999 || t_date.M_year < 1000 {
		judge = false
	} else if t_date.M_month > 12 || t_date.M_month < 1 {
		judge = false
	} else if t_date.M_day > date[t_date.M_month-1] || t_date.M_day < 1 {
		judge = false
	} else if t_date.M_hour > 23 || t_date.M_hour < 0 {
		judge = false
	} else if t_date.M_minute > 59 || t_date.M_minute < 0 {
		judge = false
	}
	return judge
}

//比较两个Date结构体的大小
func CompareDate(date1, date2 Date) int {
	if date1.M_year < date2.M_year {
		return -1
	} else if date1.M_year > date2.M_year {
		return 1
	}

	if date1.M_month < date2.M_month {
		return -1
	} else if date1.M_month > date2.M_month {
		return 1
	}

	if date1.M_day < date2.M_day {
		return -1
	} else if date1.M_day > date2.M_day {
		return 1
	}

	if date1.M_hour < date2.M_hour {
		return -1
	} else if date1.M_hour > date2.M_hour {
		return 1
	}

	if date1.M_minute < date2.M_minute {
		return -1
	} else if date1.M_minute > date2.M_minute {
		return 1
	}

	return 0
}

//从Date转为string
func (d *Date) DateToString() string {
	var str string
	str = fmt.Sprintf("%04d-%02d-%02d/%02d:%02d", d.M_year, d.M_month, d.M_day, d.M_hour, d.M_minute)
	return str
}

//从string转为Date
func StringToDate(str string) Date {
	var d Date
	fmt.Sscanf(str, "%d-%d-%d/%d:%d", &d.M_year, &d.M_month, &d.M_day, &d.M_hour, &d.M_minute)
	return d
}
