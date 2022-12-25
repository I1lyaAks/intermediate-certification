package city

import "fmt"

type CityStruct struct {
	Id         int    `json:"id"`         //(уникальный номер);
	Name       string `json:"name"`       // (название города);
	Region     string `json:"region"`     //(регион);
	District   string `json:"district"`   //(округ);
	Population int    `json:"population"` //(численность населения);
	Poundation int    `json:"poundation"` //(год основания).
}
type NewPopulation struct {
	NewNumber int `json:"newnumber"`
}

func (c CityStruct) Str() string {
	str := fmt.Sprintf("номер:%v, город:%s, регион:%s, округ:%s, численность населения:%d, год основания:%d \n",
		c.Id, c.Name, c.Region, c.District, c.Population, c.Poundation)
	return str
}
