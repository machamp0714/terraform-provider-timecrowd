package timecrowd

import (
	"context"
	"strconv"
	tc "terraform-provider-timecrowd/timecrowd_client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCategory() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCategoryCreate,
		ReadContext:   resourceCategoryRead,
		UpdateContext: resourceCategoryUpdtae,
		DeleteContext: resourceCategoryDelete,
		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"team_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"color": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"position": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"parent_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ancestry_depth": {
				Type:     schema.TypeInt,
				Computed: true,
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

func resourceCategoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*tc.Client)

	var diags diag.Diagnostics

	teamId := d.Get("team_id").(int)
	params := tc.Category{
		Title:    d.Get("title").(string),
		Color:    d.Get("color").(int),
		ParentId: d.Get("parent_id").(int),
		Position: d.Get("position").(int),
	}

	ca, err := c.CreateCategory(strconv.Itoa(teamId), params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(ca.Id))
	d.Set("team_id", ca.TeamId)

	return diags
}

func resourceCategoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*tc.Client)

	var diags diag.Diagnostics

	teamId := d.Get("team_id").(int)
	categoryId := d.Id()

	cai, _ := strconv.Atoi(categoryId)

	ca, err := c.GetCategory(strconv.Itoa(teamId), cai)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("title", ca.Title)
	d.Set("color", ca.Color)
	d.Set("ancestry_depth", ca.AncestryDepth)
	d.Set("created_at", ca.CreatedAt)
	d.Set("updated_at", ca.UpdatedAt)
	d.SetId(categoryId)

	return diags
}

func resourceCategoryUpdtae(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*tc.Client)

	if d.HasChanges("title", "color", "parent_id", "position") {
		teamId := strconv.Itoa(d.Get("team_id").(int))
		categoryId := d.Id()
		params := tc.Category{
			Title:    d.Get("title").(string),
			Color:    d.Get("color").(int),
			ParentId: d.Get("parent_id").(int),
			Position: d.Get("position").(int),
		}
		_, err := c.UpdateCategory(teamId, categoryId, params)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceCategoryRead(ctx, d, m)
}

func resourceCategoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}
