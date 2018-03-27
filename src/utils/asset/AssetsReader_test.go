package asset

import (
	"testing"
	"io/ioutil"
	"bytes"
)

func TestReadString(t *testing.T) {

	name := "9u429ehuih89wyfv4ugjtg"
	path := "assets/" + name
	//path :=  name
	content := `gdybym ci ja była słoneczkiem na niebie, to nie świeciłabym, ja tylko dla ciebie
ani na wody, ani na lasy, ale po wszystkie czasy pod twym okienkiem i tylko dla ciebie gdybym ja była słoneczkiem na niebie`
	ioutil.WriteFile(path, []byte(content), 0777)


	readed, err := ReadString(name)

	if err != nil {
		t.Errorf("%s", err)
	}

	if content !=  readed {
		t.Errorf("String are different %s : %s", content, readed )
	}

}


func TestReadBytes(t *testing.T) {

	name := "3e78yf78dshfuih34789g478fe3"
	path := "assets/" + name
	//path :=  name
	content := []byte {1,2,3,4,5,6,7,7,6,5,4,3,2,1}
	ioutil.WriteFile(path, []byte(content), 0777)

	readed, err := ReadBytes(name)

	if err != nil {
		t.Errorf("%s", err)
	}

	if bytes.Compare(content,readed) != 0 {
		t.Errorf("String are different %s : %s", content, readed )
	}

}