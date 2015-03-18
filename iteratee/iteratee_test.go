package iteratee_test
import (
    "testing"
    "github.com/skibaa/goiter/iteratee"
    "github.com/skibaa/goiter/enumerator"
)

func consume(data []int) iteratee.Iteratee {
    return iteratee.MakeCont{
        func (input iteratee.Input) iteratee.Iteratee {
            switch input := input.(type) {
                case iteratee.InputEl: return consume(append(data, input.E()))
                case iteratee.InputEOF: return iteratee.Done{data, iteratee.InputEmpty{}}
                case iteratee.InputEmpty: return consume(data)
                default: panic(nil); return nil
            }
        }}
}

func sum(acc int) iteratee.Iteratee {
    return iteratee.MakeCont{
        func (input iteratee.Input) iteratee.Iteratee {
            switch input := input.(type) {
                case iteratee.InputEl: return consume(acc + int(input.E()))
                case iteratee.InputEOF: return iteratee.Done{acc, iteratee.InputEmpty{}}
                case iteratee.InputEmpty: return consume(acc)
                default: panic(nil); return nil
            }
        }}
}

func TestConsumingIteratee(t *testing.T) {
    summer := sum(0)
    fResult := enumerator.Enumerate([]int{1, 2, 3}).Apply(summer)
    if !fResult.IsCompleted() {
        t.Error("expected completed")
    }
    result := fResult.Value()
    if result != 6 {
        t.Error("expected 6")
    }
}