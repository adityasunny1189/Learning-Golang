package practice

import (
	"encoding/json"
	"fmt"
)

type Offer struct {
	CompanyName string  `json:"company name"`
	Package     float64 `json:"package"`
	Role        string  `json:"role"`
	Fte         bool    `json:"full time"`
}

type Student struct {
	Name   string  `json:"student name"`
	Usn    string  `json:"USN"`
	Branch string  `json:"branch"`
	Cgpa   float64 `json:"cgpa"`
	Offers []Offer `json:"offers"`
}

func EncodeJSON() {
	students := []Student{
		{
			Name:   "Aditya Pathak",
			Usn:    "1SI18IS003",
			Branch: "Information Science",
			Cgpa:   7.87,
			Offers: []Offer{
				{
					CompanyName: "Nuclei",
					Package:     1600000,
					Role:        "full stack developer",
					Fte:         true,
				},
			},
		},
		{
			Name:   "Amit Shahwal",
			Usn:    "1SI18IS063",
			Branch: "Information Science",
			Cgpa:   9.14,
			Offers: []Offer{
				{
					CompanyName: "Amazon",
					Package:     3200000,
					Role:        "big data engineer",
					Fte:         false,
				},
				{
					CompanyName: "TCS Digital",
					Package:     750000,
					Role:        "associate engineer",
					Fte:         true,
				},
			},
		},
	}
	out, err := json.MarshalIndent(students, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
