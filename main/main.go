package main

import (
	"bytes"
	"fmt"

	"github.com/tuananhlai/gocsv"
)

type Sample struct {
	Foo string `csv:"foo"`
	Bar int    `csv:"bar"`
	Baz bool   `csv:"baz"`
}

func main() {
	b := bytes.NewBufferString(`foo,bar,baz
f,1,true
e,3,false`)
	gocsv.UnmarshalEachToCallbackWithError(bytes.NewReader(b.Bytes()), func(s Sample, err error) error {
		fmt.Println(s)
		return nil
	})
}
