package polite

import "testing"

func TestSalute(t *testing.T) {
	cases := []struct {
		in string
		want string
	}{
		{"Nick", "Good morning Mr. Nick"},
		{"", "Hello nice to meet you!"},
	}

	for _, c := range cases {
		got := Salute(c.in)
		if got != c.want {
			t.Errorf("Salute(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestAurevoir(t *testing.T) {
	cases := []struct {
		in string
		want string
	}{
		{"Nick", "Good night Mr. Nick"},
		{"", "Bye nice to meet you!"},
	}

	for _, c := range cases {
		got := Aurevoir(c.in)
		if got != c.want {
			t.Errorf("Aurevoir(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestIntroduce(t *testing.T) {
	cases := []struct {
		in string
		want string
	}{
		{"Nick", "Hello I am Lord Nick"},
		{"", "Hello I am Lord "},
	}

	for _, c := range cases {
		got := Introduce(c.in)
		if got != c.want {
			t.Errorf("Introduce(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}