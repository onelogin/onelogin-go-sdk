package legalvalues

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/onelogin/onelogin-go-sdk/pkg/services"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
)

const errLegalValuesContext = "legal values service"

// V2Service holds the information needed to interface with a repository
type LegalValuesService struct {
	BaseURL, ErrorContext string
	Repository            services.Repository
}

// New creates the new svc service v2.
func New(repo services.Repository, host string) LegalValuesService {
	return LegalValuesService{
		BaseURL:      fmt.Sprintf("%s/api/2", host),
		Repository:   repo,
		ErrorContext: errLegalValuesContext,
	}
}

func (svc LegalValuesService) GetLegalValues(endpoint string) ([]string, error) {
	respBytes, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		AuthMethod: "bearer",
		URL:        fmt.Sprintf("%s/%s", svc.BaseURL, endpoint),
		Headers:    map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return []string{}, err
	}

	var respStruct []map[string]string
	if err := json.Unmarshal(respBytes, &respStruct); err != nil {
		return []string{}, errors.New("Unexpected data shape given")
	}

	legalVals := make([]string, len(respStruct))
	for i, legalVal := range respStruct {
		legalVals[i] = legalVal["value"]
	}
	return legalVals, nil
}
