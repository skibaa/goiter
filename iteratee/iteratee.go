package iteratee

import (
    "github.com/skibaa/goiter/future"
)

type Input interface {
}

type InputEOF struct {}
type InputEmpty struct {}
type InputEl struct {
    e interface {}
}
func (this InputEl) E() interface{} {
    return this.e
}
func MakeInputEl(e interface{}) InputEl {return InputEl{e}}

type Step interface {
}

type Done struct {
    a interface {}
    e Input
}
func (this Done) Fold(folder Folder) future.Future {
    return folder(this)
}

type Cont struct {
    k func (input Input) Iteratee
}
func (this Cont) Fold(folder Folder) future.Future {
    return folder(this)
}
func (this Cont) K() func (input Input) Iteratee {return this.k}
func MakeCont(k func (input Input) Iteratee) Cont {return Cont{k}}

type Error struct {
    msg string
    input Input
}
func (this Error) Fold(folder Folder) future.Future {
    return folder(this)
}

type Folder func(step Step) future.Future

type Iteratee interface {
    Fold(Folder) future.Future
}

