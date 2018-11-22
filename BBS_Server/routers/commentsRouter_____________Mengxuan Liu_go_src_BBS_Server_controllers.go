package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["BBS_Server/controllers:ApplicantListController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:ApplicantListController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:ApplicantListController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:ApplicantListController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:ApplicantListController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:ApplicantListController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:ApplicantListController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:ApplicantListController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:ApplicantListController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:ApplicantListController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:BookTransactionController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:BookTransactionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:BookTransactionController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:BookTransactionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:BookTransactionController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:BookTransactionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:BookTransactionController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:BookTransactionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:BookTransactionController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:BookTransactionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:BugReportController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:BugReportController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:BugReportController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:BugReportController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:BugReportController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:BugReportController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:BugReportController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:BugReportController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:BugReportController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:BugReportController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:CampusController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:CampusController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:CampusController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:CampusController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:CampusController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:CampusController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:CampusController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:CampusController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:CampusController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:CampusController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:UserController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:UserController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:UserController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:UserController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:UserController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:UserMessageController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:UserMessageController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:UserMessageController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:UserMessageController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:UserMessageController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:UserMessageController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:UserMessageController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:UserMessageController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:UserMessageController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:UserMessageController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:UserProfileController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:UserProfileController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:UserProfileController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:UserProfileController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:UserProfileController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:UserProfileController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:UserProfileController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:UserProfileController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BBS_Server/controllers:UserProfileController"] = append(beego.GlobalControllerRouter["BBS_Server/controllers:UserProfileController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
