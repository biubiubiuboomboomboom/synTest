package helper

import (
	"awesomeProject/defs"
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"os/exec"
)

const (
	ConstHeader         = "Headers"
	ConstHeaderLength   = 7
	ConstMLength = 4
)



//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}


//封包
func Enpack(message []byte) []byte {
	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}

//解包
func Depack(buffer []byte) []byte {
	length := len(buffer)

	var i int
	data := make([]byte, 32)
	for i = 0; i < length; i = i + 1 {
		if length < i+ConstHeaderLength+ConstMLength {
			break
		}
		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
			messageLength := BytesToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstMLength])
			if length < i+ConstHeaderLength+ConstMLength+messageLength {
				break
			}
			data = buffer[i+ConstHeaderLength+ConstMLength : i+ConstHeaderLength+ConstMLength+messageLength]

		}
	}

	if i == length {
		return make([]byte, 0)
	}
	return data
}

func Get_UserCount() int64 {
	ps := exec.Command("netstat", "-nat" , "grep","-i","8080")
	grep := exec.Command("wc", "-l")
	r, w := io.Pipe() // 创建一个管道
	defer r.Close()
	defer w.Close()
	ps.Stdout = w // ps向管道的一端写
	grep.Stdin = r // grep从管道的一端读
	var buffer bytes.Buffer
	grep.Stdout = &buffer // grep的输出为buffer

	_ = ps.Start()
	_ = grep.Start()
	ps.Wait()
	w.Close()
	grep.Wait()
	i, _ := io.Copy(os.Stdout, &buffer)
	return i
}

func ChangeUserRecode(info defs.Info,time int64,count int)  {
	defs.TestUser_record.UserLastTime = time
	defs.TestUser_record.Userinfo = info
	defs.TestUser_record.UserCount = count
}