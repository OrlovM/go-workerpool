package workerpool

type Result struct {
	Data interface{}
}

func NewResult(data interface{}) *Result {
	return &Result{data}
}