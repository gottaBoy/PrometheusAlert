package mts

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

// QueryMCTemplateList invokes the mts.QueryMCTemplateList API synchronously
// api document: https://help.aliyun.com/api/mts/querymctemplatelist.html
func (client *Client) QueryMCTemplateList(request *QueryMCTemplateListRequest) (response *QueryMCTemplateListResponse, err error) {
	response = CreateQueryMCTemplateListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMCTemplateListWithChan invokes the mts.QueryMCTemplateList API asynchronously
// api document: https://help.aliyun.com/api/mts/querymctemplatelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMCTemplateListWithChan(request *QueryMCTemplateListRequest) (<-chan *QueryMCTemplateListResponse, <-chan error) {
	responseChan := make(chan *QueryMCTemplateListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMCTemplateList(request)
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

// QueryMCTemplateListWithCallback invokes the mts.QueryMCTemplateList API asynchronously
// api document: https://help.aliyun.com/api/mts/querymctemplatelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMCTemplateListWithCallback(request *QueryMCTemplateListRequest, callback func(response *QueryMCTemplateListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMCTemplateListResponse
		var err error
		defer close(result)
		response, err = client.QueryMCTemplateList(request)
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

// QueryMCTemplateListRequest is the request struct for api QueryMCTemplateList
type QueryMCTemplateListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TemplateIds          string           `position:"Query" name:"TemplateIds"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// QueryMCTemplateListResponse is the response struct for api QueryMCTemplateList
type QueryMCTemplateListResponse struct {
	*responses.BaseResponse
	RequestId    string                            `json:"RequestId" xml:"RequestId"`
	NonExistTids NonExistTidsInQueryMCTemplateList `json:"NonExistTids" xml:"NonExistTids"`
	TemplateList TemplateListInQueryMCTemplateList `json:"TemplateList" xml:"TemplateList"`
}

// CreateQueryMCTemplateListRequest creates a request to invoke QueryMCTemplateList API
func CreateQueryMCTemplateListRequest() (request *QueryMCTemplateListRequest) {
	request = &QueryMCTemplateListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryMCTemplateList", "mts", "openAPI")
	return
}

// CreateQueryMCTemplateListResponse creates a response to parse from QueryMCTemplateList response
func CreateQueryMCTemplateListResponse() (response *QueryMCTemplateListResponse) {
	response = &QueryMCTemplateListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
