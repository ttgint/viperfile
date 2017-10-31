package viperfile

import "testing"

func TestReadLocal(t *testing.T) {

	c := Configuration{}

	ReadLocal("test_config", &c)

	if c.Item.Intkey != 1 {
		t.Error("Expected 1, got ", c.Item.Intkey)
	}

	if c.Item.Stringkey != "test" {
		t.Error("Expected test, got ", c.Item.Stringkey)
	}
}
