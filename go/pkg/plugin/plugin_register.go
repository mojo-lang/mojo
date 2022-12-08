package plugin

type parserPlugin struct {
	BasicPlugin
	Parser PackageParser
}

type Register struct {
	plugins      map[string]Plugin
	groupPlugins map[string][]Plugin
}

var register Register

func GetPlugin(name string) Plugin {
	if p, ok := register.plugins[name]; ok {
		return p
	}
	return nil
}

func GetPluginGroup(name string) []Plugin {
	if len(register.groupPlugins) > 0 {
		return register.groupPlugins[name]
	}
	return nil
}

func RegisterPlugin(plugin Plugin) {
	if register.plugins == nil {
		register.plugins = make(map[string]Plugin)
		register.groupPlugins = make(map[string][]Plugin)
	}

	register.plugins[plugin.GetName()] = plugin

	if gn := plugin.GetGroup(); len(gn) > 0 {
		group := register.groupPlugins[gn]
		contained := false
		for _, p := range group {
			if p.GetName() == plugin.GetName() {
				contained = true
				break
			}
		}
		if !contained {
			group = append(group, plugin)
			register.groupPlugins[gn] = group
		}
	}
}
