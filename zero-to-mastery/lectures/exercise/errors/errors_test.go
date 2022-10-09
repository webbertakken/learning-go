package timeparse

import "testing"

func TestParseFromStringTable(t *testing.T) {
  table := []struct {
    input   string
    correct bool
  }{
    {"00:00:00", true},
    {"23:59:59", true},
    {"23:00:00", true},
    {"00:59:00", true},
    {"00:00:59", true},
    {"24:00:00", false},
    {"00:60:00", false},
    {"00:00:60", false},
    {"", false},
    {"01", false},
    {"01:", false},
    {"01:01", false},
    {"01:01:", false},
    {"01:01:01", true},
    {"01:01:01:", false},
    {"01:01:01:01", false},
    {"-01", false},
    {"-01:01:01", false},
    {"01:-01:01", false},
    {"01:01:-01", false},
    {"aa:bb:cc", false},
  }

  for _, data := range table {
    _, err := ParseFromString(data.input)
    correct := err == nil

    if correct != data.correct {
      if data.correct {
        t.Errorf("Expected %v to be correctly parsed.", data.input)
      } else {
        t.Errorf("Expected parsing %v to emit an error.", data.input)
      }

    }
  }
}
