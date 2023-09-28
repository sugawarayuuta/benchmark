package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type entry struct {
	cat  string
	impl string
	avg  float64
}

func main() {
	in := flag.String("in", "./result/result.txt", "benchmark results to read")
	out := flag.String("out", "./README.md", "markdown file to generate.")
	flag.Parse()

	fl, err := os.Open(*in)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fl.Close()

	results := make(map[string][]float64)
	buf := bufio.NewScanner(fl)
	for buf.Scan() {
		text := buf.Text()
		text, ok := strings.CutPrefix(text, "Benchmark")
		if !ok {
			continue
		}
		spl := strings.Split(text, "\t")
		mbs, ok := strings.CutSuffix(spl[3], " MB/s")
		if !ok {
			continue
		}
		num, err := strconv.ParseFloat(strings.TrimSpace(mbs), 64)
		if err != nil {
			continue
		}
		results[spl[0]] = append(results[spl[0]], num)
	}
	ents := make([]entry, 0, len(results))
	for name, result := range results {
		spl := strings.SplitN(name, "/", 3)
		if len(spl) != 3 {
			continue
		}
		var sum float64
		for _, f64 := range result {
			sum += f64
		}
		ents = append(ents, entry{
			cat:  fmt.Sprintf("%s - %s", spl[0], spl[1]),
			impl: spl[2],
			avg:  sum / float64(len(result)),
		})
	}
	slices.SortFunc(ents, func(fst, sec entry) int {
		cmp := strings.Compare(fst.cat, sec.cat)
		if cmp != 0 {
			return cmp
		}
		return strings.Compare(fst.impl, sec.impl)
	})

	fl, err = os.Create(*out)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fl.Close()

	for idx, ent := range ents {
		if idx <= 0 || ents[idx-1].cat != ent.cat {
			fmt.Fprintln(fl, "```mermaid")
			fmt.Fprintln(fl, "gantt")
			fmt.Fprintln(fl, "title", ent.cat, "(MB/s - higher is better)")
			fmt.Fprintln(fl, "dateFormat", "X")
			fmt.Fprintln(fl, "axisFormat", "%s")
			fmt.Fprintln(fl)
		}
		fmt.Fprintln(fl, "section", ent.impl)
		fmt.Fprintln(fl, int(ent.avg), ":0,", int(ent.avg))
		if idx+1 >= len(ents) || ents[idx+1].cat != ent.cat {
			fmt.Fprintln(fl, "```")
			fmt.Fprintln(fl)
		}
	}
}
