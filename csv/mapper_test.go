package csv

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type A struct {
	Name     string    `csv:"0"`
	Address  string    `csv:"1"`
	BirthDay time.Time `csv:"2"`
	Age      int       `csv:"3"`
	Salary   int64     `csv:"3"`
	Other    struct {
		Name      string    `csv:"1"`
		Weight    float64   `csv:"4"`
		BirthDay2 time.Time `csv:"5;2006-01-02"`
		Test      float64
	} `csv:"-"`
}

func (a *A) Format() {
	a.Name = "Hi! " + a.Name
}

func TestDo(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	type B []A

	var a []A

	file, err := os.Open("test.csv")
	if err != nil {
		log.Println(err)
	}

	err = Unmarshal(file, &a)
	log.Println(a)

	assert.Equal(t, nil, err, "should not error")

	file2, err := os.Open("test2.csv")
	if err != nil {
		log.Println(err)
	}

	b := B{}
	err = Unmarshal(file2, &b)

}
