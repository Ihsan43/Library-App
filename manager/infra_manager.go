package manager

import (
	"fmt"
	"library_app/config"
	"library_app/utils"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type InfraManager interface {
	Conn() *gorm.DB
}

type infraManager struct {
	db  *gorm.DB
	cfg *config.Config
}

func (i *infraManager) openConn() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
    i.cfg.Host, i.cfg.Port, i.cfg.User, i.cfg.Password, i.cfg.Name)

	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	i.db = Db

	return nil
}

func (i *infraManager) Conn() *gorm.DB {
	return i.db
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{cfg: cfg}
	if err := conn.openConn(); err != nil {
		return nil, err
	}

	utils.MigrateModels(conn.db)

	return conn, nil
}
