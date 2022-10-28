package main

import (
	"fmt"

	ac "example.com/my-go-practice/apiclient"
)

func main() {
	client, err := ac.NewApiClient()
	if err != nil {
		panic(err)
	}

	options := ac.RequestOptions{
		FieldsToGet: []ac.BasicInformationField{
			ac.FieldOrganization,
			ac.FieldSystemClass,
			ac.FieldSystemName,
		},
		FilterByFields: map[ac.BasicInformationField]string{
			ac.FieldOrganization: "内閣府",
		},
	}

	result, err := client.FetchBasicInformation(options)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Info)

	for i, v := range result.RawData {
		fmt.Println(i, v)
	}
}
