package routers

import (
	"github.com/WymA/RESTful-go-server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
