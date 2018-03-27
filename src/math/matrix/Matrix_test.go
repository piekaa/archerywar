package matrix

import (
	"testing"
	"reflect"
	"fmt"
)

//todo more tests, test 1x4 * 4*4
//todo test panics

func TestMyMatrix_Mul_2x3_3x2(t *testing.T) {

	a := New(3,2,
		[]float32 {
			1,2,3,
			4,5,6,})

	b:= New(2,3,
		[]float32{
			7,8,
			9,10,
			11,12,
		})

	expected := New(2,2, []float32{
		58,64,
		139,154,
	})


	r := a.Mul(b)

	if !reflect.DeepEqual(r, expected) {
		t.Errorf("martices are not equal")
		expected.Print()
		fmt.Println()
		r.Print()
	}
}