package timecrowd

import (
	"context"
	"strconv"
	tc "terraform-provider-timecrowd/timecrowd_client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCategory() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCategoryRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"title": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"color": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ancestry_depth": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"team_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCategoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*tc.Client)

	var diags diag.Diagnostics

	categoryId := d.Get("id").(int)
	teamId := strconv.Itoa(d.Get("team_id").(int))

	category, err := c.GetCategory(teamId, categoryId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(categoryId))
	d.Set("title", category.Title)
	d.Set("color", category.Color)
	d.Set("ancestry_depth", category.AncestryDepth)
	d.Set("team_id", category.TeamId)
	d.Set("created_at", category.CreatedAt)
	d.Set("updated_at", category.UpdatedAt)

	return diags
}
