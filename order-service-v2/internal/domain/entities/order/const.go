package order

const (
	ORDER_SYSTEM_PROCESS = 0
	ORDER_CREATED        = 1
	ORDER_PREPARED       = 2
	ORDER_DELIVERY       = 3

	ORDER_SHIPPING_FINISH    = 4
	ORDER_COMPLETED          = 5
	ORDER_REFUND             = 6
	ORDER_CANCEL_BY_USER     = -2
	ORDER_CANCEL_BY_STORE    = -3
	ORDER_CANCEL_BY_ADMIN    = -4
	ORDER_CANCEL_BY_DELI     = -5
	ORDER_CANCEL_USER_REJECT = -7
	ORDER_FAILED             = -1
)

const (
	PAYMENT_COD        = 1
	PAYMENT_VIA_PAYPAL = 2
	PAYMENT_VIA_WALLET = 3
)

const (
	COMMIS_PENDING = 0
	COMMIS_SUCCESS = 1
	COMMIS_FAILED  = -1
)
