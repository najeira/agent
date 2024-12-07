package agent

import (
	"strconv"
	"strings"
)

type Agent struct {
	Name          string
	Major         int
	Minor         int
	Revision      int
	SystemName    string // iOS
	SystemVersion string // 15.2
	Model         string // iPhone
	Device        string // iPhone12,8
}

func (v Agent) Less(major, minor, revision int) bool {
	if v.Major < major {
		return true
	} else if v.Major > major {
		return false
	} else if v.Minor < minor {
		return true
	} else if v.Minor > minor {
		return false
	} else if v.Revision < revision {
		return true
	} else if v.Revision > revision {
		return false
	}
	return false
}

// Parse parses User-Agent
// AppName/0.3.6 (iOS 15.2; iPhone; iPhone12,8)
func Parse(ua string) Agent {
	ua = strings.TrimSpace(ua)

	parts := strings.SplitN(ua, " ", 2)
	if len(parts) < 1 {
		return Agent{}
	}
	an, ma, mi, re := parseApp(parts[0])
	sn, sv, mo, de := parseDevice(parts[1])
	return Agent{
		Name:          an,
		Major:         ma,
		Minor:         mi,
		Revision:      re,
		SystemName:    sn,
		SystemVersion: sv,
		Model:         mo,
		Device:        de,
	}
}

func parseApp(s string) (string, int, int, int) {
	parts := strings.SplitN(s, "/", 2)
	if len(parts) < 2 {
		return "", 0, 0, 0
	}
	a, b, c := parseVersion(parts[1])
	return parts[0], a, b, c
}

func parseVersion(s string) (int, int, int) {
	parts := strings.SplitN(s, ".", 3)
	if len(parts) < 3 {
		return 0, 0, 0
	}
	return parseInt(parts[0]), parseInt(parts[1]), parseInt(parts[2])
}

func parseDevice(s string) (string, string, string, string) {
	s = strings.TrimSpace(s)
	s = strings.TrimLeft(s, "(")
	s = strings.TrimRight(s, ")")

	parts := strings.SplitN(s, ";", 2)
	if len(parts) < 3 {
		return "", "", "", ""
	}

	sn, sv := parseSystem(parts[0])
	return sn, sv, strings.TrimSpace(parts[1]), strings.TrimSpace(parts[2])
}

func parseSystem(s string) (string, string) {
	s = strings.TrimSpace(s)
	parts := strings.SplitN(s, " ", 2)
	if len(parts) < 2 {
		return "", ""
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
