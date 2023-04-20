package database

import (
	"testing"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func TestConnectDb(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "connect db",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConnectDb()
		})
	}
}

func TestCreateTask(t *testing.T) {
	db := ConnectDb()
    d, _ := db.DB()
	defer d.Close()

    // Create
    task := Task{Name: "John", Status: "Active"}
    result := db.Create(&task)

    // Verify
    if result.Error != nil {
        t.Errorf("failed to create user: %s", result.Error)
    }
}

func TestGetallTasks(t *testing.T) {
	db := ConnectDb()
    d, _ := db.DB()
	defer d.Close()

    // Create
    task := Task{Name: "John", Status: "Active"}
    db.Create(&task)

    //Read
    var result []Task
    db.Find(&result)

    // Verify
    if result == nil {
        t.Error("failed to read task, expected:", result)
    }
}

func TestGettask(t *testing.T) {
	db := ConnectDb()
    d, _ := db.DB()
	defer d.Close()

    // Create
    task := Task{Name: "John", Status: "Active"}
    db.Create(&task)

    // Read
    var result Task
    db.First(&result, task.ID)

    // Verify
    if result.Name != "John" {
        t.Errorf("failed to read task, expected: %s", result.Name)
    }
}

func TestDeletetask(t *testing.T) {
	db := ConnectDb()
    d, _ := db.DB()
	defer d.Close()

    // Create
    task := Task{Name: "John", Status: "Active"}
    db.Create(&task)

    // Delete
    db.Delete(&task)

    // Read
    var result Task
    db.First(&result, task.ID)

    // Verify
    if db.Error == gorm.ErrRecordNotFound {
        t.Errorf("failed to delete task: %s", result.Name)
    }
}

func TestUpdatetask(t *testing.T) {
	db := ConnectDb()
    d, _ := db.DB()
	defer d.Close()

    // Create
    task := Task{Name: "John", Status: "Active"}
    db.Create(&task)

    // Update
    task.Name = "Jack"
    db.Save(&task)

    // Read
    var result Task
    db.First(&result, task.ID)

    // Verify
    if result.Name != "Jack" {
        t.Errorf("failed to update task: %s", result.Name)
    }
}