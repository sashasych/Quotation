package storage

import (
	"gopkg.in/couchbase/gocb.v1"
	"quotation/model"
)

type DBConnection struct {
	cluster *gocb.Cluster
	bucket  *gocb.Bucket
}

func Connect(username string, password string) (*DBConnection, error) {
	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: username,
		Password: password,
	})
	bucket, err := cluster.OpenBucket("Info", "")
	if err != nil {
		return nil, err
	}
	return &DBConnection{cluster, bucket}, nil
}

func (db *DBConnection) SaveToDatabase(cbData *model.CBRResponse) {
	db.bucket.Upsert(cbData.Date.Format("02/01/2006"), cbData, 0)
}

func (db *DBConnection) GetFromDatabase(date string) (*model.CBRResponse, error) {
	cbData := new(model.CBRResponse)
	_, err := db.bucket.Get(date, cbData)
	if err != nil {
		return nil, err
	}
	return cbData, nil
}