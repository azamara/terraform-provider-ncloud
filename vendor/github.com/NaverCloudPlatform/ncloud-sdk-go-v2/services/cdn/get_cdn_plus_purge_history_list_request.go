/*
 * cdn
 *
 * <br/>https://ncloud.apigw.ntruss.com/cdn/v2
 *
 * API version: 2018-08-07T06:43:44Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cdn

type GetCdnPlusPurgeHistoryListRequest struct {

	// CDN인스턴스번호
CdnInstanceNo *string `json:"cdnInstanceNo"`

	// 퍼지ID리스트
PurgeIdList []*string `json:"purgeIdList,omitempty"`
}
