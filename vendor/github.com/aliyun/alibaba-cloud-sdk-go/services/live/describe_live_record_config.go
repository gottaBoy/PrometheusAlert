package live

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeLiveRecordConfig invokes the live.DescribeLiveRecordConfig API synchronously
// api document: https://help.aliyun.com/api/live/describeliverecordconfig.html
func (client *Client) DescribeLiveRecordConfig(request *DescribeLiveRecordConfigRequest) (response *DescribeLiveRecordConfigResponse, err error) {
	response = CreateDescribeLiveRecordConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveRecordConfigWithChan invokes the live.DescribeLiveRecordConfig API asynchronously
// api document: https://help.aliyun.com/api/live/describeliverecordconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveRecordConfigWithChan(request *DescribeLiveRecordConfigRequest) (<-chan *DescribeLiveRecordConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveRecordConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveRecordConfig(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeLiveRecordConfigWithCallback invokes the live.DescribeLiveRecordConfig API asynchronously
// api document: https://help.aliyun.com/api/live/describeliverecordconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveRecordConfigWithCallback(request *DescribeLiveRecordConfigRequest, callback func(response *DescribeLiveRecordConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveRecordConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveRecordConfig(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeLiveRecordConfigRequest is the request struct for api DescribeLiveRecordConfig
type DescribeLiveRecordConfigRequest struct {
	*requests.RpcRequest
	PageNum       requests.Integer `position:"Query" name:"PageNum"`
	AppName       string           `position:"Query" name:"AppName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	StreamName    string           `position:"Query" name:"StreamName"`
	Order         string           `position:"Query" name:"Order"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveRecordConfigResponse is the response struct for api DescribeLiveRecordConfig
type DescribeLiveRecordConfigResponse struct {
	*responses.BaseResponse
	RequestId         string                                      `json:"RequestId" xml:"RequestId"`
	PageNum           int                                         `json:"PageNum" xml:"PageNum"`
	PageSize          int                                         `json:"PageSize" xml:"PageSize"`
	Order             string                                      `json:"Order" xml:"Order"`
	TotalNum          int                                         `json:"TotalNum" xml:"TotalNum"`
	TotalPage         int                                         `json:"TotalPage" xml:"TotalPage"`
	LiveAppRecordList LiveAppRecordListInDescribeLiveRecordConfig `json:"LiveAppRecordList" xml:"LiveAppRecordList"`
}

// CreateDescribeLiveRecordConfigRequest creates a request to invoke DescribeLiveRecordConfig API
func CreateDescribeLiveRecordConfigRequest() (request *DescribeLiveRecordConfigRequest) {
	request = &DescribeLiveRecordConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveRecordConfig", "live", "openAPI")
	return
}

// CreateDescribeLiveRecordConfigResponse creates a response to parse from DescribeLiveRecordConfig response
func CreateDescribeLiveRecordConfigResponse() (response *DescribeLiveRecordConfigResponse) {
	response = &DescribeLiveRecordConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
