package database

import (
	"log"
	"os"

	"github.com/hack-caixa/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB Database

type Database struct {
	Db            *gorm.DB
	dialector     gorm.Dialector
	User          string
	Password      string
	Host          string
	Port          string
	Dbname        string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func init() {

	DB.DsnTest = os.Getenv("DSN_TEST")
	DB.User = os.Getenv("SqlServerUser")
	DB.Password = os.Getenv("SqlServerPassword")
	DB.Host = os.Getenv("SqlServerHost")
	DB.Port = os.Getenv("SqlServerPort")
	DB.Dbname = os.Getenv("SqlServerDbname")
	DB.DbTypeTest = os.Getenv("DB_TYPE_TEST")
	DB.DbType = os.Getenv("DB_TYPE")
	DB.Env = os.Getenv("ENV")
}

func NewDB() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDB()
	dbInstance.Env = "test"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.DsnTest = ":memory:"
	dbInstance.AutoMigrateDb = true
	dbInstance.Debug = true
	dbInstance.dialector = sqlite.Open(dbInstance.DsnTest)

	connection, err := dbInstance.Connect()

	if err != nil {
		log.Fatalf("Test db error %v", err)
	}

	dbInstance.Db.AutoMigrate(domain.PRODUTO{})

	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	dsn := "sqlserver://" + d.User + ":" + d.Password + "@" + d.Host + ":" + d.Port + "?database=" + d.Dbname

	if d.Env != "test" {
		d.Db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{TablePrefix: "dbo.", SingularTable: true}, SkipDefaultTransaction: true})
	} else {
		d.Db, err = gorm.Open(d.dialector, &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}

	log.Println("Connected to Database!")

	return d.Db, nil

}
