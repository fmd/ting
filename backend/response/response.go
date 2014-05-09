package response

//R is the type that all calls to Backend functions return.
type R struct {
	Data   interface{}
	Status string
	Code   int
	Error  error
}

func Error(err error) *R {
	r := &R{}
	r.Error = err
	r.Status = "error"
	r.Code = 500
	return r
}

func Success(data interface{}) *R {
	r := &R{}
	r.Status = "success"
	r.Data = data
	r.Code = 200
	return r
}

func Fail(data interface{}) *R {
	r := &R{}
	r.Status = "fail"
	r.Data = data
	r.Code = 400
	return r
}

//W is the struct that can be marshalled alongside a http code to a response.
//See more at http://labs.omniti.com/labs/jsend
type JSend struct {
	Data    interface{} `json:"data"`    //Wrapper around any returned data.
	Status  string      `json:"status"`  // "success" | "fail" | "error"
	Message string      `json:"message"` // Error message.
}

func (r *R) Wrap() (int, JSend) {
	resp := JSend{
		r.Data,
    	r.Status,
    	"",
	}

	if r.Error != nil {
		resp.Message = r.Error.Error()
	}

	return r.Code, resp
}