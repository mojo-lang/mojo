package gokit

/*
import (
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/frankee/ncraft/gengokit"
	templateFileAssets "github.com/frankee/ncraft/gengokit/templates"
	"github.com/frankee/ncraft/svcdef"

	"github.com/frankee/ncraft/gengokit/gentesthelper"
)

var gopath []string

func init() {
	gopath = filepath.SplitList(os.Getenv("GOPATH"))
}

func init() {
	//log.SetLevel(log.DebugLevel)
}

func TestTemplatePathToActual(t *testing.T) {
	pathToWants := map[string]string{
		"NAME-service/":                "package-service/",
		"NAME-service/test.gotemplate": "package-service/test.go",
		"NAME-service/NAME-server":     "package-service/package-server",
	}

	for path, want := range pathToWants {
		if got := templatePathToActual(path, "package"); got != want {
			t.Fatalf("\n`%v` got\n`%v` wanted", got, want)
		}
	}
}

func TestApplyTemplateFromPath(t *testing.T) {
	const def = `
		syntax = "proto3";

		// General package
		package general;

		import "github.com/frankee/ncraft/deftree/googlethirdparty/annotations.proto";

		// RequestMessage is so foo
		message RequestMessage {
			string input = 1;
		}

		// ResponseMessage is so bar
		message ResponseMessage {
			string output = 1;
		}

		// ProtoInterface is a service
		service ProtoInterface {
			// ProtoMethod is simple. Like a gopher.
			rpc ProtoMethod (RequestMessage) returns (ResponseMessage) {
				// No {} in path and no body, everything is in the query
				option (google.api.http) = {
					get: "/route"
				};
			}
		}
	`
	sd, err := svcdef.NewFromString(def, gopath)
	if err != nil {
		t.Fatal(err)
	}

	conf := gengokit.Config{
		Repository: "github.com/frankee/ncraft",
		PBGoPackage: "github.com/frankee/ncraft/gengokit/general-service",
	}

	te, err := gengokit.NewData(sd, conf)
	if err != nil {
		t.Fatal(err)
	}

	end, err := applyTemplateFromPath("svc/endpoints.gotemplate", te)
	if err != nil {
		t.Fatal(err)
	}

	endCode, err := ioutil.ReadAll(end)
	if err != nil {
		t.Fatal(err)
	}

	_, err = testFormat(string(endCode))
	if err != nil {
		t.Fatal(err)
	}
}

func svcMethodsNames(methods []*svcdef.InterfaceMethod) []string {
	var mNames []string
	for _, m := range methods {
		mNames = append(mNames, m.Name)
	}

	return mNames
}

func stringToTemplateExector(def, importPath string) (*gengokit.Data, error) {
	sd, err := svcdef.NewFromString(def, gopath)
	if err != nil {
		return nil, err
	}

	conf := gengokit.Config{
		Repository: importPath,
		PBGoPackage: importPath,
	}

	te, err := gengokit.NewData(sd, conf)
	if err != nil {
		return nil, err
	}

	return te, nil
}

func TestAllTemplates(t *testing.T) {
	const goPackage = "github.com/frankee/ncraft/gengokit"
	const goPBPackage = "github.com/frankee/ncraft/gengokit/general-service"

	const def = `
		syntax = "proto3";

		// General package
		package general;

		import "github.com/frankee/ncraft/deftree/googlethirdparty/annotations.proto";

		// RequestMessage is so foo
		message RequestMessage {
			string input = 1;
		}

		// ResponseMessage is so bar
		message ResponseMessage {
			string output = 1;
		}

		// ProtoInterface is a service
		service ProtoInterface {
			// ProtoMethod is simple. Like a gopher.
			rpc ProtoMethod (RequestMessage) returns (ResponseMessage) {
				// No {} in path and no body, everything is in the query
				option (google.api.http) = {
					get: "/route"
				};
			}
		}
	`

	const def2 = `
		syntax = "proto3";

		// General package
		package general;

		import "github.com/frankee/ncraft/deftree/googlethirdparty/annotations.proto";

		// RequestMessage is so foo
		message RequestMessage {
			string input = 1;
		}

		// ResponseMessage is so bar
		message ResponseMessage {
			string output = 1;
		}

		// ProtoInterface is a service
		service ProtoInterface {
			// ProtoMethod is simple. Like a gopher.
			rpc ProtoMethod (RequestMessage) returns (ResponseMessage) {
				// No {} in path and no body, everything is in the query
				option (google.api.http) = {
					get: "/route"
				};
			}
			// ProtoMethodAgain is simple. Like a gopher again.
			rpc ProtoMethodAgain (RequestMessage) returns (ResponseMessage) {
				// No {} in path and no body, everything is in the query
				option (google.api.http) = {
					get: "/route2"
				};
			}
		}
	`

	sd1, err := svcdef.NewFromString(def, gopath)
	if err != nil {
		t.Fatal(err)
	}

	sd2, err := svcdef.NewFromString(def, gopath)
	if err != nil {
		t.Fatal(err)
	}

	conf := gengokit.Config{
		Repository: "github.com/frankee/ncraft/gengokit",
		PBGoPackage: "github.com/frankee/ncraft/gengokit/general-service",
	}

	data1, err := gengokit.NewData(sd1, conf)
	if err != nil {
		t.Fatal(err)
	}

	data2, err := gengokit.NewData(sd2, conf)
	if err != nil {
		t.Fatal(err)
	}

	for _, templFP := range templateFileAssets.AssetNames() {
		var prev io.Reader

		firstCode, err := testGenerateResponseFile(templFP, data1, prev)
		if err != nil {
			t.Fatalf("%s failed to format on first generation\n\nERROR:\n\n%s\n\nCODE:\n\n%s", templFP, err, firstCode)
		}

		// store the file to pass back to testGenerateResponseFile for second generation
		prev = strings.NewReader(firstCode)

		secondCode, err := testGenerateResponseFile(templFP, data1, prev)
		if err != nil {
			t.Fatalf("%s failed to format on second identical generation\n\nERROR: %s\nCODE:\n\n%s",
				templFP, err, secondCode)
		}

		if len(firstCode) != len(secondCode) {
			t.Fatal("Generated code differs after regeneration with same definition\n" + diff(firstCode, secondCode))
		}

		// store the file to pass back to testGenerateResponseFile for third generation
		prev = strings.NewReader(secondCode)

		// pass in data2 created from def2
		addRPCCode, err := testGenerateResponseFile(templFP, data2, prev)
		if err != nil {
			t.Fatalf("%s failed to format on third generation with 1 rpc added\n\nERROR: %s\nCODE:\n\n%s",
				templFP, err, addRPCCode)
		}

		// store the file to pass back to testGenerateResponseFile for fourth generation
		prev = strings.NewReader(addRPCCode)

		// pass in data1  create from def1
		_, err = testGenerateResponseFile(templFP, data1, prev)
		if err != nil {
			t.Fatalf("%s failed to format on fourth generation with 1 rpc removed\n\nERROR: %s\nCODE:\n\n%s",
				templFP, err, addRPCCode)
		}
	}
}

func diff(a, b string) string {
	return gentesthelper.DiffStrings(
		a,
		b,
	)
}

// testGenerateResponseFile reads the output of generateTemplateFile into a
// string which it returns as this logic needs to be repeated in tests. In
// addition this function will return an error if the code fails to format,
// while generateTemplateFile will not.
func testGenerateResponseFile(templPath string, data *gengokit.Data, prev io.Reader) (string, error) {
	code, err := generateTemplateFile(templPath, data, prev)
	if err != nil {
		return "", err
	}

	// read the code off the io.Reader
	codeBytes, err := ioutil.ReadAll(code)
	if err != nil {
		return "", err
	}

	// format the code
	formatted, err := testFormat(string(codeBytes))
	if err != nil {
		return string(codeBytes), err
	}

	return formatted, nil
}

// testFormat takes a string representing golang code and attempts to return a
// formated copy of that code.
func testFormat(code string) (string, error) {
	formatted, err := format.Source([]byte(code))

	if err != nil {
		return code, err
	}

	return string(formatted), nil
}
*/
