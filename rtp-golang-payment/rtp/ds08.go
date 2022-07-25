package rtp

var listOfDS08ValidTransactionStatus = []string{
	StatusRejected, StatusAccepted, StatusAcceptedWithChange,
}

var listOfDS08ValidNegativeTransactionReasons = []string{
	RejectReasonNotAllowedCurrency,
	RejectReasonDuplication,
	RejectReasonWrongAmount,
	RejectReasonIncorrectExpiryDateTime,
	RejectReasonNotSpecifiedReasonCustomerGenerated,
	RejectReasonNonAgreedRTP,
	RejectReasonUnknownCreditor,
}
