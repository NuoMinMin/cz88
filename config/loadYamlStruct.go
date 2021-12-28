package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var viper = new(Yaml)

// GetInstance 获取实例
func GetInstance() *Yaml {
	return viper
}

type TLoadYaml struct {
	Path     string // ./ 文件路径
	Ext      string // .yaml 文件扩展名
	FileName string // config 默认配置文件名
}

func NewLoadYaml() *TLoadYaml {
	return &TLoadYaml{
		Path:     "./",
		Ext:      ".yaml",
		FileName: "config",
	}
}

func init() {
	defaultYamlConfig()
	_ = NewLoadYaml().InitConfig()
}

// Yaml config
type Yaml struct {
	Http string `yaml:"http"`
	Rpc  string `yaml:"rpc"`
	CZip CZip   `yaml:"czip"`
}

type CZip struct {
	FilePath string `yaml:"file_path"`
	Charset  string `yaml:"charset"`
}

func (me *TLoadYaml) GetFileName() string {
	return me.Path + me.FileName + me.Ext
}

func defaultYamlConfig() {
	viper = &Yaml{
		Http: "127.0.0.1:8107",
		Rpc:  "127.0.0.1:8108",
		CZip: CZip{
			FilePath: "czip.txt",
			Charset:  "gb18030",
		},
	}
}

func (me *TLoadYaml) InitConfig() error {
	yamlFile, err := ioutil.ReadFile(me.GetFileName())
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(yamlFile, GetInstance()); err != nil {
		return err
	}
	return nil
}
