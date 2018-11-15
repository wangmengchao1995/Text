package routers

import (
	"newsWeb/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {

	beego.InsertFilter("/article/*",beego.BeforeExec,funcFilter)
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandleReg")
    beego.Router("/login",&controllers.UserController{},"get:ShowLogin;post:HandleLogin")
    beego.Router("/article/articleList",&controllers.ArticleController{},"get:ShowArticleList;post:HandleSelect")
    beego.Router("/article/addArticle",&controllers.ArticleController{},"get:ShowAddArticle;post:HandeAddArticle")
    beego.Router("/article/articleDetail",&controllers.ArticleController{},"get:ShowArticleDetail")
	beego.Router("/article/updateAriticle",&controllers.ArticleController{},"get:ShowUpdateArticle;post:HandleUpdateArticle")
	beego.Router("/article/DeleteArticle",&controllers.ArticleController{},"get:DeleteArticle")
	beego.Router("/article/addType",&controllers.ArticleController{},"get:ShowAddType;post:HandleAddType")
	beego.Router("/article/deleteType",&controllers.ArticleController{},"get:DeleteType")
}

var funcFilter = func (ctx*context.Context){

	userName:=ctx.Input.Session("userName")

	if userName==nil{

		ctx.Redirect(302,"/login")
	}
}

