package main

import (
	"fmt"
	"math"
)

// Sigmoid activation function
// The sigmoid function serves as the activation function, converting the neuron's output to a range between 0 and 1.
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

type Neuron struct {
	// Weights  are coefficients that are applied to the inputs of a neuron.
	// Each input to a neuron has an associated weight.
	// The purpose of weights is to control the influence of each input on the neuron's output.
	// A higher weight means that the corresponding input has a more significant impact on the neuron's activation.
	// During the training of a neural network,
	// weights are adjusted based on the learning algorithm (like gradient descent) to minimize the error
	// between the predicted output and the actual output. This adjustment process is often referred to as learning.
	Weights []float64
	// Bias is an additional parameter in a neuron that allows the model to fit the data better.
	// Is added to the weighted sum of inputs before being passed through the activation function.
	// The bias allows the activation function to be shifted to the left or right, which can be crucial for learning.
	// It enables the model to fit the data better by providing another degree of freedom (along with weights).
	// Unlike weights, which depend on the input values, the bias is a constant term.
	// It helps the neuron make predictions even when all inputs are zero.
	Bias float64
}

type Layer struct {
	Cells []Neuron
}

func (l *Layer) FeedForward(inputs []float64) []float64 {
	outputs := make([]float64, len(l.Cells))
	for i, neuron := range l.Cells {
		outputs[i] = neuron.FeedForward(inputs)
	}
	return outputs
}

// FeedForward Calculate output from inputs
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

// Input Processing: Each input is multiplied by its corresponding weight.
// Weighted Sum: The products are summed together.
// Bias Addition: The bias is then added to the weighted sum.
// Activation Function: The result is passed through an activation function (like sigmoid, ReLU, etc.) to produce the neuron's output.
func main() {
	// Example: A Neuron with 2 inputs
	layer := Layer{
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
