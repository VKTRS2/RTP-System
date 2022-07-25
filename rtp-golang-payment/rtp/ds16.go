package rtp

var listOfDS16ValidTransactionStatus = []string{
	"",
}

var listOfDS16ValidTransactionReasons = []string{
	StatusReasonAlreadyExpiredRTP,
	StatusReasonAlreadyAcceptedRTP,
	StatusReasonAlreadyRefusedRTP,
	StatusReasonAlreadyRejectedRTP,
	StatusReasonInitialRTPNeverReceived,
	StatusReasonRTPReceivedCanBeProcessed,
}
