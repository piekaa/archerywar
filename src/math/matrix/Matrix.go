package matrix

import "fmt"

type Matrix interface {
	//right multiplication
	Mul(m Matrix) Matrix
	//left multiplication
	Mull(m Matrix) Matrix
	//right multiplication
	Mulr(m Matrix) Matrix
	//returns array
	Array() []float32
	Print()
	X() int
	Y() int
}

type myMatrix struct {
	//width
	x int
	//height
	y int
	//values
	v []float32
}

//width, height and values
func New(x, y int, v []float32) Matrix {

	if x*y != len(v) {
		panic("number of elements doesn't match size")
	}
	return &myMatrix{x, y, v}
}


func (this *myMatrix) Mul(m Matrix) Matrix {
	res := make([]float32, this.y * m.X())
	if this.x != m.Y() {
		panic("Left matrix width != right matrix height")
	}
	for i := 0 ; i < this.y ; i++ {
		for j := 0 ; j < m.X() ; j++ {
			v:=float32(0)
			for k := 0 ; k < this.x ;k++ {
				v+= this.v[k + i * this.x] * m.Array()[k*m.X() + j]
			}
			res[j + i * this.y] = v
		}
	}
	return New(this.y, m.X(), res)
}

func (*myMatrix) Mull(m Matrix) Matrix {
	panic("implement me")
}

func (*myMatrix) Mulr(m Matrix) Matrix {
	panic("implement me")
}

func (this *myMatrix) Array() []float32 {
	return this.v
}

func (this *myMatrix) X() int {
	return this.x
}

func (this *myMatrix) Y() int {
	return this.y
}

func (this *myMatrix) Print() {

	for j := 0; j < this.y; j++ {
		for i := 0; i < this.x; i++ {
			fmt.Printf("%.2f ",this.v[i + j * this.x])
		}
		fmt.Println()
	}

}
