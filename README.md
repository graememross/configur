# Configur

An opinionated configuration package to allow easy configuration of 
Go command line programs

`Configur` is a simple wrapper around github.com/knadh/koanf to allow simple implementation of a standard configuration system.

To load a configuration from a system, simply
```go
    	config := configur.ParseConfig("tester")
``` 
`config` is returned holding a `ConfigSet` which is a simple wrapper around a **koanf.Koanf* object with all the usual accessors and functions.

This configuration is loaded from the following sources (in the above example the `<key>` would be "tester")

```
 1. Read config files
   	a. /etc/
  	 	<key>.yaml
	    	<key>.json
			<key>.ini
			<key>
		b. homedir/<key>
			config.yaml
			config.json
			config.ini
			config
		c. homedir
			.<key>.yaml
			.<key>.json
			.<key>.ini
			.<key>
		d. Current working directory
			.<key>.yaml
			.<key>.json
			.<key>.ini
			.<key>
 2. Read environment variables
		parse environment variables in the form
		KEY_one_two as one.two
```
### Installation

`go get -u github.com/graememross/configur`

### Usage

Instructions to access configs with examples ... TODO

### Reading flags from the command line

This is possible, however not worked out the details yet .... TODO

## API

This is simply an extension of the `koanf` api with a couple of extra methods  