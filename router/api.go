package router

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/controller"
	"github.com/taoshihan1991/imaptool/middleware"
	"github.com/taoshihan1991/imaptool/ws"
)

func InitApiRouter(engine *gin.Engine) {
	//路由分组
	v2 := engine.Group("/2")
	{
		//获取消息
		v2.GET("/messages", controller.GetMessagesV2)
		//发送单条信息
		v2.POST("/message", middleware.Ipblack, controller.SendMessageV2)
		//关闭连接
		v2.GET("/message_close", controller.SendCloseMessageV2)
		//绑定
		v2.POST("/bindOfficial", controller.PostBindOfficial)
		//分页查询消息
		v2.GET("/messagesPages", controller.GetMessagespages)
	}
	engine.GET("/captcha", controller.GetCaptcha)
	engine.POST("/check", controller.LoginCheckPass)
	engine.POST("/check_auth", middleware.JwtApiMiddleware, controller.MainCheckAuth)
	engine.GET("/userinfo", middleware.JwtApiMiddleware, controller.GetKefuInfoAll)
	engine.POST("/register", middleware.Ipblack, controller.PostKefuRegister)
	engine.POST("/install", controller.PostInstall)
	//前后聊天
	engine.GET("/ws_kefu", middleware.JwtApiMiddleware, ws.NewKefuServer)
	engine.GET("/ws_visitor", middleware.Ipblack, ws.NewVisitorServer)

	engine.GET("/messages", controller.GetVisitorMessage)
	engine.GET("/message_notice", controller.SendVisitorNotice)
	//上传文件
	engine.POST("/uploadimg", middleware.Ipblack, controller.UploadImg)
	//上传文件
	engine.POST("/uploadfile", middleware.Ipblack, controller.UploadFile)
	//获取未读消息数
	engine.GET("/message_status", controller.GetVisitorMessage)
	//设置消息已读
	engine.POST("/message_status", controller.GetVisitorMessage)

	//获取客服信息
	engine.POST("/kefuinfo_client", middleware.JwtApiMiddleware, controller.PostKefuClient)
	engine.GET("/kefuinfo", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.GetKefuInfo)
	engine.GET("/kefuinfo_setting", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.GetKefuInfoSetting)
	engine.POST("/kefuinfo", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.PostKefuInfo)
	engine.DELETE("/kefuinfo", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.DeleteKefuInfo)
	engine.GET("/kefulist", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.GetKefuList)
	engine.GET("/other_kefulist", middleware.JwtApiMiddleware, controller.GetOtherKefuList)
	engine.GET("/trans_kefu", middleware.JwtApiMiddleware, controller.PostTransKefu)
	engine.POST("/modifypass", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.PostKefuPass)
	engine.POST("/modifyavator", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.PostKefuAvator)
	//角色列表
	engine.GET("/roles", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.GetRoleList)
	engine.POST("/role", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.PostRole)

	engine.GET("/visitors_online", controller.GetVisitorOnlines)
	engine.GET("/visitors_kefu_online", middleware.JwtApiMiddleware, controller.GetKefusVisitorOnlines)
	engine.GET("/clear_online_tcp", controller.DeleteOnlineTcp)
	engine.POST("/visitor_login", middleware.Ipblack, controller.PostVisitorLogin)
	//engine.POST("/visitor", controller.PostVisitor)
	engine.GET("/visitor", middleware.JwtApiMiddleware, controller.GetVisitor)
	engine.GET("/visitors", middleware.JwtApiMiddleware, controller.GetVisitors)
	engine.GET("/statistics", middleware.JwtApiMiddleware, controller.GetStatistics)
	//前台接口
	engine.GET("/about", controller.GetAbout)
	engine.POST("/about", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.PostAbout)
	engine.GET("/aboutpages", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.GetAbouts)
	engine.GET("/notice", controller.GetNotice)
	engine.POST("/ipblack", middleware.JwtApiMiddleware, middleware.Ipblack, controller.PostIpblack)
	engine.DELETE("/ipblack", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.DelIpblack)
	engine.GET("/ipblacks_all", middleware.JwtApiMiddleware, controller.GetIpblacks)
	engine.GET("/ipblacks", middleware.JwtApiMiddleware, controller.GetIpblacksByKefuId)
	engine.GET("/configs", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.GetConfigs)
	engine.POST("/config", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.PostConfig)
	engine.GET("/config", controller.GetConfig)
	engine.GET("/autoreply", controller.GetAutoReplys)
	engine.GET("/replys", middleware.JwtApiMiddleware, controller.GetReplys)
	engine.POST("/reply", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.PostReply)
	engine.POST("/reply_content", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.PostReplyContent)
	engine.POST("/reply_content_save", middleware.JwtApiMiddleware, controller.PostReplyContentSave)
	engine.DELETE("/reply_content", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.DelReplyContent)
	engine.DELETE("/reply", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.DelReplyGroup)
	engine.POST("/reply_search", middleware.JwtApiMiddleware, controller.PostReplySearch)
	//客服路由分组
	kefuGroup := engine.Group("/kefu")
	kefuGroup.Use(middleware.JwtApiMiddleware)
	{
		kefuGroup.GET("/chartStatistics", controller.GetChartStatistic)
	}
	//微信接口
	engine.GET("/micro_program", middleware.JwtApiMiddleware, controller.GetCheckWeixinSign)

	// ===== 課程計畫填報系統 =====
	curriculumGroup := engine.Group("/curriculum")
	curriculumGroup.Use(middleware.JwtApiMiddleware)
	{
		// 計畫管理
		curriculumGroup.GET("/plans", controller.GetCurriculumPlans)
		curriculumGroup.GET("/plan", controller.GetCurriculumPlan)
		curriculumGroup.POST("/plan", middleware.RbacAuth, controller.PostCurriculumPlan)
		curriculumGroup.DELETE("/plan", middleware.RbacAuth, controller.DeleteCurriculumPlan)

		// 計畫表單欄位
		curriculumGroup.GET("/plan_fields", controller.GetPlanFormFields)
		curriculumGroup.POST("/plan_field", middleware.RbacAuth, controller.PostPlanFormField)
		curriculumGroup.DELETE("/plan_field", middleware.RbacAuth, controller.DeletePlanFormField)

		// 填報管理
		curriculumGroup.GET("/submissions", controller.GetPlanSubmissions)
		curriculumGroup.GET("/submission", controller.GetPlanSubmission)
		curriculumGroup.POST("/submission", controller.PostPlanSubmission)
		curriculumGroup.POST("/submit", controller.PostSubmitPlan)

		// 審查管理
		curriculumGroup.GET("/reviews", controller.GetPlanReviews)
		curriculumGroup.POST("/review", controller.PostPlanReview)

		// 審查指派
		curriculumGroup.GET("/review_assignments", controller.GetReviewAssignments)
		curriculumGroup.POST("/review_assignment", middleware.RbacAuth, controller.PostReviewAssignment)

		// 統計
		curriculumGroup.GET("/statistics", controller.GetCurriculumStatistics)
	}

	// 組織管理
	orgGroup := engine.Group("/org")
	orgGroup.Use(middleware.JwtApiMiddleware)
	{
		orgGroup.GET("/list", controller.GetOrganizations)
		orgGroup.GET("/detail", controller.GetOrganization)
		orgGroup.POST("/save", middleware.RbacAuth, controller.PostOrganization)
		orgGroup.DELETE("/delete", middleware.RbacAuth, controller.DeleteOrganization)
		orgGroup.GET("/schools_by_county", controller.GetSchoolsByCounty)
	}

	// 平台使用者管理
	puGroup := engine.Group("/platform_user")
	puGroup.Use(middleware.JwtApiMiddleware)
	{
		puGroup.GET("/list", controller.GetPlatformUsers)
		puGroup.GET("/detail", controller.GetPlatformUser)
		puGroup.POST("/save", middleware.RbacAuth, controller.PostPlatformUser)
		puGroup.GET("/committee_members", controller.GetCommitteeMembers)
	}

	// 通知與日誌
	notiGroup := engine.Group("/notification")
	notiGroup.Use(middleware.JwtApiMiddleware)
	{
		notiGroup.GET("/list", controller.GetNotifications)
		notiGroup.POST("/read", controller.PostNotificationRead)
		notiGroup.GET("/audit_logs", middleware.RbacAuth, controller.GetAuditLogs)
	}
}
