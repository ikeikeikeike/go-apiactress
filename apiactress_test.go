package apiactress

import "testing"

func TestCount(t *testing.T) {
	c := NewClient()

	data, err := c.Fetch("a")
	if err != nil {
		t.Fatal(err)
	}

	if len(data.Actresses) <= 0 {
		t.Fatalf("Unexpected ApiActress.Actresses: %s length", len(data.Actresses))
	}

	if data.Count <= 0 {
		t.Fatalf("Unexpected ApiActress.Counte: %s length", data.Count)
	}
}

func TestFetch(t *testing.T) {
	c := NewClient()

	data, err := c.Fetch("i")
	if err != nil {
		t.Fatal(err)
	}

	a := data.Actresses[0]

	if a.Id <= 0 {
		t.Fatalf("Unexpected Actress.Id: %s", a.Id)
	}
	if a.Name == "" {
		t.Fatalf("Unexpected Actress.Name: %s", a.Name)
	}
	if a.Gyou == "" {
		t.Fatalf("Unexpected Actress.Gyou: %s", a.Gyou)
	}
	if a.Thumb == "" {
		t.Fatalf("Unexpected Actress.Thumb: %s", a.Thumb)
	}
	if a.Yomi == "" {
		t.Fatalf("Unexpected Actress.Yomi: %s", a.Yomi)
	}
	if a.Oto == "" {
		t.Fatalf("Unexpected Actress.Oto: %s", a.Oto)
	}
}
