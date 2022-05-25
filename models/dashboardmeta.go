package models

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Dashboard struct {
	Id        int      `json:"id"`
	Uid       string   `json:"uid"`
	Title     string   `json:"title"`
	Uri       string   `json:"uri"`
	Slug      string   `json:"slug"`
	Type      string   `json:"type"`
	Tags      []string `json:"tags"`
	IsStarred bool     `json:"isStarred"`
	SortMeta  int      `json:"sortMeta"`
}

func RemoveIndex(d []Dashboard, index int) []Dashboard {
	return append(d[:index], d[index+1:]...)
}

// godoc conectar na api do grafana recebido passando a autencicacao recebida e devolve uma lista de dashboards encontradas nesse grafana

func ListDashboards(b *Backup) []Dashboard {
	var ListDashboards []Dashboard
	req, err := http.NewRequest("GET", b.Url+"/api/search?dashboardIds", nil)
	if err != nil {
		log.Println("Error on newrequest.\n[ERROR] -", err)
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	autorizacao := base64.StdEncoding.EncodeToString([]byte(b.Login + ":" + b.Password))
	// fmt.Println(autorizacao)
	req.Header.Add("Authorization", "Basic "+autorizacao)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	// log.Println(string([]byte(body)))
	err = json.Unmarshal(body, &ListDashboards)
	if err != nil {
		log.Println("Error while on unMarshal:", err)
	}
	NewListDashboards := ListDashboards
	for i, dash := range ListDashboards {
		if dash.Type == "dash-folder" {
			NewListDashboards = RemoveIndex(ListDashboards, i)
		}
	}
	return NewListDashboards
}

func GetDashboards(b *Backup, d []Dashboard) {
	req, _ := http.NewRequest("GET", b.Url+"/api/dashboards/"+d[1].Uri, nil)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	autorizacao := base64.StdEncoding.EncodeToString([]byte(b.Login + ":" + b.Password))
	req.Header.Add("Authorization", "Basic "+autorizacao)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	file, _ := json.MarshalIndent(string([]byte(body)), "", " ")
	err = ioutil.WriteFile("dashboards/"+d[1].Title+".json", file, 0777)
	if err != nil {
		log.Println("Error on create File.\n[ERROR] -", err)
	}
}
