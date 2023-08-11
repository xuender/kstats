package kstats

import "golang.org/x/exp/constraints"

// Outliers holds mild and extreme outliers found in data.
type Outliers[N Number] struct {
	Mild    []N
	Extreme []N
}

// QuartileOutliers returns the mild and extreme outliers.
func QuartileOutliers[F constraints.Float](floats ...F) Outliers[F] {
	if len(floats) == 0 {
		return Outliers[F]{}
	}

	var (
		opf     F = 1.5
		clone     = SortedClone(floats)
		qua       = Quartile(clone...)
		iqr       = InterQuartileRange(clone...)
		lif       = qua.Q1 - (opf * iqr)
		uif       = qua.Q3 + (opf * iqr)
		lof       = qua.Q1 - (_three * iqr)
		uof       = qua.Q3 + (_three * iqr)
		mild    []F
		extreme []F
	)

	for _, v := range clone {
		if v < lof || v > uof {
			extreme = append(extreme, v)
		} else if v < lif || v > uif {
			mild = append(mild, v)
		}
	}

	return Outliers[F]{mild, extreme}
}
