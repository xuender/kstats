package kstats

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

// Holds information about the dataset provided to Describe.
type Description[F constraints.Float] struct {
	Count                  int
	Mean                   F
	Std                    F
	Max                    F
	Min                    F
	DescriptionPercentiles []descriptionPercentile[F]
	AllowedNaN             bool
}

// Specifies percentiles to be computed.
type descriptionPercentile[F constraints.Float] struct {
	Percentile float64
	Value      F
}

// Describe returns descriptive statistics about a provided dataset, similar to python's pandas.describe().
func Describe[F constraints.Float](floats []F, allowNaN bool, percentiles []float64) (*Description[F], error) {
	return DescribePercentileFunc(floats, allowNaN, percentiles, Percentile[F])
}

// Describe returns descriptive statistics about a provided dataset, similar to python's pandas.describe()
// Takes in a function to use for percentile calculation.
func DescribePercentileFunc[F constraints.Float](
	floats []F,
	allowNaN bool,
	percentiles []float64,
	percentileFunc func(float64, ...F) F,
) (*Description[F], error) {
	var description Description[F]
	description.AllowedNaN = allowNaN
	description.Count = len(floats)

	if description.Count == 0 && !allowNaN {
		return &description, ErrEmptyInput
	}

	description.Std = StandardDeviation(floats...)
	description.Max = Max(floats...)
	description.Min = Min(floats...)
	description.Mean = Mean(floats...)

	for _, percentile := range percentiles {
		value := percentileFunc(percentile, floats...)
		description.DescriptionPercentiles = append(
			description.DescriptionPercentiles,
			descriptionPercentile[F]{Percentile: percentile, Value: value},
		)
	}

	return &description, nil
}

func (d *Description[F]) String() string {
	return d.ToString(_two)
}

// ToString returns format with specified number of decimals.
func (d *Description[F]) ToString(decimals int) string {
	buf := &strings.Builder{}

	fmt.Fprintf(buf, "count\t%d\n", d.Count)
	fmt.Fprintf(buf, "count\t%d\n", d.Count)
	fmt.Fprintf(buf, "mean\t%.*f\n", decimals, d.Mean)
	fmt.Fprintf(buf, "std\t%.*f\n", decimals, d.Std)
	fmt.Fprintf(buf, "max\t%.*f\n", decimals, d.Max)
	fmt.Fprintf(buf, "min\t%.*f\n", decimals, d.Min)

	for _, percentile := range d.DescriptionPercentiles {
		fmt.Fprintf(buf, "%.2f%%\t%.*f\n", percentile.Percentile, decimals, percentile.Value)
	}

	fmt.Fprintf(buf, "NaN OK\t%t", d.AllowedNaN)

	return buf.String()
}
