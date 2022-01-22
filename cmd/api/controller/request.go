package controller

type CreateRequest struct {
	Token       string `json:"token"`
	Symbol      string `json:"symbol"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Note        string `json:"note"`
}

type ListRequest struct {
	Token       string 	`json:"token"`
	Limit		int 	`json:"limit"`
}

type RedirectRequest struct {
	Token		string `json:"token"`
	Symbol		string `json:"symbol"`
}