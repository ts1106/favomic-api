package response

type Error struct {
	Code    int      `json:"code,omitempty"`
	Message string   `json:"message,omitempty"`
	Errors  []Errors `json:"errors,omitempty"`
}

type Errors struct {
	Message string `json:"message,omitempty"`
}

func Err(code int, err error) Error {
	return Error{
		Code:    code,
		Message: err.Error(),
	}
}

func Errs(code int, errs []error) Error {
	var errors []Errors
	for _, e := range errs {
		errors = append(errors, Errors{Message: e.Error()})
	}
	return Error{
		Code:   code,
		Errors: errors,
	}
}
