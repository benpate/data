package memory

import (
	"github.com/benpate/data"
	"github.com/benpate/derp"
)

type Iterator struct {
	Data    []data.Object
	Counter int
}

func NewIterator(data []data.Object) *Iterator {
	return &Iterator{
		Data: data,
	}
}

func (iterator *Iterator) Next(output data.Object) bool {

	if iterator.Counter >= len(iterator.Data) {
		return false
	}

	populateInterface(iterator.Data[iterator.Counter], output)
	iterator.Counter = iterator.Counter + 1
	return true
}

func (iterator *Iterator) Close() *derp.Error {
	iterator.Counter = len(iterator.Data) + 1
	return nil
}
