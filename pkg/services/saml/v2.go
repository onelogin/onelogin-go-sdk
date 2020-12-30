package saml

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/onelogin/onelogin-go-sdk/pkg/services"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
)

const errSAMLV2Context = "saml v2 service"

// V2Service holds the information needed to interface with a repository
type V2Service struct {
	Endpoint, ErrorContext string
	Repository             services.Repository
}

// New creates the new svc service v2.
func New(repo services.Repository, host string) *V2Service {
	return &V2Service{
		Endpoint:     fmt.Sprintf("%s/api/2/saml_assertion", host),
		Repository:   repo,
		ErrorContext: errSAMLV2Context,
	}
}

// Create retrieves the role by id and returns it
func (svc *V2Service) Create(request *SAMLAssertionRequest) (*SAMLAssertion, error) {
	log.Println("SAML create")
	resp, err := svc.Repository.Create(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s", svc.Endpoint),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    request,
	})

	if err != nil {
		return nil, err
	}

	var samlAssertion SAMLAssertion
	json.Unmarshal(resp, &samlAssertion)
	return &samlAssertion, nil
}

// Verify retrieves the role by id and returns it
func (svc *V2Service) Verify(request *VerifyFactorRequest) (*VerifyFactor, error) {
	resp, err := svc.Repository.Create(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/verify_factor", svc.Endpoint),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    request,
	})
	if err != nil {
		return nil, err
	}

	var verifyFactor VerifyFactor
	json.Unmarshal(resp, &verifyFactor)
	return &verifyFactor, nil
}
