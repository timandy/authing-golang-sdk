package client

import (
	"authing-go-sdk/dto"
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

/*
 * @summary 获取 Management API Token
 * @description 获取 Management API Token
 * @param requestBody
 * @returns GetManagementTokenRespDto
 */
func (c *ManagementClient) GetManagementToken(reqDto *dto.GetManagementAccessTokenDto) *dto.GetManagementTokenRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-management-token", fasthttp.MethodPost, reqDto)
	var response dto.GetManagementTokenRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户信息
 * @description 通过 id、username、email、phone、email、externalId 获取用户详情
 * @param userId 用户 ID
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @param phone 手机号
 * @param email 邮箱
 * @param username 用户名
 * @param externalId 原系统 ID
 * @returns UserSingleRespDto
 */
func (c *ManagementClient) GetUser(reqDto *dto.GetUserDto) *dto.UserSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user", fasthttp.MethodGet, reqDto)
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量获取用户信息
 * @description 根据用户 id 批量获取用户信息
 * @param userIds 用户 ID 数组
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @returns UserListRespDto
 */
func (c *ManagementClient) GetUserBatch(reqDto *dto.GetUserBatchDto) *dto.UserListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-batch", fasthttp.MethodGet, reqDto)
	var response dto.UserListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户列表
 * @description 获取用户列表接口，支持分页
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @returns UserPaginatedRespDto
 */
func (c *ManagementClient) ListUsers(reqDto *dto.ListUsersDto) *dto.UserPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-users", fasthttp.MethodGet, reqDto)
	var response dto.UserPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户的外部身份源
 * @description 获取用户的外部身份源
 * @param userId 用户 ID
 * @returns IdentityListRespDto
 */
func (c *ManagementClient) GetUserIdentities(reqDto *dto.GetUserIdentitiesDto) *dto.IdentityListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-identities", fasthttp.MethodGet, reqDto)
	var response dto.IdentityListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户角色列表
 * @description 获取用户角色列表
 * @param userId 用户 ID
 * @param namespace 所属权限分组的 code
 * @returns RolePaginatedRespDto
 */
func (c *ManagementClient) GetUserRoles(reqDto *dto.GetUserRolesDto) *dto.RolePaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-roles", fasthttp.MethodGet, reqDto)
	var response dto.RolePaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户实名认证信息
 * @description 获取用户实名认证信息
 * @param userId 用户 ID
 * @returns PrincipalAuthenticationInfoPaginatedRespDto
 */
func (c *ManagementClient) GetUserPrincipalAuthenticationInfo(reqDto *dto.GetUserPrincipalAuthenticationInfoDto) *dto.PrincipalAuthenticationInfoPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-principal-authentication-info", fasthttp.MethodGet, reqDto)
	var response dto.PrincipalAuthenticationInfoPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 删除用户实名认证信息
 * @description 删除用户实名认证信息
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) ResetUserPrincipalAuthenticationInfo(reqDto *dto.ResetUserPrincipalAuthenticationInfoDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/reset-user-principal-authentication-info", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户部门列表
 * @description 获取用户部门列表
 * @param userId 用户 ID
 * @returns UserDepartmentPaginatedRespDto
 */
func (c *ManagementClient) GetUserDepartments(reqDto *dto.GetUserDepartmentsDto) *dto.UserDepartmentPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-departments", fasthttp.MethodGet, reqDto)
	var response dto.UserDepartmentPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 设置用户所在部门
 * @description 设置用户所在部门
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) SetUserDepartment(reqDto *dto.SetUserDepartmentsDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/set-user-departments", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户分组列表
 * @description 获取用户分组列表
 * @param userId 用户 ID
 * @returns GroupPaginatedRespDto
 */
func (c *ManagementClient) GetUserGroups(reqDto *dto.GetUserGroupsDto) *dto.GroupPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-groups", fasthttp.MethodGet, reqDto)
	var response dto.GroupPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 删除用户
 * @description 删除用户（支持批量删除）
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) DeleteUsersBatch(reqDto *dto.DeleteUsersBatchDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-users-batch", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户 MFA 绑定信息
 * @description 获取用户 MFA 绑定信息
 * @param userId 用户 ID
 * @returns UserMfaSingleRespDto
 */
func (c *ManagementClient) GetUserMfaInfo(reqDto *dto.GetUserMfaInfoDto) *dto.UserMfaSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-mfa-info", fasthttp.MethodGet, reqDto)
	var response dto.UserMfaSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取已归档的用户列表
 * @description 获取已归档的用户列表
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns ListArchivedUsersSingleRespDto
 */
func (c *ManagementClient) ListArchivedUsers(reqDto *dto.ListArchivedUsersDto) *dto.ListArchivedUsersSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-archived-users", fasthttp.MethodGet, reqDto)
	var response dto.ListArchivedUsersSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 强制下线用户
 * @description 强制下线用户
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) KickUsers(reqDto *dto.KickUsersDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/kick-users", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 判断用户是否存在
 * @description 根据条件判断用户是否存在
 * @param requestBody
 * @returns IsUserExistsRespDto
 */
func (c *ManagementClient) IsUserExists(reqDto *dto.IsUserExistsReqDto) *dto.IsUserExistsRespDto {
	b, err := c.SendHttpRequest("/api/v3/is-user-exists", fasthttp.MethodPost, reqDto)
	var response dto.IsUserExistsRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建用户
 * @description 创建用户，邮箱、手机号、用户名必须包含其中一个
 * @param requestBody
 * @returns UserSingleRespDto
 */
func (c *ManagementClient) CreateUser(reqDto *dto.CreateUserReqDto) *dto.UserSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-user", fasthttp.MethodPost, reqDto)
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量创建用户
 * @description 此接口将以管理员身份批量创建用户，不需要进行手机号验证码检验等安全检测。用户的手机号、邮箱、用户名、externalId 用户池内唯一。
 * @param requestBody
 * @returns UserListRespDto
 */
func (c *ManagementClient) CreateUserBatch(reqDto *dto.CreateUserBatchReqDto) *dto.UserListRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-users-batch", fasthttp.MethodPost, reqDto)
	var response dto.UserListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 修改用户资料
 * @description 修改用户资料
 * @param requestBody
 * @returns UserSingleRespDto
 */
func (c *ManagementClient) UpdateUser(reqDto *dto.UpdateUserReqDto) *dto.UserSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-user", fasthttp.MethodPost, reqDto)
	var response dto.UserSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户可访问应用
 * @description 获取用户可访问应用
 * @param userId 用户 ID
 * @returns AppListRespDto
 */
func (c *ManagementClient) GetUserAccessibleApps(reqDto *dto.GetUserAccessibleAppsDto) *dto.AppListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-accessible-apps", fasthttp.MethodGet, reqDto)
	var response dto.AppListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户授权的应用
 * @description 获取用户授权的应用
 * @param userId 用户 ID
 * @returns AppListRespDto
 */
func (c *ManagementClient) GetUserAuthorizedApps(reqDto *dto.GetUserAuthorizedAppsDto) *dto.AppListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-authorized-apps", fasthttp.MethodGet, reqDto)
	var response dto.AppListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 判断用户是否有某个角色
 * @description 判断用户是否有某个角色，支持同时传入多个角色进行判断
 * @param requestBody
 * @returns HasAnyRoleRespDto
 */
func (c *ManagementClient) HasAnyRole(reqDto *dto.HasAnyRoleReqDto) *dto.HasAnyRoleRespDto {
	b, err := c.SendHttpRequest("/api/v3/has-any-role", fasthttp.MethodPost, reqDto)
	var response dto.HasAnyRoleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户的登录历史记录
 * @description 获取用户登录历史记录
 * @param userId 用户 ID
 * @param appId 应用 ID
 * @param clientIp 客户端 IP
 * @param start 开始时间戳（毫秒）
 * @param end 结束时间戳（毫秒）
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns UserLoginHistoryPaginatedRespDto
 */
func (c *ManagementClient) GetUserLoginHistory(reqDto *dto.GetUserLoginHistoryDto) *dto.UserLoginHistoryPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-login-history", fasthttp.MethodGet, reqDto)
	var response dto.UserLoginHistoryPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户曾经登录过的应用
 * @description 获取用户曾经登录过的应用
 * @param userId 用户 ID
 * @returns UserLoggedInAppsListRespDto
 */
func (c *ManagementClient) GetUserLoggedInApps(reqDto *dto.GetUserLoggedinAppsDto) *dto.UserLoggedInAppsListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-loggedin-apps", fasthttp.MethodGet, reqDto)
	var response dto.UserLoggedInAppsListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户被授权的所有资源
 * @description 获取用户被授权的所有资源，用户被授权的资源是用户自身被授予、通过分组继承、通过角色继承、通过组织机构继承的集合
 * @param userId 用户 ID
 * @param namespace 所属权限分组的 code
 * @param resourceType 资源类型
 * @returns AuthorizedResourcePaginatedRespDto
 */
func (c *ManagementClient) GetUserAuthorizedResources(reqDto *dto.GetUserAuthorizedResourcesDto) *dto.AuthorizedResourcePaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-user-authorized-resources", fasthttp.MethodGet, reqDto)
	var response dto.AuthorizedResourcePaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取分组详情
 * @description 获取分组详情，通过 code 唯一标志用户池中的一个分组
 * @param code 分组 code
 * @returns GroupSingleRespDto
 */
func (c *ManagementClient) GetGroup(reqDto *dto.GetGroupDto) *dto.GroupSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-group", fasthttp.MethodGet, reqDto)
	var response dto.GroupSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取分组列表
 * @description 获取分组列表接口，支持分页
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns GroupPaginatedRespDto
 */
func (c *ManagementClient) ListGroups(reqDto *dto.ListGroupsDto) *dto.GroupPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-groups", fasthttp.MethodGet, reqDto)
	var response dto.GroupPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建分组
 * @description 创建分组，一个分组必须包含一个用户池全局唯一的标志符（code），此标志符必须为一个合法的英文标志符，如 developers；以及分组名称
 * @param requestBody
 * @returns GroupSingleRespDto
 */
func (c *ManagementClient) CreateGroup(reqDto *dto.CreateGroupReqDto) *dto.GroupSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-group", fasthttp.MethodPost, reqDto)
	var response dto.GroupSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量创建分组
 * @description 批量创建分组
 * @param requestBody
 * @returns GroupListRespDto
 */
func (c *ManagementClient) CreateGroupsBatch(reqDto *dto.CreateGroupBatchReqDto) *dto.GroupListRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-groups-batch", fasthttp.MethodPost, reqDto)
	var response dto.GroupListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 修改分组
 * @description 修改分组，通过 code 唯一标志用户池中的一个分组。你可以修改此分组的 code
 * @param requestBody
 * @returns GroupSingleRespDto
 */
func (c *ManagementClient) UpdateGroup(reqDto *dto.UpdateGroupReqDto) *dto.GroupSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-group", fasthttp.MethodPost, reqDto)
	var response dto.GroupSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量删除分组
 * @description 批量删除分组
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) DeleteGroupsBatch(reqDto *dto.DeleteGroupsReqDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-groups-batch", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 添加分组成员
 * @description 添加分组成员
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) AddGroupMembers(reqDto *dto.AddGroupMembersReqDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/add-group-members", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量移除分组成员
 * @description 批量移除分组成员
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) RemoveGroupMembers(reqDto *dto.RemoveGroupMembersReqDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/remove-group-members", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取分组成员列表
 * @description 获取分组成员列表
 * @param code 分组 code
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @returns UserPaginatedRespDto
 */
func (c *ManagementClient) ListGroupMembers(reqDto *dto.ListGroupMembersDto) *dto.UserPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-group-members", fasthttp.MethodGet, reqDto)
	var response dto.UserPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取分组被授权的资源列表
 * @description 获取分组被授权的资源列表
 * @param code 分组 code
 * @param namespace 所属权限分组的 code
 * @param resourceType 资源类型
 * @returns AuthorizedResourceListRespDto
 */
func (c *ManagementClient) GetGroupAuthorizedResources(reqDto *dto.GetGroupAuthorizedResourcesDto) *dto.AuthorizedResourceListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-group-authorized-resources", fasthttp.MethodGet, reqDto)
	var response dto.AuthorizedResourceListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取角色详情
 * @description 获取角色详情
 * @param code 权限分组内角色的唯一标识符
 * @param namespace 所属权限分组的 code
 * @returns RoleSingleRespDto
 */
func (c *ManagementClient) GetRole(reqDto *dto.GetRoleDto) *dto.RoleSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-role", fasthttp.MethodGet, reqDto)
	var response dto.RoleSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 分配角色
 * @description 分配角色，被分配者可以是用户，可以是部门
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) AssignRole(reqDto *dto.AssignRoleDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/assign-role", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 移除分配的角色
 * @description 移除分配的角色，被分配者可以是用户，可以是部门
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) RevokeRole(reqDto *dto.RevokeRoleDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/revoke-role", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 角色被授权的资源列表
 * @description 角色被授权的资源列表
 * @param code 权限分组内角色的唯一标识符
 * @param namespace 所属权限分组的 code
 * @param resourceType 资源类型
 * @returns RoleAuthorizedResourcePaginatedRespDto
 */
func (c *ManagementClient) GetRoleAuthorizedResources(reqDto *dto.GetRoleAuthorizedResourcesDto) *dto.RoleAuthorizedResourcePaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-role-authorized-resources", fasthttp.MethodGet, reqDto)
	var response dto.RoleAuthorizedResourcePaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取角色成员列表
 * @description 获取角色成员列表
 * @param code 权限分组内角色的唯一标识符
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @param namespace 所属权限分组的 code
 * @returns UserPaginatedRespDto
 */
func (c *ManagementClient) ListRoleMembers(reqDto *dto.ListRoleMembersDto) *dto.UserPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-role-members", fasthttp.MethodGet, reqDto)
	var response dto.UserPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取角色的部门列表
 * @description 获取角色的部门列表
 * @param code 权限分组内角色的唯一标识符
 * @param namespace 所属权限分组的 code
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns RoleDepartmentListPaginatedRespDto
 */
func (c *ManagementClient) ListRoleDepartments(reqDto *dto.ListRoleDepartmentsDto) *dto.RoleDepartmentListPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-role-departments", fasthttp.MethodGet, reqDto)
	var response dto.RoleDepartmentListPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建角色
 * @description 创建角色，可以指定不同的权限分组
 * @param requestBody
 * @returns RoleSingleRespDto
 */
func (c *ManagementClient) CreateRole(reqDto *dto.CreateRoleDto) *dto.RoleSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-role", fasthttp.MethodPost, reqDto)
	var response dto.RoleSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取角色列表
 * @description 获取角色列表
 * @param namespace 所属权限分组的 code
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns RolePaginatedRespDto
 */
func (c *ManagementClient) ListRoles(reqDto *dto.ListRolesDto) *dto.RolePaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-roles", fasthttp.MethodGet, reqDto)
	var response dto.RolePaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary （批量）删除角色
 * @description 删除角色
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) DeleteRolesBatch(reqDto *dto.DeleteRoleDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-roles-batch", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量创建角色
 * @description 批量创建角色
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) CreateRolesBatch(reqDto *dto.CreateRolesBatch) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-roles-batch", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 修改角色
 * @description 修改角色
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) UpdateRole(reqDto *dto.UpdateRoleDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-role", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取顶层组织机构列表
 * @description 获取顶层组织机构列表
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns OrganizationPaginatedRespDto
 */
func (c *ManagementClient) ListOrganizations(reqDto *dto.ListOrganizationsDto) *dto.OrganizationPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-organizations", fasthttp.MethodGet, reqDto)
	var response dto.OrganizationPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建顶层组织机构
 * @description 创建组织机构，会创建一个只有一个节点的组织机构
 * @param requestBody
 * @returns OrganizationSingleRespDto
 */
func (c *ManagementClient) CreateOrganization(reqDto *dto.CreateOrganizationReqDto) *dto.OrganizationSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-organization", fasthttp.MethodPost, reqDto)
	var response dto.OrganizationSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 修改顶层组织机构
 * @description 修改顶层组织机构
 * @param requestBody
 * @returns OrganizationSingleRespDto
 */
func (c *ManagementClient) UpdateOrganization(reqDto *dto.UpdateOrganizationReqDto) *dto.OrganizationSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-organization", fasthttp.MethodPost, reqDto)
	var response dto.OrganizationSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 删除组织机构
 * @description 删除组织机构树
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) DeleteOrganization(reqDto *dto.DeleteOrganizationReqDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-organization", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取部门信息
 * @description 获取部门信息
 * @param organizationCode 组织 code
 * @param departmentId 部门 id，根部门传 `root`
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @returns DepartmentSingleRespDto
 */
func (c *ManagementClient) GetDepartment(reqDto *dto.GetDepartmentDto) *dto.DepartmentSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-department", fasthttp.MethodGet, reqDto)
	var response dto.DepartmentSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建部门
 * @description 创建部门
 * @param requestBody
 * @returns DepartmentSingleRespDto
 */
func (c *ManagementClient) CreateDepartment(reqDto *dto.CreateDepartmentReqDto) *dto.DepartmentSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-department", fasthttp.MethodPost, reqDto)
	var response dto.DepartmentSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 修改部门
 * @description 修改部门
 * @param requestBody
 * @returns DepartmentSingleRespDto
 */
func (c *ManagementClient) UpdateDepartment(reqDto *dto.UpdateDepartmentReqDto) *dto.DepartmentSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-department", fasthttp.MethodPost, reqDto)
	var response dto.DepartmentSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 删除部门
 * @description 删除部门
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) DeleteDepartment(reqDto *dto.DeleteDepartmentReqDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-department", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 搜索部门
 * @description 搜索部门
 * @param requestBody
 * @returns DepartmentListRespDto
 */
func (c *ManagementClient) SearchDepartments(reqDto *dto.SearchDepartmentsReqDto) *dto.DepartmentListRespDto {
	b, err := c.SendHttpRequest("/api/v3/search-departments", fasthttp.MethodPost, reqDto)
	var response dto.DepartmentListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取子部门列表
 * @description 获取子部门列表
 * @param departmentId 需要获取的部门 ID
 * @param organizationCode 组织 code
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @returns DepartmentPaginatedRespDto
 */
func (c *ManagementClient) ListChildrenDepartments(reqDto *dto.ListChildrenDepartmentsDto) *dto.DepartmentPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-children-departments", fasthttp.MethodGet, reqDto)
	var response dto.DepartmentPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取部门成员列表
 * @description 获取部门成员列表
 * @param organizationCode 组织 code
 * @param departmentId 部门 id，根部门传 `root`
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @param includeChildrenDepartments 是否包含子部门的成员
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @returns UserListRespDto
 */
func (c *ManagementClient) ListDepartmentMembers(reqDto *dto.ListDepartmentMembersDto) *dto.UserListRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-department-members", fasthttp.MethodGet, reqDto)
	var response dto.UserListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取部门直属成员 ID 列表
 * @description 获取部门直属成员 ID 列表
 * @param organizationCode 组织 code
 * @param departmentId 部门 id，根部门传 `root`
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @returns UserIdListRespDto
 */
func (c *ManagementClient) ListDepartmentMemberIds(reqDto *dto.ListDepartmentMemberIdsDto) *dto.UserIdListRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-department-member-ids", fasthttp.MethodGet, reqDto)
	var response dto.UserIdListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 部门下添加成员
 * @description 部门下添加成员
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) AddDepartmentMembers(reqDto *dto.AddDepartmentMembersReqDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/add-department-members", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 部门下删除成员
 * @description 部门下删除成员
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) RemoveDepartmentMembers(reqDto *dto.RemoveDepartmentMembersReqDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/remove-department-members", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取父部门信息
 * @description 获取父部门信息
 * @param organizationCode 组织 code
 * @param departmentId 部门 id
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @returns DepartmentSingleRespDto
 */
func (c *ManagementClient) GetParentDepartment(reqDto *dto.GetParentDepartmentDto) *dto.DepartmentSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-parent-department", fasthttp.MethodGet, reqDto)
	var response dto.DepartmentSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取身份源列表
 * @description 获取身份源列表
 * @param tenantId 租户 ID
 * @returns ExtIdpListPaginatedRespDto
 */
func (c *ManagementClient) ListExtIdp(reqDto *dto.ListExtIdpDto) *dto.ExtIdpListPaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-ext-idp", fasthttp.MethodGet, reqDto)
	var response dto.ExtIdpListPaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取身份源详情
 * @description 获取身份源详情
 * @param id 身份源 id
 * @param tenantId 租户 ID
 * @returns ExtIdpDetailSingleRespDto
 */
func (c *ManagementClient) GetExtIdp(reqDto *dto.GetExtIdpDto) *dto.ExtIdpDetailSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-ext-idp", fasthttp.MethodGet, reqDto)
	var response dto.ExtIdpDetailSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建身份源
 * @description 创建身份源
 * @param requestBody
 * @returns ExtIdpSingleRespDto
 */
func (c *ManagementClient) CreateExtIdp(reqDto *dto.CreateExtIdpDto) *dto.ExtIdpSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-ext-idp", fasthttp.MethodPost, reqDto)
	var response dto.ExtIdpSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 更新身份源配置
 * @description 更新身份源配置
 * @param requestBody
 * @returns ExtIdpSingleRespDto
 */
func (c *ManagementClient) UpdateExtIdp(reqDto *dto.UpdateExtIdpDto) *dto.ExtIdpSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-ext-idp", fasthttp.MethodPost, reqDto)
	var response dto.ExtIdpSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 删除身份源
 * @description 删除身份源
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) DeleteExtIdp(reqDto *dto.DeleteExtIdpDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-ext-idp", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 在某个已有身份源下创建新连接
 * @description 在某个已有身份源下创建新连接
 * @param requestBody
 * @returns ExtIdpConnDetailSingleRespDto
 */
func (c *ManagementClient) CreateExtIdpConn(reqDto *dto.CreateExtIdpConnDto) *dto.ExtIdpConnDetailSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-ext-idp-conn", fasthttp.MethodPost, reqDto)
	var response dto.ExtIdpConnDetailSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 更新身份源连接
 * @description 更新身份源连接
 * @param requestBody
 * @returns ExtIdpConnDetailSingleRespDto
 */
func (c *ManagementClient) UpdateExtIdpConn(reqDto *dto.UpdateExtIdpConnDto) *dto.ExtIdpConnDetailSingleRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-ext-idp-conn", fasthttp.MethodPost, reqDto)
	var response dto.ExtIdpConnDetailSingleRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 删除身份源连接
 * @description 删除身份源连接
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) DeleteExtIdpConn(reqDto *dto.DeleteExtIdpConnDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-ext-idp-conn", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 身份源连接开关
 * @description 身份源连接开关
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) ChangeConnState(reqDto *dto.EnableExtIdpConnDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/enable-ext-idp-conn", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户池配置的自定义字段列表
 * @description 获取用户池配置的自定义字段列表
 * @param targetType 主体类型，目前支持用户、角色、分组和部门
 * @returns CustomFieldListRespDto
 */
func (c *ManagementClient) GetCustomFields(reqDto *dto.GetCustomFieldsDto) *dto.CustomFieldListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-custom-fields", fasthttp.MethodGet, reqDto)
	var response dto.CustomFieldListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建/修改自定义字段定义
 * @description 创建/修改自定义字段定义，如果传入的 key 不存在则创建，存在则更新。
 * @param requestBody
 * @returns CustomFieldListRespDto
 */
func (c *ManagementClient) SetCustomFields(reqDto *dto.SetCustomFieldsReqDto) *dto.CustomFieldListRespDto {
	b, err := c.SendHttpRequest("/api/v3/set-custom-fields", fasthttp.MethodPost, reqDto)
	var response dto.CustomFieldListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 设置自定义字段的值
 * @description 给用户、角色、部门设置自定义字段的值，如果存在则更新，不存在则创建。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) SetCustomData(reqDto *dto.SetCustomDataReqDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/set-custom-data", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取用户、分组、角色、组织机构的自定义字段值
 * @description 获取用户、分组、角色、组织机构的自定义字段值
 * @param targetType 主体类型，目前支持用户、角色、分组和部门
 * @param targetIdentifier 目标对象唯一标志符
 * @param namespace 所属权限分组的 code，当 targetType 为角色的时候需要填写，否则可以忽略。
 * @returns GetCustomDataRespDto
 */
func (c *ManagementClient) GetCustomData(reqDto *dto.GetCustomDataDto) *dto.GetCustomDataRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-custom-data", fasthttp.MethodGet, reqDto)
	var response dto.GetCustomDataRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建资源
 * @description 创建资源
 * @param requestBody
 * @returns ResourceRespDto
 */
func (c *ManagementClient) CreateResource(reqDto *dto.CreateResourceDto) *dto.ResourceRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-resource", fasthttp.MethodPost, reqDto)
	var response dto.ResourceRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量创建资源
 * @description 批量创建资源
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) CreateResourcesBatch(reqDto *dto.CreateResourcesBatchDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-resources-batch", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取资源详情
 * @description 获取资源详情
 * @param code 资源唯一标志符
 * @param namespace 所属权限分组的 code
 * @returns ResourceRespDto
 */
func (c *ManagementClient) GetResource(reqDto *dto.GetResourceDto) *dto.ResourceRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-resource", fasthttp.MethodGet, reqDto)
	var response dto.ResourceRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量获取资源详情
 * @description 批量获取资源详情
 * @param codeList 资源 code 列表,批量可以使用逗号分隔
 * @param namespace 所属权限分组的 code
 * @returns ResourceListRespDto
 */
func (c *ManagementClient) GetResourcesBatch(reqDto *dto.GetResourcesBatchDto) *dto.ResourceListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-resources-batch", fasthttp.MethodGet, reqDto)
	var response dto.ResourceListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 分页获取资源列表
 * @description 分页获取资源列表
 * @param namespace 所属权限分组的 code
 * @param type 资源类型
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns ResourcePaginatedRespDto
 */
func (c *ManagementClient) ListResources(reqDto *dto.ListResourcesDto) *dto.ResourcePaginatedRespDto {
	b, err := c.SendHttpRequest("/api/v3/list-resources", fasthttp.MethodGet, reqDto)
	var response dto.ResourcePaginatedRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 修改资源
 * @description 修改资源（Pratial Update）
 * @param requestBody
 * @returns ResourceRespDto
 */
func (c *ManagementClient) UpdateResource(reqDto *dto.UpdateResourceDto) *dto.ResourceRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-resource", fasthttp.MethodPost, reqDto)
	var response dto.ResourceRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 删除资源
 * @description 删除资源
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) DeleteResource(reqDto *dto.DeleteResourceDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-resource", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量删除资源
 * @description 批量删除资源
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) DeleteResourcesBatch(reqDto *dto.DeleteResourcesBatchDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-resources-batch", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 创建权限分组
 * @description 创建权限分组
 * @param requestBody
 * @returns NamespaceRespDto
 */
func (c *ManagementClient) CreateNamespace(reqDto *dto.CreateNamespaceDto) *dto.NamespaceRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-namespace", fasthttp.MethodPost, reqDto)
	var response dto.NamespaceRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量创建权限分组
 * @description 批量创建权限分组
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) CreateNamespacesBatch(reqDto *dto.CreateNamespacesBatchDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/create-namespaces-batch", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取权限分组详情
 * @description 获取权限分组详情
 * @param code 权限分组唯一标志符
 * @returns NamespaceRespDto
 */
func (c *ManagementClient) GetNamespace(reqDto *dto.GetNamespaceDto) *dto.NamespaceRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-namespace", fasthttp.MethodGet, reqDto)
	var response dto.NamespaceRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量获取权限分组详情
 * @description 批量获取权限分组详情
 * @param codeList 资源 code 列表,批量可以使用逗号分隔
 * @returns NamespaceListRespDto
 */
func (c *ManagementClient) GetNamespacesBatch(reqDto *dto.GetNamespacesBatchDto) *dto.NamespaceListRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-namespaces-batch", fasthttp.MethodGet, reqDto)
	var response dto.NamespaceListRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 修改权限分组信息
 * @description 修改权限分组信息
 * @param requestBody
 * @returns UpdateNamespaceRespDto
 */
func (c *ManagementClient) UpdateNamespace(reqDto *dto.UpdateNamespaceDto) *dto.UpdateNamespaceRespDto {
	b, err := c.SendHttpRequest("/api/v3/update-namespace", fasthttp.MethodPost, reqDto)
	var response dto.UpdateNamespaceRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 删除权限分组信息
 * @description 删除权限分组信息
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) DeleteNamespace(reqDto *dto.DeleteNamespaceDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-namespace", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 批量删除权限分组
 * @description 批量删除权限分组
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) DeleteNamespacesBatch(reqDto *dto.DeleteNamespacesBatchDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/delete-namespaces-batch", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 授权资源
 * @description 给多个主体同时授权多个资源
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) AuthorizeResources(reqDto *dto.AuthorizeResourcesDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/authorize-resources", fasthttp.MethodPost, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}

/*
 * @summary 获取某个主体被授权的资源列表
 * @description 获取某个主体被授权的资源列表
 * @param targetType 目标对象类型
 * @param targetIdentifier 目标对象唯一标志符
 * @param namespace 所属权限分组的 code
 * @param resourceType 资源类型，如数据、API、按钮、菜单
 * @returns IsSuccessRespDto
 */
func (c *ManagementClient) GetTargetAuthorizedResources(reqDto *dto.GetAuthorizedResourcesDto) *dto.IsSuccessRespDto {
	b, err := c.SendHttpRequest("/api/v3/get-authorized-resources", fasthttp.MethodGet, reqDto)
	var response dto.IsSuccessRespDto
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &response
}
