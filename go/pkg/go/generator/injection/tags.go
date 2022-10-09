package injection

import "github.com/fatih/structtag"

func HasTagOption(tag *structtag.Tag, option string) bool {
	return tag.Name == option || tag.HasOption(option)
}
