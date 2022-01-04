package timecrowd_client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Category struct {
	Id            int        `json:"id"`
	Title         string     `json:"title"`
	Color         int        `json:"color,omitempty"`
	AncestryDepth int        `json:"ancestry_depth,omitempty"`
	TeamId        int        `json:"team_id,omitempty"`
	CreatedAt     string     `json:"created_at,omitempty"`
	UpdatedAt     string     `json:"updated_at,omitempty"`
	Position      int        `json:"position,omitempty"`
	ParentId      int        `json:"parent_id,omitempty"`
	Children      []Category `json:"children,omitempty"`
}

func (c *Client) CreateCategory(teamId string, params Category) (*Category, error) {
	p, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/teams/%s/categories", c.Host, teamId), strings.NewReader(string(p)))
	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var ca Category
	if err = json.Unmarshal(body, &ca); err != nil {
		return nil, err
	}
	return &ca, nil
}

func (c *Client) UpdateCategory(teamId, categoryId string, params Category) (*Category, error) {
	p, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v1/teams/%s/categories/%s", c.Host, teamId, categoryId), strings.NewReader(string(p)))
	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)

	ca := Category{}
	if err = json.Unmarshal(body, &ca); err != nil {
		return nil, err
	}

	return &ca, nil
}

func (c *Client) DeleteCategory(teamId, categoryId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/teams/%s/categories/%s", c.Host, teamId, categoryId), nil)
	if err != nil {
		return err
	}

	_, err = c.DoRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetCategories(teamId string) (*[]Category, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/teams/%s/categories", c.Host, teamId), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var categories []Category
	err = json.Unmarshal(body, &categories)
	if err != nil {
		return nil, err
	}

	return &categories, nil
}

func (c *Client) GetCategory(teamId string, categoryId int) (*Category, error) {
	categories, err := c.GetCategories(teamId)
	if err != nil {
		return nil, err
	}
	category := findCategory(*categories, categoryId)
	if category == nil {
		return nil, fmt.Errorf("Not Found")
	}
	return category, nil
}

func findCategory(categories []Category, categoryId int) *Category {
	var category *Category
	for _, c := range categories {
		if len(c.Children) != 0 {
			fc := findCategory(c.Children, categoryId)
			if fc != nil {
				category = fc
				break
			}
		}
		if categoryId == c.Id {
			category = &c
			break
		}
	}

	if category != nil {
		return category
	}
	return nil
}
