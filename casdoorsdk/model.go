package casdoorsdk

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Model struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk unique index" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`
	UpdatedTime string `xorm:"varchar(100)" json:"updatedTime"`

	DisplayName  string  `xorm:"varchar(100)" json:"displayName"`
	Manager      string  `xorm:"varchar(100)" json:"manager"`
	ContactEmail string  `xorm:"varchar(100)" json:"contactEmail"`
	Type         string  `xorm:"varchar(100)" json:"type"`
	ParentId     string  `xorm:"varchar(100)" json:"parentId"`
	IsTopModel   bool    `xorm:"bool" json:"isTopModel"`
	Users        []*User `xorm:"-" json:"users"`

	Title    string   `json:"title,omitempty"`
	Key      string   `json:"key,omitempty"`
	Children []*Model `json:"children,omitempty"`

	IsEnabled bool `json:"isEnabled"`
}

func (c *Client) GetModels() ([]*Model, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-models", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var models []*Model
	err = json.Unmarshal(bytes, &models)
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (c *Client) GetPaginationModels(p int, pageSize int, queryMap map[string]string) ([]*Model, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-models", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	if response.Status != "ok" {
		return nil, 0, fmt.Errorf(response.Msg)
	}

	bytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var models []*Model
	err = json.Unmarshal(bytes, &models)
	if err != nil {
		return nil, 0, err
	}
	return models, int(response.Data2.(float64)), nil
}

func (c *Client) GetModel(name string) (*Model, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-model", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var model *Model
	err = json.Unmarshal(bytes, &model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (c *Client) UpdateModel(model *Model) (bool, error) {
	_, affected, err := c.modifyModel("update-model", model, nil)
	return affected, err
}

func (c *Client) AddModel(model *Model) (bool, error) {
	_, affected, err := c.modifyModel("add-model", model, nil)
	return affected, err
}

func (c *Client) DeleteModel(model *Model) (bool, error) {
	_, affected, err := c.modifyModel("delete-model", model, nil)
	return affected, err
}
