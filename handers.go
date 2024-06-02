package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/icfoxy/GoTools"
)

func TestAlive(w http.ResponseWriter, r *http.Request) {
	GoTools.RespondByJSON(w, 200, "I AM ALIVE")
}

func GetNodes(w http.ResponseWriter, r *http.Request) {
	nodeNum := 0
	GoTools.DBGet("/main", "nodeNum", &nodeNum)
	defer func() {
		GoTools.DBPut("/main", "node"+fmt.Sprint(nodeNum), r.RemoteAddr)
		GoTools.DBPut("/main", "nodeNum", nodeNum+1)
		log.Println(r.RemoteAddr, "added")
	}()
	if nodeNum == 0 {
		GoTools.RespondByErr(w, 801, "no nodes avaliable", "high")
		return
	}
	result := make([]string, nodeNum)
	for i := 0; i < nodeNum; i++ {
		GoTools.DBGet("/main", "node"+fmt.Sprint(i), &result[i])
		if result[i] == r.RemoteAddr {
			GoTools.RespondByErr(w, 801, "already in the list", "high")
		}
	}
	GoTools.RespondByJSON(w, 200, result)
}
