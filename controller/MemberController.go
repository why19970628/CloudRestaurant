package controller

import (
	"encoding/json"
	"fmt"
	"CloudRestaurant/model"
	"CloudRestaurant/param"
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	v1 := engine.Group("/v1/api")

	v1.GET("/", mc.sendSmsCode)        //发送短信验证码
	v1.POST("/login_sms", mc.smsLogin)         //手机号短信方式登录
	v1.GET("/captcha", mc.captcha)             //生产验证码
	v1.POST("/vertifycha", mc.vertifyCaptcha)  //验证验证码
	v1.POST("/login_pwd", mc.nameLogin)        //用户名和密码登陆接口
	v1.POST("/upload/avator", mc.uploadAvator) //头像上传
	v1.GET("/userinfo", mc.userInfo)           //获取用户信息

}

func (mc *MemberController)userInfo(context *gin.Context){
	cookie,err:=tool.CookieAuth(context)
	if err!=nil{
		context.Abort()
		tool.Failed(context,"还未登录，请先登录")
		return
	}
	memberService:=service.MemberService{}
	member:=memberService.GetUserInfo(cookie.Value)
	if member!=nil{
		tool.Success(context,map[string]interface{}{
			"id":            member.Id,
			"user_name":     member.UserName,
			"mobile":        member.Mobile,
			"register_time": member.RegisterTime,
			"avatar":        member.Avatar,
			"balance":       member.Balance,
			"city":          member.City,
		})
		return
	}
	tool.Failed(context,"获取用户信息失败")
}
//头像上传
func (mc *MemberController) uploadAvator(context *gin.Context) {

	//1、解析上传的参数：file、user_id
	userId := context.PostForm("user_id")//用户id
	file, err := context.FormFile("avator")
	if err != nil || userId == "" {
		tool.Failed(context, "参数解析失败")
		return
	}

	//2、判断user_id对应的用户是否已经登录
	sess := tool.GetSess(context, "user_"+userId)
	if sess == nil {
		tool.Failed(context, "参数不合法")
		return
	}
	var member model.Member
	json.Unmarshal(sess.([]byte),&member)

	//3、file保存到本地
	fileName:="./uploadfile/"+strconv.FormatInt(time.Now().Unix(),10)+file.Filename
	err=context.SaveUploadedFile(file,fileName)
	if err!=nil{
		tool.Failed(context,"头像更新失败")
		return
	}

	//3.1 将文件上传到fastDFS系统
	fileId:=tool.UploadFile(fileName)
	if fileId!="" {
		//删除本地uploadfile下的文件
		os.Remove(fileName)

		//4、将保存后的文件本地路径 保存到用户表中的头像字段
		memberService := service.MemberService{}
		path := memberService.UploadAvatar(member.Id, fileId)
		if path != "" {
			tool.Success(context,tool.FileServerAddr()+path)
			return
		}
	}
	//5、返回结果
	tool.Failed(context,"上传失败")
}

// http://localhost:8090/api/sendcode?phone=13167582436  //发送验证码
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	phone, exist := context.GetQuery("phone")
	if !exist {
		context.JSON(200, map[string]interface{}{
			"code": 0,
			"msg":  "参数解析失败",
		})
		return
	}

	ms := service.MemberService{}
	isSend := ms.Sendcode(phone)
	if isSend {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  "发送成功",
		})
		return
	}

	context.JSON(200, map[string]interface{}{
		"code": 0,
		"msg":  "发送失败",
	})

}

//手机号+短信  登录的方法
//  http://localhost:8090/api/login_sms
func (mc *MemberController) smsLogin(context *gin.Context) {

	var smsLoginParam param.SmsLoginParam
	//context.Request.Body 包含请求参数
	err := tool.Decode(context.Request.Body, &smsLoginParam)

	if err != nil {
		fmt.Printf("参数解析失败~~,原因是%v\n", err.Error())
		tool.Failed(context, "参数解析失败~~")
		return
	}

	//完成手机+验证码登录
	us := service.MemberService{}
	member := us.SmsLogin(smsLoginParam)

	if member != nil {
		sess, _ := json.Marshal(member)
		if err = tool.SetSess(context, "user_"+string(member.Id), sess); err != nil {
			tool.Failed(context, "登录失败")
		}
		context.SetCookie("cookie_user",strconv.Itoa(int(member.Id)), 10*60, "/", "localhost",true, true)
		tool.Success(context, member)
		return
	}

	tool.Failed(context, "登录失败")
}

//生成验证码
// http://localhost:8090/api/captcha
func (mc *MemberController) captcha(context *gin.Context) {
	tool.GenerateCaptcha(context)
}

//验证验证码是否正确
func (mc *MemberController) vertifyCaptcha(context *gin.Context) {
	var captcha tool.CaptchaResult
	// 验证是否符合captcha struct的值
	err := tool.Decode(context.Request.Body, &captcha)
	if err != nil {
		tool.Failed(context, " 参数解析失败 ")
		return
	}

	result := tool.VertifyCaptcha(captcha.Id, captcha.VertifyValue)
	if result {
		fmt.Println("验证通过")
	} else {
		fmt.Println("验证失败")
	}
}

//用户名+密码、验证码登录
func (mc *MemberController) nameLogin(context *gin.Context) {

	//1、解析用户登录传递参数
	var loginParam param.LoginParam
	err := tool.Decode(context.Request.Body, &loginParam)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}

	//2、验证验证码
	validate := tool.VertifyCaptcha(loginParam.Id, loginParam.Value)
	if !validate {
		tool.Failed(context, "验证码不正确，请重新验证")
		return
	}

	//3、登录
	ms := service.MemberService{}
	member := ms.Login(loginParam.Name, loginParam.Password)
	if member.Id != 0 {

		//登陆成功 用户信息保存到session
		sess, _ := json.Marshal(member)
		// fmt.Printf("未初始化数据是 %v , 初始化之后的数据是 %v\n", member, sess)
		err = tool.SetSess(context, "user_"+string(member.Id), sess)
		if err != nil {
			tool.Failed(context, "登录失败")
			return
		}
		context.SetCookie("cookie_user",strconv.Itoa(int(member.Id)), 10*60, "/", "localhost",true, true)
		tool.Success(context, &member)
		return
	}

	tool.Failed(context, "登录失败")
}
