package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bahan struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	NamaBahan   string             `bson:"namaBahan" json:"namaBahan"`
	JumlahBahan int                `bson:"jumlahBahan" json:"jumlahBahan"`
	Satuan      string             `bson:"satuan" json:"satuan"`
	Keterangan  string             `bson:"keterangan" json:"keterangan"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}
