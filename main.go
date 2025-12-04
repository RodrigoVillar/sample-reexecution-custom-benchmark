package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/ava-labs/avalanchego/utils/perms"
)

var benchmarkOutputFileArg string

func init() {
	flag.StringVar(&benchmarkOutputFileArg, "benchmark-output-file", "benchmark_output.json", "The filepath where benchmark results should be written to.")
	flag.Parse()
}

func main() {
	benchmarkTool := newBenchmarkTool("BenchmarkReexecuteRange/[1,50000]-Config-default-Runner-dev")

	benchmarkTool.addResult(1.3, "block_accept_ms/ggas")
	benchmarkTool.addResult(2.4, "block_parse_ms/ggas")
	benchmarkTool.addResult(3.5, "block_verify_ms/ggas")
	benchmarkTool.addResult(4.0, "mgas/s")

	benchmarkTool.saveToFile(benchmarkOutputFileArg)

}

// benchmarkResult represents a single benchmark measurement.
type benchmarkResult struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

// benchmarkTool collects and manages benchmark results for a named benchmark.
// It allows adding multiple results and saving them to a file in JSON format.
type benchmarkTool struct {
	name    string
	results []benchmarkResult
}

// newBenchmarkTool creates a new benchmarkTool instance with the given name.
// The name is used as the base name for all results collected by this tool.
// When results are added, the unit is appended to this base name.
func newBenchmarkTool(name string) *benchmarkTool {
	return &benchmarkTool{
		name:    name,
		results: make([]benchmarkResult, 0),
	}
}

// addResult adds a new benchmark result with the given value and unit.
// The result name is constructed by appending the unit to the benchmark name
// Calling `addResult` is analogous to calling `b.ReportMetric()`.
func (b *benchmarkTool) addResult(value float64, unit string) {
	result := benchmarkResult{
		Name:  fmt.Sprintf("%s - %s", b.name, unit),
		Value: strconv.FormatFloat(value, 'f', -1, 64),
		Unit:  unit,
	}
	b.results = append(b.results, result)
}

// saveToFile writes all collected benchmark results to a JSON file at the
// specified path. The output is formatted with indentation for readability.
// Returns an error if marshaling or file writing fails.
func (b *benchmarkTool) saveToFile(path string) error {
	output, err := json.MarshalIndent(b.results, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, output, perms.ReadWrite)
}
