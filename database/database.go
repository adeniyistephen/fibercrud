package database

import (
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Task struct {
	gorm.Model
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() *gorm.DB {
	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// 	os.Getenv("DB_PORT"),
	// )

	//================================TEST==============================
	// to run test: docker run --name postgres_db  -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=crud -d postgres:13
	// then run: go test ./...
	dsn_test := "host=172.17.0.2 user=postgres password=pass dbname=crud port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//==================================================================

	db, err := gorm.Open(postgres.Open(dsn_test), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&Task{})

	DB = Dbinstance{
		Db: db,
	}

	return DB.Db
}

func CreateTask(name string, status string) (Task, error) {
	var newTask = Task{Name: name, Status: status}

	DB.Db.Create(&Task{Name: name, Status: status})

	return newTask, nil
}

func GetallTasks() ([]Task, error) {
	var tasks []Task

	DB.Db.Find(&tasks)

	return tasks, nil
}

func Gettask(id string) (Task, error) {
	var task Task

	DB.Db.Where("ID = ?", id).First(&task)

	return task, nil
}

func Deletetask(id string) error {
	var task Task

	DB.Db.Where("ID = ?", id).Delete(&task)

	return nil

}

func Updatetask(name string, status string, id string) (Task, error) {
	var newTask = Task{Name: name, Status: status}

	DB.Db.Where("ID = ?", id).Updates(&Task{Name: newTask.Name, Status: newTask.Status})

	return newTask, nil
}
