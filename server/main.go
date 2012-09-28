/**
 * Created with IntelliJ IDEA.
 * User: yaotian
 * Date: 12-9-22
 * Time: 下午7:06
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
	"os"
)

type EasyMarkdownServer struct {
	folder string
}

func (server EasyMarkdownServer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Println("folder:",server.folder)

	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fi, err := ListDir(server.folder)
	if err != nil {
		fmt.Println("Error", err)
	}
	var result = "";
	for _, f := range fi {
		//		d := "-"
		if f.IsDir() { continue }
		result += f.Name()+"<br>"
	}

	fmt.Fprint(w,result)
}


func ListDir(dir string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dir)
}


func main() {
	var server EasyMarkdownServer
	var port ="9009"
	server.folder="."

	if len(os.Args) > 1 {
		if x := strings.Join(os.Args[1:2], ""); x != "" {
			port = x
		}

		if x := strings.Join(os.Args[2:3], ""); x != "" {
			server.folder = x
		}
	}

	http.ListenAndServe("localhost:"+port,server)
}

