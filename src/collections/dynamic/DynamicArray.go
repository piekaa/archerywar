package dynamic


type DynamicArray interface {
	Add([]byte)
	AddN([]byte, int)
}

type DynamicArrayImpl struct {
	Bytes []byte
}

func New() *DynamicArrayImpl {
	return &DynamicArrayImpl{make([]byte, 0)}
}

//n - number of bytes to add
func (arr *DynamicArrayImpl) AddN(bytes []byte, n int) {

	newSize := len(arr.Bytes) + n
	newBytes := make([]byte, newSize)

	for i := 0 ; i < len(arr.Bytes) ;i ++ {
		newBytes[i] = arr.Bytes[i]
	}
	j:=0
	for i := len(arr.Bytes); i < newSize ;i++ {
		newBytes[i] = bytes[j]
		j++
	}
	arr.Bytes = newBytes
}


func (arr *DynamicArrayImpl) Add(bytes []byte) {
	arr.AddN(bytes, len(bytes))
}