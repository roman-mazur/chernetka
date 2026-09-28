package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	// resultsFile is the path of the results file relative to the module root.
	resultsFile = ".benchmarks/results.json"
	// maxMeasurements is the number of the last measurements kept in the results file.
	maxMeasurements = 3
)

// Measurement is the results of one run of the benchmarks.
type Measurement struct {
	Time      time.Time `json:"time"`
	GoVersion string    `json:"go"`
	GOOS      string    `json:"goos"`
	GOARCH    string    `json:"goarch"`
	CPU       string    `json:"cpu"`
	Results   []Result  `json:"results"`
}

// Result is one line of the benchmark output.
type Result struct {
	Pkg        string `json:"pkg"`
	Name       string `json:"name"`
	Procs      int    `json:"procs"`
	Iterations int64  `json:"iterations"`
	// Metrics maps the units (ns/op, B/op, allocs/op, or custom ones) to their values.
	Metrics map[string]float64 `json:"metrics"`
	// Tolerance overrides the tolerated relative increase of ns/op for the benchmark.
	// It is set with a "tolerance=N%" log line in the benchmark.
	Tolerance float64 `json:"tolerance,omitempty"`
}

// Key identifies the benchmark across the measurements.
func (r Result) Key() string {
	return fmt.Sprintf("%s.%s-%d", r.Pkg, r.Name, r.Procs)
}

// sameMachine reports whether the measurements can be compared.
func (m Measurement) sameMachine(other Measurement) bool {
	return m.GOOS == other.GOOS && m.GOARCH == other.GOARCH && m.CPU == other.CPU
}

// add appends m to the measurements, dropping the oldest ones beyond maxMeasurements.
func add(ms []Measurement, m Measurement) []Measurement {
	ms = append(ms, m)
	if extra := len(ms) - maxMeasurements; extra > 0 {
		ms = ms[extra:]
	}
	return ms
}

// load reads the measurements from the results file. It returns none if the file does not exist.
func load(path string) ([]Measurement, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var ms []Measurement
	if err := json.Unmarshal(data, &ms); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return ms, nil
}

func save(path string, ms []Measurement) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(ms, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0o644)
}

// toleranceLog matches a log line of a benchmark like "    bench_test.go:12: tolerance=20%".
var toleranceLog = regexp.MustCompile(`^\s+\S+:\d+: tolerance=(\d+(?:\.\d+)?)%$`)

// parseOutput adds the results from the go test -bench output to m.
func parseOutput(r io.Reader, m *Measurement) error {
	var (
		pkg string
		// logged is the result whose log lines follow.
		logged *Result
	)
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		if name, ok := strings.CutPrefix(line, "--- BENCH: "); ok {
			logged = nil
			if n := len(m.Results); n > 0 && fmt.Sprintf("%s-%d", m.Results[n-1].Name, m.Results[n-1].Procs) == name {
				logged = &m.Results[n-1]
			}
			continue
		}
		if logged != nil {
			if sub := toleranceLog.FindStringSubmatch(line); sub != nil {
				percent, err := strconv.ParseFloat(sub[1], 64)
				if err != nil {
					return err
				}
				logged.Tolerance = percent / 100
				continue
			}
			if strings.HasPrefix(line, " ") {
				continue
			}
			logged = nil
		}
		if k, v, ok := strings.Cut(line, ": "); ok {
			switch k {
			case "goos":
				m.GOOS = v
				continue
			case "goarch":
				m.GOARCH = v
				continue
			case "cpu":
				m.CPU = v
				continue
			case "pkg":
				pkg = v
				continue
			}
		}
		if res, ok := parseResult(line); ok {
			res.Pkg = pkg
			m.Results = append(m.Results, res)
		}
	}
	return sc.Err()
}

// parseResult parses a line like "BenchmarkName-8  1000  1234 ns/op  16 B/op  1 allocs/op".
func parseResult(line string) (Result, bool) {
	fields := strings.Fields(line)
	if len(fields) < 4 || len(fields)%2 != 0 || !strings.HasPrefix(fields[0], "Benchmark") {
		return Result{}, false
	}
	n, err := strconv.ParseInt(fields[1], 10, 64)
	if err != nil {
		return Result{}, false
	}
	res := Result{Name: fields[0], Procs: 1, Iterations: n, Metrics: make(map[string]float64)}
	if i := strings.LastIndexByte(res.Name, '-'); i >= 0 {
		if procs, err := strconv.Atoi(res.Name[i+1:]); err == nil {
			res.Name, res.Procs = res.Name[:i], procs
		}
	}
	for i := 2; i < len(fields); i += 2 {
		v, err := strconv.ParseFloat(fields[i], 64)
		if err != nil {
			return Result{}, false
		}
		res.Metrics[fields[i+1]] = v
	}
	return res, true
}
