package database

import (
	"RDMS_TCP/structures"
	"errors"
)

func GetPackageByDownloadKey(key string) (structures.Package, error) {
	pkg := structures.Package{}

	err := db.Joins("INNER JOIN files ON packages.file = files.id").
		Joins("INNER JOIN download_queue as dq ON dq.file_id = files.id").
		Select("packages.id, packages.name, packages.version, packages.ord, files.path, packages.on_server").
		Where("dq.session_key = ?", key).
		Last(&pkg).Error

	if err != nil {
		return pkg, err
	}

	if pkg.Id == 0 {
		return pkg, errors.New("Invalid session key")
	}

	return pkg, nil
}

func RemoveSession(key string) (error) {
	session := structures.Session{SessionKey: key}

	err := db.Delete(session).Error

	return err
}
