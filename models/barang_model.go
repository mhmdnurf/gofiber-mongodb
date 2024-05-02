package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Barang struct {
	Id           primitive.ObjectID `bson:"_id" json:"_id"`
	NamaBarang   string             `bson:"namaBarang" json:"namaBarang"`
	JumlahBarang int                `bson:"jumlahBarang" json:"jumlahBarang"`
	Satuan       string             `bson:"satuan" json:"satuan"`
	Keterangan   string             `bson:"keterangan" json:"keterangan"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
}
