package utils

import (
	"log"
	"os"
)

var GrafanaUrl string
var AdminUser string
var AdminPassword string

func SetAdminUser() {
	if AdminUser = os.Getenv("GF_SECURITY_ADMIN_USER"); AdminUser == "" {
		AdminUser = "admin"
	}
}

func SetAdminPassword() {
	if AdminPassword = os.Getenv("GF_SECURITY_ADMIN_PASSWORD"); AdminPassword == "" {
		AdminPassword = "admin123"
	}

}

func SetGrafanaUrl() {
	if GrafanaUrl = os.Getenv("GRAFANA_URL"); GrafanaUrl == "" {
		GrafanaUrl = "http://localhost:3000"
	}
}

func MakeDir(name string) error {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(name, 0755)
		if errDir != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func CreateFile(name string) error {
	emptyFile, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println(emptyFile)
	emptyFile.Close()
	return nil
}
