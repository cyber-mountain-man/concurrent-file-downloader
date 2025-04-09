# 📊 Version 2 – Concurrent File Downloader with Progress Bars

This version improves upon the [basic concurrent downloader (v1)](../v1-basic-downloader) by adding **real-time progress bars** to each download, giving immediate visual feedback to the user.

Built using Go’s concurrency model and the [`schollz/progressbar`](https://github.com/schollz/progressbar) package.

---

## 🚀 What It Does

- Downloads multiple files **in parallel**
- Uses `goroutines` + `sync.WaitGroup` to manage concurrency
- Shows **live progress bars** in the terminal for each file being downloaded
- Automatically names and saves the downloaded files

---

## 📦 Dependencies

This version uses a Go module and one external library:

```bash
go get github.com/schollz/progressbar/v3
```

---

## 🧱 Folder Contents

- `main.go` – The full program, with progress bars and comments
- `go.mod` and `go.sum` – Go module tracking and dependency versions

---

## 🛠️ How to Run

1. Navigate to this version’s folder:

```bash
cd v2-progress-bar
```

2. Download the required dependency:

```bash
go mod tidy
```

3. Run the program:

```bash
go run main.go
```

4. Output:

You’ll see something like:

```
📥 Downloading mountain.jpg [===========------]  65%
📥 Downloading transparency.png [======--------]  42%
✅ Download completed: https://... → mountain.jpg
✅ All downloads with progress bars completed.
```

---

## 🧠 What This Teaches

| Concept                  | Description |
|--------------------------|-------------|
| External dependency use  | Using `go get` and `go.mod` for library management |
| `goroutines`             | Launching concurrent tasks |
| `sync.WaitGroup`         | Waiting for all goroutines to finish |
| `http.Head()`            | Pre-check file size for progress tracking |
| `io.MultiWriter()`       | Write to file **and** stream progress updates simultaneously |

---

## 📘 Next Version (Coming Soon)

In [v3](../v3-worker-pool), we’ll build a **worker pool** to control how many files are downloaded at once, making the program more bandwidth- and CPU-friendly.

---

## 👤 Author

Guillermo Morrison  
[GitHub: cyber-mountain-man](https://github.com/cyber-mountain-man)

---

## 🪪 License

MIT – free for personal, educational, or commercial use
```