// Run with: go run simulator.go a.txt out_a.txt
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Scan(scanner *bufio.Scanner) {
	if !scanner.Scan() {
		Check(scanner.Err())
	}
}

func ReadProblem(p string) (D, I, S, F, V int, B, E, L []int, N []string, streets map[string]int, P []int, path [][]int) {
	file, err := os.Open(p)
	Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	Scan(scanner)
	DISVF := strings.Split(scanner.Text(), " ")
	D, err = strconv.Atoi(DISVF[0])
	Check(err)
	I, err = strconv.Atoi(DISVF[1])
	Check(err)
	S, err = strconv.Atoi(DISVF[2])
	Check(err)
	V, err = strconv.Atoi(DISVF[3])
	Check(err)
	F, err = strconv.Atoi(DISVF[4])
	Check(err)
	B = make([]int, S, S)
	E = make([]int, S, S)
	L = make([]int, S, S)
	N = make([]string, S, S)
	streets = make(map[string]int)
	for i := 0; i < S; i++ {
		Scan(scanner)
		BENL := strings.Split(scanner.Text(), " ")
		B[i], err = strconv.Atoi(BENL[0])
		Check(err)
		E[i], err = strconv.Atoi(BENL[1])
		Check(err)
		N[i] = BENL[2]
		streets[N[i]] = i
		L[i], err = strconv.Atoi(BENL[3])
		Check(err)
	}
	P = make([]int, V, V)
	path = make([][]int, V, V)
	for i := 0; i < V; i++ {
		Scan(scanner)
		line := strings.Split(scanner.Text(), " ")
		P[i], err = strconv.Atoi(line[0])
		Check(err)
		path[i] = make([]int, P[i], P[i])
		for j := 0; j < P[i]; j++ {
			path[i][j] = streets[line[j+1]]
		}
	}
	return
}

type GreenLight struct {
	street int
	T      int
}

func GreenNow(schedule []GreenLight, t int) (street int) {
	total := 0
	for _, e := range schedule {
		total += e.T
	}
	if total == 0 {
		return -1
	}
	t = t % total
	total = 0
	for _, e := range schedule {
		total += e.T
		if t < total {
			return e.street
		}
	}
	return
}

func ReadSolution(p string, I int, streets map[string]int) (schedule [][]GreenLight) {
	file, err := os.Open(p)
	Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	Scan(scanner)
	A, err := strconv.Atoi(scanner.Text())
	Check(err)
	schedule = make([][]GreenLight, I, I)
	for j := 0; j < A; j++ {
		Scan(scanner)
		i, err := strconv.Atoi(scanner.Text())
		Check(err)
		Scan(scanner)
		E, err := strconv.Atoi(scanner.Text())
		Check(err)
		schedule[i] = make([]GreenLight, E, E)
		for k := 0; k < E; k++ {
			Scan(scanner)
			streetT := strings.Split(scanner.Text(), " ")
			schedule[i][k] = GreenLight{}
			schedule[i][k].street = streets[streetT[0]]
			schedule[i][k].T, err = strconv.Atoi(streetT[1])
			Check(err)
		}
	}
	return
}

func main() {
	D, I, S, _, V, _, _, L, _, streets, P, path := ReadProblem(os.Args[1])
	// log.Printf("%v %v %v %v %v %v %v %v %v %v", D, I, S, F, V, B, E, L, N, path)
	schedule := ReadSolution(os.Args[2], I, streets)
	// log.Printf("%v", schedule)

	step := make([]int, V, V)      // Index of street the car is on in its path.
	togo := make([]int, V, V)      // How much longer it has to drive straight.
	waiting := make([][]int, S, S) // Cars waiting on each street.
	// Put everyone in the wait queue for the first street.
	for v, p := range path {
		waiting[p[0]] = append(waiting[p[0]], v)
	}
	t := 0
	score := 0
	for t < D {
		// Drive on a straight line.
		for v := 0; v < V; v++ {
			if togo[v] > 0 {
				togo[v] -= 1
				if togo[v] == 0 {
					if step[v] == len(path[v])-1 {
						score += P[v] + D - t
					} else {
						waiting[path[v][step[v]]] = append(waiting[path[v][step[v]]], v)
					}
				}
			}
		}
		// Cross an intersection.
		for i := 0; i < I; i++ {
			g := GreenNow(schedule[i], t)
			if g < 0 {
				continue
			}
			if len(waiting[g]) > 0 {
				v := waiting[g][0]
				waiting[g] = waiting[g][1:]
				step[v]++
				togo[v] = L[path[v][step[v]]]
			}
		}
		t++
	}
	log.Printf("score: %v", score)
}
