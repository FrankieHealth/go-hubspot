package hubspot

const (
	companyBasePath = "companies"
)

// CompanyService is an interface of company endpoints of the HubSpot API.
// HubSpot companies store information about individuals.
// It can also be associated with other CRM objects such as deal and company.
// Reference: https://developers.hubspot.com/docs/api/crm/companies
type CompanyService interface {
	Get(companyID string, company interface{}, option *RequestQueryOption) (*ResponseResource, error)
	Create(company interface{}) (*ResponseResource, error)
	Update(companyID string, company interface{}) (*ResponseResource, error)
	AssociateAnotherObj(companyID string, conf *AssociationConfig) (*ResponseResource, error)
}

// CompanyServiceOp handles communication with the product related methods of the HubSpot API.
type CompanyServiceOp struct {
	companyPath string
	client      *Client
}

var _ CompanyService = (*CompanyServiceOp)(nil)

type Company struct {
	AboutUs                   *HsStr  `json:"about_us,omitempty"`
	AnnualRevenue             *HsStr  `json:"annualrevenue,omitempty"`
	City                      *HsStr  `json:"city,omitempty"`
	CloseDate                 *HsTime `json:"closedate,omitempty"`
	CompanySize               *HsStr  `json:"company_size,omitempty"`
	Country                   *HsStr  `json:"country,omitempty"`
	CreateDate                *HsTime `json:"createdate,omitempty"`
	DaysToClose               *HsStr  `json:"days_to_close,omitempty"`
	HsCreateDate              *HsTime `json:"hs_createdate,omitempty"`
	HsSequencesIsEnrolled     HsBool  `json:"hs_sequences_is_enrolled,omitempty"`
	HubspotOwnerAssignedDate  *HsTime `json:"hubspot_owner_assigneddate,omitempty"`
	HubspotOwnerID            *HsStr  `json:"hubspot_owner_id,omitempty"`
	HubspotTeamID             *HsStr  `json:"hubspot_team_id,omitempty"`
	HubspotScore              *HsStr  `json:"hubspotscore,omitempty"`
	HsObjectID                *HsStr  `json:"hs_object_id,omitempty"`
	Industry                  *HsStr  `json:"industry,omitempty"`
	LastModifiedDate          *HsTime `json:"lastmodifieddate,omitempty"`
	LifeCycleStage            *HsStr  `json:"lifecyclestage,omitempty"`
	Message                   *HsStr  `json:"message,omitempty"`
	Name                      *HsStr  `json:"name,omitempty"`
	NumAssociatedDeals        *HsStr  `json:"num_associated_deals,omitempty"`
	NumNotes                  *HsStr  `json:"num_notes,omitempty"`
	NumUniqueConversionEvents *HsStr  `json:"num_unique_conversion_events,omitempty"`
	NumEmployees              *HsStr  `json:"numemployees,omitempty"`
	RecentConversionDate      *HsTime `json:"recent_conversion_date,omitempty"`
	RecentConversionEventName *HsStr  `json:"recent_conversion_event_name,omitempty"`
	RecentDealAmount          *HsStr  `json:"recent_deal_amount,omitempty"`
	RecentDealCloseDate       *HsTime `json:"recent_deal_close_date,omitempty"`
	State                     *HsStr  `json:"state,omitempty"`
	TotalRevenue              *HsStr  `json:"total_revenue,omitempty"`
	Website                   *HsStr  `json:"website,omitempty"`
	Zip                       *HsStr  `json:"zip,omitempty"`
}

var defaultCompanyFields = []string{
	"about_us",
	"annualrevenue",
	"city",
	"closedate",
	"company_size",
	"country",
	"createdate",
	"days_to_close",
	"hs_createdate",
	"hs_sequences_is_enrolled",
	"hubspot_owner_assigneddate",
	"hubspot_owner_id",
	"hubspot_team_id",
	"hubspotscore",
	"industry",
	"lastmodifieddate",
	"lifecyclestage",
	"message",
	"name",
	"num_associated_deals",
	"num_notes",
	"num_unique_conversion_events",
	"numemployees",
	"recent_conversion_date",
	"recent_conversion_event_name",
	"recent_deal_amount",
	"recent_deal_close_date",
	"state",
	"total_revenue",
	"website",
	"zip",
}

// Get gets a company.
// In order to bind the get content, a structure must be specified as an argument.
// Also, if you want to gets a custom field, you need to specify the field name.
// If you specify a non-existent field, it will be ignored.
// e.g. &hubspot.RequestQueryOption{ Properties: []string{"custom_a", "custom_b"}}
func (s *CompanyServiceOp) Get(companyID string, company interface{}, option *RequestQueryOption) (*ResponseResource, error) {
	resource := &ResponseResource{Properties: company}
	if err := s.client.Get(s.companyPath+"/"+companyID, resource, option.setupProperties(defaultCompanyFields)); err != nil {
		return nil, err
	}
	return resource, nil
}

// Create creates a new company.
// In order to bind the created content, a structure must be specified as an argument.
// When using custom fields, please embed hubspot.Company in your own structure.
func (s *CompanyServiceOp) Create(company interface{}) (*ResponseResource, error) {
	req := &RequestPayload{Properties: company}
	resource := &ResponseResource{Properties: company}
	if err := s.client.Post(s.companyPath, req, resource); err != nil {
		return nil, err
	}
	return resource, nil
}

// Update updates a company.
// In order to bind the updated content, a structure must be specified as an argument.
// When using custom fields, please embed hubspot.Company in your own structure.
func (s *CompanyServiceOp) Update(companyID string, company interface{}) (*ResponseResource, error) {
	req := &RequestPayload{Properties: company}
	resource := &ResponseResource{Properties: company}
	if err := s.client.Patch(s.companyPath+"/"+companyID, req, resource); err != nil {
		return nil, err
	}
	return resource, nil
}

// AssociateAnotherObj associates company with another HubSpot objects.
// If you want to associate a custom object, please use a defined value in HubSpot.
func (s *CompanyServiceOp) AssociateAnotherObj(companyID string, conf *AssociationConfig) (*ResponseResource, error) {
	resource := &ResponseResource{Properties: &Company{}}
	if err := s.client.Put(s.companyPath+"/"+companyID+"/"+conf.makeAssociationPath(), nil, resource); err != nil {
		return nil, err
	}
	return resource, nil
}
