package strategy

import (
	tracer "advanced-algorithms/otel"
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/attribute"
)

type SortStrategy interface {
	Sort([]int) ([]int, int, int)
	Name() string
}

type Sorter struct {
	Strategy SortStrategy
	name string
}

func NewSorter(strategy SortStrategy) *Sorter {
    return &Sorter{Strategy: strategy, name: strategy.Name()}
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
	s.Strategy = strategy
	s.name = strategy.Name()
}

func (s *Sorter) ExecuteSort(arr []int) float64 {

    ctx := context.Background()
    tr := tracer.Tracer()

    ctx, rootSpan := tr.Start(ctx, fmt.Sprintf("%s_Execution", s.name))
    defer rootSpan.End()

    rootSpan.SetAttributes(
        attribute.String("algorithm", s.name),
        attribute.Int("input_size", len(arr)),
        attribute.String("start_time", time.Now().Format(time.RFC3339)),
    )

    ctx, initSpan := tr.Start(ctx, "Initialization")
    numbersCopy := make([]int, len(arr))
    copy(numbersCopy, arr)
    initSpan.End()

    ctx, sortSpan := tr.Start(ctx, "Sorting")
    start := time.Now()

    sortedArray, comparisons, swaps := s.Strategy.Sort(arr)

    duration := time.Since(start).Seconds() * 1000

    sortSpan.SetAttributes(
        attribute.Int("final_comparisons", comparisons),
        attribute.Int("final_swaps", swaps),
        attribute.Float64("duration_ms", duration),
    )
    sortSpan.End()

    _, validateSpan := tr.Start(ctx, "Validation")
    isValid := true
    for i := 1; i < len(sortedArray); i++ {
        if sortedArray[i] < sortedArray[i-1] {
            isValid = false
            break
        }
    }
    validateSpan.SetAttributes(
        attribute.Bool("is_sorted", isValid),
    )
    validateSpan.End()

    rootSpan.SetAttributes(
        attribute.Bool("success", isValid),
        attribute.Float64("total_duration_ms", duration),
        attribute.Int("total_operations", comparisons+swaps),
    )

    fmt.Printf("\n=== %s ===\n", s.name)
    fmt.Printf("Array Size: %d\n", len(arr))
    fmt.Printf("Duration: %.3fms\n", duration)
    fmt.Printf("Comparisons: %d\n", comparisons)
    fmt.Printf("Swaps: %d\n", swaps)
    fmt.Printf("Sorted: %v\n", isValid)
    fmt.Printf("============\n\n")

    return duration
}