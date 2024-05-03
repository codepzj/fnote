// Copyright 2024 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"github.com/chenmingyong0423/fnote/server/internal/comment"
	csServ "github.com/chenmingyong0423/fnote/server/internal/count_stats/service"
	"github.com/chenmingyong0423/fnote/server/internal/pkg/domain"
	apiwrap "github.com/chenmingyong0423/fnote/server/internal/pkg/web/wrap"
	"github.com/chenmingyong0423/fnote/server/internal/post_like"
	"github.com/chenmingyong0423/fnote/server/internal/visit_log/service"
	"github.com/gin-gonic/gin"
)

func NewDataAnalysisHandler(vlServ service.IVisitLogService, csServ csServ.ICountStatsService, postLikeServ post_like.Service, commentServ comment.Service) *DataAnalysisHandler {
	return &DataAnalysisHandler{
		vlServ:       vlServ,
		csServ:       csServ,
		postLikeServ: postLikeServ,
		commentServ:  commentServ,
	}
}

type DataAnalysisHandler struct {
	vlServ       service.IVisitLogService
	csServ       csServ.ICountStatsService
	postLikeServ post_like.Service
	commentServ  comment.Service
}

func (h *DataAnalysisHandler) RegisterGinRoutes(engine *gin.Engine) {
	routerGroup := engine.Group("/admin-api/data-analysis")
	routerGroup.GET("/traffic/today", apiwrap.Wrap(h.GetTodayTrafficStats))
	routerGroup.GET("/traffic", apiwrap.Wrap(h.GetWebsiteCountStats))
	routerGroup.GET("/content", apiwrap.Wrap(h.GetWebsiteContentStats))
	routerGroup.GET("/tendency", apiwrap.Wrap(h.GetTendencyStats))
}

func (h *DataAnalysisHandler) GetTodayTrafficStats(ctx *gin.Context) (*apiwrap.ResponseBody[TodayTrafficStatsVO], error) {
	// 查询当日访问量
	todayViewCount, err := h.vlServ.GetTodayViewCount(ctx)
	if err != nil {
		return nil, err
	}
	// 查询当日实际访问用户量
	userViewCount, err := h.vlServ.GetTodayUserViewCount(ctx)
	if err != nil {
		return nil, err
	}

	commentCount, err := h.commentServ.FindCommentCountOfToday(ctx)
	if err != nil {
		return nil, err
	}

	likeCount, err := h.postLikeServ.FindLikeCountToday(ctx)
	if err != nil {
		return nil, err
	}

	return apiwrap.SuccessResponseWithData(TodayTrafficStatsVO{
		ViewCount:     todayViewCount,
		UserViewCount: userViewCount,
		CommentCount:  commentCount,
		LikeCount:     likeCount,
	}), nil
}

func (h *DataAnalysisHandler) GetWebsiteCountStats(ctx *gin.Context) (*apiwrap.ResponseBody[TrafficStatsVO], error) {
	// 查询网站统计
	websiteCountStats, err := h.csServ.GetWebsiteCountStats(ctx)
	if err != nil {
		return nil, err
	}
	return apiwrap.SuccessResponseWithData(TrafficStatsVO{
		ViewCount:    websiteCountStats.WebsiteViewCount,
		CommentCount: websiteCountStats.CommentCount,
		LikeCount:    websiteCountStats.LikeCount,
	}), nil
}

func (h *DataAnalysisHandler) GetWebsiteContentStats(ctx *gin.Context) (*apiwrap.ResponseBody[ContentStatsVO], error) {
	// 查询网站统计
	websiteCountStats, err := h.csServ.GetWebsiteCountStats(ctx)
	if err != nil {
		return nil, err
	}
	return apiwrap.SuccessResponseWithData(ContentStatsVO{
		PostCount:     websiteCountStats.PostCount,
		CategoryCount: websiteCountStats.CategoryCount,
		TagCount:      websiteCountStats.TagCount,
	}), nil
}

func (h *DataAnalysisHandler) GetTendencyStats(ctx *gin.Context) (*apiwrap.ResponseBody[TendencyDataVO], error) {
	var (
		period = ctx.Query("period")
		days   = 7
	)
	if period == "month" {
		days = 30
	}
	tendencyData4PV, err := h.vlServ.GetViewTendencyStats4PV(ctx, days)
	if err != nil {
		return nil, err
	}
	tendencyData4UV, err := h.vlServ.GetViewTendencyStats4UV(ctx, days)
	if err != nil {
		return nil, err
	}

	return apiwrap.SuccessResponseWithData(TendencyDataVO{
		PV: h.tdToVO(tendencyData4PV),
		UV: h.tdToVO(tendencyData4UV),
	}), nil
}

func (h *DataAnalysisHandler) tdToVO(data []domain.TendencyData) []TendencyData {
	voList := make([]TendencyData, 0, len(data))
	for _, td := range data {
		voList = append(voList, TendencyData{
			Timestamp: td.Timestamp,
			ViewCount: td.ViewCount,
		})
	}
	return voList
}
