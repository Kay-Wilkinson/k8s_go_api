package database

import (
    "fmt"
    "os"
    "log"

    "github.com/Kay-Wilkinson/k8s_go_api/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

// Dbinstance struct to represent postgres DB instance.

type Dbinstance struct {
    Db *gorm.DB
}

// pkg level variable to hold our global db. Set to global level to ensure app has access to db.
// var has name of DB with instance of Dbinstance
var DB Dbinstance

func ConnectDb() {
    // dsn arg (data source string) required to be passed in to gorm_dialector. Using postgres driver Open() func
    dsn := fmt.Sprintf(
        "host=db user=%s password=%s dbname=%s port=5432 sslmode=disable Timezone=Europe/Amsterdam",
        os.Getenv("POSTGRES_USER"),
        os.Getenv("POSTGRES_PASSWORD"),
        os.Getenv("POSTGRES_DB"),
    )
    // gorm_options is the gorm config obj.
    // gorm.Open(gorm_dialector, gorm_options)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    // gorm.Open() returns a db obj and an error.If fatal err returned, api should exit
    if err != nil {
        log.Fatal("Failed to connect to database. \n", err)
        os.Exit(2)
    }

    log.Println("connected")
    db.Logger = logger.Default.LogMode(logger.Info)

    log.Println("running migrations")
    db.AutoMigrate(&models.Fact{})

    DB = Dbinstance{
        Db: db,
    }
}
