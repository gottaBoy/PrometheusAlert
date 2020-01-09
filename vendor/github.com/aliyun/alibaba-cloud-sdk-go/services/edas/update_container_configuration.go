package edas

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

// UpdateContainerConfiguration invokes the edas.UpdateContainerConfiguration API synchronously
// api document: https://help.aliyun.com/api/edas/updatecontainerconfiguration.html
func (client *Client) UpdateContainerConfiguration(request *UpdateContainerConfigurationRequest) (response *UpdateContainerConfigurationResponse, err error) {
	response = CreateUpdateContainerConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateContainerConfigurationWithChan invokes the edas.UpdateContainerConfiguration API asynchronously
// api document: https://help.aliyun.com/api/edas/updatecontainerconfiguration.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateContainerConfigurationWithChan(request *UpdateContainerConfigurationRequest) (<-chan *UpdateContainerConfigurationResponse, <-chan error) {
	responseChan := make(chan *UpdateContainerConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateContainerConfiguration(request)
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

// UpdateContainerConfigurationWithCallback invokes the edas.UpdateContainerConfiguration API asynchronously
// api document: https://help.aliyun.com/api/edas/updatecontainerconfiguration.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateContainerConfigurationWithCallback(request *UpdateContainerConfigurationRequest, callback func(response *UpdateContainerConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateContainerConfigurationResponse
		var err error
		defer close(result)
		response, err = client.UpdateContainerConfiguration(request)
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

// UpdateContainerConfigurationRequest is the request struct for api UpdateContainerConfiguration
type UpdateContainerConfigurationRequest struct {
	*requests.RoaRequest
	UseBodyEncoding requests.Boolean `position:"Query" name:"UseBodyEncoding"`
	MaxThreads      requests.Integer `position:"Query" name:"MaxThreads"`
	URIEncoding     string           `position:"Query" name:"URIEncoding"`
	AppId           string           `position:"Query" name:"AppId"`
	GroupId         string           `position:"Query" name:"GroupId"`
	HttpPort        requests.Integer `position:"Query" name:"HttpPort"`
	ContextPath     string           `position:"Query" name:"ContextPath"`
}

// UpdateContainerConfigurationResponse is the response struct for api UpdateContainerConfiguration
type UpdateContainerConfigurationResponse struct {
	*responses.BaseResponse
	Code                   int                    `json:"Code" xml:"Code"`
	Message                string                 `json:"Message" xml:"Message"`
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	ContainerConfiguration ContainerConfiguration `json:"ContainerConfiguration" xml:"ContainerConfiguration"`
}

// CreateUpdateContainerConfigurationRequest creates a request to invoke UpdateContainerConfiguration API
func CreateUpdateContainerConfigurationRequest() (request *UpdateContainerConfigurationRequest) {
	request = &UpdateContainerConfigurationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UpdateContainerConfiguration", "/pop/v5/app/container_config", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateContainerConfigurationResponse creates a response to parse from UpdateContainerConfiguration response
func CreateUpdateContainerConfigurationResponse() (response *UpdateContainerConfigurationResponse) {
	response = &UpdateContainerConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
