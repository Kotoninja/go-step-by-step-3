package main

import (
	"fmt"
)

// Order представляет информацию о заказе.
type Order struct {
	OrderNumber  int
	CustomerName string
	OrderAmount  float64
}

// OrderLogger представляет журнал заказов и хранит записи о заказах.
type OrderLogger struct{}

// NewOrderLogger создает новый экземпляр OrderLogger.
func NewOrderLogger() *OrderLogger {
	return &OrderLogger{}
}

func (logger *OrderLogger) AddOrder(order Order) {
	fmt.Printf("Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f\n", order.OrderNumber, order.CustomerName, order.OrderAmount)
}
