package mydict

import "errors"

type Dictionary map[string]string

var ( 
	errNotFound = errors.New("not Found")
	errKeyExists = errors.New("word already exists")
	errCantUpdate = errors.New("cant update non-existing word")
)
func (d Dictionary) Search(key string) (string, error) {
	value, exists := d[key]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(key, value string) error{
	_, err := d.Search(key)
	switch err{
	case errNotFound:
		d[key] = value
	case nil :
		return errKeyExists
	}
	return nil
}

func (d Dictionary) Update(key , value string) error{
	_,err := d.Search(key)
	switch err {
	case nil:
		d[key] = value
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(key string){
	delete(d,key)
}