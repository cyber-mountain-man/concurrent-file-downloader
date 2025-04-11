# 🔐 Version 4 – Concurrent File Downloader with SHA256 Checksum & Retry Logic

This version enhances the concurrent file downloader by adding **SHA256 hash verification** and **automatic retry logic** for failed or corrupted downloads.

It builds on [v3 (Worker Pool)](../v3-worker-pool) and introduces production-level features for **data integrity** and **resilience**.

---

## 🚀 What It Does

- Downloads files concurrently using a **worker pool**
- **Verifies SHA256 hash** after download to ensure file integrity
- Automatically **retries failed downloads** (e.g., network errors or bad hashes)
- Deletes invalid files before retrying
- Limits retry attempts with exponential backoff between tries

---

## 📁 Files

- `main.go` – Full source code with detailed comments and a walkthrough
- `go.mod` – Module file for managing dependencies (standard library only)

---

## 🛠️ How to Run

1. Navigate to the folder:

```bash
cd v4-checksum-verification
```

2. Run the program:

```bash
go run main.go
```

3. Output (example):

```
🧵 Worker 1: Attempt 1 for mountain.jpg
✅ Worker 1: Verified and saved mountain.jpg
🧵 Worker 2: Attempt 1 for transparency.png
✅ Worker 2: Verified and saved transparency.png
🚀 All downloads with checksum verification completed.
```

---

## 🧠 Key Concepts Covered

| Feature               | Description |
|------------------------|-------------|
| `DownloadJob` struct   | Encapsulates file URL, filename, expected SHA256 hash, and max retry attempts |
| SHA256 checksum        | Ensures downloaded file content is not tampered or incomplete |
| Retry mechanism        | Retries failed downloads or mismatched hashes |
| `sync.WaitGroup`       | Waits for all workers to complete |
| Manual resource cleanup | Proper `Close()` and file removal logic prevents leaks |

---

## 🧪 Example Files with Placeholder Hashes

The hashes included in `main.go` are **placeholders**. To generate your own:

```bash
sha256sum downloaded_file.jpg
```

Then update the `ExpectedHash` field accordingly.

---

## 🔄 What's Next?

In future versions, we may explore:

- ✅ Reading checksum values from a file (`checksums.txt`)
- 📦 Bundling as a CLI tool with arguments (`--url`, `--file`, `--retries`)
- 📊 Adding metrics or logging
- 🔁 Parallel downloads with individual retry backoff windows

---

## 👤 Author

Guillermo Morrison  
[GitHub: cyber-mountain-man](https://github.com/cyber-mountain-man)

---

## 🪪 License

MIT – Free for personal, educational, or commercial use
```