package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

// findIP makes a dummy connection in order to determine the user's local IP
func findIP() string {
	// Get preferred outbound ip of this machine
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ip := conn.LocalAddr().String()
	ip = ip[:strings.IndexRune(ip, ':')+1]

	return "http://" + ip + *port
}

// generateNav generates a NavContent array for use in the page header
// Ex ) /home/evan/Downlaods -> "home & /home", "evan & /home/evan",
//                              "Downloads & /home/evan/Downloads"
func generateNav(path string) []NavContent {

	path = ensureDirPath(path)

	spl := strings.Split(path, "/")
	out := make([]NavContent, len(spl))
	for i, val := range spl {
		if i == 0 {
			out[i].Name = "Base"
			out[i].Link = "/"
			continue
		}
		for j := 0; j <= i; j++ {
			out[i].Link += spl[j] + "/"
		}
		out[i].Name = val
	}
	// ! Hack to make it work. Need to rethink parsing logic.
	return out[:len(out)-1]
}

// getDirSize returns the total size of the contents of a directory
func getDirSize(path string) int64 {

	path = ensureDirPath(path)

	var size int64 = -1
	filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size
}

func getFormattedSize(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

// Ensure that the given path coresponds to a directory (trailing backslash)
// return it as is or fix it first
func ensureDirPath(path string) string {
	finalChar := path[len(path)-1]
	if finalChar == '/' {
		return path
	}

	return path + "/"
}
