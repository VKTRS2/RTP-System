package findencoding

import (
	"fmt"

	fs "github.com/ungerik/go-fs"
	"golang.org/x/text/encoding/charmap"

	"github.com/domonda/go-types/charset"
	"github.com/domonda/go-types/strutil"
)

func PrintFileWithAllEncodings(file fs.FileReader, maxBytes int) error {
	sourceData, err := file.ReadAll()
	if err != nil {
		return err
	}
	if maxBytes == 0 {
		maxBytes = 1024 * 1024
	}
	if len(sourceData) > maxBytes {
		sourceData = sourceData[:maxBytes]
	}

	print := func(charset string, str string) {
		fmt.Println("====================")
		fmt.Println(charset)
		fmt.Println("--------------------")
		fmt.Print(str, "\n")
		fmt.Println("====================")
	}

	bom, sourceData := charset.SplitBOM(sourceData)
	if bom != charset.NoBOM {
		fmt.Println("Found BOM:", bom)

		str, err := bom.DecodeString(sourceData)
		if err != nil {
			return err
		}
		print(bom.String(), str)
		return nil
	}

	sourceData = strutil.SanitizeLineEndingsBytes(sourceData)

	for _, enc := range charmap.All {
		if cm, ok := enc.(*charmap.Charmap); ok {
			decoded, err := enc.NewDecoder().Bytes(sourceData)
			if err != nil {
				return err
			}
			if cm.String() == "IBM Code Page 037" {
				continue
			}
			print(cm.String(), string(decoded))
		}
	}

	print("UTF-8", string(sourceData))

	return nil
}
