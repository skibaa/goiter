package enumerator

import (
    "github.com/skibaa/goiter/future"
    "github.com/skibaa/goiter/iteratee"
)

type Enumerator interface {
    Apply(i iteratee.Iteratee) future.Future //Future of Iteratee
}
func (this Enumerator) AndThen (e Enumerator) Enumerator {
    return andThan{this, e}
}

type andThen struct {
    first Enumerator
    second Enumerator
}
func (this andThen) Apply(i iteratee.Iteratee) future.Future {
    fFirst := this.first.Apply(i)
    fFirst.
}

type empty struct {}
func (empty) Apply(i iteratee.Iteratee) future.Future {
    return future.Successful(i)
}
var varEmpty = new(empty)
func Empty() Enumerator {return varEmpty}

func consume1(in interface {}, i iteratee.Iteratee) future.Future {
    return func (step iteratee.Step) future.Future {
    switch step := step.(type) {
        case iteratee.Cont: return step.K()(future.Successful(iteratee.MakeInputEl(in)))
        default: return future.Successful(i)
    }}
}

type enumerate1 struct {in interface{}}
func (this enumerate1) Apply(i iteratee.Iteratee) future.Future {
    return i.Fold(consume1(this.in, i))
}
func Enumerate1(i interface{}) Enumerator {return enumerate1{i}}

type enumerateMore struct {in [] interface{}}
func (this enumerateMore) Apply(i iteratee.Iteratee) future.Future {
    return enumerateSlice(this.in, i)
}

func enumerateSlice(in [] interface{}, i iteratee.Iteratee)

func EnumerateSlice(in [] interface {}) Enumerator {
    switch len(in) {
        case 0: return Empty()
        case 1: return Enumerate1(in[0])
        default: return enumerateSlice(in)
    }
}