// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"BBS_Server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/apply",
			beego.NSInclude(
				&controllers.ApplicantListController{},
			),
		),

		beego.NSNamespace("/book_transaction",
			beego.NSInclude(
				&controllers.BookTransactionController{},
			),
		),

		beego.NSNamespace("/bug_report",
			beego.NSInclude(
				&controllers.BugReportController{},
			),
		),

		beego.NSNamespace("/campus",
			beego.NSInclude(
				&controllers.CampusController{},
			),
		),

		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/messger",
			beego.NSInclude(
				&controllers.UserMessageController{},
			),
		),

		beego.NSNamespace("/user_profile",
			beego.NSInclude(
				&controllers.UserProfileController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
