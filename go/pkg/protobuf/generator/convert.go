package generator

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/mojo-lang/mojo/go/pkg/util"
)

func ConvertGeneratedFiles(files []*util.GeneratedFile) *pluginpb.CodeGeneratorResponse {
	return nil
}

func ConvertCodeGeneratorResponse(response *pluginpb.CodeGeneratorResponse) []*util.GeneratedFile {
	var files []*util.GeneratedFile
	for _, file := range response.File {
		if file.Name != nil && file.Content != nil {
			files = append(files, &util.GeneratedFile{
				Name:                *file.Name,
				Content:             *file.Content,
				SkipIfUserCodeMixed: true,
			})
		} else {
			logs.Warnw("meet an empty file!", "name", *file.Name)
		}
	}
	return files
}
