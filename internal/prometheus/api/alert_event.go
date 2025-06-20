/*
 * MIT License
 *
 * Copyright (c) 2024 Bamboo
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 */

package api

import (
	"github.com/GoSimplicity/AI-CloudOps/pkg/utils"

	"github.com/GoSimplicity/AI-CloudOps/internal/model"
	alertEventService "github.com/GoSimplicity/AI-CloudOps/internal/prometheus/service/alert"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AlertEventHandler struct {
	alertEventService alertEventService.AlertManagerEventService
	l                 *zap.Logger
}

func NewAlertEventHandler(l *zap.Logger, alertEventService alertEventService.AlertManagerEventService) *AlertEventHandler {
	return &AlertEventHandler{
		l:                 l,
		alertEventService: alertEventService,
	}
}

func (a *AlertEventHandler) RegisterRouters(server *gin.Engine) {
	monitorGroup := server.Group("/api/monitor")

	alertEvents := monitorGroup.Group("/alert_events")
	{
		alertEvents.GET("/list", a.GetMonitorAlertEventList)
		alertEvents.POST("/:id/silence", a.EventAlertSilence)
		alertEvents.POST("/:id/claim", a.EventAlertClaim)
		alertEvents.POST("/:id/unSilence", a.EventAlertUnSilence)
		alertEvents.POST("/silence", a.BatchEventAlertSilence)
		alertEvents.GET("/total", a.GetMonitorAlertEventTotal)
	}
}

// GetMonitorAlertEventList 获取告警事件列表
func (a *AlertEventHandler) GetMonitorAlertEventList(ctx *gin.Context) {
	var listReq model.ListReq

	utils.HandleRequest(ctx, &listReq, func() (interface{}, error) {
		return a.alertEventService.GetMonitorAlertEventList(ctx, &listReq)
	})
}

// EventAlertSilence 将指定告警事件设置为静默状态
func (a *AlertEventHandler) EventAlertSilence(ctx *gin.Context) {
	var req model.AlertEventSilenceRequest

	uc := ctx.MustGet("user").(utils.UserClaims)
	id, err := utils.GetParamID(ctx)
	if err != nil {
		utils.ErrorWithMessage(ctx, "参数错误")
		return
	}

	req.ID = id
	req.UserID = uc.Uid

	utils.HandleRequest(ctx, &req, func() (interface{}, error) {
		return nil, a.alertEventService.EventAlertSilence(ctx, &req)
	})
}

// EventAlertClaim 认领指定的告警事件
func (a *AlertEventHandler) EventAlertClaim(ctx *gin.Context) {
	var req model.AlertEventClaimRequest

	uc := ctx.MustGet("user").(utils.UserClaims)
	id, err := utils.GetParamID(ctx)
	if err != nil {
		utils.ErrorWithMessage(ctx, "参数错误")
		return
	}

	req.ID = id
	req.UserID = uc.Uid

	utils.HandleRequest(ctx, &req, func() (interface{}, error) {
		return nil, a.alertEventService.EventAlertClaim(ctx, &req)
	})
}

// EventAlertUnSilence 取消指定告警事件的静默状态
func (a *AlertEventHandler) EventAlertUnSilence(ctx *gin.Context) {
	var req model.AlertEventUnSilenceRequest

	uc := ctx.MustGet("user").(utils.UserClaims)
	id, err := utils.GetParamID(ctx)
	if err != nil {
		utils.ErrorWithMessage(ctx, "参数错误")
		return
	}

	req.ID = id
	req.UserID = uc.Uid

	utils.HandleRequest(ctx, &req, func() (interface{}, error) {
		return nil, a.alertEventService.EventAlertUnSilence(ctx, &req)
	})
}

// BatchEventAlertSilence 批量设置告警事件为静默状态
func (a *AlertEventHandler) BatchEventAlertSilence(ctx *gin.Context) {
	var req model.BatchEventAlertSilenceRequest

	uc := ctx.MustGet("user").(utils.UserClaims)
	req.UserID = uc.Uid

	utils.HandleRequest(ctx, &req, func() (interface{}, error) {
		return nil, a.alertEventService.BatchEventAlertSilence(ctx, &req)
	})
}

// GetMonitorAlertEventTotal 获取监控告警事件总数
func (a *AlertEventHandler) GetMonitorAlertEventTotal(ctx *gin.Context) {
	utils.HandleRequest(ctx, nil, func() (interface{}, error) {
		return a.alertEventService.GetMonitorAlertEventTotal(ctx)
	})
}
