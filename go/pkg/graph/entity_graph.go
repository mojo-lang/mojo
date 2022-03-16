package graph

import (
    "fmt"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "io"
    "text/template"
)

var graphTemplate = `digraph {
{{- if eq .direction "horizontal" -}}
rankdir=LR;
{{ end -}}
node [shape=box];
{{ range $nodeName, $node := .nodes -}}
"{{ $nodeName }}" [label="{{ $nodeName }}"];
{{ end -}}
{{- range $edgeName, $edge := .edges -}}
"{{ $edge.From.Name }}" -> "{{ $edge.To.Name }}" [label="{{$edge.Name}}",{{if $edge.Inverse}} color="grey80",{{end}} headlabel="{{if $edge.Multiple}}n{{else}}1{{end}}"];
{{end}}
}
`

type EntityGraph struct {
    Nodes map[string]*lang.EntityNode
    Edges map[string]*lang.EntityEdge
}

func NewEntityGraph(pkg *lang.Package) *EntityGraph {
    if pkg != nil && pkg.EntityRelationSet != nil {
        return &EntityGraph{
            Nodes: pkg.EntityRelationSet.Nodes,
            Edges: pkg.EntityRelationSet.Edges,
        }
    }
    return nil
}

func (m *EntityGraph) Render(w io.Writer) error {
    templ, err := template.New("graph").Parse(graphTemplate)
    if err != nil {
        return fmt.Errorf("templ.Parse: %v", err)
    }

    var direction string
    if len(m.Edges) > 15 {
        direction = "horizontal"
    }

    if err := templ.Execute(w, map[string]interface{}{
        "nodes":     m.Nodes,
        "edges":     m.Edges,
        "direction": direction,
    }); err != nil {
        return fmt.Errorf("templ.Execute: %v", err)
    }

    return nil
}
