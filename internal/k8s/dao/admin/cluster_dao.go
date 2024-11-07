package admin

import (
	"context"
	"github.com/GoSimplicity/AI-CloudOps/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ClusterDAO interface {
	ListAllClusters(ctx context.Context) ([]*model.K8sCluster, error)
	CreateCluster(ctx context.Context, cluster *model.K8sCluster) error
	UpdateCluster(ctx context.Context, cluster *model.K8sCluster) error
	DeleteCluster(ctx context.Context, id int) error
	GetClusterByID(ctx context.Context, id int) (*model.K8sCluster, error)
	GetClusterByName(ctx context.Context, name string) (*model.K8sCluster, error)
	BatchDeleteClusters(ctx context.Context, ids []int) error
}

type clusterDAO struct {
	db *gorm.DB
	l  *zap.Logger
}

func NewClusterDAO(db *gorm.DB, l *zap.Logger) ClusterDAO {
	return &clusterDAO{
		db: db,
		l:  l,
	}
}

// ListAllClusters 查询所有集群
func (c *clusterDAO) ListAllClusters(ctx context.Context) ([]*model.K8sCluster, error) {
	var clusters []*model.K8sCluster

	if err := c.db.WithContext(ctx).Find(&clusters).Error; err != nil {
		c.l.Error("ListAllClusters 查询所有集群失败", zap.Error(err))
		return nil, err
	}

	return clusters, nil
}

// CreateCluster 创建新集群
func (c *clusterDAO) CreateCluster(ctx context.Context, cluster *model.K8sCluster) error {
	if err := c.db.WithContext(ctx).Create(cluster).Error; err != nil {
		c.l.Error("CreateCluster 创建集群失败", zap.Error(err))
		return err
	}

	return nil
}

// UpdateCluster 更新集群信息
func (c *clusterDAO) UpdateCluster(ctx context.Context, cluster *model.K8sCluster) error {
	tx := c.db.WithContext(ctx).Begin()

	if err := tx.Model(cluster).Where("id = ?", cluster.ID).Updates(cluster).Error; err != nil {
		tx.Rollback()
		c.l.Error("UpdateCluster 更新集群失败", zap.Int("id", cluster.ID), zap.Error(err))
		return err
	}

	tx.Commit()
	return nil
}

// DeleteCluster 删除指定集群
func (c *clusterDAO) DeleteCluster(ctx context.Context, id int) error {
	if err := c.db.WithContext(ctx).Where("id = ?", id).Delete(&model.K8sCluster{}).Error; err != nil {
		c.l.Error("DeleteCluster 删除集群失败", zap.Int("id", id), zap.Error(err))
		return err
	}

	return nil
}

// GetClusterByID 根据ID查询集群
func (c *clusterDAO) GetClusterByID(ctx context.Context, id int) (*model.K8sCluster, error) {
	var cluster model.K8sCluster

	if err := c.db.WithContext(ctx).Where("id = ?", id).First(&cluster).Error; err != nil {
		c.l.Error("GetClusterByID 查询集群失败", zap.Int("id", id), zap.Error(err))
		return nil, err
	}

	return &cluster, nil
}

// GetClusterByName 根据名称查询集群
func (c *clusterDAO) GetClusterByName(ctx context.Context, name string) (*model.K8sCluster, error) {
	var cluster model.K8sCluster

	if err := c.db.WithContext(ctx).Where("name = ?", name).First(&cluster).Error; err != nil {
		c.l.Error("GetClusterByName 查询集群失败", zap.String("name", name), zap.Error(err))
		return nil, err
	}

	return &cluster, nil
}

// BatchDeleteClusters 批量删除集群
func (c *clusterDAO) BatchDeleteClusters(ctx context.Context, ids []int) error {
	if err := c.db.WithContext(ctx).Where("id IN ?", ids).Delete(&model.K8sCluster{}).Error; err != nil {
		c.l.Error("BatchDeleteClusters 批处理删除集群失败", zap.Error(err))
		return err
	}

	return nil
}
