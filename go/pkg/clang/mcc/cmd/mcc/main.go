package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"

	"github.com/mojo-lang/mojo/go/pkg/clang/mcc/parser"
)

// MCC_TARGET_SRC
// MCC_TARGET_PARSER
// MCC_TARGET_COMPILER

func main() {
	args := os.Args
	args = args[1:]
	if p, ok := os.LookupEnv("MCC_TARGET_PARSER"); ok && p == "true" {
		// clang -fsyntax-only -Xclang -ast-dump=json target.c
		var newArgs []string
		for i := 0; i < len(args); i++ {
			switch args[i] {
			case "-MT", "-MF", "-o": // skip the key and value
				i++
			case "-c", "-MD":
			default:
				if i == len(args)-1 && IsCSrcFile(args[i]) {
					src, _ := jsoniter.ConfigDefault.Marshal(newArgs)
					if len(src) > 0 {
						file := args[i] + ".mcc.args.json"
						if err := os.WriteFile(file, src, os.ModePerm); err != nil {
							logs.Warnw("failed to write ast file", "name", file, "error", err)
						}
					}
					newArgs = append(newArgs, "-fsyntax-only", "-Xclang", "-ast-dump=json", args[i])
				} else {
					newArgs = append(newArgs, args[i])
				}
			}
		}

		buf := bytes.NewBuffer(nil)
		out := bytes.NewBuffer(nil)
		cmd := exec.Command("clang", newArgs...)
		cmd.Stdout = buf
		cmd.Stderr = out
		err := cmd.Run()
		if err != nil {
			logs.Warnw("failed to run the clang", "args", newArgs, "error", err, "stderr", out.String())
		} else {
			logs.Debug("output length is ", buf.Len(), "args is ", newArgs)
			if buf.Len() > 0 {
				bufBytes := buf.Bytes()
				file := args[len(args)-1] + ".mcc.json"
				// if err = os.WriteFile(file, bufBytes, os.ModePerm); err != nil {
				// 	logs.Warnw("failed to write file", "name", file, "error", err)
				// }

				psr := parser.Parser{}
				source, err := psr.ParseJSON(bufBytes)
				source.Name = args[len(args)-1]
				if err != nil {
					logs.Warnw("failed to parse the clang dump ast json file", "name", file, "error", err)
				} else {
					src, err := jsoniter.ConfigDefault.MarshalIndent(source, "", "  ")
					if err != nil {
						logs.Warnw("failed to marshal the output ast to json file", "name", file, "error", err)
					} else {
						file = args[len(args)-1] + ".mcc.ast.json"
						if err = os.WriteFile(file, src, os.ModePerm); err != nil {
							logs.Warnw("failed to write ast file", "name", file, "error", err)
						}
					}
				}
			}
		}
	}

	cmd := exec.Command("clang", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		logs.Warnw("failed to run the clang", "args", strings.Join(args, " "), "error", err, "stderr", string(out))
		os.Exit(1)
		return
	}
}

func IsCSrcFile(file string) bool {
	return strings.HasSuffix(file, ".c")
}

func IsCppSrcFile(file string) bool {
	return strings.HasSuffix(file, ".c++") || strings.HasSuffix(file, ".cpp") || strings.HasSuffix(file, ".cxx") || strings.HasSuffix(file, ".cc")
}
