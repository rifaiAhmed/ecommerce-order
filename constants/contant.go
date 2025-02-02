package constants

const (
	SuccessMessage      = "success"
	ErrFailedBadRequest = "data tidak sesuai"
	ErrServerError      = "terjadi kesalahan pada server"
)

const (
	OrderStatusSuccess = "SUCCESS"
	OrderStatusPending = "PENDING"
	OrderStatusFailed  = "FAILED"
	OrderStatusRefund  = "REFUND"
)

var MappingOrderStatus = map[string]bool{
	OrderStatusSuccess: true,
	OrderStatusPending: true,
	OrderStatusFailed:  true,
	OrderStatusRefund:  true,
}

var MappingFlowOrderStatus = map[string][]string{
	OrderStatusPending: {OrderStatusFailed, OrderStatusSuccess},
	OrderStatusSuccess: {OrderStatusRefund},
}
