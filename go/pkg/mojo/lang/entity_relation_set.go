package lang

func NewEntityRelationSet() *EntityRelationSet {
	return &EntityRelationSet{
		Nodes:         make(map[string]*EntityNode),
		Edges:         make(map[string]*EntityEdge),
		NodeEdges:     make(map[string]*EntityEdges),
		NodeRelations: make(map[string]*EntityRelations),
	}
}

func (x *EntityRelationSet) GetNode(name string) *EntityNode {
	if x != nil {
		return x.Nodes[name]
	}
	return nil
}
