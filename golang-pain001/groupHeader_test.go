package pain001

import (
	"regexp"
	"testing"
)

func TestNewGroupHeader(t *testing.T) {
	expected := `<GroupHeader><MsgId>d08de8c4-0fe8-4bcd</MsgId><CreDtTm>2020-07-16T22:40:51</CreDtTm><NbOfTxs>1</NbOfTxs><CtrlSum>23.50</CtrlSum><InitgPty><Nm>George Goodman</Nm></InitgPty></GroupHeader>`
	data := NewGroupHeader("George Goodman", 1, "23.50")
	data.MessageIdentification = "d08de8c4-0fe8-4bcd"
	data.CreationDateTime = "2020-07-16T22:40:51"
	compareXmlResult(t, data, expected)
}

func TestGetShortId(t *testing.T) {
	id := getShortId()
	if len(id) != 18 {
		t.Errorf("id \"%s\" has the wrong length %d not 18", id, len(id))
	}
	re := regexp.MustCompile(`^.*-.*-.*$`)
	if !re.MatchString(id) {
		t.Errorf("id \"%s\" has not the format \"^.*-.*-.*$\"", id)
	}
}
