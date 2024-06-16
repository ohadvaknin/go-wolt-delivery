package storage

//"github.com/google/uuid"
//"github.com/ohadvaknin/go-wolt-delivery/models"


func createCourierQueue(N int) {
	//initialize the courier queue:
	courierQueue := make([]int, N)
	for i := 0; i < N; i++ {
		courierQueue[i] = i + 1
	}
}

func assignCourier() {
	
}