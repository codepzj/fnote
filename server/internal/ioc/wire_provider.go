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

package ioc

import (
	handler12 "github.com/chenmingyong0423/fnote/server/internal/backup/handler"
	service13 "github.com/chenmingyong0423/fnote/server/internal/backup/service"
	"github.com/chenmingyong0423/fnote/server/internal/category/handler"
	"github.com/chenmingyong0423/fnote/server/internal/category/repository"
	"github.com/chenmingyong0423/fnote/server/internal/category/repository/dao"
	"github.com/chenmingyong0423/fnote/server/internal/category/service"
	service7 "github.com/chenmingyong0423/fnote/server/internal/email/service"
	service8 "github.com/chenmingyong0423/fnote/server/internal/message/service"
	handler7 "github.com/chenmingyong0423/fnote/server/internal/message_template/handler"
	repository7 "github.com/chenmingyong0423/fnote/server/internal/message_template/repository"
	dao7 "github.com/chenmingyong0423/fnote/server/internal/message_template/repository/dao"
	service9 "github.com/chenmingyong0423/fnote/server/internal/message_template/service"
	handler8 "github.com/chenmingyong0423/fnote/server/internal/tag/handler"
	repository9 "github.com/chenmingyong0423/fnote/server/internal/tag/repository"
	dao9 "github.com/chenmingyong0423/fnote/server/internal/tag/repository/dao"
	service11 "github.com/chenmingyong0423/fnote/server/internal/tag/service"
	handler6 "github.com/chenmingyong0423/fnote/server/internal/visit_log/handler"
	repository6 "github.com/chenmingyong0423/fnote/server/internal/visit_log/repository"
	dao6 "github.com/chenmingyong0423/fnote/server/internal/visit_log/repository/dao"
	service6 "github.com/chenmingyong0423/fnote/server/internal/visit_log/service"
	"github.com/google/wire"
)

var (
	CategoryProviders = wire.NewSet(handler.NewCategoryHandler, service.NewCategoryService, repository.NewCategoryRepository, dao.NewCategoryDao,
		wire.Bind(new(service.ICategoryService), new(*service.CategoryService)),
		wire.Bind(new(repository.ICategoryRepository), new(*repository.CategoryRepository)),
		wire.Bind(new(dao.ICategoryDao), new(*dao.CategoryDao)),
	)

	VlProviders = wire.NewSet(handler6.NewVisitLogHandler, service6.NewVisitLogService, repository6.NewVisitLogRepository, dao6.NewVisitLogDao,
		wire.Bind(new(service6.IVisitLogService), new(*service6.VisitLogService)),
		wire.Bind(new(repository6.IVisitLogRepository), new(*repository6.VisitLogRepository)),
		wire.Bind(new(dao6.IVisitLogDao), new(*dao6.VisitLogDao)))

	EmailProviders = wire.NewSet(service7.NewEmailService, wire.Bind(new(service7.IEmailService), new(*service7.EmailService)))

	MsgProviders = wire.NewSet(service8.NewMessageService, wire.Bind(new(service8.IMessageService), new(*service8.MessageService)))

	MsgTplProviders = wire.NewSet(handler7.NewMsgTplHandler, service9.NewMsgTplService, repository7.NewMsgTplRepository, dao7.NewMsgTplDao,
		wire.Bind(new(service9.IMsgTplService), new(*service9.MsgTplService)),
		wire.Bind(new(repository7.IMsgTplRepository), new(*repository7.MsgTplRepository)),
		wire.Bind(new(dao7.IMsgTplDao), new(*dao7.MsgTplDao)))

	TagProviders = wire.NewSet(handler8.NewTagHandler, service11.NewTagService, repository9.NewTagRepository, dao9.NewTagDao,
		wire.Bind(new(service11.ITagService), new(*service11.TagService)),
		wire.Bind(new(repository9.ITagRepository), new(*repository9.TagRepository)),
		wire.Bind(new(dao9.ITagDao), new(*dao9.TagDao)),
	)

	BackupProviders = wire.NewSet(handler12.NewBackupHandler, service13.NewBackupService,
		wire.Bind(new(service13.IBackupService), new(*service13.BackupService)),
	)
)
