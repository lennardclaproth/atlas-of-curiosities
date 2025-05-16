package main

import (
	"fmt"
	"syscall"
)

const (
	HOST = "localhost"
	PORT = "9045"
	TYPE = "TCP"
)

func main() {
	socket, err := openSock()
	if err != nil {
		fmt.Println("Error creating socket:", err)
	}
	defer closeSock(socket)

	fmt.Printf("Socket returned with value %v\n", socket)
}

func openSock() (syscall_windows.Handle, error) {
	sock, err := syscall_windows.Socket(syscall_windows.AF_INET, syscall_windows.SOCK_STREAM, syscall_windows.IPPROTO_TCP)
	if err != nil {
		return 0, err
	}
	return sock, nil
}

func closeSock(socket syscall_windows.Handle) error {
	err := syscall_windows.Close(socket)
	return err
}

// func test(){
// 	var d syscall.WSAData
// 	e := syscall.WSAStartup(uint32(0x202), &d)
// 	if e != nil {
// 		initErr = e
// 	}
// }
