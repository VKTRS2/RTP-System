package sequence_type

func (s SequenceType) Prefix() string {
	switch s {
	case StandardInvoice:
		return "F"
	case SpecialInvoice:
		return "X"
	case RectificationInvoice:
		return "R"
	case Customer:
		return "C"
	default:
		return ""
	}
}
