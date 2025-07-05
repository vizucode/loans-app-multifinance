package formatting

func MapXenditStatus(status string) string {
	switch status {
	case "EXPIRED":
		return "CANCELED"
	case "FAILED":
		return "CANCELED"
	case "REFUNDED":
		return "CANCELED"
	case "PENDING":
		return "BOOKED"
	case "SETTLING":
		return "BOOKED"
	default:
		return "PAID"
	}
}
