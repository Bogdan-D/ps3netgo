package command

import (
	"encoding/binary"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"viso"
)

type FileReader interface {
	io.ReadSeeker
	io.Closer
	Stat() (os.FileInfo, error)
}

type open_file_message struct {
	Cmd     uint16
	DataLen uint16
	Pad     [12]byte
}

type open_file_result struct {
	File FileReader

	Message struct {
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

	var file FileReader

	if strings.Contains(path, "/***PS3***/") {
		path = strings.Replace(path, "/***PS3***/", "", 1)
		file, err = viso.Open(path, true)
	} else {
		file, err = os.Open(path)
	}

	if err != nil {
		log.Print("Failed to open file ", path, " | ", err)
		result.Message.FileSize = -1

		return result, nil
	}

	info, err := file.Stat()
	if err != nil {
		log.Print("Failed to stat file ", path, " | ", err)
		result.Message.FileSize = -1

		return result, nil
	}

	result.Message.FileSize = info.Size()
	result.Message.Mtime = info.ModTime().Unix()
	result.File = file

	return result, nil
}
