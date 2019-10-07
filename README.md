# Organize files by month

Copies or moves the files from given folder to destination folder by grouping them by month.

## Usage

### Options

* src  - Source Folder
* dst - Destination Folder
* copy - Copies items if true, otherwise moves.

### Copy 

```go
go run main.go -src /Users/emrekarabacak/Desktop/source -dst /Users/emrekarabacak/Desktop/dest -copy=true
```

### Move

```go
go run main.go -src /Users/emrekarabacak/Desktop/source -dst /Users/emrekarabacak/Desktop/dest -copy=false
```