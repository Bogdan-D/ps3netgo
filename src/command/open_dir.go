package command

import (
	"encoding/binary"
	"io"
	"log"
	"os"
	"path/filepath"
)

type open_dir_message struct {
	Cmd     uint16
	DataLen uint16
	Pad     [12]byte
}

type open_dir_result struct {
	Dir     *os.File
	Success int32 // 0 success, -1 error
}

func OpenDir(r io.Reader) (result open_dir_result, err error) {
	var msg = &open_dir_message{}
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

	log.Print("Opening directory ", path)
	dir, err := os.Open(path)

	if err != nil {
		log.Print("Failed to open directory ", path, " | ", err)
		result.Success = -1

		return result, nil
	}

	info, err := dir.Stat()
	if err != nil || !info.IsDir() {
		result.Success = -1
		return result, nil
	}

	result.Dir = dir
	result.Success = 0
	return result, nil
}
