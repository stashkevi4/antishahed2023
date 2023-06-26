package main

import (
	"Hackathon2023/cmd/generator/simulator"
	"path"
)

func main() {
	//producer.ProduceMessage("localhost:9094", "1", `{"name":"hello kafka"}`)
	simulator.InitNetwork([]float64{0, 0}, []float64{100, 100}, 1)
	s := simulator.NewSimulator(path.Join("shahed-vectors.json"))
	for {
		s.Simulate()
	}
}
