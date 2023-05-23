// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/hashicorp/terraform-provider-google/google/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

var BigqueryAnalyticsHubListingIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"location": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"data_exchange_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"listing_id": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type BigqueryAnalyticsHubListingIamUpdater struct {
	project        string
	location       string
	dataExchangeId string
	listingId      string
	d              tpgresource.TerraformResourceData
	Config         *transport_tpg.Config
}

func BigqueryAnalyticsHubListingIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	location, _ := getLocation(d, config)
	if location != "" {
		if err := d.Set("location", location); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
	}
	values["location"] = location
	if v, ok := d.GetOk("data_exchange_id"); ok {
		values["data_exchange_id"] = v.(string)
	}

	if v, ok := d.GetOk("listing_id"); ok {
		values["listing_id"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/dataExchanges/(?P<data_exchange_id>[^/]+)/listings/(?P<listing_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<data_exchange_id>[^/]+)/(?P<listing_id>[^/]+)", "(?P<location>[^/]+)/(?P<data_exchange_id>[^/]+)/(?P<listing_id>[^/]+)", "(?P<listing_id>[^/]+)"}, d, config, d.Get("listing_id").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &BigqueryAnalyticsHubListingIamUpdater{
		project:        values["project"],
		location:       values["location"],
		dataExchangeId: values["data_exchange_id"],
		listingId:      values["listing_id"],
		d:              d,
		Config:         config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("location", u.location); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}
	if err := d.Set("data_exchange_id", u.dataExchangeId); err != nil {
		return nil, fmt.Errorf("Error setting data_exchange_id: %s", err)
	}
	if err := d.Set("listing_id", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting listing_id: %s", err)
	}

	return u, nil
}

func BigqueryAnalyticsHubListingIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	location, _ := getLocation(d, config)
	if location != "" {
		values["location"] = location
	}

	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/dataExchanges/(?P<data_exchange_id>[^/]+)/listings/(?P<listing_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<data_exchange_id>[^/]+)/(?P<listing_id>[^/]+)", "(?P<location>[^/]+)/(?P<data_exchange_id>[^/]+)/(?P<listing_id>[^/]+)", "(?P<listing_id>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &BigqueryAnalyticsHubListingIamUpdater{
		project:        values["project"],
		location:       values["location"],
		dataExchangeId: values["data_exchange_id"],
		listingId:      values["listing_id"],
		d:              d,
		Config:         config,
	}
	if err := d.Set("listing_id", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting listing_id: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *BigqueryAnalyticsHubListingIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyListingUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := tpgresource.GetProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "POST",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
	})
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = tpgresource.Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *BigqueryAnalyticsHubListingIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyListingUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := tpgresource.GetProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "POST",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   u.d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *BigqueryAnalyticsHubListingIamUpdater) qualifyListingUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{BigqueryAnalyticsHubBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/listings/%s", u.project, u.location, u.dataExchangeId, u.listingId), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *BigqueryAnalyticsHubListingIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/listings/%s", u.project, u.location, u.dataExchangeId, u.listingId)
}

func (u *BigqueryAnalyticsHubListingIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-bigqueryanalyticshub-listing-%s", u.GetResourceId())
}

func (u *BigqueryAnalyticsHubListingIamUpdater) DescribeResource() string {
	return fmt.Sprintf("bigqueryanalyticshub listing %q", u.GetResourceId())
}
