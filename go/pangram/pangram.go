package pangram

import "strings"

func IsPangram(input string) bool {
  check := strings.ToLower(input)
  for char:='a' ; char <= 'z' ; char++ {
    if !strings.ContainsRune(check, char) {
      return false
    }
  }
  return true
}
