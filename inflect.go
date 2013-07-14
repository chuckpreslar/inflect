package inflect

import (
  "fmt"
  "regexp"
)

type Rule struct {
  Regexp  *regexp.Regexp
  Replace string
  Append  bool
}

var plurals = []Rule{
  // Rules should be ordered by precedence.
  Rule{regexp.MustCompile(`(auto|kangaroo|kilo|memo|photo|piano|pimento|pro|solo|soprano|studio|tattoo|video|zoo)$`), `s`, true},
  Rule{regexp.MustCompile(`(er)$`), `ers`, false},
  Rule{regexp.MustCompile(`(s|ss|sh|ch|x|o)$`), `es`, true},
  Rule{regexp.MustCompile(`(a|e|o)y$`), `s`, true},
  Rule{regexp.MustCompile(`(y)$`), `ies`, false},
}

func Plural(str string) string {
  for _, rule := range plurals {
    if rule.Regexp.MatchString(str) {
      if rule.Append {
        return fmt.Sprintf(`%v%v`, str, rule.Replace)
      } else {
        return rule.Regexp.ReplaceAllString(str, rule.Replace)
      }
    }
  }

  return str
}

func Singular(str string) string {
  return str
}
