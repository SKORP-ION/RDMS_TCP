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
		n, err := reader.Read(sendBuffer)

		if err == io.EOF {
			fmt.Println("EOF")
			break
		}

		_, err = conn.Write(sendBuffer[:n])
		if err != nil {
			return err
		}
	}
	fmt.Println("done")
	return nil
}
