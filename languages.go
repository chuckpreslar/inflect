package inflect

import (
  "github.com/chuckpreslar/inflect/languages"
  "github.com/chuckpreslar/inflect/types"
)

var (
  LANGUAGE = "en"
  LANGUAGES = map[string]*types.LanguageType {
    "en": languages.English,
  }
)