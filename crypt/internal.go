package gpcrypt

import (
	"encoding/base64"
	"fmt"
	"hash"
	"reflect"
	"sort"
	"time"
)

func digest(input interface{}, hasher hash.Hash) (string, error) {

	if err := process(input, &hasher); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(hasher.Sum(nil)), nil
}

func process(in interface{}, hasher *hash.Hash) error {

	if in == nil {
		_, err := fmt.Fprint(*hasher, reflect.ValueOf(nil))
		if err != nil {
			return err
		}

		return nil
	}

	if t, ok := in.(time.Time); ok {
		tStr := t.Format(time.RFC3339)
		if err := process(tStr, hasher); err != nil {
			return err
		}
	}

	fieldValue := reflect.Indirect(reflect.ValueOf(in))
	fieldKind := fieldValue.Type().Kind()
	if !fieldValue.IsValid() || fieldValue.IsZero() {
		return nil
	}

	switch fieldKind {
	case reflect.Map:
		keyHash := make([]string, len(fieldValue.MapKeys()))
		keyHashValue := make(map[string]reflect.Value)

		for i, key := range fieldValue.MapKeys() {
			kh, err := digest(key.Interface(), *hasher)
			if err != nil {
				return err
			}

			keyHash[i] = kh
			keyHashValue[kh] = fieldValue.MapIndex(key)
		}

		sort.Strings(keyHash)

		for _, kh := range keyHash {
			_, err := fmt.Fprint(*hasher, kh)
			if err != nil {
				return err
			}
			vh, err := digest(keyHashValue[kh].Interface(), *hasher)
			if err != nil {
				return err
			}

			_, err = fmt.Fprint(*hasher, vh)
			if err != nil {
				return err
			}
		}

	case reflect.Struct, reflect.Ptr:
		for i := 0; i < fieldValue.NumField(); i++ {
			fieldTag := fieldValue.Type().Field(i).Tag.Get("hash")
			fv := fieldValue.Field(i)

			if fv.IsZero() || !fv.IsValid() || fieldTag == "ignore" {
				continue
			}

			var valOf interface{}
			if reflect.Indirect(fv).CanInterface() {
				valOf = reflect.Indirect(fv).Interface()
			} else {
				return nil
			}

			if err := process(valOf, hasher); err != nil {
				return nil
			}
		}

	case reflect.Slice, reflect.Array:
		var hashesAr []string
		for it := 0; it < fieldValue.Len(); it++ {
			itH, err := digest(reflect.Indirect(fieldValue.Index(it)).Interface(), *hasher)
			if err != nil {
				return err
			}

			hashesAr = append(hashesAr, itH)
		}

		sort.Strings(hashesAr)
		for _, h := range hashesAr {
			err := process(h, hasher)
			if err != nil {
				return err
			}
		}

	default:
		if _, err := fmt.Fprint(*hasher, reflect.ValueOf(fieldValue).Interface()); err != nil {
			return nil
		}
	}

	return nil
}
