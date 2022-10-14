package management

import (
	"encoding/json"
	"fmt"
	"github.com/Authing/authing-golang-sdk/dto"
	"github.com/valyala/fasthttp"
)

/*
 * @summary 获取/搜索用户列表
 * @description
 * 此接口用于获取用户列表，支持模糊搜索，以及通过用户基础字段、用户自定义字段、用户所在部门、用户历史登录应用等维度筛选用户。
 *
 * ### 模糊搜素示例
 *
 * 模糊搜索默认会从 `phone`, `email`, `name`, `username`, `nickname` 五个字段对用户进行模糊搜索，你也可以通过设置 `options.fuzzySearchOn`
 * 决定模糊匹配的字段范围：
 *
 * ```json
 * {
     * "query": "北京",
     * "options": {
         * "fuzzySearchOn": [
             * "address"
             * ]
             * }
             * }
             * ```
             *
             * ### 高级搜索示例
             *
             * 你可以通过 `advancedFilter` 进行高级搜索，高级搜索支持通过用户的基础信息、自定义数据、所在部门、用户来源、登录应用、外部身份源信息等维度对用户进行筛选。
             * **且这些筛选条件可以任意组合。**
             *
             * #### 筛选状态为禁用的用户
             *
             * 用户状态（`status`）为字符串类型，可选值为 `Activated` 和 `Suspended`：
             *
             * ```json
             * {
                 * "advancedFilter": [
                     * {
                         * "field": "status",
                         * "operator": "EQUAL",
                         * "value": "Suspended"
                         * }
                         * ]
                         * }
                         * ```
                         *
                         * #### 筛选邮箱中包含 `@example.com` 的用户
                         *
                         * 用户邮箱（`email`）为字符串类型，可以进行模糊搜索：
                         *
                         * ```json
                         * {
                             * "advancedFilter": [
                                 * {
                                     * "field": "email",
                                     * "operator": "CONTAINS",
                                     * "value": "@example.com"
                                     * }
                                     * ]
                                     * }
                                     * ```
                                     *
                                     * #### 根据用户登录次数筛选
                                     *
                                     * 筛选登录次数大于 10 的用户：
                                     *
                                     * ```json
                                     * {
                                         * "advancedFilter": [
                                             * {
                                                 * "field": "loginsCount",
                                                 * "operator": "GREATER",
                                                 * "value": 10
                                                 * }
                                                 * ]
                                                 * }
                                                 * ```
                                                 *
                                                 * 筛选登录次数在 10 - 100 次的用户：
                                                 *
                                                 * ```json
                                                 * {
                                                     * "advancedFilter": [
                                                         * {
                                                             * "field": "loginsCount",
                                                             * "operator": "BETWEEN",
                                                             * "value": [10, 100]
                                                             * }
                                                             * ]
                                                             * }
                                                             * ```
                                                             *
                                                             * #### 根据用户上次登录时间进行筛选
                                                             *
                                                             * 筛选最近 7 天内登录过的用户：
                                                             *
                                                             * ```json
                                                             * {
                                                                 * "advancedFilter": [
                                                                     * {
                                                                         * "field": "lastLoginTime",
                                                                         * "operator": "GREATER",
                                                                         * "value": new Date(Date.now() - 7 * 24 * 60 * 60 * 1000)
                                                                         * }
                                                                         * ]
                                                                         * }
                                                                         * ```
                                                                         *
                                                                         * 筛选在某一段时间内登录过的用户：
                                                                         *
                                                                         * ```json
                                                                         * {
                                                                             * "advancedFilter": [
                                                                                 * {
                                                                                     * "field": "lastLoginTime",
                                                                                     * "operator": "BETWEEN",
                                                                                     * "value": [
                                                                                         * new Date(Date.now() - 14 * 24 * 60 * 60 * 1000),
                                                                                         * new Date(Date.now() - 7 * 24 * 60 * 60 * 1000)
                                                                                         * ]
                                                                                         * }
                                                                                         * ]
                                                                                         * }
                                                                                         * ```
                                                                                         *
                                                                                         * #### 根据用户曾经登录过的应用筛选
                                                                                         *
                                                                                         * 筛选出曾经登录过应用 `appId1` 或者 `appId2` 的用户：
                                                                                         *
                                                                                         * ```json
                                                                                         * {
                                                                                             * "advancedFilter": [
                                                                                                 * {
                                                                                                     * "field": "loggedInApps",
                                                                                                     * "operator": "IN",
                                                                                                     * "value": [
                                                                                                         * "appId1",
                                                                                                         * "appId2"
                                                                                                         * ]
                                                                                                         * }
                                                                                                         * ]
                                                                                                         * }
                                                                                                         * ```
                                                                                                         *
                                                                                                         * #### 根据用户所在部门进行筛选
                                                                                                         *
                                                                                                         * ```json
                                                                                                         * {
                                                                                                             * "advancedFilter": [
                                                                                                                 * {
                                                                                                                     * "field": "department",
                                                                                                                     * "operator": "IN",
                                                                                                                     * "value": [
                                                                                                                         * {
                                                                                                                             * "organizationCode": "steamory",
                                                                                                                             * "departmentId": "root",
                                                                                                                             * "departmentIdType": "department_id",
                                                                                                                             * "includeChildrenDepartments": true
                                                                                                                             * }
                                                                                                                             * ]
                                                                                                                             * }
                                                                                                                             * ]
                                                                                                                             * }
                                                                                                                             * ```
                                                                                                                             *
                                                                                                                             *
                                                                                                                             * @param requestBody
                                                                                                                             * @returns UserPaginatedRespDto
*/
func (client *ManagementClient) ListUsers(reqDto *dto.ListUsersRequestDto) *dto.UserPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-users", fasthttp.MethodPost, reqDto)
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
 * @deprecated
 * @summary 获取用户列表
 * @description 获取用户列表接口，支持分页，可以选择获取自定义数据、identities 等。
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param status 账户当前状态，如 已停用、已离职、正常状态、已归档
 * @param updatedAtStart 用户创建、修改开始时间，为精确到秒的 UNIX 时间戳；支持获取从某一段时间之后的增量数据
 * @param updatedAtEnd 用户创建、修改终止时间，为精确到秒的 UNIX 时间戳；支持获取某一段时间内的增量数据。默认为当前时间
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @returns UserPaginatedRespDto
 */
func (client *ManagementClient) ListUsersLegacy(reqDto *dto.ListUsersDto) *dto.UserPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-users", fasthttp.MethodGet, reqDto)
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
 * @summary 获取用户信息
 * @description 通过用户 ID，获取用户详情，可以选择获取自定义数据、identities、选择指定用户 ID 类型等。
 * @param userId 用户 ID
 * @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @returns UserSingleRespDto
 */
func (client *ManagementClient) GetUser(reqDto *dto.GetUserDto) *dto.UserSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user", fasthttp.MethodGet, reqDto)
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
 * @description 通过用户 ID 列表，批量获取用户信息，可以选择获取自定义数据、identities、选择指定用户 ID 类型等。
 * @param userIds 用户 ID 数组
 * @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @returns UserListRespDto
 */
func (client *ManagementClient) GetUserBatch(reqDto *dto.GetUserBatchDto) *dto.UserListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-batch", fasthttp.MethodGet, reqDto)
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
 * @summary 获取用户的外部身份源
 * @description 通过用户 ID，获取用户的外部身份源、选择指定用户 ID 类型。
 * @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
 * @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
 * @returns IdentityListRespDto
 */
func (client *ManagementClient) GetUserIdentities(reqDto *dto.GetUserIdentitiesDto) *dto.IdentityListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-identities", fasthttp.MethodGet, reqDto)
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
 * @description 通过用户 ID，获取用户角色列表，可以选择所属权限分组 code、选择指定用户 ID 类型等。
 * @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
 * @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
 * @param namespace 所属权限分组的 code
 * @returns RolePaginatedRespDto
 */
func (client *ManagementClient) GetUserRoles(reqDto *dto.GetUserRolesDto) *dto.RolePaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-roles", fasthttp.MethodGet, reqDto)
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
 * @description 通过用户 ID，获取用户实名认证信息，可以选择指定用户 ID 类型。
 * @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
 * @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
 * @returns PrincipalAuthenticationInfoPaginatedRespDto
 */
