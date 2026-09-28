package main

import (
	"flag"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// Record a new measurement with go generate.
//go:generate go run .

var threshold = flag.Float64("threshold", 0.1, "tolerated relative increase of ns/op in TestNoDegradation")

// TestNoDegradation compares the last measurement in the results file with the previous
// ones recorded on the same machine. It is skipped when there are no such measurements.
func TestNoDegradation(t *testing.T) {
	root, err := moduleRoot()
	if err != nil {
		t.Fatal(err)
	}
	ms, err := load(filepath.Join(root, resultsFile))
	if err != nil {
		t.Fatal(err)
	}
	if len(ms) == 0 {
		t.Skipf("no benchmark results in %s, record them with go run ./internal/cmd/benchmarks", resultsFile)
	}
	last := ms[len(ms)-1]
	base := slices.DeleteFunc(ms[:len(ms)-1], func(m Measurement) bool { return !m.sameMachine(last) })
	if len(base) == 0 {
		t.Skipf("no previous benchmark results recorded on the machine of the last measurement (%s)", last.CPU)
	}
	t.Logf("comparing the measurement of %s with %d previous ones", last.Time, len(base))
	for _, d := range compare(base, []Measurement{last}, *threshold) {
		t.Error(d)
	}
}

func TestParseOutput(t *testing.T) {
	const out = `goos: darwin
goarch: arm64
pkg: rmazur.io/chernetka/internal/editor/extlsp
cpu: Apple M1 Pro
BenchmarkDiff/800lines/start-10         	  100000	     10534 ns/op	   58880 B/op	       2 allocs/op
--- BENCH: BenchmarkDiff/800lines/start-10
    diff_test.go:120: other log
    diff_test.go:121: tolerance=20%
BenchmarkDiff/800lines/end-10         	  100000	     10534 ns/op	   58880 B/op	       2 allocs/op
--- BENCH: BenchmarkDiff/800lines/end-10
    diff_test.go:121: tolerance=12.5% is not the whole message
PASS
ok  	rmazur.io/chernetka/internal/editor/extlsp	1.234s
goos: darwin
goarch: arm64
pkg: rmazur.io/chernetka/internal/vt/escape
cpu: Apple M1 Pro
BenchmarkStyleText-10    	 5000000	       230.5 ns/op	       1.500 custom/op	      48 B/op	       1 allocs/op
BenchmarkNoProcs 	 10	 7 ns/op
Benchmark log line: not a result
PASS
ok  	rmazur.io/chernetka/internal/vt/escape	2.345s
`
	var got Measurement
	if err := parseOutput(strings.NewReader(out), &got); err != nil {
		t.Fatal(err)
	}
	want := Measurement{
		GOOS:   "darwin",
		GOARCH: "arm64",
		CPU:    "Apple M1 Pro",
		Results: []Result{
			{
				Pkg: "rmazur.io/chernetka/internal/editor/extlsp", Name: "BenchmarkDiff/800lines/start", Procs: 10, Iterations: 100000,
				Metrics: map[string]float64{"ns/op": 10534, "B/op": 58880, "allocs/op": 2}, Tolerance: 0.2,
			},
			{
				Pkg: "rmazur.io/chernetka/internal/editor/extlsp", Name: "BenchmarkDiff/800lines/end", Procs: 10, Iterations: 100000,
				Metrics: map[string]float64{"ns/op": 10534, "B/op": 58880, "allocs/op": 2},
			},
			{
				Pkg: "rmazur.io/chernetka/internal/vt/escape", Name: "BenchmarkStyleText", Procs: 10, Iterations: 5000000,
				Metrics: map[string]float64{"ns/op": 230.5, "custom/op": 1.5, "B/op": 48, "allocs/op": 1},
			},
			{
				Pkg: "rmazur.io/chernetka/internal/vt/escape", Name: "BenchmarkNoProcs", Procs: 1, Iterations: 10,
				Metrics: map[string]float64{"ns/op": 7},
			},
		},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("parseOutput() mismatch (-want +got):\n%s", diff)
	}
}

func TestMeasurementsFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), resultsFile)
	ms, err := load(path)
	if err != nil {
		t.Fatal(err)
	}
	start := time.Date(2026, 9, 28, 0, 0, 0, 0, time.UTC)
	for i := range 5 {
		ms = add(ms, Measurement{Time: start.Add(time.Duration(i) * time.Hour)})
		if err := save(path, ms); err != nil {
			t.Fatal(err)
		}
		if ms, err = load(path); err != nil {
			t.Fatal(err)
		}
	}
	want := []Measurement{
		{Time: start.Add(2 * time.Hour)},
		{Time: start.Add(3 * time.Hour)},
		{Time: start.Add(4 * time.Hour)},
	}
	if diff := cmp.Diff(want, ms); diff != "" {
		t.Errorf("load() mismatch (-want +got):\n%s", diff)
	}
}

func TestCompare(t *testing.T) {
	measurement := func(ns, allocs float64) Measurement {
		return Measurement{Results: []Result{
			{Pkg: "p", Name: "BenchmarkA", Procs: 8, Metrics: map[string]float64{"ns/op": ns, "allocs/op": allocs}},
		}}
	}
	tolerant := func(m Measurement, tolerance float64) Measurement {
		m.Results[0].Tolerance = tolerance
		return m
	}
	extra := measurement(1, 0)
	extra.Results[0].Name = "BenchmarkNew"

	for _, tc := range []struct {
		name      string
		base, cur []Measurement
		want      []Degradation
	}{
		{
			name: "noise within the threshold",
			base: []Measurement{measurement(100, 2), measurement(120, 2)},
			cur:  []Measurement{measurement(109, 2), measurement(150, 2)},
		},
		{
			name: "best of the measurements is compared",
			base: []Measurement{measurement(100, 2)},
			cur:  []Measurement{measurement(200, 2), measurement(105, 2), measurement(130, 2)},
		},
		{
			name: "faster",
			base: []Measurement{measurement(100, 2)},
			cur:  []Measurement{measurement(50, 1)},
		},
		{
			name: "slower",
			base: []Measurement{measurement(100, 2), measurement(90, 2)},
			cur:  []Measurement{measurement(120, 2), measurement(130, 2)},
			want: []Degradation{{Key: "p.BenchmarkA-8", Unit: "ns/op", Base: 90, Cur: 120}},
		},
		{
			name: "within the benchmark tolerance",
			base: []Measurement{measurement(100, 2)},
			cur:  []Measurement{tolerant(measurement(120, 2), 0.25)},
		},
		{
			name: "beyond the benchmark tolerance",
			base: []Measurement{measurement(100, 2)},
			cur:  []Measurement{tolerant(measurement(104, 2), 0.03)},
			want: []Degradation{{Key: "p.BenchmarkA-8", Unit: "ns/op", Base: 100, Cur: 104}},
		},
		{
			name: "more allocations",
			base: []Measurement{measurement(100, 2)},
			cur:  []Measurement{measurement(100, 3)},
			want: []Degradation{{Key: "p.BenchmarkA-8", Unit: "allocs/op", Base: 2, Cur: 3}},
		},
		{
			name: "new benchmark",
			base: []Measurement{measurement(100, 2)},
			cur:  []Measurement{measurement(100, 2), extra},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := compare(tc.base, tc.cur, 0.1)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("compare() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
