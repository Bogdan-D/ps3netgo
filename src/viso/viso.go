package viso

import (
	"errors"
	"os"
	"path/filepath"
	"sfo"
	"strings"
)

const (
	MAX_PATH              = 2048
	MULTIEXTENT_PART_SIZE = 0xFFFFF800
	TEMP_BUF_SIZE         = (4 * 1024 * 1024)
	FS_BUF_SIZE           = (32 * 1024 * 1024)
)

type VISO struct {
	ps3mode    bool
	path       string
	volumeName string
	title_id   string
}

func Open(path string, ps3mode bool) (*VISO, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, errors.New("VISO can open only directory, file is given: " + path)
	}

	var viso = &VISO{path: path, ps3mode: ps3mode}

	if ps3mode {
		viso.title_id, err = viso.getTitleId()
		if err != nil {
			return nil, err
		}
		viso.volumeName = "PS3VOLUME"
	} else {
		//TODO:
	}

	return viso, nil
}

func (iso *VISO) Read(buf []byte) (int, error) {
	return 0, nil
}

func (iso *VISO) Seek(offset int64, start int) (int64, error) {
	return 0, nil
}

func (iso *VISO) ReadAt(buf []byte, offset int64) (int, error) {
	return 0, nil
}

func (iso *VISO) Stat() (os.FileInfo, error) {
	return nil, nil
}

func (iso *VISO) Close() error {
	return nil
}

func (iso *VISO) getTitleId() (string, error) {
	var title_id string

	path := filepath.Join(iso.path, "PS3_GAME", "PARAM.SFO")

	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		return "", err
	}

	result, err := sfo.ParseSfo(f)

	for _, entry := range result.Entries {
		if entry.Key != "TITLE_ID" {
			continue
		}
		title_id = strings.Trim(string(entry.Data), "\x00")
		break
	}

	if title_id == "" {
		err = errors.New("TITLE_ID is not found")
	}

	return title_id, err
}
