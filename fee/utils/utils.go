package utils

import "time"

// CalYearDuration 计算两个时间之间差了多少年
func CalYearDuration(startTime, endTime time.Time) float64 {
	yearDuration := endTime.Year() - startTime.Year()
	var year float64
	year = float64(yearDuration)
	year += float64((endTime.Month()+12-startTime.Month())%12) / 12
	return year
}
