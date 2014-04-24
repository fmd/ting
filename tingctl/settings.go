package main

import (
    "fmt"
    "errors"
    "reflect"
    "github.com/fmd/goting/ting"
)

func settingsSet(key string, value string) error {
    s, err := ting.LoadSettings()
    if err != nil {
        return err
    }

    fieldname := s.FieldNameByJsonTag(key)
    if len(fieldname) == 0 {
        return errors.New(fmt.Sprintf("Could not find field in settings with key '%s'", key))
    }

    reflect.ValueOf(s).Elem().FieldByName(fieldname).SetString(value)

    if err = s.Save(); err != nil {
        return err
    }

    return nil
}

func settingsGet(key string) string {
    s, err := ting.LoadSettings()
    if err != nil {
        return ""
    }

    fieldname := s.FieldNameByJsonTag(key)
    if len(fieldname) == 0 {
        return ""
    }

    return reflect.ValueOf(s).Elem().FieldByName(fieldname).String()
}