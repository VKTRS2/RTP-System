package pain001

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

// GroupHeader represents a Group Header item.
type GroupHeader struct {
	// Unique identification of the payment document. Bank system will reject two files with the same identifier.
	MessageIdentification string `xml:"MsgId"`
	// Date and time of the creation of the file.
	CreationDateTime string `xml:"CreDtTm"`
	// The number of total transactions in the document.
	NumbersOfTransactions int `xml:"NbOfTxs"`
	// Sum of all transactions
	ControlSum string `xml:"CtrlSum"`
	// The name of the organization/person which sends the money.
	Name string `xml:"InitgPty>Nm"`
}

// NewGroupHeader takes an convert.Sender object and the number of transitions as parameters and returns a GroupHeader.
func NewGroupHeader(senderName string, numberOfTrans int, controlSum string) GroupHeader {
	return GroupHeader{
		MessageIdentification: getShortId(),
		CreationDateTime:      time.Now().Format(DateTimeFormat),
		NumbersOfTransactions: numberOfTrans,
		ControlSum:            controlSum,
		Name:                  senderName,
	}
}

func getShortId() string {
	parts := strings.Split(uuid.NewV4().String(), "-")
	return fmt.Sprintf("%s-%s-%s", parts[0], parts[1], parts[2])
}
