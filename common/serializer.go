package common

import (
	"k8s.io/apimachinery/pkg/util/json"
	"reflect"
	"strconv"
)

func StringToInt64(str string) int64 {
	parsedInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return parsedInt
}

type ModelSerializer struct {
	Model    interface{}
	IsPlural bool
}

type Serializable interface {
	Serialize() interface{}
}

func (s *ModelSerializer) Serialize() interface{} {

	if s.IsPlural {
		var resList []map[string]interface{}

		var models []interface{}

		// https://stackoverflow.com/questions/52479739/how-to-convert-interface-to-interfaces-slice
		// interface{} to []interface{}, this solution is bravo!
		rv := reflect.ValueOf(s.Model)
		if rv.Kind() == reflect.Slice {
			for i := 0; i < rv.Len(); i++ {
				models = append(models, rv.Index(i).Interface())
			}
		}

		for _, model := range models {
			var tempRes map[string]interface{}
			_json, _ := json.Marshal(model)
			_ = json.Unmarshal(_json, &tempRes)

			for k, v := range tempRes {
				if k == "id" || k == "created_by" || k == "updated_by" || k == "parent_id" {
					tempRes[k] = strconv.FormatInt(v.(int64), 10)
				}
			}

			resList = append(resList, tempRes)
		}
		return resList

	} else {
		var res map[string]interface{}
		_json, _ := json.Marshal(s.Model)
		_ = json.Unmarshal(_json, &res)

		for k, v := range res {
			if k == "id" || k == "created_by" || k == "updated_by" || k == "parent_id" {
				res[k] = strconv.FormatInt(v.(int64), 10)
			}
		}

		return res
	}
}
