package mydict

import (
	"testing"
)

func TestDict(t *testing.T) {
	t.Run("find existent key", func(t *testing.T) {
		dictionary := Dictionary{"key": "value"}

		result, err := dictionary.Find("key")

		if err != nil {
			t.Errorf("Not expected error: %s", err.Error())
		}

		expected := "value"

		if result != expected {
			t.Errorf("Result: '%s'; Expected '%s'", result, expected)
		}
	})

	t.Run("finc not existent key", func(t *testing.T) {
		dictionary := Dictionary{"key": "value"}

		_, err := dictionary.Find("xxx")

		if err == nil {
			t.Errorf("This test should FAIL, key should not exist in the map")
		}
	})

	t.Run("insert single value", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Insert("key", "value")

		if err != nil {
			t.Errorf("Not expected error: %s", err.Error())
		}
	})

	t.Run("insert existent value", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Insert("key", "value")

		if err != nil {
			t.Errorf("Not expected error: %s", err.Error())
		}

		err = dictionary.Insert("key", "value2")

		if err != ErrorExistentKey {
			t.Errorf("This test should FAIL, value of key should no bu updated")
		}
	})
}
