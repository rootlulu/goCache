package utils

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

type ConfigOption struct {
	Path string
	Name string
	Type string
}

// TODO: like the mongo options.
// type ConfigOption interface{
//     GetOption() string
//     SetOption() string
// }

type NameOption struct {
	name string
}

type PathOption struct {
	path string
}

type TypeOption struct {
	type_ string
}

func (no *NameOption) GetOption() string {
	return no.name
}

func (po *PathOption) GetOption() string {
	return po.path
}

func (to *TypeOption) GetOption() string {
	return to.type_
}

func Parse(path string, opts ...*ConfigOption) (*viper.Viper, error) {
	var name, type_, suffix string
	for _, opt := range opts {
		if opt.Path != "" {
			path = opt.Path
		}
		if opt.Name != "" {
			name = opt.Name
		}
		if opt.Type != "" {
			type_ = opt.Type
		}
	}
	if name == "" || type_ == "" {
		name = filepath.Base(path)
		suffix = filepath.Ext(path)
		path = filepath.Dir(path)
		switch suffix {
		case ".yaml":
			type_ = "yaml"
		case ".json":
			type_ = "json"
		case ".ini":
			type_ = "config"
		}
	}

	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(type_)

	if err := viper.ReadInConfig(); err != nil {
		if errors.Is(err, viper.ConfigFileNotFoundError{}) {
			return nil, fmt.Errorf("not fount the %s/%s", path, name)
		}
		return nil, fmt.Errorf("Parse the %s/%s in %s type failed: %v", path, name, type_, err)
	}
	fmt.Println(viper.AllSettings())
	return viper.GetViper(), nil
}
