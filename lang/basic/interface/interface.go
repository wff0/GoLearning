package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

var input = `
{
	"created_at": "Thu May 31 00:00:01 +0000 2012"
}
`

type TimeStamp time.Time

func (t *TimeStamp) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = TimeStamp(v)
	return nil
}

func main() {
	var val map[string]TimeStamp
	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}
	fmt.Println(val)

	for k, v := range val {
		fmt.Println(k, reflect.TypeOf(v))
	}
	fmt.Println(time.Time(val["created_at"]))
}
