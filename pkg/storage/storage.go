package storage

import (
	"encoding/json"
	"fmt"
	c "hi/pkg/city"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

var City = make(map[int]*c.CityStruct, 500)

func CityID(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println(err)
	}
	s := City[idInt]
	w.Write([]byte((s.Str())))
}
func CreateCity(w http.ResponseWriter, r *http.Request) {
	content, _ := ioutil.ReadAll(r.Body)
	var s c.CityStruct
	json.Unmarshal(content, &s)
	fmt.Println(s)
	City[s.Id] = &s
	w.Write([]byte((s.Str())))
}
func DeleteCity(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println(err)
	}
	str := fmt.Sprintf("город %s удалён из списка", City[idInt].Name)
	delete(City, idInt)
	w.Write([]byte(str))
}
func UpdatePopulation(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println(err)
	}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var p c.NewPopulation
	json.Unmarshal(content, &p)
	City[idInt].Population = p.NewNumber
	str := fmt.Sprintf("численность населения в городе %s изменена", City[idInt].Name)
	w.Write([]byte(str))
}
func CityRegion(w http.ResponseWriter, r *http.Request) {
	reg := chi.URLParam(r, "region")
	reg = strings.ReplaceAll(reg, "%20", " ")
	var regionMap = make(map[int]*c.CityStruct)
	for _, i := range City {
		if i.Region == reg {
			regionMap[i.Id] = i
		}
	}
	for _, val := range regionMap {
		s := val.Str()
		w.Write([]byte(s))
	}
}
func CityDistrict(w http.ResponseWriter, r *http.Request) {
	districtStr := chi.URLParam(r, "district")
	var regionMap = make(map[int]*c.CityStruct)
	for _, i := range City {
		if i.District == districtStr {
			regionMap[i.Id] = i
		}
	}
	for _, val := range regionMap {
		s := val.Str()
		fmt.Println(s)
		w.Write([]byte(s))
	}
}
func Population(w http.ResponseWriter, r *http.Request) {
	var populationMap = make(map[int]*c.CityStruct)
	minString := chi.URLParam(r, "min")
	min, err := strconv.Atoi(minString)
	if err != nil {
		fmt.Println(err)
	}
	maxString := chi.URLParam(r, "max")
	max, err := strconv.Atoi(maxString)
	if err != nil {
		fmt.Println(err)
	}
	if min > max {
		min, max = max, min
	}
	for _, i := range City {
		if i.Population > min && i.Population < max {
			populationMap[i.Id] = i
		}
	}
	for _, val := range populationMap {
		s := val.Str()
		fmt.Println(s)
		w.Write([]byte(s))
	}
}
func Poundation(w http.ResponseWriter, r *http.Request) {
	var poundationMap = make(map[int]*c.CityStruct)
	minString := chi.URLParam(r, "min")
	min, err := strconv.Atoi(minString)
	if err != nil {
		fmt.Println(err)
	}
	maxString := chi.URLParam(r, "max")
	max, err := strconv.Atoi(maxString)
	if err != nil {
		fmt.Println(err)
	}
	if min > max {
		min, max = max, min
	}
	for _, i := range City {
		if i.Poundation > min && i.Poundation < max {
			poundationMap[i.Id] = i
		}
	}
	for _, val := range poundationMap {
		s := val.Str()
		fmt.Println(s)
		w.Write([]byte(s))
	}
}
