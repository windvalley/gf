package gtime_test

import (
	"testing"
	"time"

	"github.com/gogf/gf/g/os/gtime"
	"github.com/gogf/gf/g/test/gtest"
)

func Test_SetTimeZone(t *testing.T) {
	gtest.Case(t, func() {
		gtime.SetTimeZone("Asia/Shanghai")
		gtest.Assert(time.Local.String(), "Asia/Shanghai")
	})
}

func Test_Nanosecond(t *testing.T) {
	gtest.Case(t, func() {
		nanos := gtime.Nanosecond()
		timeTemp := time.Unix(0, nanos)
		gtest.Assert(nanos, timeTemp.UnixNano())
	})
}

func Test_Microsecond(t *testing.T) {
	gtest.Case(t, func() {
		micros := gtime.Microsecond()
		timeTemp := time.Unix(0, micros*1e3)
		gtest.Assert(micros, timeTemp.UnixNano()/1e3)
	})
}

func Test_Millisecond(t *testing.T) {
	gtest.Case(t, func() {
		millis := gtime.Millisecond()
		timeTemp := time.Unix(0, millis*1e6)
		gtest.Assert(millis, timeTemp.UnixNano()/1e6)
	})
}

func Test_Second(t *testing.T) {
	gtest.Case(t, func() {
		s := gtime.Second()
		timeTemp := time.Unix(s, 0)
		gtest.Assert(s, timeTemp.Unix())
	})
}

func Test_Date(t *testing.T) {
	gtest.Case(t, func() {
		gtest.Assert(gtime.Date(), time.Now().Format("2006-01-02"))
	})
}

func Test_Datetime(t *testing.T) {
	gtest.Case(t, func() {
		datetime := gtime.Datetime()
		timeTemp, err := gtime.StrToTime(datetime, "Y-m-d H:i:s")
		if err != nil {
			t.Error("test fail")
		}
		gtest.Assert(datetime, timeTemp.Time.Format("2006-01-02 15:04:05"))
	})
}

/*
func Test_parseDateStr(t *testing.T) {
	gtest.Case(t, func() {
		//正常日期列表
		var testDates = []string{
			"2006-01-02",
			"2006/01/02",
			"2006.01.02",
			"06.01.02",
			"02.jan.2006",
		}

		for _, item := range testDates {
			year, month, day := parseDateStr(item)
			gtest.Assert(year, 2006)
			gtest.Assert(month, 1)
			gtest.Assert(day, 2)
		}

		//异常日期列表
		var testDatesFail = []string{
			"2006.01",
			"06..02",
		}

		for _, item := range testDatesFail {
			year, month, day := parseDateStr(item)
			gtest.Assert(year, 0)
			gtest.Assert(month, 0)
			gtest.Assert(day, 0)
		}

	})
}
*/
//
func Test_ConvertZone(t *testing.T) {
	gtest.Case(t, func() {
		//现行时间
		nowUTC := time.Now().UTC()
		testZone := "America/Los_Angeles"

		//转换为洛杉矶时间
		t1, err := gtime.ConvertZone(nowUTC.Format("2006-01-02 15:04:05"), testZone, "")
		if err != nil {
			t.Error("test fail")
		}

		//使用洛杉矶时区解析上面转换后的时间
		laStr := t1.Time.Format("2006-01-02 15:04:05")
		loc, err := time.LoadLocation(testZone)
		t2, err := time.ParseInLocation("2006-01-02 15:04:05", laStr, loc)

		//判断是否与现行时间匹配
		gtest.Assert(t2.UTC().Unix(), nowUTC.Unix())

	})

	//test err
	gtest.Case(t, func() {
		//现行时间
		nowUTC := time.Now().UTC()
		//t.Log(nowUTC.Unix())
		testZone := "errZone"

		//错误时间输入
		_, err := gtime.ConvertZone(nowUTC.Format("06..02 15:04:05"), testZone, "")
		if err == nil {
			t.Error("test fail")
		}
		//错误时区输入
		_, err = gtime.ConvertZone(nowUTC.Format("2006-01-02 15:04:05"), testZone, "")
		if err == nil {
			t.Error("test fail")
		}
		//错误时区输入
		_, err = gtime.ConvertZone(nowUTC.Format("2006-01-02 15:04:05"), testZone, testZone)
		if err == nil {
			t.Error("test fail")
		}
	})
}

func Test_StrToTime(t *testing.T) {
	gtest.Case(t, func() {
		//正常日期列表
		var testDatetimes = []string{
			"2006-01-02 15:04:05",
			"2006/01/02 15:04:05",
			"2006.01.02 15:04:05.000",
			"2006.01.02 - 15:04:05",
			"2006.01.02 15:04:05 +0800 CST",
			"2006-01-02T20:05:06+05:01:01",
			"2006-01-02T14:03:04Z01:01:01",
			"2006-01-02T15:04:05Z",
			"02-jan-2006 15:04:05",
			"02/jan/2006 15:04:05",
			"02.jan.2006 15:04:05",
			"02.jan.2006:15:04:05",
		}

		for _, item := range testDatetimes {
			timeTemp, err := gtime.StrToTime(item)
			if err != nil {
				t.Error("test fail")
			}
			gtest.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2006-01-02 15:04:05")
		}

		//正常日期列表，时间00:00:00
		var testDates = []string{
			"2006.01.02",
			"2006.01.02 00:00",
			"2006.01.02 00:00:00.000",
		}

		for _, item := range testDates {
			timeTemp, err := gtime.StrToTime(item)
			if err != nil {
				t.Error("test fail")
			}
			gtest.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2006-01-02 00:00:00")
		}

		//异常日期列表
		var testDatesFail = []string{
			"2006.01",
			"06..02",
			"20060102",
		}

		for _, item := range testDatesFail {
			_, err := gtime.StrToTime(item)
			if err == nil {
				t.Error("test fail")
			}
		}

		//test err
		_, err := gtime.StrToTime("2006-01-02 15:04:05", "aabbccdd")
		if err == nil {
			t.Error("test fail")
		}
	})
}

func Test_ParseTimeFromContent(t *testing.T) {
	gtest.Case(t, func() {
		timeTemp := gtime.ParseTimeFromContent("我是中文2006-01-02 15:04:05我也是中文", "Y-m-d H:i:s")
		gtest.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2006-01-02 15:04:05")

		timeTemp1 := gtime.ParseTimeFromContent("我是中文2006-01-02 15:04:05我也是中文")
		gtest.Assert(timeTemp1.Time.Format("2006-01-02 15:04:05"), "2006-01-02 15:04:05")

		timeTemp2 := gtime.ParseTimeFromContent("我是中文02.jan.2006 15:04:05我也是中文")
		gtest.Assert(timeTemp2.Time.Format("2006-01-02 15:04:05"), "2006-01-02 15:04:05")

		//test err
		timeTempErr := gtime.ParseTimeFromContent("我是中文", "Y-m-d H:i:s")
		if timeTempErr != nil {
			t.Error("test fail")
		}
	})
}

func Test_FuncCost(t *testing.T) {
	gtest.Case(t, func() {
		gtime.FuncCost(func() {

		})
	})
}
