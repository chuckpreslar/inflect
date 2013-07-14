package inflect

import (
  "testing"
)

func TestUpperCamelCase(t *testing.T) {
  tests := []string{"single", "lowerCamelCase", "under_scored", "hyphen-ated", "UpperCamelCase", "spaced Out"}
  results := []string{"Single", "LowerCamelCase", "UnderScored", "HyphenAted", "UpperCamelCase", "SpacedOut"}

  for index, test := range tests {
    if result := UpperCamelCase(test); result != results[index] {
      t.Errorf("Expected %v, got %v", results[index], result)
    }
  }
}

func TestLowerCamelCase(t *testing.T) {
  tests := []string{"single", "lowerCamelCase", "under_scored", "hyphen-ated", "UpperCamelCase", "spaced Out"}
  results := []string{"single", "lowerCamelCase", "underScored", "hyphenAted", "upperCamelCase", "spacedOut"}

  for index, test := range tests {
    if result := LowerCamelCase(test); result != results[index] {
      t.Errorf("Expected %v, got %v", results[index], result)
    }
  }
}

func TestUnderscore(t *testing.T) {
  tests := []string{"single", "lowerCamelCase", "under_scored", "hyphen-ated", "UpperCamelCase", "spaced Out"}
  results := []string{"single", "lower_camel_case", "under_scored", "hyphen_ated", "upper_camel_case", "spaced_out"}

  for index, test := range tests {
    if result := Underscore(test); result != results[index] {
      t.Errorf("Expected %v, got %v", results[index], result)
    }
  }
}

func TestHyphenate(t *testing.T) {
  tests := []string{"single", "lowerCamelCase", "under_scored", "hyphen-ated", "UpperCamelCase", "spaced Out"}
  results := []string{"single", "lower-camel-case", "under-scored", "hyphen-ated", "upper-camel-case", "spaced-out"}

  for index, test := range tests {
    if result := Hyphenate(test); result != results[index] {
      t.Errorf("Expected %v, got %v", results[index], result)
    }
  }
}

func TestHumanize(t *testing.T) {
  tests := []string{"single", "lowerCamelCase", "under_scored", "hyphen-ated", "UpperCamelCase", "spaced Out"}
  results := []string{"Single", "Lower camel case", "Under scored", "Hyphen ated", "Upper camel case", "Spaced out"}

  for index, test := range tests {
    if result := Humanize(test); result != results[index] {
      t.Errorf("Expected %v, got %v", results[index], result)
    }
  }
}

func TestTitleize(t *testing.T) {
  tests := []string{"single", "lowerCamelCase", "under_scored", "hyphen-ated", "UpperCamelCase", "spaced Out"}
  results := []string{"Single", "Lower Camel Case", "Under Scored", "Hyphen Ated", "Upper Camel Case", "Spaced Out"}

  for index, test := range tests {
    if result := Titleize(test); result != results[index] {
      t.Errorf("Expected %v, got %v", results[index], result)
    }
  }
}

func TestConstantize(t *testing.T) {
  tests := []string{"single", "lowerCamelCase", "under_scored", "hyphen-ated", "UpperCamelCase", "spaced Out"}
  results := []string{"SINGLE", "LOWER_CAMEL_CASE", "UNDER_SCORED", "HYPHEN_ATED", "UPPER_CAMEL_CASE", "SPACED_OUT"}

  for index, test := range tests {
    if result := Constantize(test); result != results[index] {
      t.Errorf("Expected %v, got %v", results[index], result)
    }
  }
}
