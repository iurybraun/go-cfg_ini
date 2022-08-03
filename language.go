/*
 * Copyright © 2016-2020 Iury Braun
 * Copyright © 2017-2020 Weyboo
 */

package cfg_ini

import (
    "net/http"
)

func GetLang(r *http.Request) string {
    var lang = r.FormValue("language")
    
    if lang != "" {
        return lang
    } else {
        return LoadKey_string("core", "default-language")
    }
}
