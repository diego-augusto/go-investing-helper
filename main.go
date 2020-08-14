package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
)

func main() {

	csvfile, err := os.Open("input.csv")

	checkError("Couldn't open the csv file", err)

	r := csv.NewReader(csvfile)

	total, _ := strconv.ParseFloat(os.Args[1], 64)

	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"papel", "peso", "valor", "quatidade", "total"})

	checkError("Cannot write to file", err)

	for {

		record, err := r.Read()

		if err == io.EOF {
			break
		}

		checkError("Cannot read a file", err)

		if record[0] == "script" {
			continue
		}

		value := getScrpitValue("http://www.fundamentus.com.br/detalhes.php?papel=" + record[0])

		value = strings.Replace(value, ",", ".", 1)

		percetn, _ := strconv.ParseFloat(record[1], 64)
		floatValue, _ := strconv.ParseFloat(value, 64)

		calc := int(total * percetn / 100.00 / floatValue)

		if err != nil {
			log.Fatal(err)
		}

		err = writer.Write([]string{record[0], record[1], fmt.Sprintf("%.2f", floatValue), fmt.Sprintf("%d", calc), fmt.Sprintf("%.2f", floatValue*float64(calc))})

		checkError("Cannot write to file", err)
	}
}

func getScrpitValue(url string) string {

	doc, err := htmlquery.LoadURL(url)

	checkError("Cannot load the url", err)

	xpath := "/html/body/div[1]/div[2]/table[1]/tbody/tr[1]/td[4]/span"

	span := htmlquery.FindOne(doc, xpath)

	value := htmlquery.InnerText(span)

	return value
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
