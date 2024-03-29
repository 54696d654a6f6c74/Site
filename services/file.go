package services

import (
	"io/fs"
	"syscall"
	"time"
)

func GetFileChangeTime(fileInfo fs.FileInfo) time.Time {
	ct := fileInfo.Sys().(*syscall.Stat_t).Ctim
	return time.Unix(int64(ct.Sec), int64(ct.Nsec))
}

func FilterForFiles(entries []fs.DirEntry) []fs.DirEntry {
	filtered := []fs.DirEntry{}

	for _, entry := range entries {
		if entry.IsDir() != true {
			filtered = append(filtered, entry)
		}
	}

	return filtered
}

func isHidden(folder fs.DirEntry) bool {
	return folder.Name()[0] == '.'
}

func FilterForFolders(entries []fs.DirEntry, incHidden bool) []fs.DirEntry {
	filtered := []fs.DirEntry{}

	for _, entry := range entries {
		if !incHidden && isHidden(entry) {
			continue
		}

		if entry.IsDir() != true {
			continue
		}

		filtered = append(filtered, entry)
	}

	return filtered
}
