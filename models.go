package main

import "os"

// Directory is the user's current position in the filesystem
type Directory struct {
	Name         string
	Path         string
	RelativePath string
	Nav          []NavContent
	Contents     []FileInfo
}

// FileInfo contains information about a single file
type FileInfo struct {
	Info         os.FileInfo
	Link         string
	DownloadLink string
	FDate        string
	DirSize      int64
}

// NavContent contains info for the top page NavBar
type NavContent struct {
	Name string
	Link string
}