func (client *ManagementClient) GetUserPrincipalAuthenticationInfo(reqDto *dto.GetUserPrincipalAuthenticationInfoDto) *dto.PrincipalAuthenticationInfoPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-principal-authentication-info", fasthttp.MethodGet, reqDto)
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
 * @description 通过用户 ID，删除用户实名认证信息，可以选择指定用户 ID 类型等。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) ResetUserPrincipalAuthenticationInfo(reqDto *dto.ResetUserPrincipalAuthenticationInfoDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/reset-user-principal-authentication-info", fasthttp.MethodPost, reqDto)
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
 * @description 通过用户 ID，获取用户部门列表，支持分页，可以选择获取自定义数据、选择指定用户 ID 类型、增序或降序等。
 * @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
 * @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param withCustomData 是否获取自定义数据
 * @param sortBy 排序依据，如 部门创建时间、加入部门时间、部门名称、部门标志符
 * @param orderBy 增序或降序
 * @returns UserDepartmentPaginatedRespDto
 */
func (client *ManagementClient) GetUserDepartments(reqDto *dto.GetUserDepartmentsDto) *dto.UserDepartmentPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-departments", fasthttp.MethodGet, reqDto)
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
 * @description 通过用户 ID，设置用户所在部门，可以选择指定用户 ID 类型等。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) SetUserDepartments(reqDto *dto.SetUserDepartmentsDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/set-user-departments", fasthttp.MethodPost, reqDto)
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
 * @description 通过用户 ID，获取用户分组列表，可以选择指定用户 ID 类型等。
 * @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
 * @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
 * @returns GroupPaginatedRespDto
 */
func (client *ManagementClient) GetUserGroups(reqDto *dto.GetUserGroupsDto) *dto.GroupPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-groups", fasthttp.MethodGet, reqDto)
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
 * @description 通过用户 ID 列表，删除用户，支持批量删除，可以选择指定用户 ID 类型等。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteUsersBatch(reqDto *dto.DeleteUsersBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-users-batch", fasthttp.MethodPost, reqDto)
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
 * @description 通过用户 ID，获取用户 MFA 绑定信息，可以选择指定用户 ID 类型等。
 * @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
 * @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
 * @returns UserMfaSingleRespDto
 */
func (client *ManagementClient) GetUserMfaInfo(reqDto *dto.GetUserMfaInfoDto) *dto.UserMfaSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-mfa-info", fasthttp.MethodGet, reqDto)
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
 * @description 获取已归档的用户列表，支持分页，可以筛选开始时间等。
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param startAt 开始时间，为精确到秒的 UNIX 时间戳，默认不指定
 * @returns ListArchivedUsersSingleRespDto
 */
func (client *ManagementClient) ListArchivedUsers(reqDto *dto.ListArchivedUsersDto) *dto.ListArchivedUsersSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-archived-users", fasthttp.MethodGet, reqDto)
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
 * @description 通过用户 ID、App ID 列表，强制让用户下线，可以选择指定用户 ID 类型等。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) KickUsers(reqDto *dto.KickUsersDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/kick-users", fasthttp.MethodPost, reqDto)
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
 * @description 根据条件判断用户是否存在，可以筛选用户名、邮箱、手机号、第三方外部 ID 等。
 * @param requestBody
 * @returns IsUserExistsRespDto
 */
func (client *ManagementClient) IsUserExists(reqDto *dto.IsUserExistsReqDto) *dto.IsUserExistsRespDto {
	b, err := client.SendHttpRequest("/api/v3/is-user-exists", fasthttp.MethodPost, reqDto)
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
 * @description 创建用户，邮箱、手机号、用户名必须包含其中一个，邮箱、手机号、用户名、externalId 用户池内唯一，此接口将以管理员身份创建用户因此不需要进行手机号验证码检验等安全检测。
 * @param requestBody
 * @returns UserSingleRespDto
 */
func (client *ManagementClient) CreateUser(reqDto *dto.CreateUserReqDto) *dto.UserSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-user", fasthttp.MethodPost, reqDto)
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
 * @description 批量创建用户，邮箱、手机号、用户名必须包含其中一个，邮箱、手机号、用户名、externalId 用户池内唯一，此接口将以管理员身份批量创建用户因此不需要进行手机号验证码检验等安全检测。
 * @param requestBody
 * @returns UserListRespDto
 */
func (client *ManagementClient) CreateUsersBatch(reqDto *dto.CreateUserBatchReqDto) *dto.UserListRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-users-batch", fasthttp.MethodPost, reqDto)
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
 * @description 通过用户 ID，修改用户资料，邮箱、手机号、用户名、externalId 用户池内唯一，此接口将以管理员身份修改用户资料因此不需要进行手机号验证码检验等安全检测。
 * @param requestBody
 * @returns UserSingleRespDto
 */
func (client *ManagementClient) UpdateUser(reqDto *dto.UpdateUserReqDto) *dto.UserSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-user", fasthttp.MethodPost, reqDto)
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
 * @summary 修改用户资料
 * @description 通过用户 ID，修改用户资料，邮箱、手机号、用户名、externalId 用户池内唯一，此接口将以管理员身份修改用户资料因此不需要进行手机号验证码检验等安全检测。
 * @param requestBody
 * @returns UserListRespDto
 */
func (client *ManagementClient) UpdateUserBatch(reqDto *dto.UpdateUserBatchReqDto) *dto.UserListRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-user-batch", fasthttp.MethodPost, reqDto)
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
 * @summary 获取用户可访问的应用
 * @description 通过用户 ID，获取用户可访问的应用，可以选择指定用户 ID 类型等。
 * @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
 * @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
 * @returns AppListRespDto
 */
func (client *ManagementClient) GetUserAccessibleApps(reqDto *dto.GetUserAccessibleAppsDto) *dto.AppListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-accessible-apps", fasthttp.MethodGet, reqDto)
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
 * @description 通过用户 ID，获取用户授权的应用，可以选择指定用户 ID 类型等。
 * @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
 * @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
 * @returns AppListRespDto
 */
func (client *ManagementClient) GetUserAuthorizedApps(reqDto *dto.GetUserAuthorizedAppsDto) *dto.AppListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-authorized-apps", fasthttp.MethodGet, reqDto)
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
 * @description 通过用户 ID，判断用户是否有某个角色，支持传入多个角色，可以选择指定用户 ID 类型等。
 * @param requestBody
 * @returns HasAnyRoleRespDto
 */
func (client *ManagementClient) HasAnyRole(reqDto *dto.HasAnyRoleReqDto) *dto.HasAnyRoleRespDto {
	b, err := client.SendHttpRequest("/api/v3/has-any-role", fasthttp.MethodPost, reqDto)
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
 * @description 通过用户 ID，获取用户登录历史记录，支持分页，可以选择指定用户 ID 类型、应用 ID、开始与结束时间戳等。
 * @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
 * @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
 * @param appId 应用 ID
 * @param clientIp 客户端 IP
 * @param start 开始时间戳（毫秒）
 * @param end 结束时间戳（毫秒）
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns UserLoginHistoryPaginatedRespDto
 */
func (client *ManagementClient) GetUserLoginHistory(reqDto *dto.GetUserLoginHistoryDto) *dto.UserLoginHistoryPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-login-history", fasthttp.MethodGet, reqDto)
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
 * @description 通过用户 ID，获取用户曾经登录过的应用，可以选择指定用户 ID 类型等。
 * @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
 * @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
 * @returns UserLoggedInAppsListRespDto
 */
func (client *ManagementClient) GetUserLoggedinApps(reqDto *dto.GetUserLoggedinAppsDto) *dto.UserLoggedInAppsListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-loggedin-apps", fasthttp.MethodGet, reqDto)
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
 * @summary 获取用户曾经登录过的身份源
 * @description 通过用户 ID，获取用户曾经登录过的身份源，可以选择指定用户 ID 类型等。
 * @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
 * @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
 * @returns UserLoggedInIdentitiesRespDto
 */
func (client *ManagementClient) GetUserLoggedinIdentities(reqDto *dto.GetUserLoggedInIdentitiesDto) *dto.UserLoggedInIdentitiesRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-logged-in-identities", fasthttp.MethodGet, reqDto)
	var response dto.UserLoggedInIdentitiesRespDto
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
 * @summary 用户离职
 * @description 通过用户 ID，对用户进行离职操作
 * @param requestBody
 * @returns ResignUserRespDto
 */
func (client *ManagementClient) ResignUser(reqDto *dto.ResignUserReqDto) *dto.ResignUserRespDto {
	b, err := client.SendHttpRequest("/api/v3/resign-user", fasthttp.MethodPost, reqDto)
	var response dto.ResignUserRespDto
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
 * @summary 批量用户离职
 * @description 通过用户 ID，对用户进行离职操作
 * @param requestBody
 * @returns ResignUserRespDto
 */
func (client *ManagementClient) ResignUserBatch(reqDto *dto.ResignUserBatchReqDto) *dto.ResignUserRespDto {
	b, err := client.SendHttpRequest("/api/v3/resign-user-batch", fasthttp.MethodPost, reqDto)
	var response dto.ResignUserRespDto
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
 * @description 通过用户 ID，获取用户被授权的所有资源，可以选择指定用户 ID 类型等，用户被授权的资源是用户自身被授予、通过分组继承、通过角色继承、通过组织机构继承的集合。
 * @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
 * @param userIdType 用户 ID 类型，默认值为 `user_id`，可选值为：
 * - `user_id`: Authing 用户 ID，如 `6319a1504f3xxxxf214dd5b7`
 * - `phone`: 用户手机号
 * - `email`: 用户邮箱
 * - `username`: 用户名
 * - `external_id`: 用户在外部系统的 ID，对应 Authing 用户信息的 `externalId` 字段
 * - `identity`: 用户的外部身份源信息，格式为 `<extIdpId>:<userIdInIdp>`，其中 `<extIdpId>` 为 Authing 身份源的 ID，`<userIdInIdp>` 为用户在外部身份源的 ID。
 * 示例值：`62f20932716fbcc10d966ee5:ou_8bae746eac07cd2564654140d2a9ac61`。
 *
 * @param namespace 所属权限分组的 code
 * @param resourceType 资源类型，如 数据、API、菜单、按钮
 * @returns AuthorizedResourcePaginatedRespDto
 */
func (client *ManagementClient) GetUserAuthorizedResources(reqDto *dto.GetUserAuthorizedResourcesDto) *dto.AuthorizedResourcePaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-authorized-resources", fasthttp.MethodGet, reqDto)
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
 * @summary 检查某个用户在应用下是否具备 Session 登录态
 * @description 检查某个用户在应用下是否具备 Session 登录态
 * @param requestBody
 * @returns CheckSessionStatusRespDto
 */
func (client *ManagementClient) CheckSessionStatus(reqDto *dto.CheckSessionStatusDto) *dto.CheckSessionStatusRespDto {
	b, err := client.SendHttpRequest("/api/v3/check-session-status", fasthttp.MethodPost, reqDto)
	var response dto.CheckSessionStatusRespDto
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
 * @summary 导入用户的 OTP
 * @description 导入用户的 OTP
 * @param requestBody
 * @returns CommonResponseDto
 */
func (client *ManagementClient) ImportOtp(reqDto *dto.ImportOtpReqDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/import-otp", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
 * @summary 获取组织机构详情
 * @description 获取组织机构详情
 * @param organizationCode 组织 Code（organizationCode）
 * @param withCustomData 是否获取自定义数据
 * @returns OrganizationSingleRespDto
 */
func (client *ManagementClient) GetOrganization(reqDto *dto.GetOrganizationDto) *dto.OrganizationSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-organization", fasthttp.MethodGet, reqDto)
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
 * @summary 批量获取组织机构详情
 * @description 批量获取组织机构详情
 * @param organizationCodeList 组织 Code（organizationCode）列表
 * @param withCustomData 是否获取自定义数据
 * @returns OrganizationListRespDto
 */
func (client *ManagementClient) GetOrganizationsBatch(reqDto *dto.GetOrganizationBatchDto) *dto.OrganizationListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-organization-batch", fasthttp.MethodGet, reqDto)
	var response dto.OrganizationListRespDto
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
 * @description 获取顶层组织机构列表，支持分页。
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param fetchAll 拉取所有
 * @param withCustomData 是否获取自定义数据
 * @returns OrganizationPaginatedRespDto
 */
func (client *ManagementClient) ListOrganizations(reqDto *dto.ListOrganizationsDto) *dto.OrganizationPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-organizations", fasthttp.MethodGet, reqDto)
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
 * @description 创建组织机构，会创建一个只有一个节点的组织机构，可以选择组织描述信息、根节点自定义 ID、多语言等。
 * @param requestBody
 * @returns OrganizationSingleRespDto
 */
func (client *ManagementClient) CreateOrganization(reqDto *dto.CreateOrganizationReqDto) *dto.OrganizationSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-organization", fasthttp.MethodPost, reqDto)
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
 * @description 通过组织 code，修改顶层组织机构，可以选择部门描述、新组织 code、组织名称等。
 * @param requestBody
 * @returns OrganizationSingleRespDto
 */
func (client *ManagementClient) UpdateOrganization(reqDto *dto.UpdateOrganizationReqDto) *dto.OrganizationSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-organization", fasthttp.MethodPost, reqDto)
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
 * @description 通过组织 code，删除组织机构树。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteOrganization(reqDto *dto.DeleteOrganizationReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-organization", fasthttp.MethodPost, reqDto)
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
 * @summary 搜索顶层组织机构列表
 * @description 通过搜索关键词，搜索顶层组织机构列表，支持分页。
 * @param keywords 搜索关键词，如组织机构名称
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param withCustomData 是否获取自定义数据
 * @returns OrganizationPaginatedRespDto
 */
func (client *ManagementClient) SearchOrganizations(reqDto *dto.SearchOrganizationsDto) *dto.OrganizationPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/search-organizations", fasthttp.MethodGet, reqDto)
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
 * @summary 获取部门信息
 * @description 通过组织 code 以及 部门 ID 或 部门 code，获取部门信息，可以获取自定义数据。
 * @param organizationCode 组织 code
 * @param departmentId 部门 ID，根部门传 `root`。departmentId 和 departmentCode 必传其一。
 * @param departmentCode 部门 code。departmentId 和 departmentCode 必传其一。
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @param withCustomData 是否获取自定义数据
 * @returns DepartmentSingleRespDto
 */
func (client *ManagementClient) GetDepartment(reqDto *dto.GetDepartmentDto) *dto.DepartmentSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-department", fasthttp.MethodGet, reqDto)
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
 * @description 通过组织 code、部门名称、父部门 ID，创建部门，可以设置多种参数。
 * @param requestBody
 * @returns DepartmentSingleRespDto
 */
func (client *ManagementClient) CreateDepartment(reqDto *dto.CreateDepartmentReqDto) *dto.DepartmentSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-department", fasthttp.MethodPost, reqDto)
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
 * @description 通过组织 code、部门 ID，修改部门，可以设置多种参数。
 * @param requestBody
 * @returns DepartmentSingleRespDto
 */
func (client *ManagementClient) UpdateDepartment(reqDto *dto.UpdateDepartmentReqDto) *dto.DepartmentSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-department", fasthttp.MethodPost, reqDto)
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
 * @description 通过组织 code、部门 ID，删除部门。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteDepartment(reqDto *dto.DeleteDepartmentReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-department", fasthttp.MethodPost, reqDto)
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
 * @description 通过组织 code、搜索关键词，搜索部门，可以搜索组织名称等。
 * @param requestBody
 * @returns DepartmentListRespDto
 */
func (client *ManagementClient) SearchDepartments(reqDto *dto.SearchDepartmentsReqDto) *dto.DepartmentListRespDto {
	b, err := client.SendHttpRequest("/api/v3/search-departments", fasthttp.MethodPost, reqDto)
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
 * @description 通过组织 code、部门 ID，获取子部门列表，可以选择获取自定义数据、虚拟组织等。
 * @param organizationCode 组织 code
 * @param departmentId 需要获取的部门 ID
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @param excludeVirtualNode 是否要排除虚拟组织
 * @param onlyVirtualNode 是否只包含虚拟组织
 * @param withCustomData 是否获取自定义数据
 * @returns DepartmentPaginatedRespDto
 */
func (client *ManagementClient) ListChildrenDepartments(reqDto *dto.ListChildrenDepartmentsDto) *dto.DepartmentPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-children-departments", fasthttp.MethodGet, reqDto)
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
 * @description 通过组织 code、部门 ID、排序，获取部门成员列表，支持分页，可以选择获取自定义数据、identities 等。
 * @param organizationCode 组织 code
 * @param departmentId 部门 ID，根部门传 `root`
 * @param sortBy 排序依据
 * @param orderBy 增序还是倒序
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @param includeChildrenDepartments 是否包含子部门的成员
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @returns UserPaginatedRespDto
 */
func (client *ManagementClient) ListDepartmentMembers(reqDto *dto.ListDepartmentMembersDto) *dto.UserPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-department-members", fasthttp.MethodGet, reqDto)
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
 * @summary 获取部门直属成员 ID 列表
 * @description 通过组织 code、部门 ID，获取部门直属成员 ID 列表。
 * @param organizationCode 组织 code
 * @param departmentId 部门 ID，根部门传 `root`
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @returns UserIdListRespDto
 */
func (client *ManagementClient) ListDepartmentMemberIds(reqDto *dto.ListDepartmentMemberIdsDto) *dto.UserIdListRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-department-member-ids", fasthttp.MethodGet, reqDto)
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
 * @summary 搜索部门下的成员
 * @description 通过组织 code、部门 ID、搜索关键词，搜索部门下的成员，支持分页，可以选择获取自定义数据、identities 等。
 * @param organizationCode 组织 code
 * @param departmentId 部门 ID，根部门传 `root`
 * @param keywords 搜索关键词，如成员名称
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @param includeChildrenDepartments 是否包含子部门的成员
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @returns UserPaginatedRespDto
 */
func (client *ManagementClient) SearchDepartmentMembers(reqDto *dto.SearchDepartmentMembersDto) *dto.UserPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/search-department-members", fasthttp.MethodGet, reqDto)
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
 * @summary 部门下添加成员
 * @description 通过部门 ID、组织 code，添加部门下成员。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) AddDepartmentMembers(reqDto *dto.AddDepartmentMembersReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/add-department-members", fasthttp.MethodPost, reqDto)
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
 * @description 通过部门 ID、组织 code，删除部门下成员。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) RemoveDepartmentMembers(reqDto *dto.RemoveDepartmentMembersReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/remove-department-members", fasthttp.MethodPost, reqDto)
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
 * @description 通过组织 code、部门 ID，获取父部门信息，可以选择获取自定义数据等。
 * @param organizationCode 组织 code
 * @param departmentId 部门 ID
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @param withCustomData 是否获取自定义数据
 * @returns DepartmentSingleRespDto
 */
func (client *ManagementClient) GetParentDepartment(reqDto *dto.GetParentDepartmentDto) *dto.DepartmentSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-parent-department", fasthttp.MethodGet, reqDto)
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
 * @summary 判断用户是否在某个部门下
 * @description 通过组织 code、部门 ID，判断用户是否在某个部门下，可以选择包含子部门。
 * @param userId 用户唯一标志，可以是用户 ID、用户名、邮箱、手机号、外部 ID、在外部身份源的 ID。
 * @param organizationCode 组织 code
 * @param departmentId 部门 ID，根部门传 `root`。departmentId 和 departmentCode 必传其一。
 * @param departmentIdType 此次调用中使用的部门 ID 的类型
 * @param includeChildrenDepartments 是否包含子部门
 * @returns IsUserInDepartmentRespDto
 */
func (client *ManagementClient) IsUserInDepartment(reqDto *dto.IsUserInDepartmentDto) *dto.IsUserInDepartmentRespDto {
	b, err := client.SendHttpRequest("/api/v3/is-user-in-department", fasthttp.MethodGet, reqDto)
	var response dto.IsUserInDepartmentRespDto
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
 * @description 通过分组 code，获取分组详情。
 * @param code 分组 code
 * @returns GroupSingleRespDto
 */
func (client *ManagementClient) GetGroup(reqDto *dto.GetGroupDto) *dto.GroupSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-group", fasthttp.MethodGet, reqDto)
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
 * @description 获取分组列表，支持分页。
 * @param keywords 搜索分组 code 或分组名称
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns GroupPaginatedRespDto
 */
func (client *ManagementClient) ListGroups(reqDto *dto.ListGroupsDto) *dto.GroupPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-groups", fasthttp.MethodGet, reqDto)
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
 * @description 创建分组，一个分组必须包含分组名称与唯一标志符 code，且必须为一个合法的英文标志符，如 developers。
 * @param requestBody
 * @returns GroupSingleRespDto
 */
func (client *ManagementClient) CreateGroup(reqDto *dto.CreateGroupReqDto) *dto.GroupSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-group", fasthttp.MethodPost, reqDto)
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
 * @description 批量创建分组，一个分组必须包含分组名称与唯一标志符 code，且必须为一个合法的英文标志符，如 developers。
 * @param requestBody
 * @returns GroupListRespDto
 */
func (client *ManagementClient) CreateGroupsBatch(reqDto *dto.CreateGroupBatchReqDto) *dto.GroupListRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-groups-batch", fasthttp.MethodPost, reqDto)
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
 * @description 通过分组 code，修改分组，可以修改此分组的 code。
 * @param requestBody
 * @returns GroupSingleRespDto
 */
func (client *ManagementClient) UpdateGroup(reqDto *dto.UpdateGroupReqDto) *dto.GroupSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-group", fasthttp.MethodPost, reqDto)
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
 * @description 通过分组 code，批量删除分组。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteGroupsBatch(reqDto *dto.DeleteGroupsReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-groups-batch", fasthttp.MethodPost, reqDto)
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
 * @description 添加分组成员，成员以用户 ID 数组形式传递。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) AddGroupMembers(reqDto *dto.AddGroupMembersReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/add-group-members", fasthttp.MethodPost, reqDto)
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
 * @description 批量移除分组成员，成员以用户 ID 数组形式传递。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) RemoveGroupMembers(reqDto *dto.RemoveGroupMembersReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/remove-group-members", fasthttp.MethodPost, reqDto)
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
 * @description 通过分组 code，获取分组成员列表，支持分页，可以获取自定义数据、identities、部门 ID 列表。
 * @param code 分组 code
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @returns UserPaginatedRespDto
 */
func (client *ManagementClient) ListGroupMembers(reqDto *dto.ListGroupMembersDto) *dto.UserPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-group-members", fasthttp.MethodGet, reqDto)
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
 * @description 通过分组 code，获取分组被授权的资源列表，可以通过资源类型、权限分组 code 筛选。
 * @param code 分组 code
 * @param namespace 所属权限分组的 code
 * @param resourceType 资源类型
 * @returns AuthorizedResourceListRespDto
 */
func (client *ManagementClient) GetGroupAuthorizedResources(reqDto *dto.GetGroupAuthorizedResourcesDto) *dto.AuthorizedResourceListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-group-authorized-resources", fasthttp.MethodGet, reqDto)
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
 * @description 通过权限分组内角色 code，获取角色详情。
 * @param code 权限分组内角色的唯一标识符
 * @param namespace 所属权限分组的 code
 * @returns RoleSingleRespDto
 */
func (client *ManagementClient) GetRole(reqDto *dto.GetRoleDto) *dto.RoleSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-role", fasthttp.MethodGet, reqDto)
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
 * @description 通过权限分组内角色 code，分配角色，被分配者可以是用户或部门。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) AssignRole(reqDto *dto.AssignRoleDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/assign-role", fasthttp.MethodPost, reqDto)
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
 * @description 通过权限分组内角色 code，移除分配的角色，被分配者可以是用户或部门。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) RevokeRole(reqDto *dto.RevokeRoleDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/revoke-role", fasthttp.MethodPost, reqDto)
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
 * @summary 获取角色被授权的资源列表
 * @description 通过权限分组内角色 code，获取角色被授权的资源列表。
 * @param code 权限分组内角色的唯一标识符
 * @param namespace 所属权限分组的 code
 * @param resourceType 资源类型，如 数据、API、按钮、菜单
 * @returns RoleAuthorizedResourcePaginatedRespDto
 */
func (client *ManagementClient) GetRoleAuthorizedResources(reqDto *dto.GetRoleAuthorizedResourcesDto) *dto.RoleAuthorizedResourcePaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-role-authorized-resources", fasthttp.MethodGet, reqDto)
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
 * @description 通过权限分组内内角色 code，获取角色成员列表，支持分页，可以选择或获取自定义数据、identities 等。
 * @param code 权限分组内角色的唯一标识符
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param withCustomData 是否获取自定义数据
 * @param withIdentities 是否获取 identities
 * @param withDepartmentIds 是否获取部门 ID 列表
 * @param namespace 所属权限分组的 code
 * @returns UserPaginatedRespDto
 */
func (client *ManagementClient) ListRoleMembers(reqDto *dto.ListRoleMembersDto) *dto.UserPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-role-members", fasthttp.MethodGet, reqDto)
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
 * @description 通过权限分组内角色 code，获取角色的部门列表，支持分页。
 * @param code 权限分组内角色的唯一标识符
 * @param namespace 所属权限分组的 code
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns RoleDepartmentListPaginatedRespDto
 */
func (client *ManagementClient) ListRoleDepartments(reqDto *dto.ListRoleDepartmentsDto) *dto.RoleDepartmentListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-role-departments", fasthttp.MethodGet, reqDto)
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
 * @description 通过权限分组内角色 code，创建角色，可以选择权限分组、角色描述等。
 * @param requestBody
 * @returns RoleSingleRespDto
 */
func (client *ManagementClient) CreateRole(reqDto *dto.CreateRoleDto) *dto.RoleSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-role", fasthttp.MethodPost, reqDto)
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
 * @description 获取角色列表，支持分页。
 * @param keywords 搜索角色 code
 * @param namespace 所属权限分组的 code
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns RolePaginatedRespDto
 */
func (client *ManagementClient) ListRoles(reqDto *dto.ListRolesDto) *dto.RolePaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-roles", fasthttp.MethodGet, reqDto)
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
 * @summary 删除角色
 * @description 删除角色，可以批量删除。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteRolesBatch(reqDto *dto.DeleteRoleDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-roles-batch", fasthttp.MethodPost, reqDto)
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
 * @description 批量创建角色，可以选择权限分组、角色描述等。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) CreateRolesBatch(reqDto *dto.CreateRolesBatch) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-roles-batch", fasthttp.MethodPost, reqDto)
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
 * @description 通过权限分组内角色新旧 code，修改角色，可以选择角色描述等。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) UpdateRole(reqDto *dto.UpdateRoleDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-role", fasthttp.MethodPost, reqDto)
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
 * @summary 获取身份源列表
 * @description 获取身份源列表，可以指定 租户 ID 筛选。
 * @param tenantId 租户 ID
 * @param appId 应用 ID
 * @returns ExtIdpListPaginatedRespDto
 */
func (client *ManagementClient) ListExtIdp(reqDto *dto.ListExtIdpDto) *dto.ExtIdpListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-ext-idp", fasthttp.MethodGet, reqDto)
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
 * @description 通过 身份源 ID，获取身份源详情，可以指定 租户 ID 筛选。
 * @param id 身份源 ID
 * @param tenantId 租户 ID
 * @param appId 应用 ID
 * @param type 身份源类型
 * @returns ExtIdpDetailSingleRespDto
 */
func (client *ManagementClient) GetExtIdp(reqDto *dto.GetExtIdpDto) *dto.ExtIdpDetailSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-ext-idp", fasthttp.MethodGet, reqDto)
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
 * @description 创建身份源，可以设置身份源名称、连接类型、租户 ID 等。
 * @param requestBody
 * @returns ExtIdpSingleRespDto
 */
func (client *ManagementClient) CreateExtIdp(reqDto *dto.CreateExtIdpDto) *dto.ExtIdpSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-ext-idp", fasthttp.MethodPost, reqDto)
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
 * @description 更新身份源配置，可以设置身份源 ID 与 名称。
 * @param requestBody
 * @returns ExtIdpSingleRespDto
 */
func (client *ManagementClient) UpdateExtIdp(reqDto *dto.UpdateExtIdpDto) *dto.ExtIdpSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-ext-idp", fasthttp.MethodPost, reqDto)
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
 * @description 通过身份源 ID，删除身份源。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteExtIdp(reqDto *dto.DeleteExtIdpDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-ext-idp", fasthttp.MethodPost, reqDto)
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
 * @description 在某个已有身份源下创建新连接，可以设置身份源图标、是否只支持登录等。
 * @param requestBody
 * @returns ExtIdpConnDetailSingleRespDto
 */
func (client *ManagementClient) CreateExtIdpConn(reqDto *dto.CreateExtIdpConnDto) *dto.ExtIdpConnDetailSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-ext-idp-conn", fasthttp.MethodPost, reqDto)
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
 * @description 更新身份源连接，可以设置身份源图标、是否只支持登录等。
 * @param requestBody
 * @returns ExtIdpConnDetailSingleRespDto
 */
func (client *ManagementClient) UpdateExtIdpConn(reqDto *dto.UpdateExtIdpConnDto) *dto.ExtIdpConnDetailSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-ext-idp-conn", fasthttp.MethodPost, reqDto)
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
 * @description 通过身份源连接 ID，删除身份源连接。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteExtIdpConn(reqDto *dto.DeleteExtIdpConnDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-ext-idp-conn", fasthttp.MethodPost, reqDto)
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
 * @description 身份源连接开关，可以打开或关闭身份源连接。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) ChangeConnState(reqDto *dto.EnableExtIdpConnDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/enable-ext-idp-conn", fasthttp.MethodPost, reqDto)
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
 * @summary 租户关联身份源
 * @description 租户可以关联或取消关联身份源连接。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) ChangeAssociationState(reqDto *dto.AssociationExtIdpDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/association-ext-idp", fasthttp.MethodPost, reqDto)
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
 * @summary 租户控制台获取身份源列表
 * @description 在租户控制台内获取身份源列表，可以根据 应用 ID 筛选。
 * @param tenantId 租户 ID
 * @param appId 应用 ID
 * @param type 身份源类型
 * @param page 页码
 * @param limit 每页获取的数据量
 * @returns ExtIdpListPaginatedRespDto
 */
func (client *ManagementClient) ListTenantExtIdp(reqDto *dto.ListTenantExtIdpDto) *dto.ExtIdpListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-tenant-ext-idp", fasthttp.MethodGet, reqDto)
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
 * @summary 身份源下应用的连接详情
 * @description 在身份源详情页获取应用的连接情况
 * @param id 身份源 ID
 * @param tenantId 租户 ID
 * @param appId 应用 ID
 * @param type 身份源类型
 * @returns ExtIdpListPaginatedRespDto
 */
func (client *ManagementClient) ExtIdpConnStateByApps(reqDto *dto.ExtIdpConnAppsDto) *dto.ExtIdpListPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/ext-idp-conn-apps", fasthttp.MethodGet, reqDto)
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
 * @summary 获取用户内置字段列表
 * @description 获取用户内置的字段列表
 * @returns CustomFieldListRespDto
 */
func (client *ManagementClient) GetUserBaseFields() *dto.CustomFieldListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-base-fields", fasthttp.MethodGet, nil)
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
 * @summary 修改用户内置字段配置
 * @description 修改用户内置字段配置，内置字段不允许修改数据类型、唯一性。
 * @param requestBody
 * @returns CustomFieldListRespDto
 */
func (client *ManagementClient) SetUserBaseFields(reqDto *dto.SetUserBaseFieldsReqDto) *dto.CustomFieldListRespDto {
	b, err := client.SendHttpRequest("/api/v3/set-user-base-fields", fasthttp.MethodPost, reqDto)
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
 * @summary 获取自定义字段列表
 * @description 通过主体类型，获取用户、部门或角色的自定义字段列表。
 * @param targetType 主体类型，目前支持用户、角色、分组、部门
 * @returns CustomFieldListRespDto
 */
func (client *ManagementClient) GetCustomFields(reqDto *dto.GetCustomFieldsDto) *dto.CustomFieldListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-custom-fields", fasthttp.MethodGet, reqDto)
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
 * @description 创建/修改用户、部门或角色自定义字段定义，如果传入的 key 不存在则创建，存在则更新。
 * @param requestBody
 * @returns CustomFieldListRespDto
 */
func (client *ManagementClient) SetCustomFields(reqDto *dto.SetCustomFieldsReqDto) *dto.CustomFieldListRespDto {
	b, err := client.SendHttpRequest("/api/v3/set-custom-fields", fasthttp.MethodPost, reqDto)
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
 * @description 给用户、角色或部门设置自定义字段的值，如果存在则更新，不存在则创建。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) SetCustomData(reqDto *dto.SetCustomDataReqDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/set-custom-data", fasthttp.MethodPost, reqDto)
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
 * @description 通过筛选条件，获取用户、分组、角色、组织机构的自定义字段值。
 * @param targetType 主体类型，目前支持用户、角色、分组、部门
 * @param targetIdentifier 目标对象唯一标志符
 * @param namespace 所属权限分组的 code，当 targetType 为角色的时候需要填写，否则可以忽略
 * @returns GetCustomDataRespDto
 */
func (client *ManagementClient) GetCustomData(reqDto *dto.GetCustomDataDto) *dto.GetCustomDataRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-custom-data", fasthttp.MethodGet, reqDto)
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
 * @summary 创建权限分组
 * @description 创建权限分组，可以设置分组名称与描述信息。
 * @param requestBody
 * @returns NamespaceRespDto
 */
func (client *ManagementClient) CreateNamespace(reqDto *dto.CreateNamespaceDto) *dto.NamespaceRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-namespace", fasthttp.MethodPost, reqDto)
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
 * @description 批量创建权限分组，可以分别设置分组名称与描述信息。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) CreateNamespacesBatch(reqDto *dto.CreateNamespacesBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-namespaces-batch", fasthttp.MethodPost, reqDto)
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
 * @description 通过权限分组唯一标志符，获取权限分组详情。
 * @param code 权限分组唯一标志符
 * @returns NamespaceRespDto
 */
func (client *ManagementClient) GetNamespace(reqDto *dto.GetNamespaceDto) *dto.NamespaceRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-namespace", fasthttp.MethodGet, reqDto)
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
 * @description 分别通过权限分组唯一标志符，批量获取权限分组详情。
 * @param codeList 资源 code 列表，批量可以使用逗号分隔
 * @returns NamespaceListRespDto
 */
func (client *ManagementClient) GetNamespacesBatch(reqDto *dto.GetNamespacesBatchDto) *dto.NamespaceListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-namespaces-batch", fasthttp.MethodGet, reqDto)
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
 * @description 修改权限分组信息，可以修改名称、描述信息以及新的唯一标志符。
 * @param requestBody
 * @returns UpdateNamespaceRespDto
 */
func (client *ManagementClient) UpdateNamespace(reqDto *dto.UpdateNamespaceDto) *dto.UpdateNamespaceRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-namespace", fasthttp.MethodPost, reqDto)
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
 * @description 通过权限分组唯一标志符，删除权限分组信息。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteNamespace(reqDto *dto.DeleteNamespaceDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-namespace", fasthttp.MethodPost, reqDto)
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
 * @description 分别通过权限分组唯一标志符，批量删除权限分组。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteNamespacesBatch(reqDto *dto.DeleteNamespacesBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-namespaces-batch", fasthttp.MethodPost, reqDto)
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
 * @description 将一个/多个资源授权给用户、角色、分组、组织机构等主体，且可以分别指定不同的操作权限。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) AuthorizeResources(reqDto *dto.AuthorizeResourcesDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/authorize-resources", fasthttp.MethodPost, reqDto)
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
 * @description 根据筛选条件，获取某个主体被授权的资源列表。
 * @param targetType 目标对象类型
 * @param targetIdentifier 目标对象唯一标志符
 * @param namespace 所属权限分组的 code
 * @param resourceType 限定资源类型，如数据、API、按钮、菜单
 * @param resourceList 限定查询的资源列表，如果指定，只会返回所指定的资源列表。
 * @param withDenied 是否获取被拒绝的资源
 * @returns AuthorizedResourcePaginatedRespDto
 */
func (client *ManagementClient) GetAuthorizedResources(reqDto *dto.GetAuthorizedResourcesDto) *dto.AuthorizedResourcePaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-authorized-resources", fasthttp.MethodGet, reqDto)
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
 * @summary 判断用户是否对某个资源的某个操作有权限
 * @description 判断用户是否对某个资源的某个操作有权限。
 * @param requestBody
 * @returns IsActionAllowedRespDtp
 */
func (client *ManagementClient) IsActionAllowed(reqDto *dto.IsActionAllowedDto) *dto.IsActionAllowedRespDtp {
	b, err := client.SendHttpRequest("/api/v3/is-action-allowed", fasthttp.MethodPost, reqDto)
	var response dto.IsActionAllowedRespDtp
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
 * @summary 获取资源被授权的主体
 * @description 获取资源被授权的主体
 * @param requestBody
 * @returns GetAuthorizedTargetRespDto
 */
func (client *ManagementClient) GetAuthorizedTargets(reqDto *dto.GetAuthorizedTargetsDto) *dto.GetAuthorizedTargetRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-authorized-targets", fasthttp.MethodPost, reqDto)
	var response dto.GetAuthorizedTargetRespDto
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
 * @summary 获取同步任务详情
 * @description 获取同步任务详情
 * @param syncTaskId 同步任务 ID
 * @returns SyncTaskSingleRespDto
 */
func (client *ManagementClient) GetSyncTask(reqDto *dto.GetSyncTaskDto) *dto.SyncTaskSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-sync-task", fasthttp.MethodGet, reqDto)
	var response dto.SyncTaskSingleRespDto
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
 * @summary 获取同步任务列表
 * @description 获取同步任务列表
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns SyncTaskPaginatedRespDto
 */
func (client *ManagementClient) ListSyncTasks(reqDto *dto.ListSyncTasksDto) *dto.SyncTaskPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-sync-tasks", fasthttp.MethodGet, reqDto)
	var response dto.SyncTaskPaginatedRespDto
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
 * @summary 创建同步任务
 * @description 创建同步任务
 * @param requestBody
 * @returns SyncTaskPaginatedRespDto
 */
func (client *ManagementClient) CreateSyncTask(reqDto *dto.CreateSyncTaskDto) *dto.SyncTaskPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-sync-task", fasthttp.MethodPost, reqDto)
	var response dto.SyncTaskPaginatedRespDto
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
 * @summary 修改同步任务
 * @description 修改同步任务
 * @param requestBody
 * @returns SyncTaskPaginatedRespDto
 */
func (client *ManagementClient) UpdateSyncTask(reqDto *dto.UpdateSyncTaskDto) *dto.SyncTaskPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-sync-task", fasthttp.MethodPost, reqDto)
	var response dto.SyncTaskPaginatedRespDto
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
 * @summary 执行同步任务
 * @description 执行同步任务
 * @param requestBody
 * @returns TriggerSyncTaskRespDto
 */
func (client *ManagementClient) TriggerSyncTask(reqDto *dto.TriggerSyncTaskDto) *dto.TriggerSyncTaskRespDto {
	b, err := client.SendHttpRequest("/api/v3/trigger-sync-task", fasthttp.MethodPost, reqDto)
	var response dto.TriggerSyncTaskRespDto
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
 * @summary 获取同步作业详情
 * @description 获取同步作业详情
 * @param syncJobId 同步作业 ID
 * @returns SyncJobSingleRespDto
 */
func (client *ManagementClient) GetSyncJob(reqDto *dto.GetSyncJobDto) *dto.SyncJobSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-sync-job", fasthttp.MethodGet, reqDto)
	var response dto.SyncJobSingleRespDto
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
 * @summary 获取同步作业详情
 * @description 获取同步作业详情
 * @param syncTaskId 同步任务 ID
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param syncTrigger 同步任务触发类型：
 * - `manually`: 手动触发执行
 * - `timed`: 定时触发
 * - `automatic`: 根据事件自动触发
 *
 * @returns SyncJobPaginatedRespDto
 */
func (client *ManagementClient) ListSyncJobs(reqDto *dto.ListSyncJobsDto) *dto.SyncJobPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-sync-jobs", fasthttp.MethodGet, reqDto)
	var response dto.SyncJobPaginatedRespDto
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
 * @summary 获取同步作业详情
 * @description 获取同步作业详情
 * @param syncJobId 同步作业 ID
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param success 根据是否操作成功进行筛选
 * @param action 根据操作类型进行筛选：
 * - `CreateUser`: 创建用户
 * - `UpdateUser`: 修改用户信息
 * - `DeleteUser`: 删除用户
 * - `UpdateUserIdentifier`: 修改用户唯一标志符
 * - `ChangeUserDepartment`: 修改用户部门
 * - `CreateDepartment`: 创建部门
 * - `UpdateDepartment`: 修改部门信息
 * - `DeleteDepartment`: 删除部门
 * - `MoveDepartment`: 移动部门
 * - `UpdateDepartmentLeader`: 同步部门负责人
 * - `CreateGroup`: 创建分组
 * - `UpdateGroup`: 修改分组
 * - `DeleteGroup`: 删除分组
 * - `Updateless`: 无更新
 *
 * @param objectType 操作对象类型:
 * - `department`: 部门
 * - `user`: 用户
 *
 * @returns TriggerSyncTaskRespDto
 */
func (client *ManagementClient) ListSyncJobLogs(reqDto *dto.ListSyncJobLogsDto) *dto.TriggerSyncTaskRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-sync-job-logs", fasthttp.MethodGet, reqDto)
	var response dto.TriggerSyncTaskRespDto
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
 * @summary 获取同步风险操作列表
 * @description 获取同步风险操作列表
 * @param syncTaskId 同步任务 ID
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param status 根据执行状态筛选
 * @param objectType 根据操作对象类型，默认获取所有类型的记录：
 * - `department`: 部门
 * - `user`: 用户
 *
 * @returns SyncRiskOperationPaginatedRespDto
 */
func (client *ManagementClient) ListSyncRiskOperations(reqDto *dto.ListSyncRiskOperationsDto) *dto.SyncRiskOperationPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-sync-risk-operations", fasthttp.MethodGet, reqDto)
	var response dto.SyncRiskOperationPaginatedRespDto
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
 * @summary 执行同步风险操作
 * @description 执行同步风险操作
 * @param requestBody
 * @returns TriggerSyncRiskOperationsRespDto
 */
func (client *ManagementClient) TriggerSyncRiskOperations(reqDto *dto.TriggerSyncRiskOperationDto) *dto.TriggerSyncRiskOperationsRespDto {
	b, err := client.SendHttpRequest("/api/v3/trigger-sync-risk-operations", fasthttp.MethodPost, reqDto)
	var response dto.TriggerSyncRiskOperationsRespDto
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
 * @summary 取消同步风险操作
 * @description 取消同步风险操作
 * @param requestBody
 * @returns CancelSyncRiskOperationsRespDto
 */
func (client *ManagementClient) CancelSyncRiskOperation(reqDto *dto.CancelSyncRiskOperationDto) *dto.CancelSyncRiskOperationsRespDto {
	b, err := client.SendHttpRequest("/api/v3/cancel-sync-risk-operation", fasthttp.MethodPost, reqDto)
	var response dto.CancelSyncRiskOperationsRespDto
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
 * @summary 获取用户行为日志
 * @description 可以选择请求 ID、客户端 IP、用户 ID、应用 ID、开始时间戳、请求是否成功、分页参数去获取用户行为日志
 * @param requestBody
 * @returns UserActionLogRespDto
 */
func (client *ManagementClient) GetUserActionLogs(reqDto *dto.GetUserActionLogsDto) *dto.UserActionLogRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-user-action-logs", fasthttp.MethodPost, reqDto)
	var response dto.UserActionLogRespDto
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
 * @summary 获取管理员操作日志
 * @description 可以选择请求 ID、客户端 IP、操作类型、资源类型、管理员用户 ID、请求是否成功、开始时间戳、结束时间戳、分页来获取管理员操作日志接口
 * @param requestBody
 * @returns AdminAuditLogRespDto
 */
func (client *ManagementClient) GetAdminAuditLogs(reqDto *dto.GetAdminAuditLogsDto) *dto.AdminAuditLogRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-admin-audit-logs", fasthttp.MethodPost, reqDto)
	var response dto.AdminAuditLogRespDto
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
 * @summary 获取邮件模版列表
 * @description 获取邮件模版列表
 * @returns GetEmailTemplatesRespDto
 */
func (client *ManagementClient) GetEmailTemplates() *dto.GetEmailTemplatesRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-email-templates", fasthttp.MethodGet, nil)
	var response dto.GetEmailTemplatesRespDto
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
 * @summary 修改邮件模版
 * @description 修改邮件模版
 * @param requestBody
 * @returns EmailTemplateSingleItemRespDto
 */
func (client *ManagementClient) UpdateEmailTemplate(reqDto *dto.UpdateEmailTemplateDto) *dto.EmailTemplateSingleItemRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-email-template", fasthttp.MethodPost, reqDto)
	var response dto.EmailTemplateSingleItemRespDto
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
 * @summary 预览邮件模版
 * @description 预览邮件模版
 * @param requestBody
 * @returns PreviewEmailTemplateRespDto
 */
func (client *ManagementClient) PreviewEmailTemplate(reqDto *dto.PreviewEmailTemplateDto) *dto.PreviewEmailTemplateRespDto {
	b, err := client.SendHttpRequest("/api/v3/preview-email-template", fasthttp.MethodPost, reqDto)
	var response dto.PreviewEmailTemplateRespDto
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
 * @summary 获取第三方邮件服务配置
 * @description 获取第三方邮件服务配置
 * @returns EmailProviderDto
 */
func (client *ManagementClient) GetEmailProvider() *dto.EmailProviderDto {
	b, err := client.SendHttpRequest("/api/v3/get-email-provier", fasthttp.MethodGet, nil)
	var response dto.EmailProviderDto
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
 * @summary 配置第三方邮件服务
 * @description 配置第三方邮件服务
 * @param requestBody
 * @returns EmailProviderDto
 */
func (client *ManagementClient) ConfigEmailProvider(reqDto *dto.ConfigEmailProviderDto) *dto.EmailProviderDto {
	b, err := client.SendHttpRequest("/api/v3/config-email-provier", fasthttp.MethodPost, reqDto)
	var response dto.EmailProviderDto
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
 * @summary 获取应用详情
 * @description 通过应用 ID，获取应用详情。
 * @param appId 应用 ID
 * @returns ApplicationSingleRespDto
 */
func (client *ManagementClient) GetApplication(reqDto *dto.GetApplicationDto) *dto.ApplicationSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-application", fasthttp.MethodGet, reqDto)
	var response dto.ApplicationSingleRespDto
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
 * @summary 获取应用列表
 * @description 获取应用列表
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param isIntegrateApp 是否为集成应用
 * @param isSelfBuiltApp 是否为自建应用
 * @param ssoEnabled 是否开启单点登录
 * @param keyword 模糊搜索字符串
 * @returns ApplicationPaginatedRespDto
 */
func (client *ManagementClient) ListApplications(reqDto *dto.ListApplicationsDto) *dto.ApplicationPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-applications", fasthttp.MethodGet, reqDto)
	var response dto.ApplicationPaginatedRespDto
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
 * @summary 获取应用简单信息
 * @description 通过应用 ID，获取应用简单信息。
 * @param appId 应用 ID
 * @returns ApplicationSimpleInfoSingleRespDto
 */
func (client *ManagementClient) GetApplicationSimpleInfo(reqDto *dto.GetApplicationSimpleInfoDto) *dto.ApplicationSimpleInfoSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-application-simple-info", fasthttp.MethodGet, reqDto)
	var response dto.ApplicationSimpleInfoSingleRespDto
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
 * @summary 获取应用简单信息列表
 * @description 获取应用简单信息列表
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @param isIntegrateApp 是否为集成应用
 * @param isSelfBuiltApp 是否为自建应用
 * @param ssoEnabled 是否开启单点登录
 * @param keyword 模糊搜索字符串
 * @returns ApplicationSimpleInfoSingleRespDto
 */
func (client *ManagementClient) ListApplicationSimpleInfo(reqDto *dto.ListApplicationSimpleInfoDto) *dto.ApplicationSimpleInfoSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-application-simple-info", fasthttp.MethodGet, reqDto)
	var response dto.ApplicationSimpleInfoSingleRespDto
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
 * @summary 创建应用
 * @description 创建应用
 * @param requestBody
 * @returns ApplicationPaginatedRespDto
 */
func (client *ManagementClient) CreateApplication(reqDto *dto.CreateApplicationDto) *dto.ApplicationPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-application", fasthttp.MethodPost, reqDto)
	var response dto.ApplicationPaginatedRespDto
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
 * @summary 删除应用
 * @description 通过应用 ID，删除应用。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteApplication(reqDto *dto.DeleteApplicationDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-application", fasthttp.MethodPost, reqDto)
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
 * @summary 获取应用密钥
 * @description 获取应用密钥
 * @param appId 应用 ID
 * @returns GetApplicationSecretRespDto
 */
func (client *ManagementClient) GetApplicationSecret(reqDto *dto.GetApplicationSecretDto) *dto.GetApplicationSecretRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-application-secret", fasthttp.MethodGet, reqDto)
	var response dto.GetApplicationSecretRespDto
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
 * @summary 刷新应用密钥
 * @description 刷新应用密钥
 * @param requestBody
 * @returns RefreshApplicationSecretRespDto
 */
func (client *ManagementClient) RefreshApplicationSecret(reqDto *dto.RefreshApplicationSecretDto) *dto.RefreshApplicationSecretRespDto {
	b, err := client.SendHttpRequest("/api/v3/refresh-application-secret", fasthttp.MethodPost, reqDto)
	var response dto.RefreshApplicationSecretRespDto
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
 * @summary 获取应用当前登录用户
 * @description 获取应用当前处于登录状态的用户
 * @param requestBody
 * @returns UserPaginatedRespDto
 */
func (client *ManagementClient) ListApplicationActiveUsers(reqDto *dto.ListApplicationActiveUsersDto) *dto.UserPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-application-active-users", fasthttp.MethodPost, reqDto)
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
 * @summary 获取应用默认访问授权策略
 * @description 获取应用默认访问授权策略
 * @param appId 应用 ID
 * @returns GetApplicationPermissionStrategyRespDto
 */
func (client *ManagementClient) GetApplicationPermissionStrategy(reqDto *dto.GetApplicationPermissionStrategyDto) *dto.GetApplicationPermissionStrategyRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-application-permission-strategy", fasthttp.MethodGet, reqDto)
	var response dto.GetApplicationPermissionStrategyRespDto
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
 * @summary 更新应用默认访问授权策略
 * @description 更新应用默认访问授权策略
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) UpdateApplicationPermissionStrategy(reqDto *dto.UpdateApplicationPermissionStrategyDataDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-application-permission-strategy", fasthttp.MethodPost, reqDto)
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
 * @summary 授权应用访问权限
 * @description 给用户、分组、组织或角色授权应用访问权限
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) AuthorizeApplicationAccess(reqDto *dto.AddApplicationPermissionRecord) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/add-application-permission-record", fasthttp.MethodPost, reqDto)
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
 * @summary 删除应用访问授权记录
 * @description 取消给用户、分组、组织或角色的应用访问权限授权
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) RevokeApplicationAccess(reqDto *dto.DeleteApplicationPermissionRecord) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-application-permission-record", fasthttp.MethodPost, reqDto)
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
 * @summary 检测域名是否可用
 * @description 检测域名是否可用于创建新应用或更新应用域名
 * @param requestBody
 * @returns CheckDomainAvailableSecretRespDto
 */
func (client *ManagementClient) CheckDomainAvailable(reqDto *dto.CheckDomainAvailable) *dto.CheckDomainAvailableSecretRespDto {
	b, err := client.SendHttpRequest("/api/v3/check-domain-available", fasthttp.MethodPost, reqDto)
	var response dto.CheckDomainAvailableSecretRespDto
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
 * @summary 获取安全配置
 * @description 无需传参获取安全配置
 * @returns SecuritySettingsRespDto
 */
func (client *ManagementClient) GetSecuritySettings() *dto.SecuritySettingsRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-security-settings", fasthttp.MethodGet, nil)
	var response dto.SecuritySettingsRespDto
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
 * @summary 修改安全配置
 * @description 可选安全域、Authing Token 有效时间（秒）、验证码长度、验证码尝试次数、用户修改邮箱的安全策略、用户修改手机号的安全策略、Cookie 过期时间设置、是否禁止用户注册、频繁注册检测配置、验证码注册后是否要求用户设置密码、未验证的邮箱登录时是否禁止登录并发送认证邮件、用户自助解锁配置、Authing 登录页面是否开启登录账号选择、APP 扫码登录安全配置进行修改安全配置
 * @param requestBody
 * @returns SecuritySettingsRespDto
 */
func (client *ManagementClient) UpdateSecuritySettings(reqDto *dto.UpdateSecuritySettingsDto) *dto.SecuritySettingsRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-security-settings", fasthttp.MethodPost, reqDto)
	var response dto.SecuritySettingsRespDto
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
 * @summary 获取全局多因素认证配置
 * @description 无需传参获取全局多因素认证配置
 * @returns MFASettingsRespDto
 */
func (client *ManagementClient) GetGlobalMfaSettings() *dto.MFASettingsRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-global-mfa-settings", fasthttp.MethodGet, nil)
	var response dto.MFASettingsRespDto
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
 * @summary 修改全局多因素认证配置
 * @description 传入 MFA 认证因素列表进行修改
 * @param requestBody
 * @returns MFASettingsRespDto
 */
func (client *ManagementClient) UpdateGlobalMfaSettings(reqDto *dto.MFASettingsDto) *dto.MFASettingsRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-global-mfa-settings", fasthttp.MethodPost, reqDto)
	var response dto.MFASettingsRespDto
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
 * @description 创建资源，可以设置资源的描述、定义的操作类型、URL 标识等。
 * @param requestBody
 * @returns ResourceRespDto
 */
func (client *ManagementClient) CreateResource(reqDto *dto.CreateResourceDto) *dto.ResourceRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-resource", fasthttp.MethodPost, reqDto)
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
 * @description 批量创建资源，可以设置资源的描述、定义的操作类型、URL 标识等。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) CreateResourcesBatch(reqDto *dto.CreateResourcesBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-resources-batch", fasthttp.MethodPost, reqDto)
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
 * @description 根据筛选条件，获取资源详情。
 * @param code 资源唯一标志符
 * @param namespace 所属权限分组的 code
 * @returns ResourceRespDto
 */
func (client *ManagementClient) GetResource(reqDto *dto.GetResourceDto) *dto.ResourceRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-resource", fasthttp.MethodGet, reqDto)
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
 * @description 根据筛选条件，批量获取资源详情。
 * @param codeList 资源 code 列表，批量可以使用逗号分隔
 * @param namespace 所属权限分组的 code
 * @returns ResourceListRespDto
 */
func (client *ManagementClient) GetResourcesBatch(reqDto *dto.GetResourcesBatchDto) *dto.ResourceListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-resources-batch", fasthttp.MethodGet, reqDto)
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
 * @description 根据筛选条件，分页获取资源详情列表。
 * @param namespace 所属权限分组的 code
 * @param type 资源类型
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns ResourcePaginatedRespDto
 */
func (client *ManagementClient) ListResources(reqDto *dto.ListResourcesDto) *dto.ResourcePaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-resources", fasthttp.MethodGet, reqDto)
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
 * @description 修改资源，可以设置资源的描述、定义的操作类型、URL 标识等。
 * @param requestBody
 * @returns ResourceRespDto
 */
func (client *ManagementClient) UpdateResource(reqDto *dto.UpdateResourceDto) *dto.ResourceRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-resource", fasthttp.MethodPost, reqDto)
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
 * @description 通过资源唯一标志符以及所属权限分组，删除资源。
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteResource(reqDto *dto.DeleteResourceDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-resource", fasthttp.MethodPost, reqDto)
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
 * @description 通过资源唯一标志符以及所属权限分组，批量删除资源
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) DeleteResourcesBatch(reqDto *dto.DeleteResourcesBatchDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-resources-batch", fasthttp.MethodPost, reqDto)
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
 * @summary 关联/取消关联应用资源到租户
 * @description 通过资源唯一标识以及权限分组，关联或取消关联资源到租户
 * @param requestBody
 * @returns IsSuccessRespDto
 */
func (client *ManagementClient) AssociationResources(reqDto *dto.AssociationResourceDto) *dto.IsSuccessRespDto {
	b, err := client.SendHttpRequest("/api/v3/associate-tenant-resource", fasthttp.MethodPost, reqDto)
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
 * @summary 创建 Pipeline 函数
 * @description 创建 Pipeline 函数
 * @param requestBody
 * @returns PipelineFunctionSingleRespDto
 */
func (client *ManagementClient) CreatePipelineFunction(reqDto *dto.CreatePipelineFunctionDto) *dto.PipelineFunctionSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-pipeline-function", fasthttp.MethodPost, reqDto)
	var response dto.PipelineFunctionSingleRespDto
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
 * @summary 获取 Pipeline 函数详情
 * @description 获取 Pipeline 函数详情
 * @param funcId Pipeline 函数 ID
 * @returns PipelineFunctionSingleRespDto
 */
func (client *ManagementClient) GetPipelineFunction(reqDto *dto.GetPipelineFunctionDto) *dto.PipelineFunctionSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-pipeline-function", fasthttp.MethodGet, reqDto)
	var response dto.PipelineFunctionSingleRespDto
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
 * @summary 重新上传 Pipeline 函数
 * @description 当 Pipeline 函数上传失败时，重新上传 Pipeline 函数
 * @param requestBody
 * @returns PipelineFunctionSingleRespDto
 */
func (client *ManagementClient) ReuploadPipelineFunction(reqDto *dto.ReUploadPipelineFunctionDto) *dto.PipelineFunctionSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/reupload-pipeline-function", fasthttp.MethodPost, reqDto)
	var response dto.PipelineFunctionSingleRespDto
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
 * @summary 修改 Pipeline 函数
 * @description 修改 Pipeline 函数
 * @param requestBody
 * @returns PipelineFunctionSingleRespDto
 */
func (client *ManagementClient) UpdatePipelineFunction(reqDto *dto.UpdatePipelineFunctionDto) *dto.PipelineFunctionSingleRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-pipeline-function", fasthttp.MethodPost, reqDto)
	var response dto.PipelineFunctionSingleRespDto
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
 * @summary 修改 Pipeline 函数顺序
 * @description 修改 Pipeline 函数顺序
 * @param requestBody
 * @returns CommonResponseDto
 */
func (client *ManagementClient) UpdatePipelineOrder(reqDto *dto.UpdatePipelineOrderDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/update-pipeline-order", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
 * @summary 删除 Pipeline 函数
 * @description 删除 Pipeline 函数
 * @param requestBody
 * @returns CommonResponseDto
 */
func (client *ManagementClient) DeletePipelineFunction(reqDto *dto.DeletePipelineFunctionDto) *dto.CommonResponseDto {
	b, err := client.SendHttpRequest("/api/v3/delete-pipeline-function", fasthttp.MethodPost, reqDto)
	var response dto.CommonResponseDto
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
 * @summary 获取 Pipeline 函数列表
 * @description 获取 Pipeline 函数列表
 * @param scene 通过函数的触发场景进行筛选（可选，默认返回所有）：
 * - `PRE_REGISTER`: 注册前
 * - `POST_REGISTER`: 注册后
 * - `PRE_AUTHENTICATION`: 认证前
 * - `POST_AUTHENTICATION`: 认证后
 * - `PRE_OIDC_ID_TOKEN_ISSUED`: OIDC ID Token 签发前
 * - `PRE_OIDC_ACCESS_TOKEN_ISSUED`: OIDC Access Token 签发前
 * - `PRE_COMPLETE_USER_INFO`: 补全用户信息前
 *
 * @returns PipelineFunctionPaginatedRespDto
 */
func (client *ManagementClient) ListPipelineFunctions(reqDto *dto.ListPipelineFunctionDto) *dto.PipelineFunctionPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-pipeline-function", fasthttp.MethodGet, reqDto)
	var response dto.PipelineFunctionPaginatedRespDto
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
 * @summary 获取 Pipeline 日志
 * @description 获取 Pipeline
 * @param funcId Pipeline 函数 ID
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns PipelineFunctionPaginatedRespDto
 */
func (client *ManagementClient) GetPipelineLogs(reqDto *dto.GetPipelineLogsDto) *dto.PipelineFunctionPaginatedRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-pipeline-logs", fasthttp.MethodGet, reqDto)
	var response dto.PipelineFunctionPaginatedRespDto
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
 * @summary 创建 Webhook
 * @description 你需要指定 Webhoook 名称、Webhook 回调地址、请求数据格式、用户真实名称来创建 Webhook。还可选是否启用、请求密钥进行创建
 * @param requestBody
 * @returns CreateWebhookRespDto
 */
func (client *ManagementClient) CreateWebhook(reqDto *dto.CreateWebhookDto) *dto.CreateWebhookRespDto {
	b, err := client.SendHttpRequest("/api/v3/create-webhook", fasthttp.MethodPost, reqDto)
	var response dto.CreateWebhookRespDto
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
 * @summary 获取 Webhook 列表
 * @description 获取 Webhook 列表，可选页数、分页大小来获取
 * @param page 当前页数，从 1 开始
 * @param limit 每页数目，最大不能超过 50，默认为 10
 * @returns GetWebhooksRespDto
 */
func (client *ManagementClient) ListWebhooks(reqDto *dto.ListWebhooksDto) *dto.GetWebhooksRespDto {
	b, err := client.SendHttpRequest("/api/v3/list-webhooks", fasthttp.MethodGet, reqDto)
	var response dto.GetWebhooksRespDto
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
 * @summary 修改 Webhook 配置
 * @description 需要指定 webhookId，可选 Webhoook 名称、Webhook 回调地址、请求数据格式、用户真实名称、是否启用、请求密钥参数进行修改 webhook
 * @param requestBody
 * @returns UpdateWebhooksRespDto
 */
func (client *ManagementClient) UpdateWebhook(reqDto *dto.UpdateWebhookDto) *dto.UpdateWebhooksRespDto {
	b, err := client.SendHttpRequest("/api/v3/update-webhook", fasthttp.MethodPost, reqDto)
	var response dto.UpdateWebhooksRespDto
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
 * @summary 删除 Webhook
 * @description 通过指定多个 webhookId，以数组的形式进行 webhook 的删除
 * @param requestBody
 * @returns DeleteWebhookRespDto
 */
func (client *ManagementClient) DeleteWebhook(reqDto *dto.DeleteWebhookDto) *dto.DeleteWebhookRespDto {
	b, err := client.SendHttpRequest("/api/v3/delete-webhook", fasthttp.MethodPost, reqDto)
	var response dto.DeleteWebhookRespDto
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
 * @summary 获取 Webhook 日志
 * @description 通过指定 webhookId，可选 page 和 limit 来获取 webhook 日志
 * @param requestBody
 * @returns ListWebhookLogsRespDto
 */
func (client *ManagementClient) GetWebhookLogs(reqDto *dto.ListWebhookLogs) *dto.ListWebhookLogsRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-webhook-logs", fasthttp.MethodPost, reqDto)
	var response dto.ListWebhookLogsRespDto
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
 * @summary 手动触发 Webhook 执行
 * @description 通过指定 webhookId，可选请求头和请求体进行手动触发 webhook 执行
 * @param requestBody
 * @returns TriggerWebhookRespDto
 */
func (client *ManagementClient) TriggerWebhook(reqDto *dto.TriggerWebhookDto) *dto.TriggerWebhookRespDto {
	b, err := client.SendHttpRequest("/api/v3/trigger-webhook", fasthttp.MethodPost, reqDto)
	var response dto.TriggerWebhookRespDto
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
 * @summary 获取 Webhook 详情
 * @description 根据指定的 webhookId 获取 webhook 详情
 * @param webhookId Webhook ID
 * @returns GetWebhookRespDto
 */
func (client *ManagementClient) GetWebhook(reqDto *dto.GetWebhookDto) *dto.GetWebhookRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-webhook", fasthttp.MethodGet, reqDto)
	var response dto.GetWebhookRespDto
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
 * @summary 获取 Webhook 事件列表
 * @description 返回事件列表和分类列表
 * @returns WebhookEventListRespDto
 */
func (client *ManagementClient) GetWebhookEventList() *dto.WebhookEventListRespDto {
	b, err := client.SendHttpRequest("/api/v3/get-webhook-event-list", fasthttp.MethodGet, nil)
	var response dto.WebhookEventListRespDto
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
