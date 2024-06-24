package gen

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/moyu-x/infinite-synthesis/internal/generate"
)

func Run() {
	var flags flag.FlagSet
	options := protogen.Options{
		ParamFunc: flags.Set,
	}

	options.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, file := range plugin.Files {
			if !file.Generate {
				continue
			}
			generate.GenerateFile(plugin, file)
		}
		return nil
	})
}
