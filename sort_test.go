package sleepsort

import (
	"reflect"
	"testing"
)

func TestSleepSort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test...")
	}
	var (
		input    = report(t, "Input:   ", 3, 1, 4, 1, 5, 9)
		expected = report(t, "Expected:", 1, 1, 3, 4, 5, 9)
		actual   = report(t, "Actual:  ", NewSorter(input...).Sorted()...)
	)
	if !reflect.DeepEqual(expected, actual) {
		t.Error()
	}
}
func report(t *testing.T, label string, values ...int) []int {
	t.Logf(label+" %+v", values)
	return values
}
