package hubspot

import "fmt"

const (
	crmBasePath = "crm"

	objectsBasePath = "objects"
)

type CRM struct {
	Contact ContactService
	Company CompanyService
	Deal    DealService
}

func newCRM(c *Client) *CRM {
	crmPath := fmt.Sprintf("%s/%s", crmBasePath, c.apiVersion)
	return &CRM{
		Contact: &ContactServiceOp{
			contactPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, contactBasePath),
			client:      c,
		},
		Company: &CompanyServiceOp{
			companyPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, companyBasePath),
			client:      c,
		},
		Deal: &DealServiceOp{
			dealPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, dealBasePath),
			client:   c,
		},
	}
}
