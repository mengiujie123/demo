package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
	"io"
	_"strings"
)
func main() {
	// 创建按客户端
	conn,err := net.Dial("tcp","127.0.0.1:8080")
	if err != nil {
		fmt.Println("连接server失败")
	}else{
		fmt.Println("连接server成功")
	}
	var x int
	fmt.Println("请输入进行写: 0 读: 1")
	fmt.Scanln(&x)
	if x == 0{
	reader := bufio.NewReader(os.Stdin)
	for {
	str,err := reader.ReadString('\n')
	if err == io.EOF{
		fmt.Println("读取出现错误")
	}
	// str = strings.Trim(str,"\n\r")
	// // if str == "退出"{
	// // 	break
	// // }
	// str = str + "\n"
	_,err = conn.Write([]byte(str))
	if err != nil{
		fmt.Println("发送失败",err)
	}else{
		fmt.Println("发送成功")
	}
   }
 }
}
	// }else if x == 1{
	// 	for{
	// 	data := make([]byte, 1024)
	// 	request_len,err := conn.Read(data)
	// 	if err != nil {
	// 		fmt.Println("客户端断开连接")
	// 		return
	// 	}
	// 	fmt.Println(string(data[:request_len]))
	// 	}
	// }else{
	// 	fmt.Println("输入错误")
	// }

// }