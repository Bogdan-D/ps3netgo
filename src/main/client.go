package main

import (
	"bytes"
	"command"
	"encoding/binary"
	"io"
	"log"
	"net"
	"os"
)

type client struct {
	socket  *net.TCPConn
	dir     *os.File
	ro_file *os.File
	wo_file *os.File
}

type packet struct {
	Command uint16
	Data    [14]byte
}

func (c *client) Run() {
	for {
		var raw_msg [16]byte
		var msg = &packet{}

		_, err := c.socket.Read(raw_msg[:])
		if err != nil {
			if err == io.EOF {
				return
			}

			log.Print("Error read from socket: ", err)
			return
		}
		log.Print(raw_msg)
		binary.Read(bytes.NewReader(raw_msg[:]), binary.BigEndian, msg)

		switch msg.Command {
		case command.NETISO_CMD_OPEN_DIR:
			reader := io.MultiReader(bytes.NewReader(raw_msg[:]), c.socket)
			result, err := command.OpenDir(reader)
			if err != nil {
				log.Print("Error while processing CMD_OPEN_DIR:", err)
				return
			}

			if result.Dir != nil {
				c.dir = result.Dir
			}

			err = binary.Write(c.socket, binary.BigEndian, result.Success)
			if err != nil {
				log.Print("Error writing to socket CMD_OPEN_DIR: ", err)
				return
			}
		case command.NETISO_CMD_READ_DIR:
			result := command.ReadDir(c.dir)
			err := binary.Write(c.socket, binary.BigEndian, result.Len)
			if err != nil {
				log.Print("Error writing to socket CMD_READ_DIR: ", err)
				return
			}

			err = binary.Write(c.socket, binary.BigEndian, result.Files)
			if err != nil {
				log.Print("Error writing to socket CMD_READ_DIR: ", err)
				return
			}
		case command.NETISO_CMD_STAT_FILE:
			reader := io.MultiReader(bytes.NewReader(raw_msg[:]), c.socket)
			result, err := command.StatFile(reader)
			if err != nil {
				log.Print("Error while processing CMD_STAT_FILE:", err)
				return
			}

			err = binary.Write(c.socket, binary.BigEndian, result)
			if err != nil {
				log.Print("Error writing to socket CMD_STAT_FILE: ", err)
				return
			}
		case command.NETISO_CMD_OPEN_FILE:
			reader := io.MultiReader(bytes.NewReader(raw_msg[:]), c.socket)
			result, err := command.OpenFile(reader)
			if err != nil {
				log.Print("Error while processing CMD_OPEN_FILE")
				return
			}

			c.ro_file = result.File
			err = binary.Write(c.socket, binary.BigEndian, result.Result)
			if err != nil {
				log.Print("Error writing to socket CMD_OPEN_FILE: ", err)
				return
			}

		case command.NETISO_CMD_READ_FILE:
			type read_file_message struct {
				Cmd    uint16
				Pad    uint16
				Len    uint32
				Offset uint64
			}

			msg := &read_file_message{}
			err = binary.Read(bytes.NewReader(raw_msg[:]), binary.BigEndian, msg)
			if err != nil {
				log.Print("Error while processing CMD_READ_FILE ", err)
				return
			}

			var bytes_read int32
			var buf = make([]byte, msg.Len)

			if c.ro_file == nil {
				bytes_read = -1
			} else {
				n, err := c.ro_file.ReadAt(buf, int64(msg.Offset))
				bytes_read = int32(n)
				if err != nil {
					bytes_read = -1
				}
			}

			err = binary.Write(c.socket, binary.BigEndian, bytes_read)
			if err != nil {
				log.Print("Error write to socket CMD_READ_FILE ", err)
				return
			}
			if bytes_read > 0 {
				_, err = c.socket.Write(buf)
				if err != nil {
					log.Print("Error write to socket CMD_READ_FILE2 ", err)
					return
				}
			}
		default:
			log.Printf("Unknown command: %#x", msg.Command)
		}
	}
}
