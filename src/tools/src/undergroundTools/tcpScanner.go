package undergroundTools

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func ScanIP(target string) {
	waitGroup := sync.WaitGroup{}

	TestPort := func(port int) {
		connection, err := net.DialTimeout("tcp", target+":"+strconv.Itoa(port), time.Duration(1)*time.Second)

		if err == nil {
			fmt.Fprintf(connection, "HELLO\r\n")

			buffer := make([]byte, 0, 4096)
			tmpBuffer := make([]byte, 256)

			for {
				n, err := connection.Read(tmpBuffer)

				if err != nil {
					if err != io.EOF {
						fmt.Println("Ошибка чтения:", err.Error())
					}
					break
				}

				buffer = append(buffer, tmpBuffer[:n]...)
			}

			connection.Close()
			fmt.Println(port, " open")
		}

		waitGroup.Done()
	}

	waitGroup.Add(65534)

	for i := 0; i < 65534; i++ {
		go TestPort(i)
	}

	waitGroup.Wait()
	os.Exit(0)
}