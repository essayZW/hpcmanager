package db

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// JSON 数据库JSON类型
type JSON map[string]interface{}

// Scan 数据库json格式的字符串转化为JSON类型
func (j *JSON) Scan(src interface{}) error {
	var source map[string]interface{}
	switch src.(type) {
	case nil:
		source = make(map[string]interface{})
	case []byte:
		if err := json.Unmarshal(src.([]byte), &source); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Invalid type %t for JSON", src)
	}
	*j = JSON(source)
	return nil
}

// Value JSON格式类型转成数据库类型
func (j *JSON) Value() (driver.Value, error) {
	if j == nil {
		return driver.Value(nil), nil
	}
	strbytes, err := json.Marshal(j)
	if err != nil {
		return driver.Value(""), fmt.Errorf("Encode json error: %v", err)
	}
	return driver.Value(string(strbytes)), nil
}

func (j JSON) String() string {
	res, _ := json.Marshal(j)
	return string(res)
}

// NewJSON 从字符串生成JSON数据
func NewJSON(jsonStr string) (*JSON, error) {
	if jsonStr == "" {
		return nil, nil
	}
	var res JSON
	if err := json.Unmarshal([]byte(jsonStr), &res); err != nil {
		return nil, err
	}
	return &res, nil
}
