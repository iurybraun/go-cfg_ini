/*
 * Copyright © 2016 Iury Braun
 * Copyright © 2017 Weyboo
 */

package cfg_ini

import (
    "github.com/go-ini/ini"
    "code.weyboo.com/i18n_ini"
)

var cfg *ini.File

func LoadIni() {
    
    ini, err := ini.InsensitiveLoad("conf.ini")
    if err != nil {
        panic(err.Error())
    }
    
    cfg = ini
    
}

func LoadKey_int(section, key string) int {
    
    result, err := cfg.Section(section).Key(key).Int()
    if err != nil {
        panic(err.Error())
    }
    
    return result
    
}

func LoadKey_string(section, key string) string {
    
    result := cfg.Section(section).Key(key).String()
    
    return result
    
}

func LoadKey_bool(section, key string) bool {
    
    result, err := cfg.Section(section).Key(key).Bool()
    if err != nil {
        panic(err.Error())
    }
    
    return result
    
}

func SectionMapTo(section string, dbase interface{}) {
    
    err := cfg.Section(section).MapTo(dbase)
    if err != nil {
        panic(err.Error())
    }
    
}


func init() {
    
    LoadIni()
    
    // init lang
    i18n_ini.Load_languages(LoadKey_string("core", "available-languages"))
    
}
