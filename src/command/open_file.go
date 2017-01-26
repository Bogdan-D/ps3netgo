package command

import (
	"encoding/binary"
	"io"
	"log"
	"os"
	"path/filepath"
)

type open_file_message struct {
	Cmd     uint16
	DataLen uint16
	Pad     [12]byte
}

type open_file_result struct {
	File   *os.File
	Result struct {
		FileSize int64 // -1 error
		Mtime    int64
	}
}

func OpenFile(r io.Reader) (result open_file_result, err error) {
	var msg = &open_file_message{}
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

	file, err := os.Open(path)

	if err != nil {
		log.Print("Failed to open file ", path, " | ", err)
		result.Result.FileSize = -1

		return result, nil
	}

	info, err := file.Stat()
	if err != nil {
		log.Print("Failed to stat file ", path, " | ", err)
		result.Result.FileSize = -1

		return result, nil
	}

	result.Result.FileSize = info.Size()
	result.Result.Mtime = info.ModTime().Unix()
	result.File = file

	return result, nil
}
