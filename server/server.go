package main

import (
        "net/http"
        "os"
        "strings"
)

func shareDir(dirName string,port string,ch chan bool){
        h := http.FileServer(http.Dir(dirName))
        err := http.ListenAndServe(":"+port,h)
        if err != nil {
                println("ListenAndServe : ",err.Error())
                ch <- false
        }
}

func main(){
        ch := make(chan bool)
        port := "8000"  //Default port 
        folder := "."
        if len(os.Args)>1 { 
                port = strings.Join(os.Args[1:2],"")

                if x := strings.Join(os.Args[2:3],""); x != "" {
                        folder = x
                }
                
        }
        go shareDir(folder,port,ch)
        println("Listening on port ",port,"...","folder ",folder)
        bresult := <-ch
        if false == bresult {
                println("Listening on port ",port," failed")
        }
}