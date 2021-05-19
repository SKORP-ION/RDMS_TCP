package utils

import "os"

func GetServerAddress() string {
	return os.Getenv("tcp_server_host") + ":" + os.Getenv("tcp_server_port")
}
