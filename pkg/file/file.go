package file

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"hi/pkg/city"
	"hi/pkg/storage"
	"io"
	"log"
	"os"
	"strconv"
)

func OpenFileCsv() {
	csvFile, err := os.Open("cities.csv")
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal("err!=nil ", error)
		}
		idint, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Println("ошибка парсинга", err)
		}
		populationint, err := strconv.Atoi(line[4])
		if err != nil {
			fmt.Println("ошибка парсинга", err)
		}
		poundationint, err := strconv.Atoi(line[5])
		if err != nil {
			fmt.Println("ошибка парсинга", err)
		}
		var k city.CityStruct = city.CityStruct{idint, line[1], line[2], line[3], populationint, poundationint}
		storage.City[idint] = &k
	}
}

func CloseFileCsv() {
	csvFile, _ := os.Create("cities.csv")
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	for _, record := range storage.City {
		var s []string
		s = append(s, strconv.Itoa(record.Id))
		s = append(s, record.Name)
		s = append(s, record.Region)
		s = append(s, record.District)
		s = append(s, strconv.Itoa(record.Population))
		s = append(s, strconv.Itoa(record.Poundation))
		writer.Write(s)
	}
	writer.Flush()
}
