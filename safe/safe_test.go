package safe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	b := Block{
		Try: func() error {
			return nil
		},
	}
	b2 := Block{}

	b.Do()
	b2.Do()

	assert.Equal(t, "", "", "should not dead lock")

}
