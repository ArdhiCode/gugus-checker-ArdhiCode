package migrations

import (
	"fmt"

	"github.com/ArdhiCode/gugus-checker-ArdhiCode/internal/entity"
	mylog "github.com/ArdhiCode/gugus-checker-ArdhiCode/internal/pkg/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	fmt.Println(mylog.ColorizeInfo("\n==========Start Migrating=========="))

	mylog.Infof("Migrating Tables . . .")

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		return err
	}

	if err := db.AutoMigrate(
		&entity.Mahasiswa{},
		&entity.Gugus{},
		&entity.Region{},
	); err != nil {
		return err
	}

	return nil
}
