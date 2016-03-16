package controllers

import (
	"BeegoApi01/models"
	"BeegoApi01/services"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	BaseController
}

func (uc *UserController) Prepare() {
	beego.Debug("UserController::Prepare() is called here ~~~~")
}

func (uc *UserController) Finish() {
	beego.Debug("UserController::Finish() is called here ~~~~")
}

func (uc *UserController) URLMapping() {
	beego.Debug("UserController::URLMapping() is called here ~~~~")
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (uc *UserController) Post() {
	beego.Debug("UserController::Post() is called here ~~~~")

	var user models.User
	json.Unmarshal(uc.Ctx.Input.RequestBody, &user)

	uid, err := services.AddUser(&uc.Service, &user)
	if err != nil {
		beego.Error(uc.UserID, "UserController.Post() Error ~~~~, user :", user)
		uc.ServeError(err)
		return
	}

	uc.Data["json"] = map[string]string{"uid": uid}
	uc.ServeJSON()
}

// @Title Get
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (uc *UserController) GetAll() {
	beego.Debug("UserController::GetAll() is called here ~~~~")

	users, err := services.FindAllUsers(&uc.Service, uc.UserID)
	if err != nil {
		beego.Error(uc.UserID, "UserController.GetAll() Error ~~~~, UserId :", uc.UserID)
		uc.ServeError(err)
		return
	}

	if users != nil {
		uc.Data["json"] = users
	} else {
		uc.Data["json"] = "nothing user"
	}

	uc.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (uc *UserController) Get() {
	beego.Debug("UserController::Get() is called here ~~~~")
	uid := uc.GetString(":uid")
	beego.Debug("UserController::Get(), uid : ", uid)

	users, err := services.FindUser(&uc.Service, uc.UserID)
	if err != nil {
		beego.Error(uc.UserID, "UserController.GetAll() Error ~~~~, UserId :", uc.UserID)
		uc.ServeError(err)
		return
	}

	if users != nil {
		uc.Data["json"] = users
	} else {
		uc.Data["json"] = "nothing user"
	}

	uc.ServeJSON()
}

// @Title update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (uc *UserController) Put() {
	beego.Debug("UserController::Put() is called here ~~~~")
	uid := uc.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(uc.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			uc.Data["json"] = err.Error()
		} else {
			uc.Data["json"] = uu
		}
	}
	uc.ServeJSON()
}

// @Title delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (uc *UserController) Delete() {
	beego.Debug("UserController::Delete() is called here ~~~~")
	uid := uc.GetString(":uid")
	models.DeleteUser(uid)
	uc.Data["json"] = "delete success!"
	uc.ServeJSON()
}

// @Title login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (uc *UserController) Login() {
	beego.Debug("UserController::Login() is called here ~~~~")
	username := uc.GetString("username")
	password := uc.GetString("password")
	if models.Login(username, password) {
		uc.Data["json"] = "login success"
	} else {
		uc.Data["json"] = "user not exist"
	}
	uc.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (uc *UserController) Logout() {
	beego.Debug("UserController::Logout() is called here ~~~~")
	uc.Data["json"] = "logout success"
	uc.ServeJSON()
}
