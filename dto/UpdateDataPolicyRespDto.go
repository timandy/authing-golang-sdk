package dto


type UpdateDataPolicyRespDto struct{
    PolicyId  string `json:"policyId"`
    PolicyName  string `json:"policyName"`
    Description  string `json:"description,omitempty"`
    CreatedAt  string `json:"createdAt"`
    UpdatedAt  string `json:"updatedAt"`
}

