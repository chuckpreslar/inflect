package inflect

import (
  "fmt"
  "regexp"
  "strings"
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
  Rule{regexp.MustCompile(`(ife)$`), `ives`, false},
  Rule{regexp.MustCompile(`(ef|ff|of|ay|ey|iy|oy|uy)$`), `s`, true},
  Rule{regexp.MustCompile(`(ice)$`), `ie`, false},
  Rule{regexp.MustCompile(`(house)$`), `houses`, false},
  Rule{regexp.MustCompile(`(blouse)$`), `blouses`, false},
  Rule{regexp.MustCompile(`(ouse)$`), `ice`, false},
  Rule{regexp.MustCompile(`(oo)th$`), `eeth`, false},
  Rule{regexp.MustCompile(`(oo)t$`), `eet`, false},
  Rule{regexp.MustCompile(`(oo)se$`), `eese`, false},
  Rule{regexp.MustCompile(`(f)$`), `ves`, false},
  Rule{regexp.MustCompile(`(y)$`), `ies`, false},
  Rule{regexp.MustCompile(`(s|x|z|ch|sh)$`), `es`, true},
  Rule{regexp.MustCompile(`(to|ro|ho|jo)$`), `es`, true},
  Rule{regexp.MustCompile(`(person)`), `people`, false},
}

var (
  //FIXME: This should probably be read in from a file.
  uncountables = []string{`fish`, `sheep`, `deer`, `tuna`, `salmon`, `trout`, `music`, `art`, `love`, `happiness`, `advice`, `information`, `news`, `furniture`, `luggage`, `rice`, `sugar`, `butter`, `water`, `electricity`, `gas`, `power`, `money`, `currency`, `scenery`}
  compiled     = strings.Join(uncountables, ` `)
)

// List of singularization rules in order by precedence.
var singulars = []Rule{}

// Plural returns the pluralized form of the word if a
// matched rule is found, else the original string is returned.
func Pluralize(str string) string {
  if 0 <= strings.Index(compiled, str) {
    return str
  }

  for _, rule := range plurals {
    if rule.Regexp.MatchString(str) {
      if rule.Append {
        return fmt.Sprintf(`%v%v`, str, rule.Replace)
      }

      return rule.Regexp.ReplaceAllString(str, rule.Replace)
    }
  }

  return fmt.Sprintf(`%v%v`, str, `s`)
}

// Converts a plural string to it's singular form.
// FIX ME: NOT IMPLEMENTED.
func Singularize(str string) string {
  for _, rule := range singulars {
    if rule.Regexp.MatchString(str) {
      return rule.Regexp.ReplaceAllString(str, rule.Replace)
    }
  }

  return str
}

// Split's a string so that it can be converted to a different casing.
// Splits on underscores, hyphens, spaces and camel casing.
func split(str string) (pieces []string) {
  //FIXME: Go's Regexp's annoy me.
  str = strings.Trim(str, `_`)
  str = strings.Trim(str, `-`)
  str = strings.Trim(str, ` `)

  var (
    current int
    next    int
    end     int
  )

  for 0 < len(str) {
    end = len(str)
    next = current + 1

    if end <= next {
      pieces = append(pieces, str)
      break
    }

    if isLowerCase(str[current]) && isUpperCase(str[next]) {
      pieces = append(pieces, str[:next])
      str = str[next:]
      current = 0
    } else if '-' == str[current] || '_' == str[current] || ' ' == str[current] {
      pieces = append(pieces, str[:current])
      str = str[next:]
      current = 0
    } else {
      current++
    }
  }

  return pieces
}

// Checks if a single character is upper cased.
func isUpperCase(c uint8) (matched bool) {
  matched, _ = regexp.MatchString(`[A-Z]`, string(c))
  return
}

// Checks if a single character is lower cased.
func isLowerCase(c uint8) (matched bool) {
  matched, _ = regexp.MatchString(`[a-z]`, string(c))
  return
}

// Converts a string to it's upper camel case version.
func UpperCamelCase(str string) string {
  pieces := split(str)

  for index, s := range pieces {
    pieces[index] = fmt.Sprintf(`%v%v`, strings.ToUpper(string(s[0])), s[1:])
  }

  return strings.Join(pieces, ``)
}

// Converts a string to it's lower camel case version.
func LowerCamelCase(str string) string {
  pieces := split(str)

  pieces[0] = fmt.Sprintf(`%v%v`, strings.ToLower(string(pieces[0][0])), pieces[0][1:])

  for i := 1; i < len(pieces); i++ {
    pieces[i] = fmt.Sprintf(`%v%v`, strings.ToUpper(string(pieces[i][0])), pieces[i][1:])
  }

  return strings.Join(pieces, ``)
}

// Converts a string to it's underscored version.
func Underscore(str string) string {
  pieces := split(str)

  for index, piece := range pieces {
    pieces[index] = strings.ToLower(piece)
  }

  return strings.Join(pieces, `_`)
}

// Converts a string to it's underscored version.
func Hyphenate(str string) string {
  pieces := split(str)

  for index, piece := range pieces {
    pieces[index] = strings.ToLower(piece)
  }

  return strings.Join(pieces, `-`)
}

// Converts a string to it's constantized version.
func Constantize(str string) string {
  pieces := split(str)

  for index, piece := range pieces {
    pieces[index] = strings.ToUpper(piece)
  }

  return strings.Join(pieces, `_`)
}

// Converts a string to it's humanized version.
func Humanize(str string) string {
  pieces := split(str)

  pieces[0] = fmt.Sprintf(`%v%v`, strings.ToUpper(string(pieces[0][0])), pieces[0][1:])

  for i := 1; i < len(pieces); i++ {
    pieces[i] = fmt.Sprintf(`%v`, strings.ToLower(pieces[i]))
  }

  return strings.Join(pieces, ` `)
}

// Converts a string to it's titleized version.
func Titleize(str string) string {
  pieces := split(str)

  for i := 0; i < len(pieces); i++ {
    pieces[i] = fmt.Sprintf(`%v%v`, strings.ToUpper(string(pieces[i][0])), pieces[i][1:])
  }

  return strings.Join(pieces, ` `)
}
