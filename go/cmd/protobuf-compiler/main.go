package main

import (
	"fmt"
	"github.com/mojo-lang/lang/go/pkg/lang"
	"github.com/mojo-lang/mojo/go/compiler/cmd/protobuf-compiler/templates"
	"github.com/mojo-lang/mojo/go/compiler/protobuf"
	"github.com/mojo-lang/mojo/go/compiler/util"
	"github.com/mojo-lang/mojo/go/parser/parser"
	flag "github.com/spf13/pflag"
	"os"
	"path/filepath"
	"strings"
)

var (
	fileNameFlag = flag.StringP("file", "f", "", "the mojo file to compile")
	packageFlag  = flag.StringP("package", "p", "", "the mojo package folder to compile")
	verboseFlag  = flag.BoolP("verbose", "v", false, "Verbose output")
	helpFlag     = flag.BoolP("help", "h", false, "Print usage")
)

var binName = filepath.Base(os.Args[0])
var workplace = ""

var (
	// Version is compiled into truss with the flag
	// go install -ldflags "-X main.Version=$SHA"
	Version string
	// BuildDate is compiled into truss with the flag
	// go install -ldflags "-X main.VersionDate=$VERSION_DATE"
	VersionDate string
)

func init() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	workplace = filepath.Dir(ex)
	fmt.Println(workplace)

	// If Version or VersionDate are not set, truss was not built with make
	if Version == "" || VersionDate == "" {
		//rebuild := promptNoMake()
		//if !rebuild {
		//	os.Exit(1)
		//}
		//err := makeAndRunTruss(os.Args)
		//exitIfError(errors.Wrap(err, "please install truss with make manually"))
		//os.Exit(0)
	}

	var buildInfo string
	buildInfo = fmt.Sprintf("version: %s", Version)
	buildInfo = fmt.Sprintf("%s version date: %s", buildInfo, VersionDate)

	flag.Usage = func() {
		if buildInfo != "" && (*verboseFlag || *helpFlag) {
			fmt.Fprintf(os.Stderr, "%s (%s)\n", binName, strings.TrimSpace(buildInfo))
		}
		fmt.Fprintf(os.Stderr, "\nUsage: %s [options] <mojo file>...\n", binName)
		fmt.Fprintf(os.Stderr, "\nGenerates protobuffer files using mojo.\n")
		fmt.Fprintln(os.Stderr, "\nOptions:")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	//if len(flag.Args()) == 0 {
	//	fmt.Fprintf(os.Stderr, "%s: missing .mojo file(s)\n", binName)
	//	flag.Usage()
	//	os.Exit(1)
	//}

	parser := parser.New()

	// compile one file from mojo file to protobuffer file
	if *fileNameFlag != "" {
		f, err := parser.ParseFile(*fileNameFlag)

		if err != nil {
		}
		generateProtoFile(f)
	} else if *packageFlag != "" { // compile mojo package to protobuffer files
		pkg, err := parser.ParsePackage("")
		if err != nil {
		}

		for _, file := range pkg.SourceFiles {
			generateProtoFile(file)
		}
	}
}

func generateProtoFile(f *lang.SourceFile) {
	compiler := protobuf.New()
	protoFile, err := compiler.CompileFile(f)
	if err != nil {
		return
	}

	print(protoFile)

	code, err := util.ApplyTemplate("ProtoTemplate", templates.ProtobufTemplate, protoFile, nil)
	if err != nil {
		println("err...")
		println(err)
		return
	}

	code = util.FormatCode(code)

	println("code...")
	println(code)
}
