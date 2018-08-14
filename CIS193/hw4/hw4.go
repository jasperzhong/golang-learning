// Homework 4: Concurrency
// Due February 21, 2017 at 11:59pm
package main

import (
	"fmt"
	"io"
	"os"
	"log"
	"bufio"
	"strconv"
	"sync"
)

func main() {
	// Feel free to use the main function for testing your functions
	//FileSum("file_sum.txt","sum.txt")
	/*
	f1, err := os.Open("file_sum.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.Create("sum.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer f2.Close()

	IOSum(f1, f2)*/
	d := PennDirectory{
		directory: make(map[int]string),
	}
	total := 10
	var wg sync.WaitGroup
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			switch i % 3 {
			case 0:
				d.Add(1, "golang")
			case 1:
				d.Add(2, "cpp")
			case 2:
				d.Remove(2)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("d.directory = %v\n", d.directory)

}

// Problem 1a: File processing
// You will be provided an input file consisting of integers, one on each line.
// Your task is to read the input file, sum all the integers, and write the
// result to a separate file.

// FileSum sums the integers in input and writes them to an output file.
// The two parameters, input and output, are the filenames of those files.
// You should expect your input to end with a newline, and the output should
// have a newline after the result.
func FileSum(input, output string) {
	fi, err := os.Open(input)
	if err != nil{
		log.Fatal(err)
	}
	defer fi.Close()

	sum := 0
	var tmp int
	for{
		_, err = fmt.Fscanf(fi, "%d\n", &tmp)
		if err != nil {
			if err == io.EOF{
				break
			}
			log.Fatal(err)
		}

		sum += tmp
	}

	fo, err := os.Create(output)
	if err != nil{
		log.Fatal(err)
	}
	defer fo.Close()

	fmt.Fprintf(fo, "%d", sum)
}

// Problem 1b: IO processing with interfaces
// You must do the exact same task as above, but instead of being passed 2
// filenames, you are passed 2 interfaces: io.Reader and io.Writer.
// See https://golang.org/pkg/io/ for information about these two interfaces.
// Note that os.Open returns an io.Reader, and os.Create returns an io.Writer.

// IOSum sums the integers in input and writes them to output
// The two parameters, input and output, are interfaces for io.Reader and
// io.Writer. The type signatures for these interfaces is in the Go
// documentation.
// You should expect your input to end with a newline, and the output should
// have a newline after the result.
func IOSum(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan(){
		x, err := strconv.Atoi(scanner.Text())
		if err != nil{
			log.Fatal(err)
		}
		sum += x
	}

	writer := bufio.NewWriter(output)
	writer.WriteString(strconv.Itoa(sum) + "\n")
	writer.Flush()
}

// Problem 2: Concurrent map access
// Maps in Go [are not safe for concurrent use](https://golang.org/doc/faq#atomic_maps).
// For this assignment, you will be building a custom map type that allows for
// concurrent access to the map using mutexes.
// The map is expected to have concurrent readers but only 1 writer can have
// access to the map.

// PennDirectory is a mapping from PennID number to PennKey (12345678 -> adelq).
// You may only add *private* fields to this struct.
// Hint: Use an embedded sync.RWMutex, see lecture 2 for a review on embedding
type PennDirectory struct {
	mu sync.RWMutex
	directory map[int]string
}

// Add inserts a new student to the Penn Directory.
// Add should obtain a write lock, and should not allow any concurrent reads or
// writes to the map.
// You may NOT write over existing data - simply raise a warning.
func (d *PennDirectory) Add(id int, name string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if _, ok := d.directory[id]; !ok {
		fmt.Println("Write over existing data！")
	}
	d.directory[id] = name
}

// Get fetches a student from the Penn Directory by their PennID.
// Get should obtain a read lock, and should allow concurrent read access but
// not write access.
func (d *PennDirectory) Get(id int) string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.directory[id]
}

// Remove deletes a student to the Penn Directory.
// Remove should obtain a write lock, and should not allow any concurrent reads
// or writes to the map.
func (d *PennDirectory) Remove(id int) {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.directory, id)
}
