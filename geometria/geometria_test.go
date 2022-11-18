package geometria

import ( 
	"testing"
 	"github.com/stretchr/testify/assert"
)
func TestArea(t *testing.T) {

	got := Area(12.0, 6)
	want := 36.0
	assert.Equal(t, got, want)


}

