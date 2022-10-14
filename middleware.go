package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var letters = []rune("0123456789")

func main() {

	mux := http.NewServeMux()

	endpoint := http.HandlerFunc(greet)

	mux.Handle("/", middlewarel(middleware2(endpoint)))

	fmt.Println("Listening to port 8080")

	err := http.ListenAndServe(":3000", mux)

	log.Fatal(err)

}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!"))
}

func middlewarel(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("status water dan wind \n\n")
		next.ServeHTTP(w, r)
	})
}

func middleware2(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		for i := 1; i <= 100; i++ {
			start := time.Now()

			time.Sleep(14 * time.Second)

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

			switch {
			case intbilacak1 <= 6:
				fmt.Println("status aman :", bilacak1, "m")

			case (intbilacak1 < 6) && (intbilacak1 >= 8):
				fmt.Println("status siaga :", bilacak1, "m")

			case intbilacak1 > 8:
				fmt.Println("status bahaya :", bilacak1, "m")
			}

			switch {
			case intbilacak2 <= 6:
				fmt.Println("status aman :", bilacak2, "m/s")

			case (intbilacak2 <= 7) && (intbilacak2 >= 15):
				fmt.Println("status siaga :", bilacak2, "m/s")

			case intbilacak2 > 15:
				fmt.Println("status bahaya :", bilacak2, "m/s")
			}

			fmt.Println(strings.Repeat("#", 25))
			fmt.Print("\n")

		}

		next.ServeHTTP(w, r)
	})
}

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
