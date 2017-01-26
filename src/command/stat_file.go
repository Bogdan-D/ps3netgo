package command

import (
	"encoding/binary"
	"io"
	"os"
	"path/filepath"
)

type stat_file_result struct {
	Size      int64 // Files: file size, directories: 0, error: -1
	Mtime     int64
	Ctime     int64
	Atime     int64
	Directory int8
}

type stat_file_message struct {
	Cmd     uint16
	DataLen uint16
	Pad     [12]byte
}

func StatFile(r io.Reader) (result stat_file_result, err error) {
	var info os.FileInfo
	var msg = &stat_file_message{}

	err = binary.Read(r, binary.BigEndian, msg)
	if err != nil {
		return
	}

	data := make([]byte, msg.DataLen)
	err = binary.Read(r, binary.BigEndian, &data)
	if err != nil {
		return
	}

	wd, _ := os.Getwd()
	path := filepath.Join(wd, string(data))

	info, err = os.Stat(path)
	if err != nil {
		result.Size = -1
		return result, nil
	}

	result.Size = info.Size()
	result.Mtime = info.ModTime().Unix()
	result.Ctime = info.ModTime().Unix()
	result.Atime = info.ModTime().Unix()
	if info.IsDir() {
		result.Directory = 1
		result.Size = 0
	}

	return result, nil
}
