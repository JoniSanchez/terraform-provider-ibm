package ibm

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/bluemix-go/models"
)

func resourceIBMCertificateManagerOrder() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMCertificateManagerOrderCertificate,
		Read:     resourceIBMCertificateManagerRead,
		Update:   resourceIBMCertificateManagerRenew,
		Importer: &schema.ResourceImporter{},
		Delete:   resourceIBMCertificateManagerDelete,
		Exists:   resourceIBMCertificateManagerExists,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"certificate_manager_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domains": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
			"rotate_keys": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_validation_method": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "dns - 01",
			},
			"dns_provider_instance_crn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"issuer": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"begins_on": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"expires_on": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"imported": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"has_previous": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"issuance_info": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ordered_on": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"additional_info": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceIBMCertificateManagerOrderCertificate(d *schema.ResourceData, meta interface{}) error {

	cmService, err := meta.(ClientSession).CertificateManagerAPI()
	if err != nil {
		return err
	}

	instanceID := d.Get("certificate_manager_instance_id").(string)
	name := d.Get("name").(string)
	var description string
	if desc, ok := d.GetOk("description"); ok {
		description = desc.(string)
	}
	domainValidationMethod := d.Get("domain_validation_method").(string)

	var dnsProviderInstanceCrn string
	if dnsInsCrn, ok := d.GetOk("dns_provider_instance_crn"); ok {
		dnsProviderInstanceCrn = dnsInsCrn.(string)
	}
	var domainList = make([]string, 0)
	if domains, ok := d.GetOk("domains"); ok {
		for _, domain := range domains.([]interface{}) {
			domainList = append(domainList, fmt.Sprintf("%v", domain))
		}
	}
	client := cmService.Certificate()
	payload := models.CertificateOrderData{Name: name, Description: description, Domains: domainList, DomainValidationMethod: domainValidationMethod, DNSProviderInstanceCrn: dnsProviderInstanceCrn}
	result, err := client.OrderCertificate(instanceID, payload)
	if err != nil {
		return err
	}
	d.SetId(result.ID)

	_, err = waitForCertificateOrder(d, meta)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for Ordering Certificate (%s) to be succeeded: %s", d.Id(), err)
	}

	return resourceIBMCertificateManagerRead(d, meta)
}
func resourceIBMCertificateManagerRead(d *schema.ResourceData, meta interface{}) error {
	cmService, err := meta.(ClientSession).CertificateManagerAPI()
	if err != nil {
		return err
	}
	certID := d.Id()
	certificatedata, err := cmService.Certificate().GetMetaData(certID)
	if err != nil {
		return err
	}
	cminstanceid := strings.Split(certID, ":certificate:")
	d.Set("certificate_manager_instance_id", cminstanceid[0]+"::")
	d.Set("name", certificatedata.Name)
	d.Set("domains", certificatedata.Domains)
	d.Set("domain_validation_method", "dns-01")
	d.Set("rotate_keys", certificatedata.RotateKeys)
	d.Set("description", certificatedata.Description)
	d.Set("begins_on", certificatedata.BeginsOn)
	d.Set("expires_on", certificatedata.ExpiresOn)
	d.Set("imported", certificatedata.Imported)
	d.Set("status", certificatedata.Status)
	d.Set("algorithm", certificatedata.Algorithm)
	d.Set("key_algorithm", certificatedata.KeyAlgorithm)
	d.Set("issuer", certificatedata.Issuer)
	d.Set("has_previous", certificatedata.HasPrevious)

	if certificatedata.IssuanceInfo != nil {
		issuanceinfo := map[string]interface{}{}
		issuanceinfo["status"] = certificatedata.IssuanceInfo.Status
		issuanceinfo["code"] = certificatedata.IssuanceInfo.Code
		issuanceinfo["additional_info"] = certificatedata.IssuanceInfo.AdditionalInfo
		issuanceinfo["ordered_on"] = certificatedata.IssuanceInfo.OrderedOn
		d.Set("issuance_info", issuanceinfo)
	}
	return nil
}

func resourceIBMCertificateManagerRenew(d *schema.ResourceData, meta interface{}) error {
	cmService, err := meta.(ClientSession).CertificateManagerAPI()
	if err != nil {
		return err
	}
	certID := d.Id()
	client := cmService.Certificate()

	if d.HasChange("rotate_keys") {
		rotateKeys := d.Get("rotate_keys").(bool)
		payload := models.CertificateRenewData{RotateKeys: rotateKeys}

		_, err := client.RenewCertificate(certID, payload)
		if err != nil {
			return err
		}
	}
	if d.HasChange("name") || d.HasChange("description") {
		name := d.Get("name").(string)
		description := d.Get("description").(string)
		payload := models.CertificateMetadataUpdate{Name: name, Description: description}

		err := client.UpdateCertificateMetaData(certID, payload)
		if err != nil {
			return err
		}
	}
	_, err = waitForCertificateRenew(d, meta)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for Renew Certificate (%s) to be succeeded: %s", d.Id(), err)
	}
	return resourceIBMCertificateManagerRead(d, meta)
}
func waitForCertificateOrder(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	cmService, err := meta.(ClientSession).CertificateManagerAPI()
	if err != nil {
		return false, err
	}
	certID := d.Id()

	stateConf := &resource.StateChangeConf{
		Pending: []string{"pending"},
		Target:  []string{"valid"},
		Refresh: func() (interface{}, string, error) {
			getcert, err := cmService.Certificate().GetMetaData(certID)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return nil, "", fmt.Errorf("The certificate %s does not exist anymore: %v", d.Id(), err)
				}
				return nil, "", err
			}
			if getcert.Status == "failed" {
				return getcert, getcert.Status, fmt.Errorf("The certificate %s failed: %v", d.Id(), err)
			}
			return getcert, getcert.Status, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      60 * time.Second,
		MinTimeout: 60 * time.Second,
	}

	return stateConf.WaitForState()
}
func waitForCertificateRenew(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	cmService, err := meta.(ClientSession).CertificateManagerAPI()
	if err != nil {
		return false, err
	}
	certID := d.Id()

	stateConf := &resource.StateChangeConf{
		Pending: []string{"pending"},
		Target:  []string{"valid"},
		Refresh: func() (interface{}, string, error) {
			getcert, err := cmService.Certificate().GetMetaData(certID)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return nil, "", fmt.Errorf("The certificate %s does not exist anymore: %v", d.Id(), err)
				}
				return nil, "", err
			}
			if getcert.Status == "failed" {
				return getcert, getcert.Status, fmt.Errorf("The certificate %s failed: %v", d.Id(), err)
			}
			return getcert, getcert.Status, nil
		},
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      60 * time.Second,
		MinTimeout: 60 * time.Second,
	}

	return stateConf.WaitForState()
}
