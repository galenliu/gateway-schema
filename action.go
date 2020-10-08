package gateway_schema

//Action Affordance

type Action struct {
	ID           string      `json:"id"`
	AtType       string      `json:"@type"`
	Title        string      `json:"title"`
	Titles       string      `json:"titles"`
	Description  string      `json:"description"`
	Descriptions []string    `json:"descriptions"`
	UriVariables interface{} `json:"uri_variables"`
}
