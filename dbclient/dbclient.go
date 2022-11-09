package dbclient

import (
	"encoding/binary"
	"math"
	"net"
	"strconv"
)

type Db struct {
	Port   string
	DbHost string
}

func (t *Db) GetInt(key string) (int, error) {
	command := []byte{'G', 'E', 'T'}
	keySize := make([]byte, 4)
	binary.LittleEndian.PutUint32(keySize, uint32(len(key)))

	command = append(command, []byte{' '}...)
	command = append(command, keySize...)
	command = append(command, []byte{' '}...)
	command = append(command, []byte(key)...)

	response, err := makeRequest(t.DbHost, t.Port, command)
	if err != nil {
		return -1, err
	}

	res, err := strconv.Atoi(string(response))
	if err != nil {
		return -1, err
	}

	return res, nil
}

func (t *Db) SetInt(key string, val int) (string, error) {
	command := []byte{'S', 'E', 'T'}

	keySize := make([]byte, 4)
	binary.LittleEndian.PutUint32(keySize, uint32(len(key)))
	numDigsInVal := numOfDigitsInInt(float64(val))
	valSize := make([]byte, 4)
	binary.LittleEndian.PutUint32(valSize, uint32(numDigsInVal))

	command = append(command, []byte{' '}...)
	command = append(command, keySize...)
	command = append(command, []byte{' '}...)
	command = append(command, valSize...)
	command = append(command, []byte{' '}...)
	command = append(command, []byte(key)...)
	command = append(command, []byte{' '}...)
	command = append(command, []byte(strconv.Itoa(val))...)
	response, err := makeRequest(t.DbHost, t.Port, command)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func (t *Db) Del(key string) (string, error) {
	command := []byte{'D', 'E', 'L'}

	keySize := make([]byte, 4)
	binary.LittleEndian.PutUint32(keySize, uint32(len(key)))

	command = append(command, []byte{' '}...)
	command = append(command, keySize...)
	command = append(command, []byte{' '}...)
	command = append(command, []byte(key)...)

	response, err := makeRequest(t.DbHost, t.Port, command)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func makeRequest(dbHost string, port string, request []byte) ([]byte, error) {
	servAddr := dbHost + ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		return nil, err
	}

	_, err = conn.Write(request)
	if err != nil {
		println("Write to server failed:", err.Error())
		return nil, err
	}

	response := make([]byte, 1024)

	_, err = conn.Read(response)
	if err != nil {
		println("Write to server failed:", err.Error())
		return nil, err
	}

	// parse response
	conn.Close()
	// read size from the response and return response based on header size
	size := bytesToInt(response)
	return response[5 : 5+size], nil
}

func bytesToInt(data []byte) int {
	return int(binary.LittleEndian.Uint32(data[0:4]))
}

func numOfDigitsInInt(num float64) int {
	res := 0.0
	for num/math.Pow(10, res) > 1 {
		res++
	}
	return int(res)
}
