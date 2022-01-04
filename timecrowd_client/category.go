package timecrowd_client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Category struct {
	Id            int        `json:"id"`
	Title         string     `json:"title"`
	Color         int        `json:"color"`
	AncestryDepth int        `json:"ancestry_depth"`
	TeamId        int        `json:"team_id"`
	CreatedAt     string     `json:"created_at"`
	UpdatedAt     string     `json:"updated_at"`
	Children      []Category `json:"children"`
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
