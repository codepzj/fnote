// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package post_asset

import (
	"github.com/chenmingyong0423/fnote/server/internal/post_asset/internal/repository"
	"github.com/chenmingyong0423/fnote/server/internal/post_asset/internal/repository/dao"
	"github.com/chenmingyong0423/fnote/server/internal/post_asset/internal/service"
	"github.com/chenmingyong0423/fnote/server/internal/post_asset/internal/web"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injectors from wire.go:

func InitPostAssetModule(mongoDB *mongo.Database) *Module {
	postAssetDao := dao.NewPostAssetDao(mongoDB)
	postAssetRepository := repository.NewPostAssetRepository(postAssetDao)
	postAssetService := service.NewPostAssetService(postAssetRepository)
	assetFolderDao := dao.NewAssetFolderDao(mongoDB)
	assetFolderRepository := repository.NewAssetFolderRepository(assetFolderDao)
	assetFolderService := service.NewAssetFolderService(assetFolderRepository)
	postAssetHandler := web.NewPostAssetHandler(postAssetService, assetFolderService)
	module := &Module{
		Svc: postAssetService,
		Hdl: postAssetHandler,
	}
	return module
}

// wire.go:

var PostAssetProviders = wire.NewSet(web.NewPostAssetHandler, service.NewPostAssetService, repository.NewPostAssetRepository, dao.NewPostAssetDao, wire.Bind(new(service.IPostAssetService), new(*service.PostAssetService)), wire.Bind(new(repository.IPostAssetRepository), new(*repository.PostAssetRepository)), wire.Bind(new(dao.IPostAssetDao), new(*dao.PostAssetDao)), service.NewAssetFolderService, repository.NewAssetFolderRepository, dao.NewAssetFolderDao, wire.Bind(new(service.IAssetFolderService), new(*service.AssetFolderService)), wire.Bind(new(repository.IAssetFolderRepository), new(*repository.AssetFolderRepository)), wire.Bind(new(dao.IAssetFolderDao), new(*dao.AssetFolderDao)))
