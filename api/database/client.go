package database

import (
	"encoding/json"
	"log"
	"time"

	"github.com/CherryDock/CherryDock/api/docker/monitoring"
	"github.com/boltdb/bolt"
)

var DbClient BoltDb

// BoltDb Client interface
type BoltDb interface {
	Init()
	RetrieveData() *[]DataMonitoring
	AddMonitoringInfo(info *monitoring.GlobalStats) error
}

// Client bolt db
type Client struct {
	boltDb *bolt.DB
}

// DataMonitoring is the format of data stored in db
type DataMonitoring struct {
	Date *time.Time
	monitoring.GlobalStats
}

// Init bolt db, create buckets if not exists
func (client *Client) Init() {
	var err error
	client.boltDb, err = bolt.Open("cherrydock.db", 384, nil)
	if err != nil {
		log.Printf("fail to open db %v", err)
	}

	err = client.boltDb.Update(func(tx *bolt.Tx) error {
		// Create root bucket
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			log.Printf("fail to create root bucket: %v\n", err)
		}
		// Create monitoring bucket
		_, err = root.CreateBucketIfNotExists([]byte("MONITORING"))
		if err != nil {
			log.Printf("fail to create monitoring bucket %v\n", err)
		}
		return nil
	})

	if err != nil {
		log.Printf("fail to setup buckets, %v\n", err)
	}
	log.Println("database successfully initialized")
}

// AddMonitoringInfo insert monitoring data into database
func (client *Client) AddMonitoringInfo(info *monitoring.GlobalStats) error {
	date := time.Now()
	encodedInfo, err := json.Marshal(info)
	if err != nil {
		log.Println("fail to marshal monitoring info")
	}

	err = client.boltDb.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte("MONITORING")).Put([]byte(date.Format(time.RFC3339)), encodedInfo)
		if err != nil {
			log.Println("could not add monitoring info to database")
		}
		return nil
	})
	return err
}

// RetrieveData return all data from monitoring bucket
func (client *Client) RetrieveData() *[]DataMonitoring {
	var databaseContent []DataMonitoring
	var data monitoring.GlobalStats
	var date time.Time

	client.boltDb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("MONITORING"))
		b.ForEach(func(k, v []byte) error {
			// Unmarshal date format RFC3339
			date, _ = time.Parse(time.RFC3339, string(k))
			// Unmarshal containers info
			json.Unmarshal(v, &data)
			databaseContent = append(databaseContent, DataMonitoring{&date, data})
			return nil
		})
		return nil
	})
	return &databaseContent
}
