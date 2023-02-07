package response

import (
	"go-blog/model"
	"go-blog/utils"
)

type DictionaryTreeData struct {
	*model.Dictionary
	Children []*DictionaryTreeData `json:"children"`
}

func (c *DictionaryTreeData) GetID() string {
	return c.ID
}

func (c *DictionaryTreeData) GetParentID() string {
	return c.ParentID
}

func (c *DictionaryTreeData) Append(dictionary utils.TreeNode) {
	c.Children = append(c.Children, dictionary.(*DictionaryTreeData))
}
