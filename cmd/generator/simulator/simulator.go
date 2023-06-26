package simulator

import (
	"Hackathon2023/cmd/generator/producer"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type Simulator struct {
	Name          string
	Speed         float64
	SpeedKmPerSec float64
	Vectors       [][]float64
	Pace          float64
}

func getFile(name string) (*os.File, error) {
	_, filename, _, _ := runtime.Caller(0)
	rootIx := strings.Index(filename, "/cmd/") // a little bit hacky, but will work for any test in cmd repo
	return os.Open(path.Join(filename[:rootIx], "configs", name))
}

func NewSimulator(fileName string) *Simulator {
	file, err := getFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}
	res := &Simulator{}
	err = json.Unmarshal(bytes, &res)
	res.SpeedKmPerSec = res.Speed / 60 / 60
	if err != nil {
		panic(err.Error())
	}
	return res
}

func (s *Simulator) BuildLineEquation(point1, point2 []float64) (float64, float64, error) {
	// Calculate the slope (m)
	if point1[0] == point2[0] {
		return 0, 0, fmt.Errorf("cannot build a line equation, x-coordinates are the same")
	}
	slope := (point2[1] - point1[1]) / (point2[0] - point1[0])

	// Calculate the y-intercept (b)
	yIntercept := point1[1] - slope*point1[0]

	// Return the line equation as a string
	return slope, yIntercept, nil
}

func (s *Simulator) Simulate() {
	for i := 0; i < len(s.Vectors)-1; i++ {
		from := s.Vectors[i]
		to := s.Vectors[i+1]
		fmt.Printf("from: %v, to: %v\n", from, to)
		slope, intercept, _ := s.BuildLineEquation(from, to)
		s.move(from, to, slope, intercept)
	}
	fmt.Println(s.Vectors)
}

func (s *Simulator) move(from []float64, to []float64, slope float64, intercept float64) {
	currentPoint := from
	for math.RoundToEven(currentPoint[0]) != to[0] && math.RoundToEven(currentPoint[1]) != to[1] {
		err := s.updateSoundPoints(currentPoint)
		if err != nil {
			return
		}
		nextPointX := currentPoint[0] + s.SpeedKmPerSec*s.Pace/1000
		if from[0] > to[0] {
			nextPointX = currentPoint[0] - s.SpeedKmPerSec*s.Pace/1000
		}
		if math.RoundToEven(nextPointX-to[0]) == 0 {
			nextPointX = to[0]
		}
		nextPointY := slope*nextPointX + intercept
		if math.RoundToEven(nextPointY-to[1]) == 0 {
			nextPointY = to[1]
		}
		currentPoint = []float64{nextPointX, nextPointY}
		time.Sleep(time.Millisecond * time.Duration(s.Pace))
	}
}

func (s *Simulator) updateSoundPoints(point []float64) error {
	nearestPoint := fmt.Sprintf("%v:%v", math.RoundToEven(point[0]), math.RoundToEven(point[1]))
	sp, ok := SoundPointsNetwork[nearestPoint]
	if !ok {
		fmt.Println("didn't find point")
		return fmt.Errorf("didn't find point")
	}
	s.sendMessage(sp, 1000)
	fmt.Println(fmt.Sprintf(`{"id": "%v", "point": [%v,%v], "level": %v}`, sp.Id, sp.Point[0], sp.Point[1], 1000))
	s.updateClosestSoundPoints(sp)
	return nil
}
func (s *Simulator) sendMessage(sp *SoundPoint, level int) {
	producer.ProduceMessage("localhost:9094", sp.Id, fmt.Sprintf(`{"id": "%v", "point": [%v,%v], "level": %v}`, sp.Id, sp.Point[0], sp.Point[1], level))
}
func (s *Simulator) updateClosestSoundPoints(sp *SoundPoint) {
	s.updateCoordinates(sp, 0, 1, 700)
	s.updateCoordinates(sp, 0, -1, 700)
	s.updateCoordinates(sp, 0, 2, 400)
	s.updateCoordinates(sp, 0, -2, 400)
	s.updateCoordinates(sp, -1, 0, 700)
	s.updateCoordinates(sp, -1, 1, 600)
	s.updateCoordinates(sp, -1, -1, 600)
	s.updateCoordinates(sp, -1, 2, 300)
	s.updateCoordinates(sp, -1, -2, 300)
	s.updateCoordinates(sp, 1, 1, 600)
	s.updateCoordinates(sp, 1, 0, 700)
	s.updateCoordinates(sp, 1, -1, 600)
	s.updateCoordinates(sp, 1, 2, 300)
	s.updateCoordinates(sp, -2, -2, 200)
	s.updateCoordinates(sp, -2, 1, 300)
	s.updateCoordinates(sp, -2, 0, 400)
	s.updateCoordinates(sp, -2, -1, 300)
	s.updateCoordinates(sp, -2, 2, 200)
	s.updateCoordinates(sp, 2, -2, 200)
	s.updateCoordinates(sp, 2, 1, 300)
	s.updateCoordinates(sp, 2, 0, 400)
	s.updateCoordinates(sp, 2, -1, 300)
	s.updateCoordinates(sp, 2, 2, 200)
}

func (s *Simulator) updateCoordinates(sp *SoundPoint, x int, y int, level int) {
	nextSp, ok := SoundPointsNetwork[fmt.Sprintf("%v:%v", math.RoundToEven(sp.Point[0]+float64(x)), math.RoundToEven(sp.Point[1]+float64(y)))]
	if ok {
		s.sendMessage(nextSp, level)
	}
}
