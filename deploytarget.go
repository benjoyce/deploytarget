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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var env = r.URL.Query().Get("env")
		var app = r.URL.Query().Get("app")

		fmt.Printf("env="+env+"\n")
		fmt.Printf("app="+app+"\n")

		file, e := ioutil.ReadFile("./servers.json")
		if e != nil {
			fmt.Printf("File error: %v\n", e)
			os.Exit(1)
		}

		var servers ServerList
		json.Unmarshal(file, &servers)

		serverResult := make([]Server,0)

		for i := range servers.Servers {
			//fmt.Printf("%s\n",servers.Servers[i].Hostname)
			for j := range servers.Servers[i].Applications {
				//fmt.Printf("-%s\n",servers.Servers[i].Applications[j].Id)
				for k := range servers.Servers[i].Applications[j].Environments {
					//fmt.Printf("--%s\n",servers.Servers[i].Applications[j].Environments[k].Id)
					if servers.Servers[i].Applications[j].Id == app {
						if servers.Servers[i].Applications[j].Environments[k].Id == env {
							//fmt.Printf("Target:%s\n",servers.Servers[i].Applications[j].Environments[k].Id)
							applications := make([]Application,0)
							applications = append(applications,servers.Servers[i].Applications[j])
							newServer := Server{Hostname:servers.Servers[i].Hostname,Applications:applications}
							serverResult = append(serverResult,newServer)
						}
					}
				}
			}
		}

		json.NewEncoder(w).Encode(serverResult)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
