package main

import (
	"fmt"
	"log"
	"time"

	"go.etcd.io/bbolt"
)

func main() {
	// Открытие базы данных
	db, err := bbolt.Open("my.db", 0600, &bbolt.Options{Timeout: 1 *
		time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Создание "bucket" для заметок
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("notes"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	// Добавление заметки
	noteKey := []byte("note1")
	noteValue := []byte("This is my first note")
	err = db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("notes"))
		err := bucket.Put(noteKey, noteValue)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	// Чтение заметки
	err = db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("notes"))
		value := bucket.Get(noteKey)
		fmt.Printf("The note is: %s\n", value)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	// Удаление заметки
	err = db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("notes"))
		err := bucket.Delete(noteKey)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
}
