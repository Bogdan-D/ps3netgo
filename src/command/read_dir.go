package command

import (
	"log"
	"os"
)

type read_dir_file struct {
	Filesize    int64
	Mtime       int64
	IsDirectory int8
	Name        [512]byte
}

type read_dir_result struct {
	Len   int64
	Files []read_dir_file
}

func ReadDir(dir *os.File) (result read_dir_result) {
	log.Print("Reading directory ", dir.Name())
	var files, err = dir.Readdir(-1)
	if err != nil {
		log.Print(err)
		result.Len = 0
		return
	}

	result.Len = int64(len(files))

	for _, info := range files {
		f := read_dir_file{}
		f.Filesize = info.Size()
		if info.IsDir() {
			f.IsDirectory = 1
			f.Filesize = 0
		}
		f.Mtime = info.ModTime().Unix()
		copy(f.Name[:], info.Name())
		result.Files = append(result.Files, f)
	}

	return result
}
