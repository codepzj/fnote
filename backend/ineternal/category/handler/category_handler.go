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

package handler

import (
	"github.com/chenmingyong0423/fnote/backend/ineternal/category/service"
	"github.com/chenmingyong0423/fnote/backend/ineternal/pkg/api"
	"github.com/chenmingyong0423/fnote/backend/ineternal/pkg/domain"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func NewCategoryHandler(engine *gin.Engine, serv service.ICategoryService) *CategoryHandler {
	ch := &CategoryHandler{
		serv: serv,
	}
	engine.GET("/categories", ch.GetCategoriesAndTags)
	engine.GET("/category/:name/tags", ch.GetTagsByName)
	engine.GET("/menus", ch.GetMenus)
	return ch
}

type CategoryHandler struct {
	serv service.ICategoryService
}

func (h *CategoryHandler) GetCategoriesAndTags(ctx *gin.Context) {
	listVO, err := h.serv.GetCategoriesAndTags(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "category", err.Error())
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, api.SuccessResponseWithData[api.ListVO[domain.SearchCategoryVO]](listVO))
}

func (h *CategoryHandler) GetMenus(ctx *gin.Context) {
	listVO, err := h.serv.GetMenus(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "menu", err.Error())
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, api.SuccessResponseWithData[api.ListVO[domain.MenuVO]](listVO))
}

func (h *CategoryHandler) GetTagsByName(ctx *gin.Context) {
	name := ctx.Param("name")
	listVO, err := h.serv.GetTagsByName(ctx, name)
	if err != nil {
		slog.ErrorContext(ctx, "tag", err.Error())
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, api.SuccessResponseWithData[api.ListVO[string]](listVO))
}
