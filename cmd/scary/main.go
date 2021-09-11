package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	scary "github.com/jackdoe/go-scary/pkg"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func max(a, b *int) int {
	if *a > *b {
		return *a
	}
	return *b
}

func main() {
	overlayMax := flag.Int("overlay-max", 0, "overlay from min to N")
	overlayMin := flag.Int("overlay-min", 0, "overlay from N to max")
	aboveMin := flag.Int("above-min", 5, "above from N to max")
	aboveMax := flag.Int("above-max", 10, "above from min to N")
	belowMin := flag.Int("below-min", 5, "below from N to max")
	belowMax := flag.Int("below-max", 10, "below from min to N")

	flag.Parse()

	settings := []scary.Settings{
		{Runes: scary.Above, Min: *aboveMin, Max: max(aboveMax, aboveMin)},
		{Runes: scary.Below, Min: *belowMin, Max: max(belowMax, belowMin)},
		{Runes: scary.Overlay, Min: *overlayMin, Max: max(overlayMax, overlayMin)},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print(scary.ScaryWithSettings(scanner.Text(), settings))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
