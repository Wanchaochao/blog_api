package collection

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"sync"
	"time"
)

var (
	ScanTagName = "json"
)

var _ctime time.Time
var _ctime_type = reflect.TypeOf(_ctime)

var (
	errScanStructValue    = errors.New("ScanStruct: value must be non-nil pointer to a struct")
	ErrScanStructEmptySrc = errors.New("ScanStruct:src must not be empty")
)

func cannotConvert(d reflect.Value, s interface{}) error {
	var sname string
	switch s.(type) {
	case string:
		sname = "simple string"
	case int64:
		sname = "integer"
	case []byte:
		sname = "bulk string"
	case []interface{}:
		sname = "array"
	default:
		sname = reflect.TypeOf(s).String()
	}
	return fmt.Errorf("cannot convert from %s to %s", sname, d.Type())
}

func convertAssignBulkString(d reflect.Value, s []byte) (err error) {
	ss := string(s)
	if ss == "" {
		return
	}

	fieldType := d.Type()
	switch fieldType.Kind() {
	case reflect.Float32, reflect.Float64:
		var x float64
		x, err = strconv.ParseFloat(ss, d.Type().Bits())
		d.SetFloat(x)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var x int64
		x, err = strconv.ParseInt(ss, 10, d.Type().Bits())
		d.SetInt(x)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var x uint64
		x, err = strconv.ParseUint(ss, 10, d.Type().Bits())
		d.SetUint(x)
	case reflect.Bool:
		var x bool
		x, err = strconv.ParseBool(ss)
		d.SetBool(x)
	case reflect.String:
		d.SetString(ss)
	case reflect.Slice:
		if d.Type().Elem().Kind() != reflect.Uint8 {
			err = cannotConvert(d, s)
		} else {
			d.SetBytes(s)
		}
	case reflect.Struct:
		if fieldType.ConvertibleTo(_ctime_type) {
			t, err := time.ParseInLocation("2006-01-02 15:04:05", ss, time.Local)
			if err != nil {
				return err
			}
			d.Set(reflect.ValueOf(t).Convert(fieldType))
		}

	default:
		err = cannotConvert(d, s)
	}
	return
}

func convertAssignInt(d reflect.Value, s int64) (err error) {
	switch d.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		d.SetInt(s)
		if d.Int() != s {
			err = strconv.ErrRange
			d.SetInt(0)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if s < 0 {
			err = strconv.ErrRange
		} else {
			x := uint64(s)
			d.SetUint(x)
			if d.Uint() != x {
				err = strconv.ErrRange
				d.SetUint(0)
			}
		}
	case reflect.Bool:
		d.SetBool(s != 0)
	default:
		err = cannotConvert(d, s)
	}
	return
}

func convertAssignValue(d reflect.Value, s interface{}) (err error) {
	if d.Kind() == reflect.Ptr && d.CanInterface() {
		// Already a reflect.Ptr
		if d.IsNil() {
			d.Set(reflect.New(d.Type().Elem()))
		}
	}

	switch s := s.(type) {
	case string:
		err = convertAssignBulkString(d, []byte(s))
	case []byte:
		err = convertAssignBulkString(d, s)
	case int64:
		err = convertAssignInt(d, s)
	default:
		err = cannotConvert(d, s)
	}
	return err
}

var (
	structSpecMutex sync.RWMutex
	structSpecCache = make(map[reflect.Type]map[string]string)
)

func structSpecForType(t reflect.Type) map[string]string {
	structSpecMutex.RLock()
	col, found := structSpecCache[t]
	structSpecMutex.RUnlock()
	if found {
		return col
	}

	structSpecMutex.Lock()
	defer structSpecMutex.Unlock()
	col, found = structSpecCache[t]
	if found {
		return col
	}
	//fmt.Println("CACHE MISS")

	col = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		val := t.Field(i).Tag.Get(ScanTagName)
		name := t.Field(i).Name
		if val != "-" {
			col[val] = name
		}
	}

	structSpecCache[t] = col
	return col
}

func ScanStruct(src map[string]string, dest interface{}) error {

	if src == nil || len(src) == 0 {
		return ErrScanStructEmptySrc
	}

	d := reflect.ValueOf(dest)
	if d.Kind() != reflect.Ptr || d.IsNil() {
		return errScanStructValue
	}

	d = d.Elem()
	if d.Kind() != reflect.Struct {
		return errScanStructValue
	}
	dd := reflect.TypeOf(dest).Elem() //通过反射获取type定义
	col := structSpecForType(dd)
	//fmt.Println(col)

	for k, c := range col {
		if err := convertAssignValue(d.FieldByName(c), src[k]); err != nil {
			return fmt.Errorf("ScanStruct: cannot assign field %s: %v", c, err)
		}
	}

	return nil
}
