package main

import (
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"go-pluspoint/sqldata"
	"go-pluspoint/utils"
	"go-pluspoint/utils/kdb"
	"time"
)

type PlusPointsConfig struct {
	CalcPoint     string
	Point         string
	Aggr          string
	MinValue      string
	MaxValue      string
	SamplingUnit  string
	SamplingValue string
}

type PlusPointsConfigSlice struct {
	Points           []PlusPointsConfig
	OneMinutePoints  []PlusPointsConfig
	FiveMinutePoints []PlusPointsConfig
	TenMinutePoints  []PlusPointsConfig
	OneHourPoints    []PlusPointsConfig
	OneDayPoints     []PlusPointsConfig
}

func ReadPlusPointsConfig() (*PlusPointsConfigSlice, error) {
	var ps PlusPointsConfigSlice
	ctx := gctx.New()
	data, err := g.Cfg().Data(ctx)
	if err != nil {
		return &ps, err
	}
	for calcpoint, v := range data["pluspoints"].(map[string]any) {
		if calcpoint == "example" {
			continue
		}
		var condition map[string]any
		condition = v.(map[string]any)

		var p PlusPointsConfig
		p.CalcPoint = calcpoint
		p.Point = condition["point"].(string)
		p.Aggr = condition["aggr"].(string)
		p.MinValue = condition["minValue"].(string)
		p.MaxValue = condition["maxValue"].(string)
		p.SamplingUnit = condition["samplingUnit"].(string)
		p.SamplingValue = condition["samplingValue"].(string)

		ps.Points = append(ps.Points, p)
		if p.SamplingUnit == "1" || p.SamplingValue == "minutes" {
			ps.OneMinutePoints = append(ps.OneMinutePoints, p)
		} else if p.SamplingUnit == "5" || p.SamplingValue == "minutes" {
			ps.FiveMinutePoints = append(ps.FiveMinutePoints, p)
		} else if p.SamplingUnit == "10" || p.SamplingValue == "minutes" {
			ps.TenMinutePoints = append(ps.TenMinutePoints, p)
		} else if p.SamplingUnit == "1" || p.SamplingValue == "hours" {
			ps.OneHourPoints = append(ps.OneHourPoints, p)
		} else if p.SamplingUnit == "1" || p.SamplingValue == "days" {
			ps.OneDayPoints = append(ps.OneDayPoints, p)
		} else {
			return &ps, errors.New("配置文件数据有误")
		}
	}
	return &ps, nil
}

func (ps *PlusPointsConfigSlice) OneMinute(nowTime time.Time) {
	if len(ps.OneMinutePoints) <= 0 {
		return
	}
	beginTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), nowTime.Hour(), nowTime.Minute(), 0, 0, utils.TimeLoc())
	endTime := beginTime.Add(-time.Minute)
	for _, v := range ps.OneMinutePoints {
		result := kdb.QueryKdb(v.Point, sqldata.GetSqlDataInstance().CodeSlice, v.Aggr, beginTime, endTime, "end", v.MinValue, v.MaxValue, v.SamplingValue, v.SamplingUnit)
		kdb.PushKdb(v.CalcPoint, result)
	}
}
func (ps *PlusPointsConfigSlice) FiveMinutes(nowTime time.Time) {
	if len(ps.FiveMinutePoints) <= 0 {
		return
	}
	beginTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), nowTime.Hour(), nowTime.Minute(), 0, 0, utils.TimeLoc())
	endTime := beginTime.Add(-time.Minute * 5)
	for _, v := range ps.FiveMinutePoints {
		result := kdb.QueryKdb(v.Point, sqldata.GetSqlDataInstance().CodeSlice, v.Aggr, beginTime, endTime, "end", v.MinValue, v.MaxValue, v.SamplingValue, v.SamplingUnit)
		kdb.PushKdb(v.CalcPoint, result)
	}
}
func (ps *PlusPointsConfigSlice) TenMinutes(nowTime time.Time) {
	if len(ps.TenMinutePoints) <= 0 {
		return
	}
	beginTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), nowTime.Hour(), nowTime.Minute(), 0, 0, utils.TimeLoc())
	endTime := beginTime.Add(-time.Minute * 10)
	for _, v := range ps.TenMinutePoints {
		result := kdb.QueryKdb(v.Point, sqldata.GetSqlDataInstance().CodeSlice, v.Aggr, beginTime, endTime, "end", v.MinValue, v.MaxValue, v.SamplingValue, v.SamplingUnit)
		kdb.PushKdb(v.CalcPoint, result)
	}
}
func (ps *PlusPointsConfigSlice) OneHour(nowTime time.Time) {
	if len(ps.OneHourPoints) <= 0 {
		return
	}
	beginTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), nowTime.Hour(), nowTime.Minute(), 0, 0, utils.TimeLoc())
	endTime := beginTime.Add(-time.Hour)
	for _, v := range ps.OneHourPoints {
		result := kdb.QueryKdb(v.Point, sqldata.GetSqlDataInstance().CodeSlice, v.Aggr, beginTime, endTime, "end", v.MinValue, v.MaxValue, v.SamplingValue, v.SamplingUnit)
		kdb.PushKdb(v.CalcPoint, result)
	}
}
func (ps *PlusPointsConfigSlice) OneDay(nowTime time.Time) {
	if len(ps.OneDayPoints) <= 0 {
		return
	}
	beginTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), nowTime.Hour(), nowTime.Minute(), 0, 0, utils.TimeLoc())
	endTime := beginTime.Add(-time.Hour * 24)
	for _, v := range ps.OneDayPoints {
		result := kdb.QueryKdb(v.Point, sqldata.GetSqlDataInstance().CodeSlice, v.Aggr, beginTime, endTime, "end", v.MinValue, v.MaxValue, v.SamplingValue, v.SamplingUnit)
		kdb.PushKdb(v.CalcPoint, result)
	}
}
