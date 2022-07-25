package pain001

// Address represents a strict address structure.
type Address struct {
	// Name of the organisation or person.
	Name string `xml:"Nm"`
	// Rest of the address
	PostalAddress PostalAddress `xml:"PstlAdr"`
}

// PostalAddress contains all address information except for the name.
type PostalAddress struct {
	StreetName   string `xml:"StrtNm"`
	StreetNumber int    `xml:"BldgNb"`
	PostalCode   int    `xml:"PstCd"`
	TownName     string `xml:"TwnNm"`
	// Use two char abbreviation in upper case (ex.: CH, DE, FR)
	Country string `xml:"Ctry"`
}
