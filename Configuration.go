package wemvc
import "strings"

type ConnConfig interface {
	GetName() string
	GetType() string
	GetConnString() string
}

type Setting interface {
	GetKey() string
	GetValue() string
}

type MimeConfig interface {
	GetExt() string
	GetMime() string
}

type Configuration interface {
	GetPort() int
	GetDefaultUrl() string
	GetConnConfig(string) ConnConfig
	GetSetting(string) string
	GetMIME(string) string
}

type connConfig struct {
	Name       string `xml:"name,attr"`
	Type       string `xml:"type,attr"`
	ConnString string `xml:"connString,attr"`
}

func (this *connConfig) GetName() string {
	return this.Name
}

func (this *connConfig) GetType() string {
	return this.Type
}

func (this *connConfig) GetConnString() string {
	return this.ConnString
}

type connGroup struct {
	ConfigSource string       `xml:"configSource,attr"`
	ConnStrings  []connConfig `xml:"add"`
}

type setting struct {
	Key   string `xml:"key,attr"`
	Value string `xml:"value,attr"`
}

func (this *setting) GetKey() string {
	return this.Key
}

func (this *setting) GetValue() string {
	return this.Value
}

type settingGroup struct {
	ConfigSource string    `xml:"configSource,attr"`
	Settings     []setting `xml:"add"`
}

type mimeSetting struct {
	FileExe string `xml:"ext,attr"`
	Mime    string `xml:"mime,attr"`
}

func (this *mimeSetting) GetExt() string {
	return this.FileExe
}

func (this *mimeSetting) GetMime() string {
	return this.Mime
}

type mimeGroup struct {
	ConfigSource string        `xml:"configSource,attr"`
	Mimes        []mimeSetting `xml:"add"`
}

type configuration struct {
	Port        int          `xml:"port"`
	DefaultUrl  string       `xml:"defaultUrl"`
	ConnStrings connGroup    `xml:"connStrings"`
	Settings    settingGroup `xml:"settings"`
	Mimes       mimeGroup    `xml:"mime"`

	mimeColl map[string]string
}

func (this *configuration) GetPort() int {
	return this.Port
}

func (this *configuration) GetConnConfig(connName string) ConnConfig {
	for i := 0; i < len(this.ConnStrings.ConnStrings); i++ {
		if this.ConnStrings.ConnStrings[i].Name == connName {
			return &(this.ConnStrings.ConnStrings[i])
		}
	}
	return nil
}

func (this *configuration) GetSetting(key string) string {
	for i := 0; i < len(this.Settings.Settings); i++ {
		if this.Settings.Settings[i].Key == key {
			return this.Settings.Settings[i].Value
		}
	}
	return ""
}

func (this *configuration)GetMIME(ext string) string {
	if len(ext) < 1 {
		return ""
	}
	if this.mimeColl == nil {
		this.mimeColl = make(map[string]string)
		for _, mime := range this.Mimes.Mimes {
			if len(mime.FileExe) < 1 || len(mime.Mime) < 1{
				continue
			}
			this.mimeColl[strings.ToLower(mime.FileExe)] = mime.Mime
		}
	}
	return this.mimeColl[strings.ToLower(ext)]
}

func (this *configuration)GetDefaultUrl() string {
	return this.DefaultUrl
}