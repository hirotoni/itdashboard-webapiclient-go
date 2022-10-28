package apiclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/hirotoni/itdashboard-webapiclient-go/config"
)

type BasicInformationField string

// json field names
const (
	FieldSystemId         = BasicInformationField("system_id")
	FieldSystemName       = BasicInformationField("system_name")
	FieldSystemClassCode  = BasicInformationField("system_class_code")
	FieldSystemClass      = BasicInformationField("system_class")
	FieldOrganizationCode = BasicInformationField("organization_code")
	FieldOrganization     = BasicInformationField("organization")
	FieldYear             = BasicInformationField("year")
)

type RequestOptions struct {
	FieldsToGet    []BasicInformationField
	FilterByFields map[BasicInformationField]string
}

type BasicInformationResponse struct {
	Info    Info               `json:"info"`
	RawData []BasicInformation `json:"raw_data"`
}

type BasicInformation struct {
	SystemId         *string `json:"system_id,omitempty"`         // 情報システムID string or null
	SystemName       *string `json:"system_name,omitempty"`       // 情報システム名 string or null
	SystemClassCode  *string `json:"system_class_code,omitempty"` // 情報システム区分コード string or null
	SystemClass      *string `json:"system_class,omitempty"`      // 情報システム区分 string or null
	OrganizationCode *string `json:"organization_code,omitempty"` // 組織コード string or null
	Organization     *string `json:"organization,omitempty"`      // 組織名 string or null
	Year             *int    `json:"year,omitempty"`              // 年（西暦） int or null
}

func (bi BasicInformation) String() string {
	v := reflect.ValueOf(bi)
	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		temp := v.Field(i)

		switch {
		case temp.IsNil():
			values[i] = "<nil>"
		case temp.Kind() == reflect.Ptr:
			values[i] = temp.Elem()
		}
	}

	return fmt.Sprint(values)
}

/* Fetch BasicInformation data */
func (a *ApiClient) FetchBasicInformation(opts RequestOptions) (*BasicInformationResponse, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	q := url.Values{}
	q.Set("dataset", config.BasicInformation)

	// fields to get option
	fields := []string{} // convert []BasicInformationField -> []string
	for _, v := range opts.FieldsToGet {
		fields = append(fields, string(v))
	}

	fieldsToGet := strings.Join(fields, ",")
	q.Set("field", fieldsToGet)

	// filter options
	for k, v := range opts.FilterByFields {
		key := string(k) // convert BasicInformationField -> string
		q.Set(key, v)
	}

	url := a.url + "?" + q.Encode()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	// req.Header.Add("someheader", "this is value")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// if not OK, return zero value struct
	var bs BasicInformationResponse
	if resp.StatusCode != 200 {
		return &bs, nil
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &bs); err != nil {
		panic(err)
	}

	return &bs, nil
}
