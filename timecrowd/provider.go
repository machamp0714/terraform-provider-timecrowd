package timecrowd

import (
	"context"
	"terraform-provider-timecrowd/timecrowd_client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("TIMECROWD_ACCESS_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"timecrowd_category": resourceCategory(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"timecrowd_team":     dataSourceTeam(),
			"timecrowd_category": dataSourceCategory(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctc context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)
	client, err := timecrowd_client.NewClient(&token)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return client, nil
}
