package main

import (
	"context"
	"fmt"
	"sync"
)

type resultStruct struct {
	index int
	value int
}

func ParallelMapCtx(ctx context.Context, inputs []int, fn func(int) int, workers int) ([]int, error) {
	var wg sync.WaitGroup

	jobsChan := make(chan int, len(inputs))
	resultChan := make(chan resultStruct, len(inputs))
	result := make([]int, len(inputs))

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for index := range jobsChan {
				newResult := &resultStruct{index: index, value: fn(inputs[index])}
				resultChan <- *newResult
			}
			wg.Done()
		}()
	}

	for i := 0; i < len(inputs); i++ {
		jobsChan <- i
	}

	close(jobsChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for i := 0; i < len(inputs); i++ {
		select {
		case <-ctx.Done():
			return []int{}, ctx.Err()
		case res := <-resultChan:
			result[res.index] = res.value
		}
	}

	return result, nil

}

func main() {
	fmt.Println(ParallelMapCtx(context.Background(), []int{1, 2, 3, 4}, func(i int) int { return i * i }, 1))
}
