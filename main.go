package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
)


var server, _ = net.ResolveTCPAddr("tcp4", "45.141.101.58:502")
//var server, _ = net.ResolveTCPAddr("tcp4", "0.0.0.0:9090")

func main() {
	m := []byte("my name is han rui da !")
	sendMessage(m)
	//read(10 * 8)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func sendMessage(m []byte) int {
	//建立tcp连接
	conn, err := net.DialTCP("tcp4", nil, server)
	checkError(err)
	defer conn.Close()

	byteLen := len(m)
	//TODO max len judge
	if byteLen > 0xFF{
		return 0
	}
	ceils := byteLen
	if byteLen% 2 !=0 {
		ceils ++
	}//5/2 = 2 +1 =3
	buf := make([]byte, binary.MaxVarintLen64)

	binary.PutUvarint(buf, uint64(ceils))
	//向服务端发送数据
	temp := []byte{
		0xFF,
		0x10,  // 功能码
		0x00, 0x00,
		buf[1], buf[0], // 多少个 byte
		byte(byteLen), //5个byte
	}

	totalLen := byteLen + len(temp)
	totalLenBuf := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(totalLenBuf, uint64(totalLen))

	send := []byte{
		0x01, 0x01,
		0x00, 0x00,
		totalLenBuf[1], totalLenBuf[0],
	}
	send = append(append(send, temp...), m...)
	fmt.Println(send)
	_, err = conn.Write(send)
	checkError(err)
	resp := make([]byte, 12)
	_, err = conn.Read(resp)
	checkError(err)
	fmt.Println(resp)
	return totalLen
}

func read(length int) {
	//建立tcp连接
	conn, err := net.DialTCP("tcp4", nil, server)
	checkError(err)
	defer conn.Close()

	totalLenBuf := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(totalLenBuf, uint64(length))

	send := []byte{
		0x01, 0x02,
		0x00, 0x00,
		0x00, 0x06,
		0xFF,
		0x01,
		0x00, 0x00,
		totalLenBuf[1], totalLenBuf[0],
	}
	fmt.Println(send)
	_, err = conn.Write(send)
	checkError(err)

	resp := make([]byte, 9 + length/8)
	_, err = conn.Read(resp)
	checkError(err)
	fmt.Println(resp)
	fmt.Println(string(resp[9:]))
}
