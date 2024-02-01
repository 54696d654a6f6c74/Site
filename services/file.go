package services

import (
	"io/fs"
	"syscall"
	"time"
)

func GetFileCreationTime(fileInfo fs.FileInfo) time.Time {
		ct := fileInfo.Sys().(*syscall.Stat_t).Ctim
		return time.Unix(int64(ct.Sec), int64(ct.Nsec))
}
