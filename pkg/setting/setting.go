package setting

import (
	"fmt"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

var sections = make(map[string]interface{})

//读取配置
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config") // name of config file (without extension)
	vp.AddConfigPath("config") // path to look for the config file in
	vp.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	err := vp.ReadInConfig()   // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		fmt.Errorf("Fatal error config file: %s \n", err)
		return nil, err
	}
	s := &Setting{vp: vp}
	return s, err
}

//读取指定的一段
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		fmt.Errorf("vp.UnmarshalKey: %s \n", err)
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}
