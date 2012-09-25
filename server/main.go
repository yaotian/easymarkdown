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
	"log"
	"io/ioutil"
	"os"
)

func ListDir(dir string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dir)
}


func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fi, err := ListDir(".")
	if err != nil {
		fmt.Println("Error", err)
	}
	var result = "";
	for _, f := range fi {
//		d := "-"
		if f.IsDir() { continue }
		result += f.Name()+"<br>"
//		fmt.Print(result)
//		fmt.Fprintf("=========:"+f.Name()+"\n")
//		fmt.Printf("%s %o %d %s %s\n", f.Mode() & 0777, f.Size(), f.ModTime().Format("Jan 2 15:04"), f.Name())
	}

	fmt.Fprint(w,result)

}

func main() {
	port := "9090"
//	folder :="."

	if len(os.Args) > 1 {
		port = strings.Join(os.Args[1:2], "")

//		if x := strings.Join(os.Args[2:3], ""); x != "" {
//			folder = x
//		}
	}

	http.HandleFunc("/", index) //设置访问的路由
	err := http.ListenAndServe(":"+port, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

