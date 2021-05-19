package database

import (
	"RDMS_TCP/structures"
	"errors"
)

func GetPackageByDownloadKey(key string) (structures.Package, error) {
	pkg := structures.Package{}

	err := db.Last(&pkg).Error

	if err != nil {
		return pkg, err
	}

	if pkg.Id == 0 {
		return pkg, errors.New("Invalid session key")
	}

	return pkg, nil
}
