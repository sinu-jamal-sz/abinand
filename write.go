package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	records1 := [][]string{}
	records2 := [][]string{}
	f, err := os.OpenFile("users.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	m, err := os.OpenFile("dep.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("enter employee id employee name")
	var id, did, name, dname string
	fmt.Scanln(&id)
	fmt.Scanln(&name)
	fmt.Println("enter department id and name")
	fmt.Scanln(&did)
	fmt.Scanln(&dname)
	values1 := []string{id, name, "0"}
	values2 := []string{did, dname}
	records1 = append(records1, values1)
	records2 = append(records2, values2)

	w := csv.NewWriter(f)
	x := csv.NewWriter(m)
	defer w.Flush()
	defer x.Flush()

	for _, record := range records1 {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	for _, record := range records2 {
		if err := x.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

}
