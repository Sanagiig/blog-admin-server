package response

import (
	"go-blog/model"
	"go-blog/utils"
)

type CategoryTreeData struct {
	*model.Category
	Children []*CategoryTreeData `json:"children"`
}

func (c *CategoryTreeData) GetID() string {
	return c.ID
}

func (c *CategoryTreeData) GetParentID() string {
	return c.ParentID
}

func (c *CategoryTreeData) Append(category utils.TreeNode) {
	c.Children = append(c.Children, category.(*CategoryTreeData))
}
