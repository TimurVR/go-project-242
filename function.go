package goproject242
import (
	"fmt"
	"os"
	"path/filepath"
	"io/fs"
	"strings"
)
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return "", err
	}
	var totalSize int64
	if fileInfo.Mode().IsRegular(){
		totalSize = fileInfo.Size()
	} else {
		if all{
		    err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir()==false  {
				info, err := d.Info()
				if err != nil {
					return err
				}
				totalSize += info.Size()
			}
			
			return nil
		})
		}else{
			entries, err := os.ReadDir(path)
			if err != nil {
				return "", err
			}

			for _, entry := range entries {
				if strings.HasPrefix(entry.Name(), ".") {
					continue
				}
				fullPath := filepath.Join(path, entry.Name())
				fileInfo, err := os.Stat(fullPath)
				if err != nil {
					continue 
				}
				if fileInfo.Mode().IsRegular() {
					totalSize += fileInfo.Size()
				}
			}
		}
	}
	if human {
		return fmt.Sprintf("%s\t%s", formatHumanReadable(totalSize), path), err
	}
	return fmt.Sprintf("%dB\t%s", totalSize, path), err
}
func formatHumanReadable(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%dB", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%cB", float64(size)/float64(div), "KMGTPE"[exp])
}