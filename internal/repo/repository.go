package repo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type GitignoreLang string

const (
	Go     GitignoreLang = "Go"
	Rust   GitignoreLang = "Rust"
	Python GitignoreLang = "Python"
)

type License string

const (
	MIT       License = "mit"
	LGPLV3    License = "lgpl-3.0"
	MPLV2     License = "mpl-2.0"
	APLV3     License = "agpl-3.0"
	Unlicense License = "unlicense"
	ApacheV2  License = "apache-2.0"
	GPLV3     License = "gpl-3.0"
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
	Name              string        `json:"name"`
	Desc              string        `json:"description"`
	AutoInit          bool          `json:"auto_init"`
	LicenseTemplate   License       `json:"license_template"`
	GitignoreTemplate GitignoreLang `json:"gitignore_template"`
}

func NewRepository(name, desc string, autoInit bool, license License, gitignore GitignoreLang) *Repository {
	return &Repository{
		Name:              name,
		Desc:              desc,
		AutoInit:          autoInit,
		LicenseTemplate:   license,
		GitignoreTemplate: gitignore,
	}
}

func (r Repository) marshal() ([]byte, error) {
	payload, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (r Repository) Publish() (*RepoResponse, error) {
	data, err := r.marshal()
	if err != nil {
		return nil, err
	}

	// TODO: add token to this request
	resp, err := http.Post(UserRepoEP, AppJSON.Value(), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response RepoResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
