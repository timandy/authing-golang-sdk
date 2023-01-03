package dto

type CheckUserSameLevelPermissionDto struct {
	Resource              string        `json:"resource"`
	Action                string        `json:"action"`
	UserId                string        `json:"userId"`
	NamespaceCode         string        `json:"namespaceCode"`
	ResourceNodeCodes     []string      `json:"resourceNodeCodes,omitempty"`
	JudgeConditionEnabled bool          `json:"judgeConditionEnabled,omitempty"`
	AuthEnvParams         AuthEnvParams `json:"authEnvParams,omitempty"`
}
