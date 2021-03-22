package goff

import "testing"
import "github.com/google/go-cmp/cmp"

func TesDictionary(t *testing.T) {
	cases := []struct {
		name      string
		input     map[string]string
		wantCount int
	}{
		{
			name: "non-empty",
			input: map[string]string{
				"a": "value_a",
				"b": "value_b",
			},
			wantCount: 2,
		},
		{
			name:      "empty",
			input:     map[string]string{},
			wantCount: 0,
		},
		{
			name:      "nil",
			input:     nil,
			wantCount: 0,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			input := map[string]string{
				"a": "value_a",
				"b": "value_b",
			}
			d, err := NewDictionary(input)
			if err != nil {
				t.Fatalf("got error parsing dictionary: %s", err)
			}
			defer d.free()

			if d.Size() != 2 {
				t.Errorf("Size() got %d wanted %d", d.Size(), 2)
			}

			got := d.toMap()
			if diff := cmp.Diff(input, got); diff != "" {
				t.Errorf("got: %s want %s\ndiff %s", got, input, diff)
			}
		})
	}
}

func TestDictionaryError(t *testing.T) {
	cases := []struct {
		name    string
		input   map[string]string
		wantErr bool
		wantKeys []string
	}{
		{
			name: "non-empty",
			input: map[string]string{
				"a": "value_a",
				"b": "value_b",
			},
			wantErr:  true,
			wantKeys: []string{"a", "b"},
		},
		{
			name: "non-empty2",
			input: map[string]string{
				"a": "value_a",
				"b": "value_b",
				"c": "value_b",
			},
			wantErr:  true,
			wantKeys: []string{"a", "b", "c"},
		},
		{
			name:  "empty",
			input: map[string]string{},
		},
		{
			name:  "nil",
			input: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d, err := NewDictionary(tc.input)
			if err != nil {
				t.Fatalf("got error parsing dictionary: %s", err)
			}
			defer d.free()

			err, gotOk := d.toUnavailableOptionsErr()
			if gotOk != tc.wantErr {
				t.Errorf("toUnavailableOptionsErr() got %t wanted %t", gotOk, tc.wantErr)
			}
			if gotOk {
				gotErr, ok := err.(*UnavailableOptionsErr)
				if !ok {
					t.Errorf("toUnavailableOptionsErr() got error type: %T wanted *UnavailableOptionsErr", err)
				}
				if diff := cmp.Diff(tc.wantKeys, gotErr.Keys); diff != "" {
					t.Errorf("got: %s want %s\ndiff %s", gotErr.Keys, tc.wantKeys, diff)
				}
			}
		})
	}
}
