package dto

type ResponseParams struct {
	StatusCode int
	Message    string `json:"message"`
	Paginate   *Paginate
	Data       any
}
