// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/chenmingyong0423/fnote/server/internal/aggregate_post"
	handler7 "github.com/chenmingyong0423/fnote/server/internal/backup/handler"
	service11 "github.com/chenmingyong0423/fnote/server/internal/backup/service"
	handler2 "github.com/chenmingyong0423/fnote/server/internal/category/handler"
	repository2 "github.com/chenmingyong0423/fnote/server/internal/category/repository"
	dao2 "github.com/chenmingyong0423/fnote/server/internal/category/repository/dao"
	service3 "github.com/chenmingyong0423/fnote/server/internal/category/service"
	"github.com/chenmingyong0423/fnote/server/internal/comment/hanlder"
	repository4 "github.com/chenmingyong0423/fnote/server/internal/comment/repository"
	dao4 "github.com/chenmingyong0423/fnote/server/internal/comment/repository/dao"
	service4 "github.com/chenmingyong0423/fnote/server/internal/comment/service"
	handler6 "github.com/chenmingyong0423/fnote/server/internal/count_stats/handler"
	repository3 "github.com/chenmingyong0423/fnote/server/internal/count_stats/repository"
	dao3 "github.com/chenmingyong0423/fnote/server/internal/count_stats/repository/dao"
	service2 "github.com/chenmingyong0423/fnote/server/internal/count_stats/service"
	"github.com/chenmingyong0423/fnote/server/internal/data_analysis"
	service5 "github.com/chenmingyong0423/fnote/server/internal/email/service"
	"github.com/chenmingyong0423/fnote/server/internal/file/handler"
	"github.com/chenmingyong0423/fnote/server/internal/file/repository"
	"github.com/chenmingyong0423/fnote/server/internal/file/repository/dao"
	"github.com/chenmingyong0423/fnote/server/internal/file/service"
	hanlder2 "github.com/chenmingyong0423/fnote/server/internal/friend/hanlder"
	repository6 "github.com/chenmingyong0423/fnote/server/internal/friend/repository"
	dao6 "github.com/chenmingyong0423/fnote/server/internal/friend/repository/dao"
	service8 "github.com/chenmingyong0423/fnote/server/internal/friend/service"
	"github.com/chenmingyong0423/fnote/server/internal/global"
	"github.com/chenmingyong0423/fnote/server/internal/ioc"
	service7 "github.com/chenmingyong0423/fnote/server/internal/message/service"
	handler4 "github.com/chenmingyong0423/fnote/server/internal/message_template/handler"
	repository5 "github.com/chenmingyong0423/fnote/server/internal/message_template/repository"
	dao5 "github.com/chenmingyong0423/fnote/server/internal/message_template/repository/dao"
	service6 "github.com/chenmingyong0423/fnote/server/internal/message_template/service"
	"github.com/chenmingyong0423/fnote/server/internal/post"
	"github.com/chenmingyong0423/fnote/server/internal/post_draft"
	"github.com/chenmingyong0423/fnote/server/internal/post_index"
	"github.com/chenmingyong0423/fnote/server/internal/post_like"
	handler5 "github.com/chenmingyong0423/fnote/server/internal/tag/handler"
	repository8 "github.com/chenmingyong0423/fnote/server/internal/tag/repository"
	dao8 "github.com/chenmingyong0423/fnote/server/internal/tag/repository/dao"
	service10 "github.com/chenmingyong0423/fnote/server/internal/tag/service"
	handler3 "github.com/chenmingyong0423/fnote/server/internal/visit_log/handler"
	repository7 "github.com/chenmingyong0423/fnote/server/internal/visit_log/repository"
	dao7 "github.com/chenmingyong0423/fnote/server/internal/visit_log/repository/dao"
	service9 "github.com/chenmingyong0423/fnote/server/internal/visit_log/service"
	"github.com/chenmingyong0423/fnote/server/internal/website_config"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func initializeApp() (*gin.Engine, error) {
	database := ioc.NewMongoDB()
	fileDao := dao.NewFileDao(database)
	fileRepository := repository.NewFileRepository(fileDao)
	fileService := service.NewFileService(fileRepository)
	fileHandler := handler.NewFileHandler(fileService)
	categoryDao := dao2.NewCategoryDao(database)
	categoryRepository := repository2.NewCategoryRepository(categoryDao)
	countStatsDao := dao3.NewCountStatsDao(database)
	countStatsRepository := repository3.NewCountStatsRepository(countStatsDao)
	countStatsService := service2.NewCountStatsService(countStatsRepository)
	model := website_config.InitWebsiteConfigModule(database)
	iWebsiteConfigService := model.Svc
	categoryService := service3.NewCategoryService(categoryRepository, countStatsService, iWebsiteConfigService)
	categoryHandler := handler2.NewCategoryHandler(categoryService)
	commentDao := dao4.NewCommentDao(database)
	commentRepository := repository4.NewCommentRepository(commentDao)
	commentService := service4.NewCommentService(commentRepository)
	post_likeModel := post_like.InitPostLikeModule(database)
	postModel := post.InitPostModule(database, model, countStatsService, fileService, post_likeModel)
	iPostService := postModel.Svc
	emailService := service5.NewEmailService()
	msgTplDao := dao5.NewMsgTplDao(database)
	msgTplRepository := repository5.NewMsgTplRepository(msgTplDao)
	msgTplService := service6.NewMsgTplService(msgTplRepository)
	messageService := service7.NewMessageService(iWebsiteConfigService, emailService, msgTplService)
	commentHandler := hanlder.NewCommentHandler(commentService, iWebsiteConfigService, iPostService, messageService, countStatsService)
	websiteConfigHandler := model.Hdl
	friendDao := dao6.NewFriendDao(database)
	friendRepository := repository6.NewFriendRepository(friendDao)
	friendService := service8.NewFriendService(friendRepository)
	friendHandler := hanlder2.NewFriendHandler(friendService, messageService, iWebsiteConfigService)
	postHandler := postModel.Hdl
	visitLogDao := dao7.NewVisitLogDao(database)
	visitLogRepository := repository7.NewVisitLogRepository(visitLogDao)
	visitLogService := service9.NewVisitLogService(visitLogRepository)
	visitLogHandler := handler3.NewVisitLogHandler(visitLogService, countStatsService)
	msgTplHandler := handler4.NewMsgTplHandler(msgTplService)
	tagDao := dao8.NewTagDao(database)
	tagRepository := repository8.NewTagRepository(tagDao)
	tagService := service10.NewTagService(tagRepository, countStatsService)
	tagHandler := handler5.NewTagHandler(tagService)
	module := data_analysis.InitDataAnalysisModule(database, visitLogService, countStatsService)
	dataAnalysisHandler := module.Hdl
	countStatsHandler := handler6.NewCountStatsHandler(countStatsService)
	backupService := service11.NewBackupService(database)
	backupHandler := handler7.NewBackupHandler(backupService)
	writer := ioc.InitLogger()
	v, err := global.IsWebsiteInitializedFn(database)
	if err != nil {
		return nil, err
	}
	v2 := ioc.InitMiddlewares(writer, v)
	validators := ioc.InitGinValidators()
	post_indexModel := post_index.InitPostIndexModule(model)
	postIndexHandler := post_indexModel.Hdl
	post_draftModel := post_draft.InitPostDraftModule(database)
	postDraftHandler := post_draftModel.Hdl
	aggregate_postModel := aggregate_post.InitAggregatePostModule(postModel, post_draftModel)
	aggregatePostHandler := aggregate_postModel.Hdl
	postLikeHandler := post_likeModel.Hdl
	engine, err := ioc.NewGinEngine(fileHandler, categoryHandler, commentHandler, websiteConfigHandler, friendHandler, postHandler, visitLogHandler, msgTplHandler, tagHandler, dataAnalysisHandler, countStatsHandler, backupHandler, v2, validators, postIndexHandler, postDraftHandler, aggregatePostHandler, postLikeHandler)
	if err != nil {
		return nil, err
	}
	return engine, nil
}
