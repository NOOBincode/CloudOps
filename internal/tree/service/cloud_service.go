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

package service

import (
	"context"

	"github.com/GoSimplicity/AI-CloudOps/internal/model"
	"github.com/GoSimplicity/AI-CloudOps/internal/tree/dao"
	"go.uber.org/zap"
)

type CloudService interface {
	ListCloudProviders(ctx context.Context) ([]model.CloudProviderResp, error)
	ListRegions(ctx context.Context, req *model.ListRegionsReq) ([]model.RegionResp, error)
	ListZones(ctx context.Context, req *model.ListZonesReq) ([]model.ZoneResp, error)
	ListInstanceTypes(ctx context.Context, req *model.ListInstanceTypesReq) ([]model.InstanceTypeResp, error)
	ListImages(ctx context.Context, req *model.ListImagesReq) ([]model.ImageResp, error)
	ListVpcs(ctx context.Context, req *model.ListVpcsReq) ([]model.ResourceVpc, error)
	ListSecurityGroups(ctx context.Context, req *model.ListSecurityGroupsReq) ([]model.SecurityGroupResp, error)
	GetTreeStatistics(ctx context.Context) (*model.TreeStatisticsResp, error)
}

type cloudService struct {
	logger *zap.Logger
	dao    dao.CloudDAO
}

// GetTreeStatistics implements CloudService.
func (c *cloudService) GetTreeStatistics(ctx context.Context) (*model.TreeStatisticsResp, error) {
	panic("unimplemented")
}

// ListCloudProviders implements CloudService.
func (c *cloudService) ListCloudProviders(ctx context.Context) ([]model.CloudProviderResp, error) {
	panic("unimplemented")
}

// ListImages implements CloudService.
func (c *cloudService) ListImages(ctx context.Context, req *model.ListImagesReq) ([]model.ImageResp, error) {
	panic("unimplemented")
}

// ListInstanceTypes implements CloudService.
func (c *cloudService) ListInstanceTypes(ctx context.Context, req *model.ListInstanceTypesReq) ([]model.InstanceTypeResp, error) {
	panic("unimplemented")
}

// ListRegions implements CloudService.
func (c *cloudService) ListRegions(ctx context.Context, req *model.ListRegionsReq) ([]model.RegionResp, error) {
	panic("unimplemented")
}

// ListSecurityGroups implements CloudService.
func (c *cloudService) ListSecurityGroups(ctx context.Context, req *model.ListSecurityGroupsReq) ([]model.SecurityGroupResp, error) {
	panic("unimplemented")
}

// ListVpcs implements CloudService.
func (c *cloudService) ListVpcs(ctx context.Context, req *model.ListVpcsReq) ([]model.ResourceVpc, error) {
	panic("unimplemented")
}

// ListZones implements CloudService.
func (c *cloudService) ListZones(ctx context.Context, req *model.ListZonesReq) ([]model.ZoneResp, error) {
	panic("unimplemented")
}

func NewCloudService(logger *zap.Logger, dao dao.CloudDAO) CloudService {
	return &cloudService{
		logger: logger,
		dao:    dao,
	}
}
