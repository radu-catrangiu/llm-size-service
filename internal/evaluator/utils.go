package evaluator

import (
	"fmt"
	"math"
)

type unit struct {
	Bytes  uint64
	Suffix string
}

var units = []unit{
	{1 << 30, "Gi"},
	{1 << 20, "Mi"},
	// {1 << 10, "Ki"},
}

func convertBytesToK8sSize(sizeInBytes uint64, rounded bool) string {
	var unit unit
	for _, u := range units {
		if sizeInBytes >= unit.Bytes {
			unit = u
			break
		}
	}

	var result string
	floatSize := float64(sizeInBytes) / float64(unit.Bytes)

	if rounded {
		size := uint64(math.Ceil(floatSize))
		result = fmt.Sprintf("%d%s", size, unit.Suffix)
	} else {
		result = fmt.Sprintf("%f%s", floatSize, unit.Suffix)
	}

	return result
}
