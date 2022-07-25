package rtp

var listOfDS12ValidTransactionStatus = []string{
	StatusCancelledAsPerRequest, StatusRejectedCancellationRequest,
}

var listOfDS12ValidNegativeTransactionReasons = []string{
	RejectCancellationReasonAlreadyCancelledRTP,
	StatusReasonAlreadyExpiredRTP,
	StatusReasonAlreadyRefusedRTP,
	StatusReasonAlreadyRejectedRTP,
	RejectCancellationReasonPaymentAlreadyTransmittedExecution,
	RejectReasonRegulatoryReason,
	RejectCancellationReasonUnknownRTP,
}
