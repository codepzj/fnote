// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build wireinject

package main

import (
	"github.com/chenmingyong0423/fnote/server/internal/global"
	"github.com/chenmingyong0423/fnote/server/internal/ioc"
	"github.com/chenmingyong0423/fnote/server/internal/post_index"
	"github.com/chenmingyong0423/fnote/server/internal/website_config"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func initializeApp() (*gin.Engine, error) {
	panic(wire.Build(
		ioc.InitLogger,
		ioc.NewMongoDB,
		ioc.InitMiddlewares,
		ioc.InitGinValidators,
		ioc.NewGinEngine,
		global.IsWebsiteInitializedFn,

		ioc.CategoryProviders,
		ioc.CommentProviders,
		website_config.InitWebsiteConfigModule,
		wire.FieldsOf(new(website_config.Model), "Svc"),
		wire.FieldsOf(new(website_config.Model), "Hdl"),
		post_index.InitPostIndexModule,
		wire.FieldsOf(new(post_index.Model), "Hdl"),
		ioc.FriendProviders,
		ioc.PostProviders,
		ioc.VlProviders,
		ioc.EmailProviders,
		ioc.MsgProviders,
		ioc.MsgTplProviders,
		ioc.CountStatsProviders,
		ioc.TagProviders,
		ioc.FileProviders,
		ioc.DataAnalysisProviders,
		ioc.BackupProviders,
	))
}
