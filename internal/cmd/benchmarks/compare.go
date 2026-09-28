package main

import (
	"fmt"
	"maps"
	"slices"
)

// Degradation is a benchmark metric that got worse than tolerated.
type Degradation struct {
	Key       string
	Unit      string
	Base, Cur float64
}

func (d Degradation) String() string {
	return fmt.Sprintf("%s: %s is %.4g, was %.4g (%+.1f%%)", d.Key, d.Unit, d.Cur, d.Base, (d.Cur/d.Base-1)*100)
}

// compare finds the benchmarks that are slower in cur than in base by more than
// the threshold (a relative increase of ns/op), or allocate more often.
// The tolerance of a benchmark in cur overrides the threshold.
// The best value from all the measurements is taken for each metric to reduce the noise.
// Benchmarks missing in one of the measurement sets are not compared.
func compare(base, cur []Measurement, threshold float64) []Degradation {
	baseBest, curBest := best(base), best(cur)
	tolerance := make(map[string]float64)
	for _, m := range cur {
		for _, r := range m.Results {
			if r.Tolerance > 0 {
				tolerance[r.Key()] = r.Tolerance
			}
		}
	}
	var res []Degradation
	for _, key := range slices.Sorted(maps.Keys(curBest)) {
		b, ok := baseBest[key]
		if !ok {
			continue
		}
		c := curBest[key]
		t, ok := tolerance[key]
		if !ok {
			t = threshold
		}
		if bv, cv, ok := metric(b, c, "ns/op"); ok && cv > bv*(1+t) {
			res = append(res, Degradation{Key: key, Unit: "ns/op", Base: bv, Cur: cv})
		}
		if bv, cv, ok := metric(b, c, "allocs/op"); ok && cv > bv {
			res = append(res, Degradation{Key: key, Unit: "allocs/op", Base: bv, Cur: cv})
		}
	}
	return res
}

func metric(base, cur map[string]float64, unit string) (float64, float64, bool) {
	bv, bok := base[unit]
	cv, cok := cur[unit]
	return bv, cv, bok && cok
}

// best returns the minimal value of every metric for every benchmark.
func best(ms []Measurement) map[string]map[string]float64 {
	res := make(map[string]map[string]float64)
	for _, m := range ms {
		for _, r := range m.Results {
			key := r.Key()
			if res[key] == nil {
				res[key] = maps.Clone(r.Metrics)
				continue
			}
			for unit, v := range r.Metrics {
				if prev, ok := res[key][unit]; !ok || v < prev {
					res[key][unit] = v
				}
			}
		}
	}
	return res
}
