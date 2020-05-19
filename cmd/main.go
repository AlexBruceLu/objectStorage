package main

import (
	"log"
	"net/http"
	"os"
	"test/objectStorage/objects"
)

/*
* 在服务端运行HTTP服务提供的REST接口
* 该服务通过访问本地磁盘来进行对象的存取
* GET/PUT   /objects/<object_name>
 */
func main() {
	// url /objects/开头的都使用函数objects.Handler
	http.HandleFunc("/objects/", objects.Handler)
	// 监听地址从系统变量中获取
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
