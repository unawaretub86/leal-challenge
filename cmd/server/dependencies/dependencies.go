package dependencies

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"

	"github.com/unawaretub86/leal-challenge/src/domain/entities/domain"
)

const (
	DBMaxIdleConns      = 25
	DBMaxOpenConns      = 50
	ConnMaxLifetimeSecs = 300
	DBReadTimeout       = 10
	DBWriteTimeout      = 10
	DBTimeout           = 10
)

type (
	Database interface {
		Connection() *gorm.DB
	}

	service struct {
		db *gorm.DB
	}
)

func ConnectDB(connectionString string) (Database, error) {

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect to db")
	}

	err = db.Use(GetPlugin())
	if err != nil {
		return nil, err
	}

	service := &service{
		db: db,
	}

	db.Migrator().DropTable(&domain.User{},
		&domain.Commerce{},
		&domain.Points{},
		&domain.Cashback{},
		&domain.Branch{},
		&domain.Campaign{},
		&domain.Purchase{},
		&domain.Award{},
		&domain.Redeem{},
	)

	db.AutoMigrate(
		&domain.User{},
		&domain.Commerce{},
		&domain.Points{},
		&domain.Cashback{},
		&domain.Branch{},
		&domain.Campaign{},
		&domain.Purchase{},
		&domain.Award{},
		&domain.Redeem{},
	)

	db.AutoMigrate(
		&domain.User{},
		&domain.Commerce{},
		&domain.Points{},
		&domain.Cashback{},
		&domain.Branch{},
		&domain.Campaign{},
		&domain.Purchase{},
		&domain.Award{},
	)

	return service, nil
}

func (service *service) Connection() *gorm.DB {
	return service.db
}

func GetPlugin() gorm.Plugin {
	connMaxLifetime := ConnMaxLifetimeSecs * time.Second

	return dbresolver.Register(dbresolver.Config{}).
		SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(connMaxLifetime).
		SetMaxIdleConns(DBMaxIdleConns).
		SetMaxOpenConns(DBMaxOpenConns)
}
