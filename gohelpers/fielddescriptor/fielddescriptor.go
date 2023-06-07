package fielddescriptor

import (
	dt "github.com/Dagosu/BookingApp/datatypes"
)

// SliceToMap will convert a slice of field descriptors to a map (index being the field descriptor path)
func SliceToMap(fds []*dt.FieldDescriptor) map[string]*dt.FieldDescriptor {
	fdsMap := make(map[string]*dt.FieldDescriptor, len(fds))
	for _, fd := range fds {
		fdsMap[NormalizePath(fd.GetPath())] = fd
	}

	return fdsMap
}
