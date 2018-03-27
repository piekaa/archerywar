package dynamic

import "testing"

func TestDynamicArrayImpl_Add(t *testing.T) {

	d := New()

	d.Add([]byte{0,1,2,3,4})
	d.Add([]byte{5,6,7,8})
	d.Add([]byte{9,10,11,12})

	arr := d.Bytes

	for i := 0 ; i <= 12 ;i++ {
		if byte(i) != arr[i] {
			t.Errorf("%s != %s", i, arr[i])
		}
	}
}

func TestDynamicArrayImpl_Add2(t *testing.T) {

	d := New()

	d.Add([]byte{0,1,2,3,4})
	d.AddN([]byte{5,6,7,8,9,10,11,12,13}, 4)
	d.Add([]byte{9,10,11,12})

	arr := d.Bytes

	for i := 0 ; i <= 12 ;i++ {
		if byte(i) != arr[i] {
			t.Errorf("%s != %s", i, arr[i])
		}
	}

}