package orders

import bricklink "github.com/funwithbots/go-bricklink-api"

const (
	pathGetOrders        = "/orders"
	pathGetOrder         = "/orders/%d"
	pathGetOrderFeedback = "/orders/%d/feedback"
	pathGetOrderItems    = "/orders/%d/items"
	pathGetMessages      = "/orders/%d/messages"
	pathGetFeedback      = "/feedback/%d"
	pathGetFeedbackList  = "/feedback"
	pathPostFeedback     = "/feedback"
	pathReplyFeedback    = "/feedback/%d/reply"
	pathGetMemberRating  = "/members/%s/ratings"

	// Member notes documentation is incorrect.
	pathGetMemberNote      = "/members/%s/my_notes"
	pathCreateMemberNote   = "/members/%s/my_notes"
	pathUpdateMemberNote   = "/members/%s/my_notes"
	pathDeleteMemberNote   = "/members/%s/my_notes"
	pathUpdate             = "/orders/%d"
	pathUpdateOrder        = "/orders/%d"
	pathUpdateStatus       = "/orders/%d/status"
	pathUpdatePayment      = "/orders/%d/payment_status"
	pathSendDriveThru      = "/orders/%d/drive_thru"
	pathGetShippingMethods = "/settings/shipping_methods"
)

type Orders struct {
	*bricklink.Bricklink

	ShippingMethods map[int]ShippingMethod
}

// New creates a new Orders instance.
func New(b *bricklink.Bricklink) (*Orders, error) {
	o := Orders{}
	o.Bricklink = b

	if err := o.loadShippingMethods(); err != nil {
		return nil, err
	}

	return &o, nil
}
