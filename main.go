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

const fundamentusURL = "http://www.fundamentus.com.br/detalhes.php?papel="

func main() {
	csvfile, err := os.Open(os.Args[2])
	checkError("Couldn't open the csv file", err)
	defer csvfile.Close()
	total, err := strconv.ParseFloat(os.Args[1], 64)
	checkError("Couldn't parse the argument value", err)
	file, err := os.Create("result.csv")
	checkError("Couldn't  create file", err)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	csvHeader := []string{"stock", "weight", "value", "quantity", "total"}
	err = writer.Write(csvHeader)
	checkError("Couldn't write to file", err)
	sumToInvest := 0.0
	reader := csv.NewReader(csvfile)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		checkError("Couldn't read a csv line", err)
		script, weight := record[0], record[1]
		url := fmt.Sprintf("%s%s", fundamentusURL, script)
		price := getScrpitValue(url)
		fWeight, err := strconv.ParseFloat(weight, 64)
		checkError("Couldn't parse weight value", err)
		fPrice, err := strconv.ParseFloat(price, 64)
		checkError("Couldn't parse price value", err)
		quantity := int(total * fWeight / 100.00 / fPrice)
		total := fPrice * float64(quantity)
		resultLine := []string{
			script,
			weight,
			fmt.Sprintf("%.2f", fPrice),
			fmt.Sprintf("%d", quantity),
			fmt.Sprintf("%.2f", total),
		}
		err = writer.Write(resultLine)
		checkError("Cannot write to file", err)
		sumToInvest = sumToInvest + total
	}
	csvFooter := []string{"-", "-", "-", "-", fmt.Sprintf("%.2f", sumToInvest)}
	err = writer.Write(csvFooter)
	checkError("Couldn't write to file", err)
}

func getScrpitValue(url string) string {
	doc, err := htmlquery.LoadURL(url)
	checkError("Cannot load the url", err)
	xpath := "/html/body/div[1]/div[2]/table[1]/tbody/tr[1]/td[4]/span"
	span := htmlquery.FindOne(doc, xpath)
	value := htmlquery.InnerText(span)
	return strings.Replace(value, ",", ".", 1)
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
