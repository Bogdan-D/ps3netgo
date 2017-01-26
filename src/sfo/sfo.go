package sfo

import (
	"bufio"
	"encoding/binary"
	"io"
	"log"
	"strings"
)

type sfo_header struct {
	Magic            [4]byte /** Always \x00PSF */
	Version          [4]byte /** Usually 1100 */
	Key_table_start  uint32  /** Start offset of key_table */
	Data_table_start uint32  /** Start offset of data_table */
	Tables_entries   uint32  /** Number of entries in all tables */
}

type sfo_index_table_entry struct {
	Key_offset   uint16 /*** param_key offset (relative to start offset of key_table) */
	Data_fmt     uint16 /*** param_data data type */
	Data_len     uint32 /*** param_data used bytes */
	Data_max_len uint32 /*** param_data total bytes */
	Data_offset  uint32 /*** param_data offset (relative to start offset of data_table) */
}

type sfo_entry struct {
	Key  string
	Data []byte
}

type sfo struct {
	Header  sfo_header
	Entries []sfo_entry
}

func ParseSfo(r io.ReadSeeker) (*sfo, error) {
	var result = &sfo{}
	var err error

	err = binary.Read(r, binary.LittleEndian, &result.Header)
	if err != nil {
		return nil, err
	}

	var index_count = int(result.Header.Tables_entries)
	var indexes = make([]sfo_index_table_entry, index_count)

	err = binary.Read(r, binary.LittleEndian, &indexes)
	if err != nil {
		return nil, err
	}

	for _, index := range indexes {
		entry, err := readSfoEntry(r, result.Header, index)
		if err != nil {
			break
		}
		result.Entries = append(result.Entries, entry)
	}

	return result, err
}

func readSfoEntry(r io.ReadSeeker, header sfo_header, index sfo_index_table_entry) (entry sfo_entry, err error) {
	var key_offset = header.Key_table_start + uint32(index.Key_offset)
	var data_offset = header.Data_table_start + index.Data_offset

	r.Seek(int64(key_offset), 0)

	var buf = bufio.NewReader(r)
	entry.Key, err = buf.ReadString('\x00')
	if err != nil {
		return
	}
	entry.Key = strings.Trim(entry.Key, "\x00")

	r.Seek(int64(data_offset), 0)

	log.Print(index.Data_fmt)

	entry.Data = make([]byte, index.Data_len)
	_, err = r.Read(entry.Data)
	return entry, err
}
