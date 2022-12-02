package managerparam

import (
	"gorm.io/gorm"

	accesstokenmanager "g.hz.netease.com/horizon/pkg/accesstoken/manager"
	applicationmanager "g.hz.netease.com/horizon/pkg/application/manager"
	applicationregionmanager "g.hz.netease.com/horizon/pkg/applicationregion/manager"
	clustermanager "g.hz.netease.com/horizon/pkg/cluster/manager"
	envmanager "g.hz.netease.com/horizon/pkg/environment/manager"
	environmentregionmanager "g.hz.netease.com/horizon/pkg/environmentregion/manager"
	eventManager "g.hz.netease.com/horizon/pkg/event/manager"
	groupmanager "g.hz.netease.com/horizon/pkg/group/manager"
	idpmanager "g.hz.netease.com/horizon/pkg/idp/manager"
	membermanager "g.hz.netease.com/horizon/pkg/member"
	prmanager "g.hz.netease.com/horizon/pkg/pipelinerun/manager"
	pipelinemanager "g.hz.netease.com/horizon/pkg/pipelinerun/pipeline/manager"
	regionmanager "g.hz.netease.com/horizon/pkg/region/manager"
	registrymanager "g.hz.netease.com/horizon/pkg/registry/manager"
	tagmanager "g.hz.netease.com/horizon/pkg/tag/manager"
	templatemanager "g.hz.netease.com/horizon/pkg/template/manager"
	trmanager "g.hz.netease.com/horizon/pkg/templaterelease/manager"
	templateschematagmanager "g.hz.netease.com/horizon/pkg/templateschematag/manager"
	trtmanager "g.hz.netease.com/horizon/pkg/templateschematag/manager"
	usermanager "g.hz.netease.com/horizon/pkg/user/manager"
	linkmanager "g.hz.netease.com/horizon/pkg/userlink/manager"
	webhookManager "g.hz.netease.com/horizon/pkg/webhook/manager"
)

type Manager struct {
	UserManager              usermanager.Manager
	UserLinksManager         linkmanager.Manager
	ApplicationManager       applicationmanager.Manager
	TemplateReleaseManager   trmanager.Manager
	TemplateSchemaTagManager trtmanager.Manager
	ClusterMgr               clustermanager.Manager
	MemberManager            membermanager.Manager
	ClusterSchemaTagMgr      templateschematagmanager.Manager
	ApplicationRegionManager applicationregionmanager.Manager
	EnvironmentRegionMgr     environmentregionmanager.Manager
	TagManager               tagmanager.Manager
	TemplateMgr              templatemanager.Manager
	EnvRegionMgr             environmentregionmanager.Manager
	RegionMgr                regionmanager.Manager
	PipelinerunMgr           prmanager.Manager
	PipelineMgr              pipelinemanager.Manager
	EnvMgr                   envmanager.Manager
	GroupManager             groupmanager.Manager
	RegistryManager          registrymanager.Manager
	IdpManager               idpmanager.Manager
	AccessTokenManager       accesstokenmanager.Manager
	WebhookManager           webhookManager.Manager
	EventManager             eventManager.Manager
}

func InitManager(db *gorm.DB) *Manager {
	return &Manager{
		UserManager:              usermanager.New(db),
		UserLinksManager:         linkmanager.New(db),
		ApplicationManager:       applicationmanager.New(db),
		TemplateReleaseManager:   trmanager.New(db),
		TemplateSchemaTagManager: trtmanager.New(db),
		ClusterMgr:               clustermanager.New(db),
		MemberManager:            membermanager.New(db),
		ClusterSchemaTagMgr:      templateschematagmanager.New(db),
		ApplicationRegionManager: applicationregionmanager.New(db),
		EnvironmentRegionMgr:     environmentregionmanager.New(db),
		TagManager:               tagmanager.New(db),
		TemplateMgr:              templatemanager.New(db),
		EnvRegionMgr:             environmentregionmanager.New(db),
		RegionMgr:                regionmanager.New(db),
		PipelinerunMgr:           prmanager.New(db),
		PipelineMgr:              pipelinemanager.New(db),
		EnvMgr:                   envmanager.New(db),
		GroupManager:             groupmanager.New(db),
		RegistryManager:          registrymanager.New(db),
		IdpManager:               idpmanager.NewManager(db),
		AccessTokenManager:       accesstokenmanager.New(db),
		WebhookManager:           webhookManager.New(db),
		EventManager:             eventManager.New(db),
	}
}
