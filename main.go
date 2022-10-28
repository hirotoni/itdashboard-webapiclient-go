package apiclient

func init() {}

const (
	defaultURL = "https://itdashboard.cio.go.jp/PublicApi/getData.json"
)

type ApiClient struct {
	url string
}

type Info struct {
	ApiVersion string `json:"api_version"`
	Dataset    string `json:"dataset"`
}

func NewApiClient() (*ApiClient, error) {
	client := ApiClient{
		url: defaultURL,
	}
	return &client, nil
}
