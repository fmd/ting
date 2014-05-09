package response

//R is the struct that all calls to Backend functions returns.
type R struct {
	Data   interface{}
	Status string
	Error  error
}

func Error(err error) *R {
	r := &R{}
	r.Error = err
	r.Status = "error"
	return r
}

func Success(data interface{}) *R {
	r := &R{}
	r.Status = "success"
	r.Data = data
	return r
}

func Fail(data interface{}) *R {
	r := &R{}
	r.Status = "fail"
	r.Data = data
	return r
}
