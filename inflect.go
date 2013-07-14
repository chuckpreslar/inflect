package inflect

import (
  "fmt"
  "regexp"
)

// Rule struct represents a linguistic pluralization/singularization rule.
type Rule struct {
  Regexp  *regexp.Regexp // Regular expression the string must match.
  Replace string         // String to replace a match with.
  Append  bool           // Append to the end of a matched string?
}

// List of pluralization rules in order by precedence.
var plurals = []Rule{
  Rule{regexp.MustCompile(`(auto|kangaroo|kilo|memo|photo|piano|pimento|pro|solo|soprano|studio|tattoo|video|zoo)$`), `s`, true},
  Rule{regexp.MustCompile(`(ex|ix)$`), `ices`, false},
  Rule{regexp.MustCompile(`(oo)`), `ee`, false},
  Rule{regexp.MustCompile(`(er)$`), `ers`, false},
  Rule{regexp.MustCompile(`(s|ss|sh|ch|x|o|is)$`), `es`, true},
  Rule{regexp.MustCompile(`(a|e|o)y$`), `s`, true},
  Rule{regexp.MustCompile(`(y)$`), `ies`, false},
  Rule{regexp.MustCompile(`(on)$`), `a`, false},
}

// Plural returns the pluralized form of the word if a
// matched rule is found, else the original string is returned.
func Plural(str string) string {
  for _, rule := range plurals {
    if rule.Regexp.MatchString(str) {
      if rule.Append {
        return fmt.Sprintf(`%v%v`, str, rule.Replace)
      }

      return rule.Regexp.ReplaceAllString(str, rule.Replace)
    }
  }

  return str
}

func Singular(str string) string {
  return str
}
