package csv

import (
	"log"
	"testing"
	"time"

	// "mime/multipart"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	type A struct {
		Name     string    `csv:"0"`
		Address  string    `csv:"1"`
		BirthDay time.Time `csv:"2"`
		Age      int       `csv:"3"`
		Other    struct {
			Name      string    `csv:"1"`
			Weight    float64   `csv:"4"`
			BirthDay2 time.Time `csv:"5;2006-01-02"`
		} `csv:"-"`
	}

	type B []A

	a := A{}

	err := New().Unmarshal(&a)
	log.Println(a)

	assert.Equal(t, nil, err, "should not dead lock")

}
