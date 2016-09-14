package main

import "fmt"
import "flag"
import "os"
import "strings"
import "bufio"
import "strconv"

var filename = flag.String("f", "x", "Path to graph file.")

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
	graph := 1
	size := 0
	node := 0
	pre := "a"

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
				state = 1
			}
		case 1:
			{
				if graph == 2 {
					pre = "b"
				}
				for i, s := range []rune(l) {
					if s == rune('1') {
						fmt.Printf("edge(%v,%v%v,%v%v).\n", graph, pre, node, pre, i)
					}
				}
				node++
				if size == node {
					graph = 2
					node = 0
				}
			}
			//		case 2:  {
			//		}
			//		default:  {
			//		}
		}
	}
}
