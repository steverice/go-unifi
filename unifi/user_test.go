package unifi_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/filipowm/go-unifi/unifi"
)

func TestUserUnmarshalJSON(t *testing.T) {
	t.Parallel()

	for n, c := range map[string]struct {
		expected func(u *unifi.User)
		json     string
	}{
		"int dev_id_override": {
			expected: func(u *unifi.User) { u.DevIdOverride = 2779 },
			json:     `{ "dev_id_override": 2779 }`,
		},
		// UniFi Network 11 returns dev_id_override as a string.
		"string dev_id_override": {
			expected: func(u *unifi.User) { u.DevIdOverride = 2779 },
			json:     `{ "dev_id_override": "2779" }`,
		},
		"empty string dev_id_override": {
			expected: func(u *unifi.User) { u.DevIdOverride = 0 },
			json:     `{ "dev_id_override": "" }`,
		},

		"int last_seen": {
			expected: func(u *unifi.User) { u.LastSeen = 1 },
			json:     `{ "last_seen": 1 }`,
		},
		"string last_seen": {
			expected: func(u *unifi.User) { u.LastSeen = 1 },
			json:     `{ "last_seen": "1" }`,
		},
	} {
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			var expected unifi.User
			c.expected(&expected)
			var actual unifi.User
			err := json.Unmarshal(([]byte)(c.json), &actual)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(expected, actual) {
				t.Fatalf("not equal:\nexpected: %#v\nactual: %#v", expected, actual)
			}
		})
	}
}
