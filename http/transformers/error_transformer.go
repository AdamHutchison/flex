package transformers

type ErrorTransformer struct {
	Message string `json:"message"`
	Errors map[string]string `json:"errors"`
}
