package main

import (
	"crypto/sha256" // For calculating SHA256 file hashes
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// DownloadJob defines a download task with the source URL,
// the local filename to save as, expected checksum, and retry count.
type DownloadJob struct {
	URL          string
	Filename     string
	ExpectedHash string
	MaxRetries   int
}

// worker processes download jobs by repeatedly attempting to download
// and verify the file until success or max retries are reached.
func worker(id int, jobs <-chan DownloadJob, wg *sync.WaitGroup) {
	defer wg.Done() // Notify WaitGroup when this worker finishes

	for job := range jobs {
		success := false // Tracks if the download eventually succeeds

		// Retry loop
		for attempt := 1; attempt <= job.MaxRetries; attempt++ {
			fmt.Printf("🧵 Worker %d: Attempt %d for %s\n", id, attempt, job.Filename)

			// Step 1: Attempt download
			err := downloadFile(job.URL, job.Filename)
			if err != nil {
				fmt.Printf("❌ Worker %d: Download failed for %s: %v\n", id, job.Filename, err)
				time.Sleep(2 * time.Second) // ⏱ Delay before retrying
				continue
			}

			// Step 2: Check integrity
			match, err := verifyChecksum(job.Filename, job.ExpectedHash)
			if err != nil {
				fmt.Printf("❌ Worker %d: Checksum error for %s: %v\n", id, job.Filename, err)
				continue
			}

			if match {
				fmt.Printf("✅ Worker %d: Verified and saved %s\n", id, job.Filename)
				success = true
				break
			} else {
				// If checksum fails, delete the file and retry
				fmt.Printf("⚠️ Worker %d: Checksum mismatch for %s. Retrying...\n", id, job.Filename)
				os.Remove(job.Filename)
				time.Sleep(2 * time.Second)
			}
		}

		if !success {
			fmt.Printf("❌ Worker %d: Failed to download %s after %d attempts\n", id, job.Filename, job.MaxRetries)
		}
	}
}

// downloadFile downloads a file from a URL and writes it to the local filesystem.
func downloadFile(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// verifyChecksum calculates the SHA256 hash of the downloaded file and
// compares it to the expected hash provided in the job.
func verifyChecksum(filename, expected string) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return false, err
	}
	actual := fmt.Sprintf("%x", hash.Sum(nil))
	return actual == expected, nil
}

func main() {
	// Define jobs with URL, filename, and known SHA256 hash
	jobs := []DownloadJob{
		{
			URL:          "https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg",
			Filename:     "mountain.jpg",
			ExpectedHash: "3f4f36a92296d008b9e3897fc02fdb0190d2db66317c38c8e0f292d0f86959a4", // Placeholder
			MaxRetries:   3,
		},
		{
			URL:          "https://upload.wikimedia.org/wikipedia/commons/4/47/PNG_transparency_demonstration_1.png",
			Filename:     "transparency.png",
			ExpectedHash: "fb3f03c9816d2a4985d6d942b1e4a1cf655f5306d861126317b07c4a2c12b3ed", // Placeholder
			MaxRetries:   3,
		},
	}

	const workerCount = 2
	jobChan := make(chan DownloadJob, len(jobs)) // Buffered job queue
	var wg sync.WaitGroup

	// Spin up worker goroutines
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobChan, &wg)
	}

	// Enqueue jobs
	for _, job := range jobs {
		jobChan <- job
	}
	close(jobChan) // No more jobs coming

	wg.Wait() // Wait for all downloads to finish
	fmt.Println("🚀 All downloads with checksum verification completed.")
}

/*
===============================================
🧠 Walkthrough – Version 4: Checksum + Retries
===============================================

This version enhances the concurrent downloader by ensuring **file integrity** and **resilience**.

🔹 Key Features:
----------------
- 🔐 SHA256 checksum validation ensures the downloaded file matches the expected content.
- 🔁 Retry logic allows up to N attempts if a download fails or checksum mismatches.
- 🧵 Workers pull jobs from a buffered channel (like a job queue).
- ✅ Resources (files and response bodies) are properly closed after use.
- 🧼 Invalid files are removed before retrying to avoid corrupt or partial data.

🔹 Flow Summary:
----------------
1. Each `DownloadJob` defines a URL, output filename, expected SHA256 hash, and retry count.
2. Workers run concurrently, each looping through jobs in the channel.
3. For each job:
   - The worker downloads the file.
   - It computes and compares the SHA256 hash.
   - If the hash matches: success 🎉
   - If not: delete and retry
4. The program exits only when all workers finish processing their jobs.

===============================
🏁 Result:
- Production-grade file downloader with data validation
- A strong foundation for CLI tools, deployment updaters, or secure mirrors
- Future-proofed for things like config files, JSON-based job definitions, or parallel logging
===============================
*/
