package entity

type ClusterQueryParams struct {
	ID                  uint64 `gorm:"column:id"`
	AddonID             uint64 `gorm:"column:addon_id"`
	AddonClusterVersion string `gorm:"column:addoncluster_version"`
	TopoName            string `gorm:"column:topo_name" json:"topoName"`
	K8sClusterConfigID  uint64 `gorm:"column:k8s_cluster_config_id" json:"k8sClusterConfigId"`
	ClusterName         string `gorm:"column:cluster_name" json:"clusterName"`
	ClusterAlias        string `gorm:"column:cluster_alias" json:"clusterAlias"`
	Namespace           string `gorm:"column:namespace" json:"namespace"`
	BkBizID             uint64 `gorm:"column:bk_biz_id" json:"bkBizId"`
	BkBizName           string `gorm:"column:bk_biz_name" json:"bkBizName"`
	BkAppAbbr           string `gorm:"column:bk_app_abbr" json:"bkAppAbbr"`
	BkAppCode           string `gorm:"column:bk_app_code" json:"bkAppCode"`
}
