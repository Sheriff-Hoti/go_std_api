package dto

type Request struct {
	First  string `json:"first"`
	Second string `json:"second"`
}

type Response struct {
	Result string `json:"result"`
}
