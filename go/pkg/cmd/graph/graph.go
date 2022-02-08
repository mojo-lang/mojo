package graph

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path"
	"runtime"

	"github.com/goccy/go-graphviz"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/cmd/build/builder"
	graph2 "github.com/mojo-lang/mojo/go/pkg/graph"
)

type Builder struct {
	builder.Builder
}

func (b Builder) Build() error {
	logs.Infow("begin to build go.", "pwd", b.PWD, "path", b.Path)

	entityGraph := graph2.NewEntityGraph(b.Package)
	writer := bytes.NewBuffer(nil)
	if err := entityGraph.Render(writer); err != nil {
		return err
	}

	g := graphviz.New()
	graph, err := graphviz.ParseBytes(writer.Bytes())

	imageFile, err := ioutil.TempDir("", "")
	if err != nil {
		return err
	}

	format := "png"
	//if runtime.GOOS == "windows" {
	//	format = "png"
	//}
	imageFile = path.Join(imageFile, "modv."+format)
	if err := g.RenderFilename(graph, graphviz.Format(format), imageFile); err != nil {
		log.Fatal(err)
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", "-a", "/System/Applications/Preview.app", imageFile)
	case "linux":
		cmd = exec.Command("xdg-open", imageFile)
	case "windows":
		cmd = exec.Command("powershell", "-c", imageFile)
	default:
		return fmt.Errorf("unsupported os: %s", runtime.GOOS)
	}

	return cmd.Run()
}
