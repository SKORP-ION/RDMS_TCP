package handler

import (
	"RDMS_TCP/structures"
	"RDMS_TCP/utils"
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

const BUFFERSIZE = 4096

func SendPackage(conn net.Conn, pkg structures.Package) error {
	file, err := utils.OpenFile(pkg.PathToFile)

	reader := bufio.NewReader(file)

	if err != nil {
		log.Println(err)
		return err
	}

	sendBuffer := make([]byte, BUFFERSIZE)
	n := 0
	for {
		_, err = reader.Read(sendBuffer)
		n++

		if err == io.EOF {
			fmt.Print(n)
			break
		}

		_, err = conn.Write(sendBuffer)
		if err != nil {
			return err
		}
	}
	_, err = conn.Write([]byte("EOF"))

	if err != nil {
		return err
	}

	return nil
}
