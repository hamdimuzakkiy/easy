package csv

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"time"

	_csv "encoding/csv"
	"mime/multipart"

	"github.com/hamdimuzakkiy/easy/safe"
)

type Module struct{}

func New() Module {
	return Module{}
}

func (m Module) Unmarshal(file multipart.File, dest interface{}) (res error) {
	res = safe.Block{
		Try: func() (err error) {
			rows, err := _csv.NewReader(file).ReadAll()
			if err != nil {
				return err
			}

			for _, row := range rows {
				m.chooser(row, dest)
			}

			return res
		}, Catch: func(e safe.Exception) error {
			return errors.New("panic")
		},
	}.Do()
	return res
}

func (m Module) chooser(data []string, dest interface{}) (res error) {
	value := reflect.Indirect(reflect.ValueOf(dest))
	if value.Kind() == reflect.Slice {
		return m.assignedSlice(data, value)
	} else if value.Kind() == reflect.Struct {
		return m.assignedStruct(data, value)
	}

	return errors.New("parameter should be struct or slice")
}

func (m Module) assignedSlice(data []string, value reflect.Value) (err error) {
	_t := reflect.Indirect(value)

	newVal := reflect.Indirect(reflect.New(value.Type().Elem()))
	m.assigning(data, newVal)
	_t = reflect.Append(_t, newVal)
	if reflect.Indirect(value).CanSet() {
		reflect.Indirect(value).Set(_t)
	}

	return err
}

func (m Module) assignedStruct(data []string, v reflect.Value) (err error) {
	uu := reflect.Indirect(reflect.New(v.Type()))
	m.assigning(data, uu)
	reflect.Indirect(v).Set(uu)
	return nil
}

func (m Module) assigning(data []string, v reflect.Value) (err error) {
	typeOf := v.Type()
	fields := typeOf.NumField()
	for i := 0; i < fields; i++ {
		fieldType := typeOf.Field(i)

		tag := fieldType.Tag.Get("csv")
		tags := strings.Split(tag, ";")

		if tag == "" {
			continue
		}

		if tag == "-" {
			switch v.Field(i).Type().Kind() {
			case reflect.Struct:
				m.assignedStruct(data, v.Field(i))
			case reflect.Slice:
				m.assignedStruct(data, v.Field(i))
			}
			continue
		}

		idx, _ := strconv.ParseInt(tags[0], 10, 32)
		format := time.RFC3339
		if len(tags) > 1 {
			format = tags[1]
		}

		assigned(v.Field(i), Data{
			Value:  data[idx],
			Format: format,
		}, fieldType.Type)
	}
	return nil
}

type Data struct {
	Value  string
	Format string
}

func assigned(v reflect.Value, data Data, types reflect.Type) {
	switch types.String() {
	case "string":
		v.SetString(data.Value)
	case "int":
		i, _ := strconv.ParseInt(data.Value, 10, 64)
		v.SetInt(i)
	case "float64":
		i, _ := strconv.ParseFloat(data.Value, 64)
		v.SetFloat(i)
	case "time.Time":
		t, _ := time.Parse(data.Format, data.Value)
		v.Set(reflect.ValueOf(t))
	}
}
