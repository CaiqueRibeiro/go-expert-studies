package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
)

const (
	numRecords = 500000          // total nunber of registers
	chunkSize  = 1000            // registers per chunk
	bufferSize = 2 * 1024 * 1024 // 2MB buffer
	serverURL  = "http://localhost:8080/upload"
	numWorkers = 50 // number of workers to the send
)

type Record struct {
	ID       int               `json:"id"`
	Values   []float64         `json:"values"`
	Metadata map[string]string `json:"metadata"`
}

func main() {
	var wg sync.WaitGroup
	chunks := make(chan *bytes.Buffer, numRecords/chunkSize)

	// Generate JSON chunks in goroutines
	wg.Add(1)
	go func() {
		defer wg.Done()
		generateJSONChunks(chunks)
	}()

	// Send JSON chunks to the server in goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sendChunks(chunks)
		}()
	}

	wg.Wait()

	printGCStats()
}

func generateJSONChunks(chunks chan<- *bytes.Buffer) {
	for i := 0; i < numRecords; i += chunkSize {
		buffer := bytes.NewBuffer(make([]byte, 0, bufferSize))

		for j := 0; j < i+chunkSize && j < numRecords; j++ {
			record := &Record{
				ID:       j,
				Values:   generateRandomValues(100),
				Metadata: generateRandomMetadata(10),
			}

			data, _ := json.Marshal(record)
			buffer.Write(data)
			buffer.WriteByte('\n')
		}

		// send the generated chunk
		chunks <- buffer
	}
	close(chunks)
}

func sendChunks(chunks <-chan *bytes.Buffer) {
	for buffer := range chunks {
		// send the chunk to the server
		resp, err := http.Post(serverURL, "application/json", buffer)
		if err != nil {
			fmt.Println("Error sending chunk:", err)
		}
		resp.Body.Close()
	}
}

func generateRandomValues(size int) []float64 {
	// generate slice of random floats
	values := make([]float64, size)
	for i := 0; i < size; i++ {
		values[i] = float64(i)
	}
	return values
}

func generateRandomMetadata(size int) map[string]string {
	// generate a map of random strings
	metadata := make(map[string]string, size)
	for i := 0; i < size; i++ {
		metadata[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", rand.Int())
	}
	return metadata
}

func printGCStats() {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	// auxiliary functions to convert bytes
	toMB := func(bytes uint64) float64 {
		return float64(bytes) / 1024 / 1024
	}

	fmt.Printf("Alloc = %v MB", toMB(stats.Alloc))
	fmt.Printf("\tTotalAlloc = %v MB", toMB(stats.TotalAlloc))
	fmt.Printf("\tSys = %v MB", toMB(stats.Sys))
	fmt.Printf("\tNumGC = %v\n", stats.NumGC)
}
