# 📦 Concurrent File Downloader (Go)

A multi-version Go project demonstrating **real-world concurrency patterns** through file downloading.

Each version builds upon the previous one, introducing new features like progress bars, worker pools, file verification, and error recovery — all using idiomatic Go.

---

## 📚 Versions

| Version | Folder | Description |
|---------|--------|-------------|
| ✅ v1 | [`v1-basic-downloader`](./v1-basic-downloader) | Download multiple files concurrently using `goroutines` and `sync.WaitGroup` |
| ✅ v2 | [`v2-progress-bar`](./v2-progress-bar)         | Adds real-time progress bars using `github.com/schollz/progressbar/v3` |
| ✅ v3 | [`v3-worker-pool`](./v3-worker-pool)           | Adds concurrency control using a worker pool and job queue |
| ✅ v4 | [`v4-checksum-verification`](./v4-checksum-verification) | Adds SHA256 checksum validation and retry logic for download failures |

---

## 🛠️ Requirements

- Go 1.20 or newer
- Internet connection for downloading files

---

## 🧠 Topics Covered

- Goroutines
- WaitGroups
- Channels
- HTTP client usage
- File I/O
- Hashing and data integrity (SHA256)
- Retry logic and error handling
- External Go modules
- Real-world concurrency modeling

---

## 📂 Project Structure

```
concurrent-file-downloader/
├── v1-basic-downloader/         # Basic concurrent file downloading
│   ├── main.go
│   └── README.md
├── v2-progress-bar/             # Adds progress bars to concurrent downloads
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   └── README.md
├── v3-worker-pool/              # Controls max concurrent downloads (TBD)
├── v4-checksum-verification/   # Adds file integrity check and retries (TBD)
└── README.md                    # This file
```
---

## 🔍 Why This Project Exists

Go is built for concurrency — this project helps demystify it through small, understandable programs that solve real problems.

---

## 👤 Author

**Guillermo Morrison**  
[GitHub: cyber-mountain-man](https://github.com/cyber-mountain-man)  
Graduate Student – Cybersecurity & Software Engineering

---

## 🪪 License

MIT – Free to use, modify, or contribute.

---

*Want to follow along as we build each version? Star the repo and check back weekly!*
```