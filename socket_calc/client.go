package main

import (
	//"bytes"
	"encoding/json"
	"net"
	"fmt"
)
const (
    addr = "127.0.0.1:8081"
)
type coordinate struct {
    OP string 
    A float64 
    B float64
}
func main() {
	
	conn, err := net.Dial("tcp", addr)
    if err != nil {
        fmt.Println("连接服务端失败:", err.Error())
        return
    }
    defer conn.Close()
	for{
		fmt.Println("client ready")
		n, err := conn.Read(buffer)
		if err!=nil {
			fmt.Println(n)
		}
		switch op:= client_op; op{
			case "add": ans = add(a,b)
			case "sub": ans = sub(a,b)
			case "mul": ans = mul(a,b)
			case "div": ans = div(a,b)
			default: add(a,b)
		}
		//conn.Write([]byte(json.NewEncoder(buf).Encode(ans)))	
    	fmt.Println(ans)    
	}
}
func Client(conn net.Conn){
    sms := make([]byte, 1024)
    for {
        fmt.Print("client ready")
        _, err := fmt.Scan(&sms)
        if err != nil {
            fmt.Println("数据输入异常:", err.Error())
        }
        conn.Write(sms)
        buf := make([]byte, 1024)
        c, err := conn.Read(buf)
        if err != nil {
            fmt.Println("读取服务器数据异常:", err.Error())
        }
        fmt.Println(string(buf[0:c]))
    }
}
func add(a,b float64)float64{return a+b}
func sub(a,b float64)float64{return a-b}
func mul(a,b float64)float64{return a*b}
func div(a,b float64)float64{return a/b}