package ncloud

import (
	"fmt"
	"regexp"
	"time"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/server"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceNcloudAccessControlRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNcloudAccessControlRulesRead,

		Schema: map[string]*schema.Schema{
			"access_control_group_configuration_no": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Access control group setting number to search",
			},
			"source_access_control_rule_name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateRegexp,
				Description:  "A regex string to apply to the ACG rule list returned by ncloud",
			},
			"access_control_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_control_rule_configuration_no": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocol_type": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     commonCodeSchemaResource,
						},
						"source_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"destination_port": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_access_control_rule_configuration_no": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_access_control_rule_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"access_control_rule_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"output_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceNcloudAccessControlRulesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NcloudAPIClient)

	d.SetId(time.Now().UTC().String())

	id := d.Get("access_control_group_configuration_no").(string)
	reqParams := server.GetAccessControlRuleListRequest{AccessControlGroupConfigurationNo: ncloud.String(id)}

	logCommonRequest("GetAccessControlRuleList", reqParams)

	resp, err := client.server.V2Api.GetAccessControlRuleList(&reqParams)
	if err != nil {
		logErrorResponse("GetAccessControlRuleList", err, id)
		return err
	}
	logCommonResponse("GetAccessControlRuleList", GetCommonResponse(resp))

	allAccessControlRuleList := resp.AccessControlRuleList
	var filteredAccessControlRuleList []*server.AccessControlRule
	nameRegex, nameRegexOk := d.GetOk("source_access_control_rule_name_regex")
	if nameRegexOk {
		r := regexp.MustCompile(nameRegex.(string))
		for _, rule := range allAccessControlRuleList {
			if r.MatchString(ncloud.StringValue(rule.SourceAccessControlRuleName)) {
				filteredAccessControlRuleList = append(filteredAccessControlRuleList, rule)
			}
		}
	} else {
		filteredAccessControlRuleList = allAccessControlRuleList[:]
	}

	if len(filteredAccessControlRuleList) < 1 {
		return fmt.Errorf("no results. please change search criteria and try again")
	}

	return accessControlRulesAttributes(d, filteredAccessControlRuleList)
}

func accessControlRulesAttributes(d *schema.ResourceData, accessControlRules []*server.AccessControlRule) error {
	var ids []string

	for _, accessControlRule := range accessControlRules {
		ids = append(ids, ncloud.StringValue(accessControlRule.AccessControlRuleConfigurationNo))
	}
	d.SetId(dataResourceIdHash(ids))

	if err := d.Set("access_control_rules", flattenAccessControlRules(accessControlRules)); err != nil {
		return err
	}

	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		writeToFile(output.(string), d.Get("access_control_rules"))
	}

	return nil
}
