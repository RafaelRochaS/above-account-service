package repository

import "github.com/RafaelRochaS/above-account-service/models"

type IDBConn interface {
	InstantiateDbConn() error
	Migrate()
	Create(*models.User) error
	Update(models.User) error
	GetOne(int) (error, models.User)
	GetMany(int) (error, models.User)
	Delete(int) (error, models.User)
	GetByName(string) (error, models.User)
}
