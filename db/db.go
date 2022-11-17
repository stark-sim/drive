package db

import (
	"ariga.io/atlas/sql/sqltool"
	"context"
	"drive/config"
	"drive/ent"
	"drive/ent/migrate"
	"entgo.io/ent/dialect/sql/schema"
	"errors"
	"fmt"
	golangMigrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var url = ""

func Init() {
	url = fmt.Sprintf("%s://%s:%s@%s:%v/%s?sslmode=disable&TimeZone=Asia/Shanghai",
		config.Conf.DBConfig.Driver,
		config.Conf.DBConfig.Username,
		config.Conf.DBConfig.Password,
		config.Conf.DBConfig.Host,
		config.Conf.DBConfig.Port,
		config.Conf.DBConfig.Database)
	generateMigrationFiles()
	migrateWithMigrationFiles()
}

func generateMigrationFiles() {
	ctx := context.Background()
	dir, err := sqltool.NewGolangMigrateDir("migrations")
	if err != nil {
		logrus.Errorf("failed creatiing atlas migration directory: %v", err)
	}

	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		schema.WithFormatter(sqltool.GolangMigrateFormatter),
		schema.WithMigrationMode(schema.ModeInspect),
		schema.WithDialect(config.Conf.Driver),
		schema.WithForeignKeys(false),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
		//schema.WithGlobalUniqueID(true),
		migrate.WithGlobalUniqueID(true),
	}

	err = migrate.NamedDiff(ctx, url, "update", opts...)
	if err != nil {
		logrus.Errorf("failed generating migration file: %v", err)
	}
}

func migrateWithMigrationFiles() {
	m, err := golangMigrate.New("file://migrations", url)
	if err != nil {
		logrus.Errorf("failed at new migrate, err: %v", err)
	}
	if err = m.Up(); err != nil && !errors.Is(err, golangMigrate.ErrNoChange) {
		logrus.Errorf("failed at up migrate: %v", err)
	}
}

func NewDBClient() *ent.Client {
	dataSourceName := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.Conf.DBConfig.Host, config.Conf.DBConfig.Port, config.Conf.Username, config.Conf.Password, config.Conf.Database)
	logrus.Debugf("dsn: %v", dataSourceName)

	client, err := ent.Open(config.Conf.DBConfig.Driver, dataSourceName)
	if err != nil {
		logrus.Errorf("failed at creating ent client with db %s, err: %v", dataSourceName, err)
		return nil
	}
	return client
}
