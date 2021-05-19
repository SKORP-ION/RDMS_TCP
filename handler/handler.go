package handler

import (
	"RDMS_TCP/database"
	"bufio"
	"log"
	"net"
)

//Обработчик всех подключений
func Accept(conn net.Conn, err error) {
	defer conn.Close()

	if err != nil {
		log.Println(err)
		return
	}

	session_key, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		log.Println(err)
		return
	}

	pkg, err := database.GetPackageByDownloadKey(session_key)

	if err != nil {
		log.Println(err)
		return
	}

	err = SendPackage(conn, pkg)

	if err != nil {
		log.Println(err)
		return
	}

	err = database.RemoveSession(session_key)
}
