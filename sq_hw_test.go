package square

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcSquare(t *testing.T) {
	sidesNums := []int{0, 3, 4, 5}
	sideLens := []float64{1, 2, 3, 4, 5, 6.0, 7.1, 5.99}
	w := []float64{1.5707963267948966, 6.283185307179586, 14.137166941154069, 25.132741228718345, 39.269908169872416, 56.548667764616276, 79.18384283373074, 56.36032928503358, 0.4330127018922193, 0.8660254037844386, 1.299038105676658, 1.7320508075688772, 2.1650635094610964, 2.598076211353316, 3.074390183434757, 2.5937460843343936, 2, 4, 6, 8, 10, 12, 14.2, 11.98, 0, 0, 0, 0, 0, 0, 0, 0}
	i := 0
	var myFync func(float64, int) float64
	assert.NotEqual(t, reflect.TypeOf(myFync), reflect.TypeOf(CalcSquare))
	for _, n := range sidesNums {
		for _, l := range sideLens {
			want := w[i]
			got := CalcSquare(l, n)
			assert.Equal(t, want, got)
			i++
		}
	}

}
