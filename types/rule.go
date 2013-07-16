package types

import (
  "regexp"
)

type RuleType struct {
  Regexp   *regexp.Regexp
  Replacer string
}

type RulesType []*RuleType

func Rule(matcher, replacer string) (rule *RuleType) {
  rule = new(RuleType)
  rule.Regexp = regexp.MustCompile(matcher)
  rule.Replacer = replacer

  return
}
