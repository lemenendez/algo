package main

import (
	"fmt"
	"math"
	"sync"
)

// Sigmoid activation function
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

type Neuron struct {
	Weights []float64
	Bias    float64
}

type Brain struct {
	mu     sync.RWMutex
	Energy uint64
	Gender int
	Age    int
	Cells  []Neuron
}

func (l *Brain) FeedForward(inputs []float64) []float64 {
	outputs := make([]float64, len(l.Cells))
	for i, neuron := range l.Cells {
		outputs[i] = neuron.FeedForward(inputs)
	}
	return outputs
}

// Calculate output from inputs
func (n *Neuron) FeedForward(inputs []float64) float64 {
	if len(inputs) != len(n.Weights) {
		panic("Number of inputs must match number of weights")
	}

	var total float64
	for i, input := range inputs {
		total += input * n.Weights[i]
	}
	total += n.Bias
	return sigmoid(total)
}

func main() {
	// Example: A Neuron with 2 inputs
	layer := Brain{
		Cells: []Neuron{
			{Weights: []float64{0.5, -0.6}, Bias: 0.1},
			{Weights: []float64{0.3, 0.8}, Bias: -0.2},
			{Weights: []float64{-0.7, 0.5}, Bias: 0.3},
		},
	}

	// Simulating an input
	inputs := []float64{1.0, 0.0} // Example inputs
	outputs := layer.FeedForward(inputs)

	fmt.Println("Layer outputs:")
	for i, output := range outputs {
		fmt.Printf("Neuron %d output: %f\n", i, output)
	}
}
