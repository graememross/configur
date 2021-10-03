package configur

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/parsers/hcl"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/spf13/pflag"
)

type ConfigSet struct {
	*koanf.Koanf
}

var k = &ConfigSet{
	koanf.New("."),
}

// Parse configuration
// Do it in this order
// 1. Read config files
//   	a. /etc/
//  	 	<key>.yaml
//	    	<key>.json
//			<key>.ini
//			<key>
//		b. homedir/<key>
//			config.yaml
//			config.json
//			config.ini
//			config
//		c. homedir
//			.<key>.yaml
//			.<key>.json
//			.<key>.ini
//			.<key>
//		d. Current working directory
//			.<key>.yaml
//			.<key>.json
//			.<key>.ini
//			.<key>
// 2. Read environment variables
//		parse environment variables in the form
//		KEY_one_two as one.two
// 3. Read command lines
//		???
func ParseConfig(prefix string) *ConfigSet {
	// Read in all the config files in an opinionated order
	scanFiles(prefix)
	// Override with environment variables
	scanEnvironment(prefix)
	// Override with command line variables
	return k
}

var parserMap = map[string]koanf.Parser{
	"yaml": yaml.Parser(),
	"json": json.Parser(),
	"env":  dotenv.Parser(),
	"hcl":  hcl.Parser(true),
	"toml": toml.Parser(),
}

type ParsePair struct {
	fName string
	pType string
}

func scanFiles(name string) {
	homedir, _ := os.UserHomeDir()
	searchPaths := []ParsePair{
		{fmt.Sprintf("%s/%s.%s", "/etc", name, "yaml"), "yaml"},
		{fmt.Sprintf("%s/%s.%s", "/etc", name, "json"), "json"},
		{fmt.Sprintf("%s/%s.%s", "/etc", name, "env"), "env"},
		{fmt.Sprintf("%s/%s", "/etc", name), "env"},
		{fmt.Sprintf("%s/.%s/config.%s", homedir, name, "yaml"), "yaml"},
		{fmt.Sprintf("%s/.%s/config.%s", homedir, name, "json"), "json"},
		{fmt.Sprintf("%s/.%s/config.%s", homedir, name, "env"), "env"},
		{fmt.Sprintf("%s/.%s/config", homedir, name), "env"},
		{fmt.Sprintf("%s/.%s.%s", homedir, name, "yaml"), "yaml"},
		{fmt.Sprintf("%s/.%s.%s", homedir, name, "json"), "json"},
		{fmt.Sprintf("%s/.%s.%s", homedir, name, "env"), "env"},
		{fmt.Sprintf("%s/.%s", homedir, name), "env"},
		{fmt.Sprintf("%s.%s", name, "yaml"), "yaml"},
		{fmt.Sprintf("%s.%s", name, "json"), "json"},
		{fmt.Sprintf("%s.%s", name, "env"), "env"},
		{name, "env"},
	}
	for _, p := range searchPaths {
		readFile(p)
	}

}
func readFile(name ParsePair) {
	b, err := ioutil.ReadFile(name.fName)
	if err != nil {
		fmt.Println("WARNING: ", err)
		return
	}
	k.Load(rawbytes.Provider(b), parserMap[name.pType])
}
func scanEnvironment(prefix string) {
	key := fmt.Sprintf("%s_", strings.ToUpper(prefix))
	k.Load(env.Provider(key, ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, key)), "_", ".", -1)
	}), nil)

}

func (c *ConfigSet) BindValues(aFlagset *pflag.FlagSet) {
	if aFlagset != nil {
		c.Load(posflag.Provider(aFlagset, ".", c.Koanf), nil)
	} else {
		fmt.Println("Warning attempt to BindValue on a nil Flag")
	}
}

func GetConfigSet() *ConfigSet {
	return k
}
