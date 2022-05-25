package models

import (
	"gopkg.in/validator.v2"
)

type Backup struct {
	Dashboard string `json:"dashboard" validate:"nonzero"`
	Url       string `json:"url" validate:"nonzero`
	Login     string `json:"login" validate:"nonzero,min=8,max=8"`
	Password  string `json:"password" validate:"nonzero,min=5"`
	Sigla     string `json:"sigla" validate:"min=3,max=3"`
}

func ValidaBackup(backup *Backup) error {
	if err := validator.Validate(backup); err != nil {
		return err
	}
	return nil
}
