package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RespBody struct {
	Slaves []Slave `json:"slaves"`
}

type Slave struct {
	ID       string `json:"id"`
	Hostname string `json:"hostname"`
}

// GET http://master-node:80/mesos/slaves
func mesosHosts(masterIP string) (hostnames []string, err error) {
	apiURL := fmt.Sprintf("http://%s/mesos/slaves", masterIP)

	data, err := httpGetSlaves(apiURL)
	if err != nil {
		return
	}

	respBody := RespBody{}
	err = json.Unmarshal(data, &respBody)

	for _, s := range respBody.Slaves {
		hostnames = append(hostnames, s.Hostname)
	}
	return
}

func httpGetSlaves(apiURL string) (body []byte, err error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	return
}
