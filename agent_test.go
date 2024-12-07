package agent

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	tcs := []struct {
		UA       string
		NA       string
		Major    int
		Minor    int
		Revision int
		SN       string
		SV       string
		MO       string
		DV       string
	}{
		{
			UA:       "AppName/1.3.6 (iOS 15.2; iPhone; iPhone12,8)",
			NA:       "AppName",
			Major:    1,
			Minor:    3,
			Revision: 6,
			SN:       "iOS",
			SV:       "15.2",
			MO:       "iPhone",
			DV:       "iPhone12,8",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.UA, func(t *testing.T) {
			v := Parse(tc.UA)
			if v.Major != tc.Major {
				t.Errorf("AppVersion.Major %d != %d", v.Major, tc.Major)
			}
			if v.Minor != tc.Minor {
				t.Errorf("AppVersion.Minor %d != %d", v.Minor, tc.Minor)
			}
			if v.Revision != tc.Revision {
				t.Errorf("AppVersion.Revision %d != %d", v.Revision, tc.Revision)
			}
		})
	}
}

func TestAppVersionLess(t *testing.T) {
	tcs := []struct {
		UA       string
		Major    int
		Minor    int
		Revision int
		Less     bool
	}{
		{
			UA:       "AppName/1.3.6 (iOS 15.2; iPhone; iPhone12,8)",
			Major:    1,
			Minor:    3,
			Revision: 6,
			Less:     false,
		},
		{
			UA:       "AppName/1.3.6 (iOS 15.2; iPhone; iPhone12,8)",
			Major:    1,
			Minor:    3,
			Revision: 5,
			Less:     false,
		},
		{
			UA:       "AppName/1.3.6 (iOS 15.2; iPhone; iPhone12,8)",
			Major:    1,
			Minor:    2,
			Revision: 8,
			Less:     false,
		},
		{
			UA:       "AppName/1.3.6 (iOS 15.2; iPhone; iPhone12,8)",
			Major:    0,
			Minor:    5,
			Revision: 8,
			Less:     false,
		},
		{
			UA:       "AppName/1.3.6 (iOS 15.2; iPhone; iPhone12,8)",
			Major:    1,
			Minor:    3,
			Revision: 8,
			Less:     true,
		},
		{
			UA:       "AppName/1.3.6 (iOS 15.2; iPhone; iPhone12,8)",
			Major:    1,
			Minor:    5,
			Revision: 3,
			Less:     true,
		},
		{
			UA:       "AppName/1.3.6 (iOS 15.2; iPhone; iPhone12,8)",
			Major:    2,
			Minor:    0,
			Revision: 3,
			Less:     true,
		},
	}
	for _, tc := range tcs {
		name := fmt.Sprintf("%d.%d.%d", tc.Major, tc.Minor, tc.Revision)
		t.Run(name, func(t *testing.T) {
			v := Parse(tc.UA)
			res := v.Less(tc.Major, tc.Minor, tc.Revision)
			if res != tc.Less {
				t.Errorf("AppVersion.Less %v != %v (%v)", res, tc.Less, v)
			}
		})
	}
}
