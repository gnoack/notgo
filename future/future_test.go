package future_test

import (
	"testing"
	"time"

	"github.com/gnoack/notgo/future"
)

func CalculateExpensive(a, b int) *future.Future[int] {
	f, fulfil := future.New[int]()
	go func() {
		time.Sleep(time.Second)
		fulfil(a+b)
	}()
	return f
}

func TestNewFuture(t*testing.T) {
	f := CalculateExpensive(30, 12)
	got := f.Value()
	want := 42
	if got != want {
		t.Errorf("CalculateExpensive(30, 12).Value() = %v, want %v", got, want)
	}
}

func TestGoFuture(t*testing.T) {
	f := future.Go(func() int {
		time.Sleep(1)
		return 42
	})
	got := f.Value()
	want := 42
	if got != want {
		t.Errorf("result from future; got %v, want %v", got, want)
	}
}
