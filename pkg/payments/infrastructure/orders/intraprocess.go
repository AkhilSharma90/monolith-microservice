package orders

import "github.com/akhilsharma90/monolith-microservice/pkg/orders/interfaces/private/intraprocess"

type IntraprocessService struct {
	paymentsInterface intraprocess.OrdersInterface
}

func NewIntraprocessService(paymentsInterface intraprocess.OrdersInterface) IntraprocessService {
	return IntraprocessService{paymentsInterface}
}

func (o IntraprocessService) MarkOrderAsPaid(orderID string) error {
	return o.paymentsInterface.MarkOrderAsPaid(orderID)
}
