package main
import(
	"fmt"
	"net"
)
func client_write(conn net.Conn){
	defer conn.Close()
	fmt.Println("分配给协程client:",conn.RemoteAddr())
	
}
// func client_read(conn net.Conn){
// 	defer conn.Close()
// 	fmt.Println("分配给协程client:",conn.RemoteAddr())
// 	for{
// 		select {
// 		case str := <- Strchan1:
// 			_,err:=conn.Write([]byte(str))
// 			if err != nil {
// 				fmt.Println("发送失败")
// 			}
// 		default:
// 			continue			
// 		}
// 		if len(Strchan1) == 0{
// 			fmt.Println("回话结束")
// 			return
// 		}
// 		}	
// 	}
func Dealwith_client(conn net.Conn){
	defer conn.Close()	
	for{
		data := make([]byte, 1024)
		request_len,err := conn.Read(data)
		if err != nil {
			fmt.Println("客户端断开连接")
			return
		}
		Strchan1 <- string(data[:request_len])
		fmt.Println(len(Strchan1))
	}
}
// else if  string(data[:request_len]) == "1"{  
// 	client_read(conn)
// }
var Strchan1 = make(chan string,1000)
func main() {
	fmt.Println("开始设置监听")
	listener,err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil{
		fmt.Println("server监听设置出现问题")
	}
	fmt.Println("设置监听成功")
	for {
		conn,err := listener.Accept()
		if err != nil{
			fmt.Println("等待accept失败")
		}
		
		go Dealwith_client(conn)
	}
}