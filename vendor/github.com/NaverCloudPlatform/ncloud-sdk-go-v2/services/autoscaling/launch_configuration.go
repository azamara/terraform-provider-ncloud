/*
 * autoscaling
 *
 * <br/>https://ncloud.apigw.ntruss.com/autoscaling/v2
 *
 * API version: 2018-08-07T06:47:31Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package autoscaling

type LaunchConfiguration struct {

	// 론치설정명
LaunchConfigurationName *string `json:"launchConfigurationName,omitempty"`

	// 서버이미지상품코드
ServerImageProductCode *string `json:"serverImageProductCode,omitempty"`

	// 서버상품코드
ServerProductCode *string `json:"serverProductCode,omitempty"`

	// 회원서버이미지번호
MemberServerImageNo *string `json:"memberServerImageNo,omitempty"`

	// 로그인키명
LoginKeyName *string `json:"loginKeyName,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate,omitempty"`

	// 사용자데이터
UserData *string `json:"userData,omitempty"`

	// ACG리스트
AccessControlGroupList []*AccessControlGroup `json:"accessControlGroupList,omitempty"`
}
