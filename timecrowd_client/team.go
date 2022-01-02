package timecrowd_client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

func (c *Client) GetTeam(teamId string) (*Team, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/teams/%s", c.Host, teamId), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	team := Team{}
	err = json.Unmarshal(body, &team)
	if err != nil {
		return nil, err
	}

	return &team, nil
}
