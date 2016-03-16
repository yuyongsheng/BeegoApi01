package services

import (
	"github.com/astaxie/beego"
	"github.com/goinggo/beego-mgo/utilities/helper"
	"github.com/goinggo/beego-mgo/utilities/mongo"
	"gopkg.in/mgo.v2"
)

//** TYPES

type (
	// Service contains common properties for all baseService.
	Service struct {
		MongoSession *mgo.Session
		UserID       string
	}
)

//** PUBLIC FUNCTIONS

// Prepare is called before any controller.
func (this *Service) Prepare() (err error) {
	beego.Info("baseService.Prepare() is called ~~~")
	this.MongoSession, err = mongo.CopyMonotonicSession(this.UserID)
	if err != nil {
		beego.Info("baseService.Prepare() is called ~~~, but Failed !!!")
		return err
	}

	return err
}

// Finish is called after the controller.
func (this *Service) Finish() (err error) {
	beego.Info("baseService.Finish() is called ~~~")
	defer helper.CatchPanic(&err, this.UserID, "Service.Finish")

	if this.MongoSession != nil {
		mongo.CloseSession(this.UserID, this.MongoSession)
		this.MongoSession = nil
	}

	return err
}

// DBAction executes the MongoDB literal function
func (this *Service) DBAction(databaseName string, collectionName string, dbCall mongo.DBCall) (err error) {
	beego.Info("baseService.DBAction() is called ~~~, ", databaseName, collectionName, dbCall)
	return mongo.Execute(this.UserID, this.MongoSession, databaseName, collectionName, dbCall)
}
