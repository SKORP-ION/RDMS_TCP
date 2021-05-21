package handler

import (
	"RDMS_TCP/structures"
	"RDMS_TCP/utils"
	"bufio"
	"io"
	"log"
	"net"
)

const BUFFERSIZE = 4096

func SendPackage(conn net.Conn, pkg structures.Package) error {
	file, err := utils.OpenFile(pkg.PathToFile)
	defer file.Close()

	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)

	if err != nil {
		log.Println(err)
		return err
	}

	sendBuffer := make([]byte, BUFFERSIZE)
	for {
		_, err = reader.Read(sendBuffer)

		if err == io.EOF {
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
