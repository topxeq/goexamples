package main

import (
	"fmt"

	tg "github.com/galeone/tfgo"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func main() {

	model := tg.LoadModel("/root/py/export", []string{"tag"}, nil)

	inputArray := [][]float32{{1, 2, 3, 4}}

	fmt.Printf("input: %v\n", inputArray)

	fakeInput, _ := tf.NewTensor(inputArray)
	results := model.Exec([]tf.Output{
		model.Op("y", 0),
	}, map[tf.Output]*tf.Tensor{
		model.Op("x", 0): fakeInput,
	})

	predictions := results[0].Value().([]float32)
	fmt.Printf("predict: %v\n", predictions)

}
