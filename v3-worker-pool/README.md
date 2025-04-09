# 🧵 Version 3 – Concurrent File Downloader with Worker Pool

This version adds a **worker pool architecture** to control the number of files downloaded at once, making the downloader more resource-conscious and scalable.

It builds on the concurrency foundation from [v1](../v1-basic-downloader) and sets the stage for more advanced features in future versions.

---

## 🚀 What It Does

- Uses a **fixed number of workers** to download files concurrently
- Introduces a **job queue** via buffered channels
- Ensures resource cleanup and error resilience
- Explicitly manages file/network resources (no deferred leaks!)

---

## 🧠 Why Worker Pools Matter

In real-world applications, spawning one goroutine per task can quickly overwhelm memory, bandwidth, or CPU. A **worker pool** prevents this by:

- Limiting concurrency
- Distributing work efficiently
- Providing a scalable pattern for background tasks, job queues, and batch processors

---

## 📁 Files

- `main.go` – Full implementation with detailed comments and walkthrough
- `go.mod` – Module file (uses only standard library)

---

## 🛠️ How to Run

1. Navigate to the version folder:

```bash
cd v3-worker-pool
```

2. Run the program:

```bash
go run main.go
```

3. Output:

```
🧵 Worker 1 started: mountain.jpg
🧵 Worker 2 started: transparency.png
✅ Worker 1 completed: mountain.jpg
✅ Worker 2 completed: transparency.png
🚀 All downloads completed via worker pool.
```

---

## 💡 Key Concepts in This Version

| Feature             | Description |
|---------------------|-------------|
| `DownloadJob` struct | Defines a unit of work (URL + filename) |
| Buffered channel     | Serves as a job queue |
| `worker()` goroutines | Controlled concurrency, reusable logic |
| `sync.WaitGroup`     | Ensures all workers complete before exiting |
| Manual `Close()` calls | Proper cleanup inside loops (no `defer` misuse) |

---

## 🔄 Next Version

In [v4](../v4-checksum-verification) *(coming soon)*:
- Add **file integrity checks** using SHA256
- Implement basic **retry logic** for failed downloads

---

## 👤 Author

Guillermo Morrison  
[GitHub: cyber-mountain-man](https://github.com/cyber-mountain-man)

---

## 🪪 License

MIT – Free for personal, educational, or commercial use
```