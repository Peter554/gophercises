package db

import (
	"encoding/json"

	bolt "go.etcd.io/bbolt"
)

var bucketKey = []byte("tasks")

type Task struct {
	ID        uint64
	Text      string
	Completed bool
}

type TasksService interface {
	Close() error
	GetAll() ([]Task, error)
	Add(text string) error
	Complete(id uint64) error
}

func NewTasksService() (TasksService, error) {
	db, e := bolt.Open("tasks.db", 0666, nil)
	if e != nil {
		return nil, e
	}
	e = db.Update(func(tx *bolt.Tx) error {
		_, e := tx.CreateBucketIfNotExists(bucketKey)
		return e
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return &tasksService{db: db}, nil
}

type tasksService struct {
	db *bolt.DB
}

func (o *tasksService) Close() error {
	return o.db.Close()
}

func (o *tasksService) GetAll() ([]Task, error) {
	tasks := make([]Task, 0)
	e := o.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketKey)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			t := &Task{}
			e := json.Unmarshal(v, t)
			if e != nil {
				return e
			}
			tasks = append(tasks, *t)
		}
		return nil
	})
	if e != nil {
		return nil, e
	}
	return tasks, nil
}

func (o *tasksService) Add(text string) error {
	t := Task{Text: text, Completed: false}
	return o.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketKey)
		id, e := b.NextSequence()
		if e != nil {
			return e
		}
		t.ID = id
		j, e := json.Marshal(t)
		if e != nil {
			return e
		}
		return b.Put(itob(id), j)
	})
}

func (o *tasksService) Complete(id uint64) error {
	return o.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketKey)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if btoi(k) == id {
				t := &Task{}
				json.Unmarshal(v, t)
				t.Completed = true
				s, e := json.Marshal(t)
				if e != nil {
					return e
				}
				return b.Put(k, s)
			}
		}
		return nil
	})
}
