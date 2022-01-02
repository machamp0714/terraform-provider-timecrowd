package timecrowd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

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
	c := m.(*Config)

	var diags diag.Diagnostics

	teamId := strconv.Itoa(d.Get("id").(int))

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://timecrowd.net/api/v1/teams/%s", teamId), nil)
	if err != nil {
		return diag.FromErr(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	res, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return diag.FromErr(err)
	}
	if res.StatusCode != http.StatusOK {
		return diag.FromErr(err)
	}

	team := Team{}
	err = json.Unmarshal(body, &team)
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

type Team struct {
	ID                int    `json:"id"`
	AvatarUrl         string `json:"avatar_url"`
	TimeLimit         int    `json:"time_limit"`
	Rounding          string `json:"rounding"`
	Name              string `json:"name"`
	Personal          bool   `json:"personal"`
	Capacity          int    `json:"capacity"`
	Hierarchized      bool   `json:"hierarchized"`
	DefaultDuration   int    `json:"default_duration"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	HtmlUrl           string `json:"html_url"`
	InvitationUrl     string `json:"invitation_url"`
	ExpiresAt         string `json:"expires_at"`
	Expired           bool   `json:"expired"`
	IsPersonal        bool   `json:"is_personal"`
	CanManage         bool   `json:"can_manage"`
	CanManageEmploy   bool   `json:"can_manage_employ"`
	IsPaymentRequired bool   `json:"is_payment_required"`
	PaymentRequired   bool   `json:"payment_required"`
	PriceUsersCount   int    `json:"price_users_count"`
	Payable           bool   `json:"payable"`
}
