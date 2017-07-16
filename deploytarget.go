package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ServerList struct {
	Servers []Server
}

type Server struct {
	Hostname     string
	Applications []Application
}

type Application struct {
	Id           string
	Environments []Environment
}

type Environment struct {
	Id    string
	State string
}

func main() {

	file, e := ioutil.ReadFile("./moreservers.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))

	var servers ServerList
	json.Unmarshal(file, &servers)
	fmt.Printf("Results: %v\n", servers)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var env = r.URL.Query().Get("env")
		var app = r.URL.Query().Get("app")
		//cmd := exec.Command("cmd", "/C", "c:/go/bin/jq-win64.exe", "-r", "[.[] | select(.environments[] | contains(\""+env+"\")) | select(.applications[] | contains(\""+app+"\") ) | .] | unique", "servers.json")
		//var out bytes.Buffer
		//var stderr bytes.Buffer
		//cmd.Stdout = &out
		//cmd.Stderr = &stderr
		//err := cmd.Run()
		//if err != nil {
		//		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		//		return
		//	}

		for i := range servers.Servers {
			for j := range servers.Servers[i].Applications {
				for k := range servers.Servers[i].Applications[j].Environments {
					if servers.Servers[i].Applications[j].Id == app {
						if servers.Servers[i].Applications[j].Environments[k].Id == env {
							fmt.Fprintf(w, servers.Servers[i].Hostname)
						}
					}
				}
			}
		}

		//fmt.Fprintf(w, out.String())
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
