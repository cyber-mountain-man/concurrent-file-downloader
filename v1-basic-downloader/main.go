package main

import (
	"fmt"      // For printing messages to the console
	"io"       // For reading from the HTTP response and writing to a file
	"net/http" // For making HTTP requests
	"os"       // For creating local files
	"sync"     // For using WaitGroup to wait for all goroutines to finish
)

// downloadFile downloads a file from the given URL and saves it as the given filename.
// It uses a WaitGroup to notify when it's done.
func downloadFile(url, filename string, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when function ends

	// Make a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to download %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	// Create the local file where data will be saved
	out, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Failed to create file %s: %v\n", filename, err)
		return
	}
	defer out.Close()

	// Copy the HTTP response body directly into the local file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("Failed to save %s: %v\n", filename, err)
		return
	}

	fmt.Printf("Downloaded: %s → %s\n", url, filename)
}

func main() {
	// Map of download URLs to filenames
	urls := map[string]string{
		"https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg": "mountain.jpg",
		"https://upload.wikimedia.org/wikipedia/commons/4/47/PNG_transparency_demonstration_1.png": "transparency.png",
	}

	var wg sync.WaitGroup // Create a WaitGroup to wait for all downloads

	// Start one goroutine per file download
	for url, filename := range urls {
		wg.Add(1)                        // Tell WaitGroup we're starting a new task
		go downloadFile(url, filename, &wg) // Start download in its own goroutine
	}

	wg.Wait() // Wait for all downloads to finish
	fmt.Println("All downloads completed.")
}

/*
=====================================
🧠 Program Walkthrough (Concurrent File Downloader)
=====================================

1. We define a map called `urls` that holds:
   - The file's download URL (as the key)
   - The name we want to save it as locally (as the value)

2. We initialize a `sync.WaitGroup` named `wg`:
   - This helps us keep track of how many downloads are still running
   - We’ll wait for them all to finish before exiting the program

3. For each entry in the `urls` map:
   - We call `wg.Add(1)` to increment the counter
   - We launch a new goroutine using `go downloadFile(...)`
     → Each file download runs in parallel, not one after another

4. Inside `downloadFile`:
   - We defer `wg.Done()` so the counter is decremented when the function finishes
   - We make an HTTP GET request to fetch the file
   - If successful, we create a local file and write the downloaded content to it
   - Any errors are logged using `fmt.Printf`

5. Back in `main()`:
   - We call `wg.Wait()` to block execution until ALL downloads finish

6. Once all downloads are complete:
   - We print a final message to the console

===============================
🏁 Result:
- All files in the map will be downloaded at the same time (concurrently)
- Each file is saved with its own name in the current directory
===============================
*/
