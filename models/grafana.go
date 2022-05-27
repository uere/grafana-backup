package models

import (
	"gopkg.in/validator.v2"
)

type Grafana struct {
	Url        string `json:"url" validate:"nonzero`
	Login      string `json:"login" validate:"nonzero,min=8,max=8"`
	Password   string `json:"password" validate:"nonzero,min=5"`
	Project    string `json:"project"`
	Dashboards int    `json:"dashboards"`
}

func ValidateGrafana(g *Grafana) error {
	if err := validator.Validate(g); err != nil {
		return err
	}
	return nil
}
