package handler

import (
	"RDMS_TCP/structures"
	"RDMS_TCP/utils"
	"io"
	"log"
	"net"
)

const BUFFERSIZE = 4096

func SendPackage(conn net.Conn, pkg structures.Package) error {
	data, err := utils.ReadFile(pkg.PathToFile)

	if err != nil {
		log.Println(err)
		return err
	}

	sendBuffer := make([]byte, BUFFERSIZE)

	for {
		_, err := data.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		_, err = conn.Write(sendBuffer)
		if err != nil {
			return err
		}
	}
	return nil
}
