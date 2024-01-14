package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func main() {
	http.HandleFunc("/temp", func(w http.ResponseWriter, r *http.Request) {
		type Input struct {
			Data string `json:"data"`
		}

		var input Input

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Decode error! please check your JSON formating.")
			return
		}

		fmt.Println("Time:", time.Now())
		fmt.Println("Inputed data:", input.Data)

		lines := strings.Split(input.Data, "\n")
		re := regexp.MustCompile(`^(\d+)°C$`)
		roomTemps := make(map[string]interface{})

		for i := 0; i < len(lines); i++ {
			matches := re.FindStringSubmatch(lines[i])
			if matches != nil {
				v, _ := strconv.Atoi(matches[1])
				roomTemps[lines[i-1]] = v
			}
		}

		fmt.Println("room temps:", roomTemps)
		writeToInflux(roomTemps)
		w.WriteHeader(200)
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func writeToInflux(params map[string]interface{}) {
	client := influxdb2.NewClient(os.Getenv("INFLUXDB_URL"), os.Getenv("INFLUXDB_TOKEN"))
	writeAPI := client.WriteAPIBlocking(os.Getenv("INFLUXDB_ORG"), os.Getenv("INFLUXDB_BUCKET"))

	tags := map[string]string{}
	point := write.NewPoint("temp", tags, params, time.Now())

	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		log.Fatal(err)
	}
}
