package db

import (
	"bytes"
	"encoding/binary"
	"log"
	"time"

	bolt "go.etcd.io/bbolt"
)

var (
	taskBucket     = []byte("tasks")
	completeBucket = []byte("completed")
	db             *bolt.DB
)

type Task struct {
	Key   int
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0o600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		if err != nil {
			log.Fatal("error creating task bucket:", err)
			return err
		}
		_, err = tx.CreateBucketIfNotExists(completeBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

func CompleteTask(key int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		b2 := tx.Bucket(completeBucket)
		if b2 == nil {
			log.Println("we didn't even initialize this bucket")
		}
		task := b.Get(itob(key))
		completedKey := time.Now().Local().Format(time.RFC3339)
		b.Delete(itob(key))
		return b2.Put([]byte(completedKey), task)
	})
	return err
}

// GetCompletedTasks returns all values of completed tasks for the current day
func GetCompletedTasks() ([]string, error) {
	var completedTasks []string
	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(completeBucket).Cursor()
		min := time.Now().Local().Format(time.RFC3339)
		max := time.Now().Local().AddDate(0, 0, 1).Format(time.RFC3339)
		for k, v := c.Seek([]byte(min)); k != nil && bytes.Compare(k, []byte(max)) <= 0; k, v = c.Next() {
			completedTasks = append(completedTasks, string(v))
		}
		return nil
	})
	if err != nil {
		log.Println("failed to get completed tasks", err)
	}
	return completedTasks, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
