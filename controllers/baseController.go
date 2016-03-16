package controllers

import (
	"reflect"
	"runtime"

	"fmt"

	"BeegoApi01/services"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/goinggo/beego-mgo/localize"
	"github.com/goinggo/beego-mgo/utilities/mongo"
)

//** TYPES

type (
	// BaseController composes all required types and behavior.
	BaseController struct {
		beego.Controller
		services.Service
	}
)

//** INTERCEPT FUNCTIONS
// Prepare is called prior to the baseController method.
func (this *BaseController) Prepare() {
	beego.Info("baseController.Prepare() is called ~~~, %p", this)

	this.UserID = this.GetString("userID")
	if this.UserID == "" {
		this.UserID = this.GetString(":userID")
	}
	if this.UserID == "" {
		this.UserID = "Unknown"
	}

	if err := this.Service.Prepare(); err != nil {
		beego.Error("baseController.Prepare() is called ~~~, but Failed !!")
		this.ServeError(err)
		return
	}
}

// Finish is called once the baseController method completes.
func (this *BaseController) Finish() {
	defer func() {
		if this.MongoSession != nil {
			mongo.CloseSession(this.UserID, this.MongoSession)
			this.MongoSession = nil
		}
	}()

	beego.Info("baseController.Prepare() is called ~~~", this.Ctx.Request.URL.Path)
}

//** VALIDATION

// ParseAndValidate will run the params through the validation framework and then
// response with the specified localized or provided message.
func (this *BaseController) ParseAndValidate(params interface{}) bool {
	// This is not working anymore :(
	if err := this.ParseForm(params); err != nil {
		this.ServeError(err)
		return false
	}

	var valid validation.Validation
	ok, err := valid.Valid(params)
	if err != nil {
		this.ServeError(err)
		return false
	}

	if ok == false {
		// Build a map of the Error messages for each field
		messages2 := make(map[string]string)

		val := reflect.ValueOf(params).Elem()
		for i := 0; i < val.NumField(); i++ {
			// Look for an Error tag in the field
			typeField := val.Type().Field(i)
			tag := typeField.Tag
			tagValue := tag.Get("Error")

			// Was there an Error tag
			if tagValue != "" {
				messages2[typeField.Name] = tagValue
			}
		}

		// Build the Error response
		var errors []string
		for _, err := range valid.Errors {
			// Match an Error from the validation framework Errors
			// to a field name we have a mapping for
			message, ok := messages2[err.Field]
			if ok == true {
				// Use a localized message if one exists
				errors = append(errors, localize.T(message))
				continue
			}

			// No match, so use the message as is
			errors = append(errors, err.Message)
		}

		this.ServeValidationErrors(errors)
		return false
	}

	return true
}

//** EXCEPTIONS

// ServeError prepares and serves an Error exception.
func (this *BaseController) ServeError(err error) {
	this.Data["json"] = struct {
		Error string `json:"Error"`
	}{err.Error()}
	this.Ctx.Output.SetStatus(500)
	this.ServeJSON()
}

// ServeValidationErrors prepares and serves a validation exception.
func (this *BaseController) ServeValidationErrors(Errors []string) {
	this.Data["json"] = struct {
		Errors []string `json:"Errors"`
	}{Errors}
	this.Ctx.Output.SetStatus(409)
	this.ServeJSON()
}

//** CATCHING PANICS

// CatchPanic is used to catch any Panic and log exceptions. Returns a 500 as the response.
func (this *BaseController) CatchPanic(functionName string) {
	if r := recover(); r != nil {
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		beego.Warn(this.Service.UserID, functionName, "PANIC Defered [%v] : Stack Trace : ", r, string(buf))

		this.ServeError(fmt.Errorf("%v", r))
	}
}

//** AJAX SUPPORT

// AjaxResponse returns a standard ajax response.
func (this *BaseController) AjaxResponse(resultCode int, resultString string, data interface{}) {
	response := struct {
		Result       int
		ResultString string
		ResultObject interface{}
	}{
		Result:       resultCode,
		ResultString: resultString,
		ResultObject: data,
	}

	this.Data["json"] = response
	this.ServeJSON()
}
