// +build ignore

package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

var notes = [...]string{
	"C",
	"Cs",
	"D",
	"Ds",
	"E",
	"F",
	"Fs",
	"G",
	"Gs",
	"A",
	"As",
	"B",
}

var same = map[string]string{
	"Cs": "Db",
	"Ds": "Eb",
	"Fs": "Gb",
	"As": "Bb",
}

func e(message string, err error) {
	if err != nil {
		fmt.Println(message, err)
		os.Exit(1)
	}
}

var spaces = "                    "

func main() {
	var (
		key    uint
		pitch  float64
		num    uint
		output string
	)
	flag.UintVar(&key, "k", 49, "reference key number")
	flag.Float64Var(&pitch, "p", 440, "reference key pitch")
	flag.StringVar(&output, "o", "notes.go", "output file")
	flag.UintVar(&num, "n", 108, "number of notes")
	flag.Parse()
	f, err := os.Create(output)
	e("unable to create output file:", err)
	_, err = fmt.Fprintf(f, "//go:generate go run notes_gen.go -k %d -p %f -n %d\n"+
		"// generated with notes_gen.go\n"+
		"\n"+
		"package music\n"+
		"\n"+
		"type Note float64\n"+
		"\n"+
		"const (\n",
		key, pitch, num)
	e("error writing file header", err)

	spaces = spaces[:numSize(num)-1]

	semitone := math.Pow(2, float64(1)/12)
	keyF := float64(key)
	for i := uint(0); i < num; i++ {
		note := pitch * math.Pow(semitone, float64(i)-float64(keyF)-8)
		_, err = fmt.Fprintf(f, "	%s%d %sNote = %f\n", notes[i%12], i/12, spaces[len(notes[i%12]):], note)
		e("error writing note", err)
		if sn, ok := same[notes[i%12]]; ok {
			_, err = fmt.Fprintf(f, "	%s%d %sNote = %f\n", sn, i/12, spaces[len(notes[i%12]):], note)
			e("error writing note", err)
		}
	}
	_, err = f.WriteString(")\n")
	e("error writing file trailer", err)
	e("error closing file", f.Close())
}

func numSize(n uint) uint {
	var m uint = 1
	for n >= 10 {
		n /= 10
		m++
	}
	return m
}
