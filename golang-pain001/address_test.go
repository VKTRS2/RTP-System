package pain001

import (
	"testing"
)

func TestAddress(t *testing.T) {
	expected := `<Address><Nm>George Goodman</Nm><PstlAdr><StrtNm>Fnord Street</StrtNm><BldgNb>23</BldgNb><PstCd>2323</PstCd><TwnNm>Fnord Town</TwnNm><Ctry></Ctry></PstlAdr></Address>`
	data := Address{
		Name: "George Goodman",
		PostalAddress: PostalAddress{
			StreetName:   "Fnord Street",
			StreetNumber: 23,
			PostalCode:   2323,
			TownName:     "Fnord Town"}}
	compareXmlResult(t, data, expected)
}

func TestPostalAddress(t *testing.T) {
	expected := `<PostalAddress><StrtNm>Fnord Street</StrtNm><BldgNb>23</BldgNb><PstCd>2323</PstCd><TwnNm>Fnord Town</TwnNm><Ctry></Ctry></PostalAddress>`
	data := PostalAddress{
		StreetName:   "Fnord Street",
		StreetNumber: 23,
		PostalCode:   2323,
		TownName:     "Fnord Town",
	}
	compareXmlResult(t, data, expected)
}
