package handlers

/*
import (
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/testdata"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func init() {
	gopath = filepath.SplitList(os.Getenv("GOPATH"))
}

func TestRenderPrevEndpoints(t *testing.T) {
	var wantEndpoints = `
		package middlewares

		import (
			"github.com/go-kit/kit/endpoint"
			"github.com/frankee/ncraft/gengokit/general-service/svc"
		)

		// WrapEndpoint will be called individually for all endpoints defined in
		// the service. Implement this with the middlewares you want applied to
		// every endpoint.
		func WrapEndpoint(in endpoint.Endpoint) endpoint.Endpoint {
			return in
		}

		// WrapEndpoints takes the service's entire collection of endpoints. This
		// function can be used to apply middlewares selectively to some endpoints,
		// but not others, like protecting some endpoints with authentication.
		func WrapEndpoints(in svc.Endpoints) svc.Endpoints {
			return in
		}

		func BarFoo(err error) bool {
			if err != nil {
				return true
			}
			return false
		}
	`

	_, data, err := generalInterface()
	if err != nil {
		t.Fatal(err)
	}

	middleware := NewMiddlewares()

	middleware.Load(strings.NewReader(wantEndpoints))

	endpoints, err := middleware.Render(MiddlewaresPath, data)
	if err != nil {
		t.Fatal(err)
	}

	endpointsBytes, err := ioutil.ReadAll(endpoints)
	if err != nil {
		t.Fatal(err)
	}

	wantFormatted, endpointFormatted, diff := testdata.DiffGoCode(wantEndpoints, string(endpointsBytes))
	if wantFormatted != endpointFormatted {
		t.Fatalf("Endpoints middleware modified unexpectedly:\n\n%s", diff)
	}
}

func TestRenderUnknownFile(t *testing.T) {
	_, data, err := generalInterface()
	if err != nil {
		t.Fatal(err)
	}

	middleware := NewMiddlewares()

	_, err = middleware.Render("not/valid/file.go", data)
	if err == nil {
		t.Fatalf("This should have produced an error, but didn't")
	}
}

func generalInterface() (*types.Service, *types.Data, error) {
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
		return nil, nil, err
	}
	conf := gengokit.Config{
		Repository: "github.com/frankee/ncraft/gengokit/general-service",
		PBGoPackage: "github.com/frankee/ncraft/gengokit/general-service",
	}

	data, err := gengokit.NewData(sd, conf)
	if err != nil {
		return nil, nil, err
	}

	return sd, data, nil
}
*/
