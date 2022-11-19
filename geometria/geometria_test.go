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

func TestPerimetro(t *testing.T)  {
	triangulo := Triangulo{
		Lado_1: 12.0,
		Lado_2: 6.0,
		Lado_3: 6.0,

	}
	got:= Perimetro(triangulo)
	want := 24.0
	assert.Equal(t, got, want)
}