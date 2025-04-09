package main

import (
	"fmt"      // For console output
	"io"       // For streaming download data
	"net/http" // For making HTTP requests
	"os"       // For creating and writing to files
	"sync"     // For coordinating goroutines using WaitGroup
)

// DownloadJob holds the URL and target filename for a download task
type DownloadJob struct {
	URL      string
	Filename string
}

// worker is a concurrent goroutine that pulls jobs from the shared job channel
func worker(id int, jobs <-chan DownloadJob, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this worker is finished when it exits

	for job := range jobs {
		fmt.Printf("🧵 Worker %d started: %s\n", id, job.Filename)

		// Step 1: Make HTTP GET request
		resp, err := http.Get(job.URL)
		if err != nil {
			fmt.Printf("❌ Worker %d failed to download %s: %v\n", id, job.URL, err)
			continue
		}

		// Step 2: Create output file locally
		out, err := os.Create(job.Filename)
		if err != nil {
			fmt.Printf("❌ Worker %d failed to create file %s: %v\n", id, job.Filename, err)
			resp.Body.Close() // Clean up HTTP resource
			continue
		}

		// Step 3: Stream response body directly into the file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			fmt.Printf("❌ Worker %d failed to write file %s: %v\n", id, job.Filename, err)
		} else {
			fmt.Printf("✅ Worker %d completed: %s\n", id, job.Filename)
		}

		// ✅ Explicitly close resources (don't defer inside loop)
		out.Close()
		resp.Body.Close()
	}
}

func main() {
	// Define the list of download jobs
	jobs := []DownloadJob{
		{
			URL:      "https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg",
			Filename: "mountain.jpg",
		},
		{
			URL:      "https://upload.wikimedia.org/wikipedia/commons/4/47/PNG_transparency_demonstration_1.png",
			Filename: "transparency.png",
		},
	}

	const workerCount = 2
	jobChan := make(chan DownloadJob, len(jobs)) // Buffered channel to queue jobs
	var wg sync.WaitGroup

	// Spin up a fixed number of worker goroutines
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobChan, &wg)
	}

	// Push jobs into the channel (producer)
	for _, job := range jobs {
		jobChan <- job
	}
	close(jobChan) // Signal that no more jobs are coming

	wg.Wait() // Wait for all workers to finish
	fmt.Println("🚀 All downloads completed via worker pool.")
}

/*
=======================================
🧠 Walkthrough – Version 3: Worker Pool
=======================================

This version demonstrates how to implement a **worker pool** to control
the number of concurrent downloads.

🔹 Core Concepts:
-----------------
- We define a struct `DownloadJob` to hold download tasks.
- We use a **channel** (`jobChan`) to queue these jobs.
- We start **N workers** (in this case, 2) that pull from the channel.
- Each worker is a goroutine that loops through jobs, processes them, and closes resources manually.

🔹 Why a Worker Pool?
---------------------
- Instead of launching 1 goroutine per file (which may be hundreds),
  a worker pool **limits concurrency** and **prevents resource exhaustion**.
- It mimics real-world systems like task queues, background job servers, and batch processors.

🔹 Key Features:
----------------
- `sync.WaitGroup` tracks when all workers are done
- `make(chan DownloadJob, len(jobs))` buffers the job queue
- `close(jobChan)` is critical to allow workers to exit cleanly
- Explicit calls to `Close()` prevent long-lived open files/sockets
- Error handling ensures failed downloads don’t crash the program

===============================
🏁 Result:
- Two workers download files concurrently from a shared queue
- The system finishes only when all jobs are complete
- Easily extendable to more workers or retry logic in future versions
===============================
*/
