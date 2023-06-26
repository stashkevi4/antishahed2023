package simulator

import (
	"fmt"
	"strconv"
)

var SoundPointsNetwork = map[string]*SoundPoint{}

type SoundPoint struct {
	Id     string
	Point  []float64
	Radius float64
}

func InitNetwork(startPoint []float64, endPoint []float64, density float64) {
	counter := 1
	for i := startPoint[0]; i < endPoint[0]; i += density {
		for j := startPoint[1]; j < endPoint[1]; j += density {
			SoundPointsNetwork[fmt.Sprintf("%v:%v", i, j)] = NewSoundPoint(
				strconv.Itoa(counter),
				[]float64{i, j},
				2)
			counter++
		}
	}
}

func NewSoundPoint(id string, point []float64, radius float64) *SoundPoint {
	return &SoundPoint{
		Id:     id,
		Point:  point,
		Radius: radius,
	}
}
