package stats

import (
	"math"
	"sync"
)

type SummaryInfo struct {
	Min       float64   // 最小数
	Max       float64   // 最大数
	Median    float64   // 中位数, Q2
	Mean      float64   // 平均数
	Quartiles Quartiles // 四分位数
	Mode      []float64 // 众数

	PTP      float64 // 极差
	Variance float64 // 方差
	STD      float64 // 标准差
	CV       float64 // 变异系数

	// 当标准差不为0且不为较接近于0的数时，z-分数是有意义的
	// 通常来说，z-分数的绝对值大于3将视为异常。
	ZScores []float64 // 偏差程度（z-分数）
}

func Summary(input Float64Data) (ret SummaryInfo) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if v, err := Min(input); err == nil {
			ret.Min = v
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if v, err := Max(input); err == nil {
			ret.Max = v
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if v, err := Median(input); err == nil {
			ret.Median = v
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if v, err := Mean(input); err == nil {
			ret.Mean = v
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if v, err := Quartile(input); err == nil {
			ret.Quartiles = v
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if v, err := Mode(input); err == nil {
			ret.Mode = v
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if v, err := Variance(input); err == nil {
			ret.Variance = v
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if v, err := ZScore(input); err == nil {
			ret.ZScores = v
		}
	}()

	wg.Wait()

	ret.PTP = ret.Max - ret.Min
	ret.STD = math.Pow(ret.Variance, 0.5)
	if ret.Mean != 0 {
		ret.CV = ret.STD / ret.Mean
	}

	return
}
