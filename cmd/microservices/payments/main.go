package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akhilsharma90/monolith-microservice/pkg/common/cmd"
	payments_app "github.com/akhilsharma90/monolith-microservice/pkg/payments/application"
	payments_infra_orders "github.com/akhilsharma90/monolith-microservice/pkg/payments/infrastructure/orders"
	"github.com/akhilsharma90/monolith-microservice/pkg/payments/interfaces/amqp"
)

func main() {
	log.Println("Starting payments microservice")
	defer log.Println("Closing payments microservice")

	ctx := cmd.Context()

	paymentsInterface := createPaymentsMicroservice()
	if err := paymentsInterface.Run(ctx); err != nil {
		panic(err)
	}
}

func createPaymentsMicroservice() amqp.PaymentsInterface {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	paymentsService := payments_app.NewPaymentsService(
		payments_infra_orders.NewHTTPClient(os.Getenv("SHOP_ORDERS_SERVICE_ADDR")),
	)

	paymentsInterface, err := amqp.NewPaymentsInterface(
		fmt.Sprintf("amqp://%s/", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
		paymentsService,
	)
	if err != nil {
		panic(err)
	}

	return paymentsInterface
}
