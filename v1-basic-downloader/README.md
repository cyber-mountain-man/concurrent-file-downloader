# 🧩 Version 1 – Basic Concurrent File Downloader

This is the first version of a multi-part Go project demonstrating **real-world concurrency**.  
It downloads multiple files concurrently using `goroutines` and `sync.WaitGroup`.

---

## 🚀 What It Does

- Takes a list of file URLs and filenames
- Spawns a goroutine for each file download
- Saves each file locally
- Waits for all downloads to complete using a `WaitGroup`

---

## 📁 Files

- `main.go` – The complete Go program with detailed inline comments
- Example URLs:
  - [Mountain Photo (Wikipedia)](https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg)
  - [PNG Transparency Demo](https://upload.wikimedia.org/wikipedia/commons/4/47/PNG_transparency_demonstration_1.png)

---

## 🛠️ How to Run

1. Navigate to this version’s folder:

```bash
cd v1-basic-downloader
```

2. Run the program:

```bash
go run main.go
```

3. You should see output like:

```bash
Downloaded: https://... → mountain.jpg
Downloaded: https://... → transparency.png
All downloads completed.
```

The files will appear in the same folder.

---

## 🧠 Concepts Demonstrated

| Feature            | Description |
|--------------------|-------------|
| `goroutines`       | Run each file download concurrently |
| `sync.WaitGroup`   | Wait for all concurrent tasks to complete |
| `http.Get`         | Make an HTTP GET request |
| `io.Copy`          | Stream data from the response to the file |
| `defer`            | Ensure cleanup (closing files and HTTP responses) |

---

## 🔄 Next Steps in Version 2

In the next version, we will:
- Add a progress indicator for each file
- Show how much time each download takes

---

## 👤 Author

Guillermo Morrison  
[GitHub: cyber-mountain-man](https://github.com/cyber-mountain-man)

---

## 🪪 License

MIT – Free for personal, educational, or commercial use
```