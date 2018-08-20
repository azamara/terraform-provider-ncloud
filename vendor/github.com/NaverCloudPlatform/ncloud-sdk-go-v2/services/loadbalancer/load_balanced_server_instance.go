/*
 * loadbalancer
 *
 * <br/>https://ncloud.apigw.ntruss.com/loadbalancer/v2
 *
 * API version: 2018-08-07T06:52:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package loadbalancer

type LoadBalancedServerInstance struct {

	// 서버인스턴스
ServerInstance *ServerInstance `json:"serverInstance,omitempty"`

	// 서버헬스체크상태리스트
ServerHealthCheckStatusList []*ServerHealthCheckStatus `json:"serverHealthCheckStatusList,omitempty"`
}
