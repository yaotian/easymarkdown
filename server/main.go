/**
 *  This is the easy markdown server
 */
package main

import (
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
	"os"
	"html/template"
)

type EasyServer struct {
	folder string
}

func (server *EasyServer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	 if r.URL.Path == "/" {
        index(w,r,server.folder)
        return
    }

	 if r.URL.Path == "/admin" {
        admin(w,r,server.folder)
        return
    }
    
    if strings.HasPrefix(r.URL.Path,"/open/") {
    	fmt.Printf("%q\n",strings.Split(r.URL.Path,"/open/")[1])
    	contents,_ := ioutil.ReadFile(server.folder+"/"+strings.Split(r.URL.Path,"/open/")[1])
    	fmt.Printf("contents"+string(contents))
    	fmt.Fprint(w,string(contents))
    	return
    }
    
    http.NotFound(w, r)
    return
}


func admin(w http.ResponseWriter, r *http.Request,folder string){
	fmt.Fprint(w,"admin")
}

func index(w http.ResponseWriter, r *http.Request,folder string){
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fi, err := ListDir(folder)
	if err != nil {
		fmt.Println("Error", err)
	}
	var result = "";
	for _, f := range fi {
		//		d := "-"
		if f.IsDir() { continue }
		result += "<a href=/open/"+template.HTMLEscapeString(f.Name())+">"+template.HTMLEscapeString(f.Name())+"</a><br>"
	}

	fmt.Fprint(w,result)
}


func ListDir(dir string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dir)
}


func main() {
	var port ="9009"
	es := &EasyServer{"."}
	if len(os.Args) > 1 {
		if x := strings.Join(os.Args[1:2], ""); x != "" {
			port = x
		}

		if x := strings.Join(os.Args[2:3], ""); x != "" {
			es.folder = x
		}
	}
	http.ListenAndServe("0.0.0.0:"+port,es)
}

