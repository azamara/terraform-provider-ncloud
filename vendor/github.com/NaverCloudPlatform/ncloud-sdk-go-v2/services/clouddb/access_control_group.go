/*
 * clouddb
 *
 * Cloud DB<br/>https://ncloud.apigw.ntruss.com/clouddb/v2
 *
 * API version: 2018-11-01T03:57:11Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package clouddb

type AccessControlGroup struct {

	// 접근제어그룹설정번호
AccessControlGroupConfigurationNo *string `json:"accessControlGroupConfigurationNo,omitempty"`

	// 접근제어그룹명
AccessControlGroupName *string `json:"accessControlGroupName,omitempty"`

	// 접근제어그룹설명
AccessControlGroupDescription *string `json:"accessControlGroupDescription,omitempty"`

	// 디폴트그룹여부
IsDefault *bool `json:"isDefault,omitempty"`

	// 생성일자
CreateDate *string `json:"createDate,omitempty"`
}
