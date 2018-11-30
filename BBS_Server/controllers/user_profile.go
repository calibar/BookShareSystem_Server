package controllers

import (
	"BBS_Server/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/gomail.v2"
	"math/rand"
	"strconv"
	"strings"
)

// UserProfileController operations for UserProfile
type UserProfileController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserProfileController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create UserProfile
// @Param	body		body 	models.UserProfile	true		"body for UserProfile content"
// @Success 201 {int} models.UserProfile
// @Failure 403 body is empty
// @router / [post]
func (c *UserProfileController) Post() {
	var v models.UserProfile
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddUserProfile(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get UserProfile by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UserProfile
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserProfileController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUserProfileById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get UserProfile

// @Param   singleuser query string false   "username"
// @Param   verify_confirm  query   string  false   "username|verify code"
// @Param	verify_email	query	string	false	"Filter. e.g. username|email"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.UserProfile
// @Failure 403
// @router / [get]
func (c *UserProfileController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64
	// singleuser : username find user profile with ranking
	if v:=c.GetString("verify_confirm");v!=""{
		strs:=strings.Split(v,"|")
		username:=strs[0]
		verifyCode,_:=strconv.Atoi(strs[1])
		result:=models.VerifyVCode(username,verifyCode)
		if result==true {
			c.Ctx.ResponseWriter.WriteHeader(200)
			body:=`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Jumping Page</title>
</head>
<body>
<div align="center"><img src="https://images.chamaileon.io/5bff50b172eff2003915899e/CBS.png" alt="" width="200" height="200" />
    <h1 style="color: #97690c;">Congratulations!</h1>
    <br />
    <div>
        <h2 style="color: #97690c; text-align: center;">`+username+`</h2>
        <p>&nbsp;</p>
        <p style="color: #97690c;">You have successfully verified your email in Campus Book Sharing.&nbsp;</p>
        <p style="color: #97690c;">Click&nbsp;<a href="http://www.baidu.com"><span style="background-color: #97690c; color: #fff; display: inline-block; padding: 3px 10px; font-weight: bold; border-radius: 5px;">Here</span></a>&nbsp;to start your&nbsp;sharing now.</p>
    </div>
</div>
</body>
</html>`

			c.Ctx.ResponseWriter.Header().Set("Content-Type","text/html")
			c.Ctx.ResponseWriter.Write([]byte(body))
		}else{
			c.Ctx.ResponseWriter.WriteHeader(408)
			c.Ctx.ResponseWriter.Write([]byte("conflict"))
		}
	}else {
		if v:= c.GetString("verify_email");v!=""{
			strs:=strings.Split(v,"|")
			username:=strs[0]
			strs1:=strings.Split(strs[1],"@")
			email:=strs[1]
			fmt.Println(username)
			fmt.Println(email)
			existence:=models.CheckCampusEmail(strs1[1])
			if existence {
				verifyCode:=rand.Intn(100000)
				models.GenerateVerifyCode(username,verifyCode)
				sendEmail(username,email,verifyCode)
			}else {
				c.Ctx.ResponseWriter.WriteHeader(404)
				c.Ctx.ResponseWriter.Write([]byte("no such campus domain"))
			}
		}else {
			if v := c.GetString("singleuser");v!=""{
				up,err:=models.GetSingleUserProfileByUsername(v)
				if err != nil {
					c.Data["json"] = err.Error()
					c.ServeJSON()
				} else {
					c.Data["json"] = up
					c.ServeJSON()
				}
			}else {
				// fields: col1,col2,entity.col3
				if v := c.GetString("fields"); v != "" {
					fields = strings.Split(v, ",")
				}
				// limit: 10 (default is 10)
				if v, err := c.GetInt64("limit"); err == nil {
					limit = v
				}
				// offset: 0 (default is 0)
				if v, err := c.GetInt64("offset"); err == nil {
					offset = v
				}
				// sortby: col1,col2
				if v := c.GetString("sortby"); v != "" {
					sortby = strings.Split(v, ",")
				}
				// order: desc,asc
				if v := c.GetString("order"); v != "" {
					order = strings.Split(v, ",")
				}
				// query: k:v,k:v
				if v := c.GetString("query"); v != "" {
					for _, cond := range strings.Split(v, ",") {
						kv := strings.SplitN(cond, ":", 2)
						if len(kv) != 2 {
							c.Data["json"] = errors.New("Error: invalid query key/value pair")
							c.ServeJSON()
							return
						}
						k, v := kv[0], kv[1]
						query[k] = v
					}
				}

				l, err := models.GetAllUserProfile(query, fields, sortby, order, offset, limit)
				if err != nil {
					c.Data["json"] = err.Error()
					c.ServeJSON()
				} else {
					c.Data["json"] = l
					c.ServeJSON()
				}
			}
		}
	}




}

// Put ...
// @Title Put
// @Description update the UserProfile
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.UserProfile	true		"body for UserProfile content"
// @Success 200 {object} models.UserProfile
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserProfileController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.UserProfile{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateUserProfileById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the UserProfile
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserProfileController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUserProfile(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
func sendEmail(username string,verify_email string,verifyCode int)  {
	vCode:=strconv.Itoa(verifyCode)
	m := gomail.NewMessage()
	m.SetHeader("From", "campusbooksharingverifier@gmail.com")
	m.SetHeader("To", verify_email)
	m.SetHeader("Subject", "Do-not-reply,Campus Book Sharing email verification.")
	body:=`
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta name="viewport" content="width=device-width">
<meta name="HandheldFriendly" content="true" />
<meta http-equiv="X-UA-Compatible" content="IE=edge" />
<!--[if gte IE 7]><html class="ie8plus" xmlns="http://www.w3.org/1999/xhtml"><![endif]-->
<!--[if IEMobile]><html class="ie8plus" xmlns="http://www.w3.org/1999/xhtml"><![endif]-->
<meta name="format-detection" content="telephone=no">
<meta name="generator" content="EDMdesigner, www.edmdesigner.com">
<title>Campus Book Sharing Email Comfirmation</title>

<link href="https://fonts.googleapis.com/css?family=Merriweather+Sans" rel="stylesheet" type="text/css">
<link href="https://fonts.googleapis.com/css?family=Merriweather" rel="stylesheet" type="text/css">

<style type="text/css" media="screen">
* {line-height: inherit;}
.ExternalClass * { line-height: 100%; }
body, p{margin:0; padding:0; margin-bottom:0; -webkit-text-size-adjust:none; -ms-text-size-adjust:none;} img{line-height:100%; outline:none; text-decoration:none; -ms-interpolation-mode: bicubic;} a img{border: none;} a, a:link, .no-detect-local a, .appleLinks a{color:#5555ff !important; text-decoration: underline;} .ExternalClass {display: block !important; width:100%;} .ExternalClass, .ExternalClass p, .ExternalClass span, .ExternalClass font, .ExternalClass td, .ExternalClass div { line-height: inherit; } table td {border-collapse:collapse;mso-table-lspace: 0pt; mso-table-rspace: 0pt;} sup{position: relative; top: 4px; line-height:7px !important;font-size:11px !important;} .mobile_link a[href^="tel"], .mobile_link a[href^="sms"] {text-decoration: default; color: #5555ff !important;
pointer-events: auto; cursor: default;} .no-detect a{text-decoration: none; color: #5555ff; pointer-events: auto; cursor: default;} {color: #5555ff;} span {color: inherit; border-bottom: none;} span:hover { background-color: transparent; }

.nounderline {text-decoration: none !important;}
h1, h2, h3 { margin:0; padding:0; }
p {Margin: 0px !important; }

table[class="email-root-wrapper"] { width: 600px !important; }

body {
background-color: #ffffff;
background: #ffffff;
}
body { min-width: 280px; width: 100%;}

</style>
<style>
@media only screen and (max-width: 599px),
only screen and (max-device-width: 599px),
only screen and (max-width: 400px),
only screen and (max-device-width: 400px) {
 .email-root-wrapper { width: 100% !important; }
 .full-width { width: 100% !important; height: auto !important; text-align:center;}
 .fullwidthhalfleft {width:100% !important;}
 .fullwidthhalfright {width:100% !important;}
 .fullwidthhalfinner {width:100% !important; margin: 0 auto !important; float: none !important; margin-left: auto !important; margin-right: auto !important; clear:both !important; }
 .hide { display:none !important; width:0px !important;height:0px !important; overflow:hidden; }
 .desktop-hide { display:block !important; width:100% !important;height:auto !important; overflow:hidden; max-height: inherit !important; }
	
}
</style>
<style>
@media only screen and (min-width: 600px) {
  
}
@media only screen and (max-width: 599px),
only screen and (max-device-width: 599px),
only screen and (max-width: 400px),
only screen and (max-device-width: 400px) {
  table[class="email-root-wrapper"] { width: 100% !important; }
  td[class="wrap"] .full-width { width: 100% !important; height: auto !important;}

  td[class="wrap"] .fullwidthhalfleft {width:100% !important;}
  td[class="wrap"] .fullwidthhalfright {width:100% !important;}
  td[class="wrap"] .fullwidthhalfinner {width:100% !important; margin: 0 auto !important; float: none !important; margin-left: auto !important; margin-right: auto !important; clear:both !important; }
  td[class="wrap"] .hide { display:none !important; width:0px;height:0px; overflow:hidden; }

  
}


</style>

<!--[if (gte IE 7) & (vml)]>
<style type="text/css">
html, body {margin:0 !important; padding:0px !important;}
img.full-width { position: relative !important; }

.img200x200 { width: 200px !important; height: 200px !important;}

</style>
<![endif]-->

<!--[if gte mso 9]>
<style type="text/css">
.mso-font-fix-arial { font-family: Arial, sans-serif;}
.mso-font-fix-georgia { font-family: Georgia, sans-serif;}
.mso-font-fix-tahoma { font-family: Tahoma, sans-serif;}
.mso-font-fix-times_new_roman { font-family: 'Times New Roman', sans-serif;}
.mso-font-fix-trebuchet_ms { font-family: 'Trebuchet MS', sans-serif;}
.mso-font-fix-verdana { font-family: Verdana, sans-serif;}
</style>
<![endif]-->

<!--[if gte mso 9]>
<style type="text/css">
table, td {
border-collapse: collapse !important;
mso-table-lspace: 0px !important;
mso-table-rspace: 0px !important;
}

.email-root-wrapper { width 600px !important;}
.imglink { font-size: 0px; }
.edm_button { font-size: 0px; }
</style>
<![endif]-->

<!--[if gte mso 15]>
<style type="text/css">
table {
font-size:0px;
mso-margin-top-alt:0px;
}

.fullwidthhalfleft {
width: 49% !important;
float:left !important;
}

.fullwidthhalfright {
width: 50% !important;
float:right !important;
}
</style>
<![endif]-->
<STYLE type="text/css" media="(pointer) and (min-color-index:0)">
html, body {background-image: none !important; background-color: transparent !important; margin:0 !important; padding:0 !important;}
</STYLE>

</head>
<body leftmargin="0" marginwidth="0" topmargin="0" marginheight="0" offset="0" style="font-family:Arial, sans-serif; font-size:0px;margin:0;padding:0;background: #ffffff !important;" bgcolor="#ffffff">
<!--[if t]><![endif]--><!--[if t]><![endif]--><!--[if t]><![endif]--><!--[if t]><![endif]--><!--[if t]><![endif]--><!--[if t]><![endif]-->
  <table align="center" border="0" cellpadding="0" cellspacing="0" height="100%" width="100%"  bgcolor="#ffffff" style="margin:0; padding:0; width:100% !important; background: #ffffff !important;">
    <tr>
        <td class="wrap" align="center" valign="top" width="100%">
          <center>
<!-- content -->
<div  style="padding:0px"><table cellpadding="0" cellspacing="0" border="0" width="100%"><tr><td valign="top"  style="padding:0px"><table cellpadding="0" cellspacing="0" width="600" align="center"  style="max-width:600px;min-width:240px;margin:0 auto" class="email-root-wrapper"><tr><td valign="top"  style="padding:0px"><table cellpadding="0" cellspacing="0" border="0" width="100%"  style="border:0px none"><tr><td valign="top"  style="padding:0px"><table cellpadding="0" cellspacing="0" width="100%"><tr><td  style="padding:0px"><table cellpadding="0" cellspacing="0" width="100%"><tr><td align="center"  style="padding:0px"><table cellpadding="0" cellspacing="0" border="0" align="center" width="200" height="200"  style="border:0px none;height:auto"><tr><td valign="top" 
style="padding:0px"><img
src="https://images.chamaileon.io/5bff50b172eff2003915899e/CBS.png" width="200" height="200" alt="" border="0"  style="display:block" class="img200x200"  /></td>
</tr>
</table>
</td>
</tr>
</table>
<table cellpadding="0" cellspacing="0" border="0" width="100%"><tr><td valign="top"  style="padding:17px"><div  style="text-align:left;font-family:Merriweather Sans, Helvetica Neue, Helvetica, Arial, sans-serif;font-size:15px;color:#756735;line-height:17px;mso-line-height:exactly;mso-text-raise:1px"><p style="padding: 0; margin: 0;"><strong>Hello `+username+`!</strong></p><p style="padding: 0; margin: 0;">&nbsp;</p><p style="padding: 0; margin: 0;"><strong><span style="background-color: transparent;">We are so excited to have you. Hope you have fun on Campus Book Sharing.</span></strong></p><p style="padding: 0; margin: 0;">&nbsp;</p><p style="padding: 0; margin: 0;"><strong><span style="background-color: transparent;">Best regards.</span></strong></p><p style="padding: 0; margin: 0;">&nbsp;</p><p
style="padding: 0; margin: 0;"><strong><span style="background-color: transparent;">Team Campus Book Sharing</span></strong></p></div></td>
</tr>
</table>
<table cellpadding="0" cellspacing="0" width="100%"><tr><td align="center"  style="padding:0px"><table cellpadding="0" cellspacing="0" border="0" align="center"  style="text-align:center;color:#000"><tr><td valign="top" align="center"  style="padding:22px"><table cellpadding="0" cellspacing="0" border="0" bgcolor="#22552a"  style="border:2px solid #97690c;border-radius:15px;border-collapse:separate !important;background-color:#ffffff"><tr><td valign="top" align="center"  style="padding:16px"><a 
href="http://d4de8068.ngrok.io/v1/user_profile?verify_confirm=`+username+`|`+vCode+`" target="_blank"  style="text-decoration:none" class="edm_button"><span  style="font-family:Arial, Helvetica Neue, Helvetica, sans-serif;font-size:26px;color:#97690c;line-height:26px;text-decoration:none"><span class="mso-font-fix-arial">Confirm</span></span>
</a></td>
</tr>
</table>
</td>
</tr>
</table>
</td>
</tr>
</table>
</td>
</tr>
</table>
</td>
</tr>
</table>
</td>
</tr>
</table>
</td>
</tr>
</table>
</div>
<!-- content end -->
          </center>
        </td>
    </tr>
  </table>
</body>
</html>
`
	m.SetBody("text/html", body)

	// Send the email to Bob
	d := gomail.NewPlainDialer("smtp.gmail.com", 587, "campusbooksharingverifier@gmail.com", "lmx1993917")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}