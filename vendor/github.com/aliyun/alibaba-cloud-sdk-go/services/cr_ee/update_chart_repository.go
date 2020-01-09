package cr_ee

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

// UpdateChartRepository invokes the cr.UpdateChartRepository API synchronously
// api document: https://help.aliyun.com/api/cr/updatechartrepository.html
func (client *Client) UpdateChartRepository(request *UpdateChartRepositoryRequest) (response *UpdateChartRepositoryResponse, err error) {
	response = CreateUpdateChartRepositoryResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateChartRepositoryWithChan invokes the cr.UpdateChartRepository API asynchronously
// api document: https://help.aliyun.com/api/cr/updatechartrepository.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateChartRepositoryWithChan(request *UpdateChartRepositoryRequest) (<-chan *UpdateChartRepositoryResponse, <-chan error) {
	responseChan := make(chan *UpdateChartRepositoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateChartRepository(request)
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

// UpdateChartRepositoryWithCallback invokes the cr.UpdateChartRepository API asynchronously
// api document: https://help.aliyun.com/api/cr/updatechartrepository.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateChartRepositoryWithCallback(request *UpdateChartRepositoryRequest, callback func(response *UpdateChartRepositoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateChartRepositoryResponse
		var err error
		defer close(result)
		response, err = client.UpdateChartRepository(request)
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

// UpdateChartRepositoryRequest is the request struct for api UpdateChartRepository
type UpdateChartRepositoryRequest struct {
	*requests.RpcRequest
	RepoType          string `position:"Query" name:"RepoType"`
	Summary           string `position:"Query" name:"Summary"`
	InstanceId        string `position:"Query" name:"InstanceId"`
	RepoNamespaceName string `position:"Query" name:"RepoNamespaceName"`
	RepoName          string `position:"Query" name:"RepoName"`
}

// UpdateChartRepositoryResponse is the response struct for api UpdateChartRepository
type UpdateChartRepositoryResponse struct {
	*responses.BaseResponse
	UpdateChartRepositoryIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                           string `json:"Code" xml:"Code"`
	RequestId                      string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateChartRepositoryRequest creates a request to invoke UpdateChartRepository API
func CreateUpdateChartRepositoryRequest() (request *UpdateChartRepositoryRequest) {
	request = &UpdateChartRepositoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "UpdateChartRepository", "cr", "openAPI")
	return
}

// CreateUpdateChartRepositoryResponse creates a response to parse from UpdateChartRepository response
func CreateUpdateChartRepositoryResponse() (response *UpdateChartRepositoryResponse) {
	response = &UpdateChartRepositoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
