package sdata

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"reflect"
	"strconv"
	"time"

	"github.com/pachyderm/pachyderm/v2/src/internal/errors"
)

func convert(dst, x interface{}) error {
	dv := reflect.ValueOf(dst)
	if dv.Kind() != reflect.Ptr {
		panic("dest must be pointer")
	}

	switch dst := dst.(type) {
	case *bool:
		return asBool(dst, x)
	case *byte:
		return asByte(dst, x)
	case *int8:
		return asInt8(dst, x)
	case *int16:
		return asInt16(dst, x)
	case *int32:
		return asInt32(dst, x)
	case *int64:
		return asInt64(dst, x)
	case *float64:
		return asFloat64(dst, x)
	case *string:
		return asString(dst, x)
	case *[]byte:
		return asBytes(dst, x)
	case *time.Time:
		return asTime(dst, x)
	case *sql.NullBool:
		return asNullBool(dst, x)
	case *sql.NullByte:
		return asNullByte(dst, x)
	case *sql.NullInt16:
		return asNullInt16(dst, x)
	case *sql.NullInt32:
		return asNullInt32(dst, x)
	case *sql.NullInt64:
		return asNullInt64(dst, x)
	case *sql.NullFloat64:
		return asNullFloat64(dst, x)
	case *sql.NullString:
		return asNullString(dst, x)
	case *sql.NullTime:
		return asNullTime(dst, x)
	default:
		return ErrCannotConvert{Dest: dst, Value: x}
	}
}

func asBool(dst *bool, x interface{}) error {
	switch x := x.(type) {
	case bool:
		*dst = x
	case string:
		b, err := strconv.ParseBool(x)
		if err != nil {
			return errors.EnsureStack(err)
		}
		*dst = b
	default:
		return ErrCannotConvert{Dest: dst, Value: x}
	}
	return nil
}

func asByte(dst *byte, x interface{}) error {
	switch x := x.(type) {
	case byte:
		*dst = x
	case string:
		b, err := strconv.ParseUint(x, 10, 8)
		if err != nil {
			return errors.EnsureStack(err)
		}
		*dst = byte(b)
	default:
		return ErrCannotConvert{Dest: dst, Value: x}
	}
	return nil
}

func asInt8(dst *int8, x interface{}) error {
	switch x := x.(type) {
	case int8:
		*dst = x
	case float64:
		*dst = int8(x)
	case string:
		i, err := strconv.ParseInt(x, 10, 8)
		if err != nil {
			return errors.EnsureStack(err)
		}
		*dst = int8(i)
	case json.Number:
		i, err := x.Int64()
		if err != nil {
			return errors.EnsureStack(err)
		}
		*dst = int8(i)
	default:
		return ErrCannotConvert{Dest: dst, Value: x}
	}
	return nil
}

func asInt16(dst *int16, x interface{}) error {
	switch x := x.(type) {
	case int64:
		*dst = int16(x)
	case float64:
		*dst = int16(x)
	case string:
		i, err := strconv.ParseInt(x, 10, 16)
		if err != nil {
			return errors.EnsureStack(err)
		}
		*dst = int16(i)
	case json.Number:
		i, err := x.Int64()
		if err != nil {
			return errors.EnsureStack(err)
		}
		*dst = int16(i)
	default:
		return ErrCannotConvert{Dest: dst, Value: x}
	}
	return nil
}

func asInt32(dst *int32, x interface{}) error {
	switch x := x.(type) {
	case int64:
		*dst = int32(x)
	case string:
		i, err := strconv.ParseInt(x, 10, 32)
		if err != nil {
			return errors.EnsureStack(err)
		}
		*dst = int32(i)
	case json.Number:
		i, err := x.Int64()
		if err != nil {
			return errors.EnsureStack(err)
		}
		*dst = int32(i)
	default:
		return ErrCannotConvert{Dest: dst, Value: x}
	}
	return nil
}

func asInt64(dst *int64, x interface{}) error {
	switch x := x.(type) {
	case int64:
		*dst = x
	case float64:
		*dst = int64(x)
	case string:
		i, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			return errors.EnsureStack(err)
		}
		*dst = i
	case json.Number:
		i, err := x.Int64()
		if err != nil {
			return errors.EnsureStack(err)
		}
		*dst = int64(i)
	default:
		return ErrCannotConvert{Dest: dst, Value: x}
	}
	return nil
}

func asFloat64(dst *float64, x interface{}) error {
	switch x := x.(type) {
	case int64:
		*dst = float64(x)
	case float64:
		*dst = x
	case string:
		f, err := strconv.ParseFloat(x, 64)
		if err != nil {
			return errors.EnsureStack(err)
		}
		*dst = f
	case json.Number:
		f, err := x.Float64()
		if err != nil {
			return errors.EnsureStack(err)
		}
		*dst = float64(f)
	default:
		return ErrCannotConvert{Dest: dst, Value: x}
	}
	return nil
}

