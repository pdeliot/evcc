package settings

import (
	"encoding/json"
	"time"

	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/config"
	"github.com/spf13/cast"
)

var _ Settings = (*ConfigSettings)(nil)

type ConfigSettings struct {
	log  *util.Logger
	conf *config.Config
}

func NewConfigSettingsAdapter(log *util.Logger, conf *config.Config) *ConfigSettings {
	return &ConfigSettings{log, conf}
}

func (s *ConfigSettings) get(key string) any {
	return s.conf.Named().Other[key]
}

func (s *ConfigSettings) set(key string, val any) {
	data := s.conf.Named().Other
	data[key] = val
	if err := s.conf.Update(data); err != nil {
		s.log.ERROR.Println(err)
	}
}

func (s *ConfigSettings) SetString(key string, val string) {
	if s == nil {
		return
	}
	s.set(key, val)
}

func (s *ConfigSettings) SetInt(key string, val int64) {
	if s == nil {
		return
	}
	s.set(key, val)
}

func (s *ConfigSettings) SetFloat(key string, val float64) {
	if s == nil {
		return
	}
	s.set(key, val)
}

func (s *ConfigSettings) SetTime(key string, val time.Time) {
	if s == nil {
		return
	}
	s.set(key, val)
}

func (s *ConfigSettings) SetBool(key string, val bool) {
	if s == nil {
		return
	}
	s.set(key, val)
}

func (s *ConfigSettings) SetJson(key string, val any) error {
	if s == nil {
		return nil
	}
	s.set(key, val)
	return nil
}

func (s *ConfigSettings) String(key string) (string, error) {
	if s == nil {
		return "", nil
	}
	return cast.ToStringE(s.get(key))
}

func (s *ConfigSettings) Int(key string) (int64, error) {
	if s == nil {
		return 0, nil
	}
	return cast.ToInt64E(s.get(key))
}

func (s *ConfigSettings) Float(key string) (float64, error) {
	if s == nil {
		return 0, nil
	}
	return cast.ToFloat64E(s.get(key))
}

func (s *ConfigSettings) Time(key string) (time.Time, error) {
	if s == nil {
		return time.Time{}, nil
	}
	return cast.ToTimeE(s.get(key))
}

func (s *ConfigSettings) Bool(key string) (bool, error) {
	if s == nil {
		return false, nil
	}
	return cast.ToBoolE(s.get(key))
}

func (s *ConfigSettings) Json(key string, res any) error {
	str, err := s.String(key)
	if str == "" || err != nil {
		return err
	}
	return json.Unmarshal([]byte(str), &res)
}
