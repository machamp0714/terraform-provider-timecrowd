package timecrowd

import (
	"context"
	"strconv"

	tc "terraform-provider-timecrowd/timecrowd_client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTeam() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTeamRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"avatar_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_limit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"rounding": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"personal": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"capacity": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"hierarchized": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"default_duration": {
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
			"html_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"invitation_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expires_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expired": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_personal": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"can_manage": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"can_manage_employ": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_payment_required": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"payment_required": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"price_users_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"payable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceTeamRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*tc.Client)

	var diags diag.Diagnostics

	teamId := strconv.Itoa(d.Get("id").(int))

	team, err := c.GetTeam(teamId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("avatar_url", team.AvatarUrl)
	d.Set("time_limit", team.TimeLimit)
	d.Set("rounding", team.Rounding)
	d.Set("name", team.Name)
	d.Set("personal", team.Personal)
	d.Set("capacity", team.Capacity)
	d.Set("hierarchized", team.Hierarchized)
	d.Set("default_duration", team.DefaultDuration)
	d.Set("created_at", team.CreatedAt)
	d.Set("updated_at", team.UpdatedAt)
	d.Set("html_url", team.HtmlUrl)
	d.Set("invitation_url", team.InvitationUrl)
	d.Set("expires_at", team.ExpiresAt)
	d.Set("expired", team.Expired)
	d.Set("is_personal", team.IsPersonal)
	d.Set("can_manage", team.CanManage)
	d.Set("can_manage_employ", team.CanManageEmploy)
	d.Set("is_payment_required", team.IsPaymentRequired)
	d.Set("payment_required", team.PaymentRequired)
	d.Set("price_users_count", team.PriceUsersCount)
	d.Set("payable", team.Payable)
	d.SetId(teamId)

	return diags
}
