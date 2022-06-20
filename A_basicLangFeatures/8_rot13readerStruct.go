package A_basicLangFeatures

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reada rot13Reader) Read(toFill []byte) (int, error) {

	holder := make([]byte, cap(toFill))

	bytes, err := reada.r.Read(holder)

	if err != nil {
		return 0, err
	}

	for i, v := range holder {

		margin := byte(96)
		if v < 97 {
			margin = 64
		}

		// pretty much asking if it is NOT a letter
		if v < 65 || (v > 90 && v < 97) || v > 122 {
			toFill[i] = v
		} else { // applying rot13 algorithm if it was a letter
			toFill[i] = (((v + 13) - margin) % 26) + margin
		}
	}

	return bytes, nil
}

func Rot13readerStruct_main() {
	stringie := "Lbh penpxrq gur pbqr!"
	s := strings.NewReader(stringie)
	r := rot13Reader{s}
	fmt.Println(stringie)
	// leveraging the reader by using copy from src (the reader) to dst (os.Stdout)
	io.Copy(os.Stdout, &r)
}
