package services

import (
	"github.com/astaxie/beego"
	"github.com/goinggo/beego-mgo/utilities/mongo"

	"BeegoApi01/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//** TYPES

type (
	// buoyConfiguration contains settings for running the buoy service.
	buoyConfig struct {
		Database string
	}
)

//** PACKAGE VARIABLES

// Config provides buoy configuration.
var (
	Config           buoyConfig
	dbCollectionName string
)

//** INIT

func init() {
	// Pull in the configuration.
	Config.Database = "testdb2"
	dbCollectionName = "users"
}

//** PUBLIC FUNCTIONS

// FindUser retrieves the specified station
func FindUser(service *Service, userId string) (*models.User, error) {
	beego.Info("userService.FindUser() is called ~~~", "FindUser: ", userId)

	var usr models.User
	f := func(collection *mgo.Collection) error {
		queryMap := bson.M{"userId": userId}

		beego.Info(service.UserID, "FindUser, MGO :", mongo.ToString(queryMap))
		return collection.Find(queryMap).One(&usr)
	}

	if err := service.DBAction(Config.Database, dbCollectionName, f); err != nil {
		if err != mgo.ErrNotFound {
			beego.Error(err, service.UserID, "FindUser")
			return nil, err
		}
	}

	return &usr, nil
}

func FindAllUsers(service *Service, userId string) ([]models.User, error) {
	beego.Info("userService.FindAllUsers() is called ~~~", "UserId: ", userId)

	var users []models.User
	f := func(collection *mgo.Collection) error {
		queryMap := bson.M{"userId": userId}

		beego.Info(service.UserID, "FindAllUsers, MGO : ", mongo.ToString(queryMap))
		return collection.Find(queryMap).All(&users)
	}

	if err := service.DBAction(Config.Database, dbCollectionName, f); err != nil {
		if err != mgo.ErrNotFound {
			beego.Error(err, service.UserID, "FindAllUsers")
			return nil, err
		}
	}

	return users, nil
}

func AddUser(service *Service, user *models.User) (string, error) {
	beego.Info("userService.AddUser() is called ~~~, User: ", user)

	uid := "uid12345678"
	return uid, nil
}
