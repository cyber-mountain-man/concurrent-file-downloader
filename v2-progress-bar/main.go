package main

import (
	"fmt"      // For console output
	"io"       // For streaming download data
	"net/http" // For making HTTP requests
	"os"       // For creating and writing to files
	"sync"     // For coordinating goroutines using WaitGroup

	"github.com/schollz/progressbar/v3" // External library for terminal progress bars
)

// downloadFileWithProgress downloads a file from the given URL and displays a real-time progress bar.
func downloadFileWithProgress(url, filename string, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as complete when function exits

	// Step 1: Make a HEAD request to determine the file size
	resp, err := http.Head(url)
	if err != nil {
		fmt.Printf("❌ Failed to get headers for %s: %v\n", url, err)
		return
	}
	size := resp.ContentLength // Total number of bytes to expect

	// Step 2: Perform a GET request to start downloading the file
	resp, err = http.Get(url)
	if err != nil {
		fmt.Printf("❌ Failed to download %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	// Step 3: Create a local file to write the downloaded data
	out, err := os.Create(filename)
	if err != nil {
		fmt.Printf("❌ Failed to create file %s: %v\n", filename, err)
		return
	}
	defer out.Close()

	// Step 4: Create a progress bar bound to the file size
	bar := progressbar.DefaultBytes(
		size,
		fmt.Sprintf("📥 Downloading %s", filename),
	)

	// Step 5: Stream the response body into the file and the progress bar simultaneously
	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	if err != nil {
		fmt.Printf("❌ Failed to save %s: %v\n", filename, err)
		return
	}

	fmt.Printf("\n✅ Download completed: %s → %s\n", url, filename)
}

func main() {
	// URLs to download mapped to filenames
	urls := map[string]string{
		"https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg": "mountain.jpg",
		"https://upload.wikimedia.org/wikipedia/commons/4/47/PNG_transparency_demonstration_1.png": "transparency.png",
	}

	var wg sync.WaitGroup // WaitGroup to track all concurrent downloads

	// Launch one goroutine per file download
	for url, filename := range urls {
		wg.Add(1)
		go downloadFileWithProgress(url, filename, &wg)
	}

	wg.Wait() // Block until all downloads are complete
	fmt.Println("✅ All downloads with progress bars completed.")
}

/*
=====================================
🧠 Version 2 Walkthrough: Progress Bar Downloader
=====================================

This version enhances the basic concurrent file downloader by showing a
real-time progress bar for each file being downloaded.

🔹 Step-by-step summary:

1. The program uses a `map[string]string` to define what to download and what to name it.
2. A `sync.WaitGroup` tracks when all downloads are finished.
3. Each download runs in a separate `goroutine` for concurrency.
4. We use `http.Head()` to get the size of the file (needed for the progress bar).
5. The `progressbar` package displays a smooth, real-time visual indicator of download progress.
6. We use `io.MultiWriter()` to write to the file **and** update the bar simultaneously.

This is a solid, production-style example of combining:
- Concurrency
- User feedback
- Real-world I/O and error handling

It also introduces external libraries (`progressbar/v3`) and `go.mod` for dependency tracking.

===============================
🏁 Result:
- Each file download is visible in real time.
- Program exits only when all downloads complete.
- A great starting point for CLI tools or download managers.
===============================
*/
