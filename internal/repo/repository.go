package repo

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ContentType string

const (
	AppJSON ContentType = "application/json"
)

func (ct ContentType) Value() string {
	return string(ct)
}

const UserRepoEP string = "https://github.com/user/repos"

type Repository struct {
	Name     string `json:"name"`
	Desc     string `json:"description"`
	AutoInit bool   `json:"auto_init"`
}

func NewRepository(name, desc string) *Repository {
	return &Repository{
		Name: name,
		Desc: desc,
	}
}

func (r Repository) marshal() ([]byte, error) {
	payload, err := json.Marshal(&r)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (r Repository) Publish() error {
	data, err := r.marshal()
	if err != nil {
		return err
	}

	http.Post(UserRepoEP, AppJSON.Value(), bytes.NewBuffer(data))
	return nil
}
