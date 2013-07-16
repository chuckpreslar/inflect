package types

type UncountablesType []string

func (self UncountablesType) Contains(str string) bool {
  for _, word := range self {
    if word == str {
      return true
    }
  }

  return false
}
