package models

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/uere/grafana-backup/utils"
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

type metaDados struct {
	CanAdmin              bool   `json:"canAdmin"`
	CanEdit               bool   `json:"canEdit"`
	CanSave               bool   `json:"canSave"`
	CanStar               bool   `json:"canStar"`
	Created               string `json:"created"`
	CreatedBy             string `json:"createdBy"`
	Expires               string `json:"expires"`
	FolderId              string `json:"folderId"`
	FolderTitle           string `json:"folderTitle"`
	FolderUrl             string `json:"folderUrl"`
	HasAcl                bool   `json:"hasAcl"`
	IsFolder              bool   `json:"isFolder"`
	Provisioned           string `json:"provisioned"`
	ProvisionedExternalId string `json:"provisionedExternalId"`
	Slug                  string `json:"slug"`
	Type                  string `json:"type"`
	Updated               string `json:"updated"`
	UpdatedBy             string `json:"updatedBy"`
	Url                   string `json:"url"`
	Version               string `json:"version"`
}

// godoc conectar na api do grafana recebido passando a autencicacao recebida e devolve uma lista de dashboards encontradas nesse grafana

func ListDashboards(g *Grafana) []Dashboard {
	var ListDashboards []Dashboard
	req, err := http.NewRequest("GET", g.Url+"/api/search?dashboardIds", nil)
	if err != nil {
		log.Println("Error on newrequest.\n[ERROR] -", err)
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	autorizacao := base64.StdEncoding.EncodeToString([]byte(g.Login + ":" + g.Password))
	fmt.Println(autorizacao)
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
	return ListDashboards
}

func GetDashboards(g *Grafana, d []Dashboard) {

	for _, v := range d {

		req, _ := http.NewRequest("GET", g.Url+"/api/dashboards/"+v.Uri, nil)
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		autorizacao := base64.StdEncoding.EncodeToString([]byte(g.Login + ":" + g.Password))
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
		if err := utils.MakeDir("dashboards/general"); err != nil {
			log.Println("Error on create directory", err)
		}
		nomearquivo := strings.Split(v.Uri, "/")
		if v.Type != "dash-folder" {
			if err := utils.CreateFile("dashboards/" + nomearquivo[1] + ".json"); err != nil {
				log.Println("Error on create directory", err)
			}

		}
		arquivo := "dashboards/" + nomearquivo[1] + ".json"
		if err = ioutil.WriteFile(arquivo, []byte(body), 0644); err != nil {
			log.Println("Error on save File.\n[ERROR] -", err)
		}
		//ler o arquivo e move-lo para pasta
		var dash map[string]interface{}
		file, _ := ioutil.ReadFile(arquivo)
		json.Unmarshal([]byte(file), &dash)
		// fmt.Println(x)
		for k, v := range dash {
			if k == "meta" {
				stringmeta, _ := json.MarshalIndent(v, "", "  ")
				var meta metaDados
				// fmt.Println(string(stringmeta))
				json.Unmarshal(stringmeta, &meta)
				if !meta.IsFolder {
					if err := utils.MakeDir("dashboards/" + meta.FolderTitle); err != nil {
						log.Println("Error on create directory", err)
					}
					os.Rename(arquivo, "dashboards/"+meta.FolderTitle+"/"+nomearquivo[1]+"-"+meta.Version+".json")
				} else {
					os.Remove(arquivo)
				}
			}
		}
	}
}
