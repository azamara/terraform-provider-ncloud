/*
 * monitoring
 *
 * <br/>https://ncloud.apigw.ntruss.com/monitoring/v2
 *
 * API version: 2018-06-25T02:38:27Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package monitoring

type GetListMetricsResponse struct {

RequestId *string `json:"requestId,omitempty"`

ReturnCode *string `json:"returnCode,omitempty"`

ReturnMessage *string `json:"returnMessage,omitempty"`

Metrics []*Metric `json:"metrics,omitempty"`
}
