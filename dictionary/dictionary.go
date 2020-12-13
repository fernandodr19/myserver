package mydict

import "errors"

type Dictionary map[string]string

var (
	ErrorKeyNotFound = errors.New("Key not found")
	ErrorExistentKey = errors.New("Key already in use")
)

func Jack() string { return "Jack da VK" }

func (d Dictionary) Insert(key, value string) error {
	// Why don't I need to use pointer here o inset on original (not copy) dict?
	_, err := d.Find(key)
	switch err {
	case ErrorKeyNotFound:
		d[key] = value
	case nil:
		return ErrorExistentKey
	default:
		return err
	}

	return nil
}

func (d Dictionary) Find(key string) (string, error) {
	value, existent := d[key]

	if !existent {
		return "", ErrorKeyNotFound
	}

	return value, nil
}
