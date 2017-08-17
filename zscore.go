package stats

import (
	"github.com/kataras/go-errors"
	"math"
)

func ZScore(input Float64Data) (zscores []float64, err error) {
	std, err := StandardDeviationPopulation(input)
	if err != nil {
		return nil, err
	}
	if math.Abs(std) <= 0.0001 {
		return nil, errors.New("too small std")
	}
	mean, err := Mean(input)
	if err != nil {
		return nil, err
	}

	for _, v := range input {
		zscores = append(zscores, (v-mean)/std)
	}

	return zscores, nil
}
