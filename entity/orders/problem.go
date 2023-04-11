package orders

var ProblemReasonIDs map[int]string = map[int]string{
	6:   "Buyer did not respond to emails",
	7:   "Buyer email address bounced",
	8:   "Seller did not receive payment",
	9:   "Buyer found items elsewhere",
	10:  "Buyer no longer wants to purchase items",
	11:  "Buyer does not have enough funds to pay",
	12:  "Buyer did not comply with store policies",
	13:  "Buyer submitted a bogus orders",
	16:  "Buyer demanded a lower price on items",
	57:  "Package returned with incorrect address",
	60:  "Seller did not have items after orders was submitted",
	63:  "Buyer is underage",
	70:  "System Problem",
	72:  "Buyer is no longer registered",
	97:  "Package cannot be delivered to buyer's address",
	100: "Buyer did not pay for orders",
	107: "Mutual agreement between buyer and seller to cancel",
}

type Problem struct {
	Type     string `json:"type"`
	Message  string `json:"message"`
	ReasonID string `json:"reason_id"`
}
