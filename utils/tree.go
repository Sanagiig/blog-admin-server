package utils

type TreeNode interface {
	GetID() string
	GetParentID() string
	Append(node TreeNode)
}

type TreeData struct {
	ID       string     `json:"id"`
	ParentID string     `json:"parentID"`
	Children []TreeNode `json:"children"`
}

func (t *TreeData) GetID() string {
	return t.ID
}

func (t *TreeData) Append(node TreeNode) {
	t.Children = append(t.Children, node)
}

func (t *TreeData) GetParentID() string {
	return t.ParentID
}

func List2TreeMap(list []TreeNode) []TreeNode {
	var roots []TreeNode = nil
	for i := 0; i < len(list); i++ {
		count := 0
		nodeI := list[i]
		for j := 0; j < len(list); j++ {

			nodeJ := list[j]
			iParent := nodeI.GetParentID()
			jID := nodeJ.GetID()
			//if nodeI.GetParentID() == nodeJ.GetID() {
			if iParent == jID {
				nodeJ.Append(nodeI)
				count++
				break
			}
		}
		if count == 0 {
			roots = append(roots, nodeI)
		}
	}

	return roots
}

func TT(t interface{}) TreeData {
	return t.([]TreeData)[0]
}
