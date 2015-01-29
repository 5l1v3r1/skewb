package main

import (
	"fmt"
	"github.com/unixpickle/skewb"
	"os"
)

var identity skewb.Skewbs
var heuristic skewb.COHeuristic

func main() {
	identity = skewb.Skewbs(skewb.NewSkewb().AllRotations())
	puzzle, err := skewb.ReadPuzzle()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Generating heuristic...")
	heuristic = skewb.MakeCOHeuristic(skewb.AllMoves())
	for depth := 0; depth < 20; depth++ {
		fmt.Println("Exploring depth", depth, "...")
		solution := solve(puzzle, '_', depth)
		if solution != nil {
			fmt.Println("Got solution:", solution)
			return
		}
	}
}

func solve(s *skewb.Skewb, last rune, remaining int) []skewb.Move {
	if remaining == 0 {
		if !identity.Contains(s) {
			return nil
		} else {
			return []skewb.Move{}
		}
	} else {
		if heuristic.MinMoves(s) > remaining {
			return nil
		}
	}

	for _, face := range "BLRU" {
		if face == last {
			continue
		}
		for i := 0; i < 2; i++ {
			m := skewb.Move{face, i == 0}
			newS := *s
			newS.Move(m)
			if solution := solve(&newS, face, remaining-1); solution != nil {
				return append([]skewb.Move{m}, solution...)
			}
		}
	}

	return nil
}