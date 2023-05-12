package orders

import bricklink "github.com/funwithbots/go-bricklink-api"

const (
	pathGetOrders          = "/orders"
	pathGetOrder           = "/orders/%d"
	pathGetOrderFeedback   = "/orders/%d/feedback"
	pathGetOrderItems      = "/orders/%d/items"
	pathGetMessages        = "/orders/%d/messages"
	pathGetFeedback        = "/feedback/%d"
	pathGetFeedbackList    = "/feedback"
	pathPostFeedback       = "/feedback"
	pathReplyFeedback      = "/feedback/%d/reply"
	pathGetMemberRating    = "/members/%s/ratings"
	pathGetMemberNote      = "/members/%s/notes"
	pathCreateMemberNote   = "/members/%s/notes"
	pathUpdateMemberNote   = "/members/%s/notes"
	pathDeleteMemberNote   = "/members/%s/notes"
	pathUpdateOrder        = "/orders/%d"
	pathUpdateStatus       = "/orders/%d/status"
	pathUpdatePayment      = "/orders/%d/payment_status"
	pathSendDriveThru      = "/orders/%d/drive_thru"
	pathGetShippingMethods = "/settings/shipping_methods"
)

type Orders struct {
	bricklink.Bricklink

	ShippingMethods map[int]ShippingMethod
}

func New(bl bricklink.Bricklink) (*Orders, error) {
	o := Orders{}
	o.Bricklink = bl

	err := o.loadShippingMethods()
	if err != nil {
		return nil, err
	}

	return &o, nil
}
