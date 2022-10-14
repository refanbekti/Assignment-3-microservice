package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Status struct {
	Water        int
	Wind         int
	Status_Wind  string
	Status_water string
}

var letters = []rune("0123456789")
var status = []Status{}
var PORT = ":8080"

func main() {
	http.HandleFunc("/", windandwater)
	fmt.Print("Application is listening on port", PORT)

	http.ListenAndServe(PORT, nil)

}

func windandwater(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-Type", "application/json")
	if r.Method == "POST" {

		for i := 1; i <= 100; i++ {
			start := time.Now()

			time.Sleep(2 * time.Second)

			time.Since(start)
			rand.Seed(time.Now().UTC().UnixNano())

			bilacak1 := randomString(2)
			bilacak2 := randomString(2)
			intbilacak1, err := strconv.Atoi(bilacak1)
			intbilacak2, err2 := strconv.Atoi(bilacak2)
			if err != nil {
				// ... handle error
				panic(err)
			}
			if err2 != nil {
				// ... handle error
				panic(err)
			}

			fmt.Println("water:", bilacak1, "m")
			fmt.Println("wind:", bilacak2, "m/s")

			var newStatus Status

			newStatus.Water = intbilacak1
			newStatus.Wind = intbilacak2

			switch {
			case intbilacak1 <= 5:
				newStatus.Status_water = "Aman"
				fmt.Println("status aman :", bilacak1, "m")

			case (intbilacak1 >= 6) && (intbilacak1 >= 8):
				newStatus.Status_water = "Siaga"
				fmt.Println("status siaga :", bilacak1, "m")

			case intbilacak1 > 8:
				newStatus.Status_water = "Bahaya"
				fmt.Println("status bahaya :", bilacak1, "m")

			}

			switch {
			case intbilacak2 <= 6:
				newStatus.Status_Wind = "Aman"
				fmt.Println("status aman :", bilacak2, "m/s")

			case (intbilacak2 <= 7) && (intbilacak2 >= 15):
				newStatus.Status_Wind = "Siaga"
				fmt.Println("status siaga :", bilacak2, "m/s")

			case intbilacak2 >= 15:
				newStatus.Status_Wind = "Bahya"
				fmt.Println("status bahaya :", bilacak2, "m/s")

			}

			// newStatus := Status{

			// 	Water:        intbilacak1,
			// 	Wind:         intbilacak2,
			// 	Status_water: statwater,

			// 	Status_Wind: statwind,
			// }

			status = append(status, newStatus)

			json.NewEncoder(w).Encode(newStatus)

			return

		}
	}

}

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