func asBytes(dst *[]byte, x interface{}) error {
	switch x := x.(type) {
	case string:
		codec := base64.StdEncoding
		data, err := codec.DecodeString(x)
		if err != nil {
			return errors.EnsureStack(err)
		}
		*dst = append((*dst)[:0], data...)
	default:
		return ErrCannotConvert{Dest: dst, Value: x}
	}
	return nil
}

func asString(dst *string, x interface{}) error {
	switch x := x.(type) {
	case string:
		*dst = x
	default:
		return ErrCannotConvert{Dest: dst, Value: x}
	}
	return nil
}

func asTime(dst *time.Time, x interface{}) error {
	switch x := x.(type) {
	case time.Time:
		*dst = x
	case string:
		t, err := parseTime(x)
		if err != nil {
			return errors.EnsureStack(err)
		}
		*dst = t
	default:
		return ErrCannotConvert{Dest: dst, Value: x}
	}
	return nil
}

func asNullBool(dst *sql.NullBool, x interface{}) error {
	switch x := x.(type) {
	case nil:
		dst.Valid = false
	case string:
		if isNullString(x) {
			dst.Valid = false
			break
		}
		if err := asBool(&dst.Bool, x); err != nil {
			return err
		}
		dst.Valid = true
	default:
		if err := asBool(&dst.Bool, x); err != nil {
			return err
		}
		dst.Valid = true
	}
	return nil
}

func asNullByte(dst *sql.NullByte, x interface{}) error {
	switch x := x.(type) {
	case nil:
		dst.Valid = false
	case string:
		if isNullString(x) {
			dst.Valid = false
			break
		}
		if err := asByte(&dst.Byte, x); err != nil {
			return err
		}
		dst.Valid = true
	default:
		if err := asByte(&dst.Byte, x); err != nil {
			return err
		}
		dst.Valid = true
	}
	return nil
}

func asNullInt16(dst *sql.NullInt16, x interface{}) error {
	switch x := x.(type) {
	case nil:
		dst.Valid = false
	case string:
		if isNullString(x) {
			dst.Valid = false
			break
		}
		if err := asInt16(&dst.Int16, x); err != nil {
			return err
		}
		dst.Valid = true
	default:
		if err := asInt16(&dst.Int16, x); err != nil {
			return err
		}
		dst.Valid = true
	}
	return nil
}

func asNullInt32(dst *sql.NullInt32, x interface{}) error {
	switch x := x.(type) {
	case nil:
		dst.Valid = false
	case string:
		if isNullString(x) {
			dst.Valid = false
			break
		}
		if err := asInt32(&dst.Int32, x); err != nil {
			return err
		}
		dst.Valid = true
	default:
		if err := asInt32(&dst.Int32, x); err != nil {
			return err
		}
		dst.Valid = true
	}
	return nil
}

func asNullInt64(dst *sql.NullInt64, x interface{}) error {
	switch x := x.(type) {
	case nil:
		dst.Valid = false
	case string:
		if isNullString(x) {
			dst.Valid = false
			break
		}
		if err := asInt64(&dst.Int64, x); err != nil {
			return err
		}
		dst.Valid = true
	default:
		if err := asInt64(&dst.Int64, x); err != nil {
			return err
		}
		dst.Valid = true
	}
	return nil
}

func asNullFloat64(dst *sql.NullFloat64, x interface{}) error {
	switch x := x.(type) {
	case nil:
		dst.Valid = false
	case string:
		if isNullString(x) {
			dst.Valid = false
			break
		}
		if err := asFloat64(&dst.Float64, x); err != nil {
			return err
		}
		dst.Valid = true
	default:
		if err := asFloat64(&dst.Float64, x); err != nil {
			return err
		}
		dst.Valid = true
	}
	return nil
}

func asNullString(dst *sql.NullString, x interface{}) error {
	switch x := x.(type) {
	case nil:
		dst.Valid = false
	default:
		if err := asString(&dst.String, x); err != nil {
			return err
		}
		dst.Valid = true
	}
	return nil
}

func asNullTime(dst *sql.NullTime, x interface{}) error {
	switch x := x.(type) {
	case nil:
		dst.Valid = false
	case string:
		if isNullString(x) {
			dst.Valid = false
			break
		}
		if err := asTime(&dst.Time, x); err != nil {
			return err
		}
		dst.Valid = true
	default:
		if err := asTime(&dst.Time, x); err != nil {
			return err
		}
		dst.Valid = true
	}
	return nil
}

// parseTime attempts to parse the time using every format and returns the first one.
func parseTime(x string) (t time.Time, err error) {
	for _, layout := range []string{
		time.RFC3339Nano,
		time.RFC1123Z,
		time.RFC822Z,
		time.Kitchen,
		time.ANSIC,
	} {
		t, err := time.Parse(layout, x)
		if err == nil {
			return t, errors.EnsureStack(err)
		}
	}
	return t, err
}

func isNullString(x string) bool {
	return x == "null" || x == "nil" || len(x) == 0
}
