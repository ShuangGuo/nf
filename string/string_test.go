package string

import "testing"

func Test(t *testing.T) {
  var tests = []struct {
    s, want string
  }{
    {"Backward", "drawkcaB"},
    {"Hello, 世界", "界世 ,olleH"},
    {"",""},
  }
  for _,c := range tests {
    got := Reverse(c.s)
    if got != c.want {
      t.Errorf("Reverse(%q)== %1, want %q", c.s, got, c.want)
    }
  }
}
