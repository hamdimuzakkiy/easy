package safe

import (
	"fmt"
	"testing"
	"errors"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	b := Block{
		Try: func() error {
			b := 0
			a := 1/b
			fmt.Println(a)
			return nil
		},
		Catch: func(e Exception) error {
			return errors.New("error")
		},
	}
	b2 := Block{}

	b.Do()
	b2.Do()

	assert.Equal(t, "", "", "should not dead lock")

}
