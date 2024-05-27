package expanduser

import (
	"os"
	"os/user"
	"strings"
)

func ExpandUser(s string) (string, error) {
	separator := string(os.PathSeparator)
	tildeSeparator := "~" + separator
	tildeSeparatorLength := len(tildeSeparator)

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	if s == "~" {
		return home, nil
	} else if strings.HasPrefix(s, tildeSeparator) {
		return home + separator + s[tildeSeparatorLength:], nil
	} else if strings.HasPrefix(s, "~") {
		name := s[1:]
		u, err := user.Lookup(name)
		if err != nil {
			panic(err)
			// return "", err
		}
		return u.HomeDir, nil
	}
	return s, nil
}
