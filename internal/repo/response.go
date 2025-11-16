package repo

type RepoResponse struct {
	ID    int       `json:"id"`
	Name  string    `json:"name"`
	Desc  string    `json:"description"`
	URL   string    `json:"html_url"`
	Owner OwnerInfo `json:"owner"`
}

type OwnerInfo struct {
	Name string `json:"login"`
	URL  string `json:"html_url"`
}
