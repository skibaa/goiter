package future
import "sync/atomic"

type CompletionHandler func(interface{})
type Future interface {
    OnComplete(CompletionHandler)
    IsCompleted() bool
    Value() interface{}
}

type successful struct {
    result interface {}
}
func (this successful) OnComplete(f CompletionHandler) {
    f(this.result)
}
func (this successful) IsCompleted() bool {return true}
func (this successful) Value() interface{} {return this.result}


func Successful(value interface{}) Future {
    return successful{value}
}

type atomicBacked struct {
    isComplete int32
    value interface {}
    handlers [] CompletionHandler
}

func Async(f func () interface {}) Future {
    res := atomicBacked{}
    go func () {
        fRes := f
        atomic.StoreInt32(&res.isComplete, 1)
        for handler := range res.handlers
    } ()
    return res
}