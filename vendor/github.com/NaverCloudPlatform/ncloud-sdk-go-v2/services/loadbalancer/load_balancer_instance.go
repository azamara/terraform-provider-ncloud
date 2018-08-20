/*
 * loadbalancer
 *
 * <br/>https://ncloud.apigw.ntruss.com/loadbalancer/v2
 *
 * API version: 2018-08-07T06:52:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package loadbalancer

type LoadBalancerInstance struct {

	// 로드밸런서인스턴스번호
LoadBalancerInstanceNo *string `json:"loadBalancerInstanceNo,omitempty"`

	// virtualIp
VirtualIp *string `json:"virtualIp,omitempty"`

	// 로드밸런서명
LoadBalancerName *string `json:"loadBalancerName,omitempty"`

	// 로드밸런서알고리즘구분코
LoadBalancerAlgorithmType *CommonCode `json:"loadBalancerAlgorithmType,omitempty"`

	// 로드밸런서설명
LoadBalancerDescription *string `json:"loadBalancerDescription,omitempty"`

	// 생성일자
CreateDate *string `json:"createDate,omitempty"`

	// 도메인명
DomainName *string `json:"domainName,omitempty"`

	// 인터넷회선구분
InternetLineType *CommonCode `json:"internetLineType,omitempty"`

	// 로드밸런서인스턴스상태명
LoadBalancerInstanceStatusName *string `json:"loadBalancerInstanceStatusName,omitempty"`

	// 로드밸런서인스턴스상태
LoadBalancerInstanceStatus *CommonCode `json:"loadBalancerInstanceStatus,omitempty"`

	// 로드밸런서인스턴스OP
LoadBalancerInstanceOperation *CommonCode `json:"loadBalancerInstanceOperation,omitempty"`

	// 네트워크사용구분
NetworkUsageType *CommonCode `json:"networkUsageType,omitempty"`

	// httpKeepAlive사용여부
IsHttpKeepAlive *bool `json:"isHttpKeepAlive,omitempty"`

	// 커넥션타임아웃
ConnectionTimeout *int32 `json:"connectionTimeout,omitempty"`

	// SSL인증명
CertificateName *string `json:"certificateName,omitempty"`

LoadBalancerRuleList []*LoadBalancerRule `json:"loadBalancerRuleList,omitempty"`

LoadBalancedServerInstanceList []*LoadBalancedServerInstance `json:"loadBalancedServerInstanceList,omitempty"`
}
