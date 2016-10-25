package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

import "os"

import "bufio"

var filename = flag.String("f", "x", "Path to graph file.")
var di = flag.Bool("di", false, "Degree Invariant.")
var ddi = flag.Bool("ddi", false, "Degree Depth Invariant.")
var dd2 = flag.Bool("dd2", false, "Degree Depth Invariant Version 2.")
var dot = flag.Bool("dot", false, "Print to std. in .dot format.")

type graph struct {
	pre  string
	size int
	adj  [][]bool
	deg  []int
}

func newGraph(pre string) (g graph) {
	g.pre = pre
	return g
}

type invariant func(int, graph) string

func degInv(node int, g graph) string {
	return "degree(" + strconv.FormatInt(int64(g.deg[node]), 10) + ")"
}

func degDepth2Inv(node int, g graph) (s string) {

	s += "deg2Depth("
	done := make([]bool, g.size)
	done[node] = true

	nodes := []int{node}

	for len(nodes) > 0 {

		degs := make([]int, len(nodes))

		for i, n := range nodes {
			c := 0
			for _, m := range g.succs(n) {
				if !done[m] {
					c++
				}
			}
			degs[i] = c
		}

		sort.Ints(degs)
		s += format(degs)

		nextNodes := []int{}

		for _, n := range nodes {
			for _, m := range g.succs(n) {
				if !done[m] {
					nextNodes = append(nextNodes, m)
					done[m] = true
				}
			}
		}
		nodes = nextNodes
	}

	s += ")"
	return
}

func degDepthInv(node int, g graph) (s string) {

	s += "degDepth("
	done := make([]bool, g.size)
	done[node] = true

	nodes := []int{node}

	for len(nodes) > 0 {

		degs := make([]int, len(nodes))

		for i, n := range nodes {
			degs[i] = g.deg[n]
		}

		sort.Ints(degs)
		s += format(degs)

		nextNodes := []int{}

		for _, n := range nodes {
			for _, m := range g.succs(n) {
				if !done[m] {
					nextNodes = append(nextNodes, m)
					done[m] = true
				}
			}
		}
		nodes = nextNodes
	}

	s += ")"
	return
}

func format(ints []int) (s string) {
	s += "["
	for i, n := range ints {
		if i != 0 {
			s += ","
		}
		s += strconv.FormatInt(int64(n), 10)
	}
	return s + "]"
}

func (g *graph) succs(i int) (s []int) {
	for j := 0; j < g.size; j++ {
		if g.adj[i][j] {
			s = append(s, j)
		}
	}
	return
}

func main() {
	flag.Parse()
	input, err2 := os.Open(*filename)
	defer input.Close()
	if err2 != nil {
		panic("Could not read file")
		return
	}
	scanner := bufio.NewScanner(input)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	state := 0 // 0: read size, 1: read graph 1, 2: read graph 2
	size := 0
	node := 0

	g1 := newGraph("a")
	g2 := newGraph("b")

	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " ")
		if l == "" || strings.HasPrefix(l, "%") || strings.HasPrefix(l, "*") {
			continue
		}

		switch state {
		case 0: // deprecated: for parsing the "header" of pb files, now parser is flexible
			{
				var b error
				size, b = strconv.Atoi(l)
				if b != nil {
					panic("bad conversion of numbers")
				}
				g1.adj = make([][]bool, size)
				g1.deg = make([]int, size)
				g1.size = size
				g2.adj = make([][]bool, size)
				g2.deg = make([]int, size)
				g2.size = size
				for i := 0; i < size; i++ {
					g1.adj[i] = make([]bool, size)
					g2.adj[i] = make([]bool, size)
				}
				state = 1
			}
		case 1:
			{
				d := 0
				for i, s := range []rune(l) {
					if s == rune('1') {
						g1.adj[node][i] = true
						d++
					}
				}
				g1.deg[node] = d
				node++
				if size == node {
					node = 0
					state = 2
				}
			}
		case 2:
			{
				d := 0
				for i, s := range []rune(l) {
					if s == rune('1') {
						g2.adj[node][i] = true
						d++
					}
				}
				g2.deg[node] = d
				node++
				if size == node {
					node = 0
					state = 2
				}
			}
		}

	}
	var invariants []invariant

	if *di {
		invariants = append(invariants, degInv)
	}
	if *ddi {
		invariants = append(invariants, degDepthInv)
	}
	if *dd2 {
		invariants = append(invariants, degDepth2Inv)
	}

	if *dot {
		fmt.Println("graph test {")
		printMappings(g1, g2, invariants)
		g1.printDot()
		g2.printDot()
		fmt.Println("}")
	} else {
		g1.printGringo()
		g2.printGringo()
		encodeMappings(g1, g2, invariants)
	}
}

func (g *graph) printDot() {
	for i := 0; i < g.size; i++ {
		for j := i + 1; j < g.size; j++ {
			if g.adj[i][j] {
				fmt.Printf("%v%v -- %v%v ; \n", g.pre, i, g.pre, j)
			}
		}
	}
}

func (g *graph) printGringo() {
	for i := 0; i < g.size; i++ {
		for j := 0; j < g.size; j++ {
			if g.adj[i][j] {
				fmt.Printf("edge(%v,%v%v,%v%v).\n", g.pre, g.pre, i, g.pre, j)
			}
		}
	}
}

func ids(node int, g graph, invariants []invariant) (s string) {
	for i, inv := range invariants {
		if i != 0 {
			s += "-"
		}
		s += inv(node, g)
	}
	return s
}

func printMappings(g1, g2 graph, invariants []invariant) {

	mapId := make(map[string]int)
	nextId := 0

	for i := 0; i < g1.size; i++ {

		s1 := ids(i, g1, invariants)
		s2 := ids(i, g2, invariants)

		if _, ok := mapId[s1]; !ok {
			mapId[s1] = nextId
			nextId++
		}

		if _, ok := mapId[s2]; !ok {
			mapId[s2] = nextId
			nextId++
		}

		fmt.Printf("%v%v [style=filled,fillcolor=\"/paired12/%v\"]\n", g1.pre, i, mapId[s1]+1)
		fmt.Printf("%v%v [style=filled,fillcolor=\"/paired12/%v\"]\n", g2.pre, i, mapId[s2]+1)

	}

}

func encodeMappings(g1, g2 graph, invariants []invariant) {

	g1map := make(map[string][]int)
	g2map := make(map[string][]int)

	for i := 0; i < g1.size; i++ {

		s1 := ids(i, g1, invariants)
		s2 := ids(i, g2, invariants)

		g1map[s1] = append(g1map[s1], i)
		g2map[s2] = append(g2map[s2], i)
	}

	conflict := false

	for s, n1s := range g1map {
		n2s := g2map[s]
		if len(n1s) != len(n2s) {
			conflict = true
			fmt.Println("a. :- a. % naive invariant checks show that there is no isomorphism")
		}
	}

	if !conflict {
		for s, n1s := range g1map {
			n2s := g2map[s]
			fmt.Println("% set:", s, " size:", len(n1s))
			fmt.Println("% graph1:", n1s)
			fmt.Println("% graph2:", n2s)
			for _, i := range n1s {
				for _, j := range n2s {
					fmt.Printf("mapping(%v%v,%v%v).\n", g1.pre, i, g2.pre, j)
				}
			}
		}
	}

	return
}
