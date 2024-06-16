package storage

import (
	"log"
	"sync"
	"time"

	"github.com/ohadvaknin/go-wolt-delivery/models"
)

var deliveryQueue []models.Delivery
var courierQueue []models.Courier
var id int
var mu sync.Mutex

// InitializeServer initializes the courier queue and starts the courier assignment process.
func InitializeServer(N int) {
	courierQueue = make([]models.Courier, N)
	for i := 0; i < N; i++ {
		courierQueue[i].ID = i + 1
	}
	id = 1
	// Start the courier assignment process
	go assignCourier()
}

// CreateOrder creates a new delivery order and adds it to the delivery queue.
func CreateOrder(order models.Delivery) models.Delivery {
	mu.Lock()
	defer mu.Unlock()
	order.ID = id
	id++
	order.ReceivedAt = time.Now()
	deliveryQueue = append(deliveryQueue, order)
	log.Printf("Received order id:%d", order.ID)
	return order
}

// assignCourier assigns couriers to deliveries in the queue.
func assignCourier() {
	for {
		mu.Lock()
		if len(courierQueue) > 0 && len(deliveryQueue) > 0 {
			courier := courierQueue[0]
			courierQueue = courierQueue[1:]
			delivery := deliveryQueue[0]
			deliveryQueue = deliveryQueue[1:]
			mu.Unlock()

			go func() {
				duration := time.Duration(delivery.Distance * float64(time.Minute))
				arrivalTime := time.Now().Add(duration)
				log.Printf("Order %d was picked up by courier %d. The order will arrive at %v\n", delivery.ID, courier.ID, arrivalTime.Format("15:04:05"))
				time.Sleep(duration)
				log.Printf("Order %d arrived.\n", delivery.ID)
				mu.Lock()
				courierQueue = append(courierQueue, courier)
				mu.Unlock()
			}()
		} else {
			mu.Unlock()
			time.Sleep(time.Second)
		}
	}
}
