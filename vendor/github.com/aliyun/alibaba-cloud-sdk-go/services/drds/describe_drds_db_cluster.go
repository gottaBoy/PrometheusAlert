package drds

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

// DescribeDrdsDBCluster invokes the drds.DescribeDrdsDBCluster API synchronously
// api document: https://help.aliyun.com/api/drds/describedrdsdbcluster.html
func (client *Client) DescribeDrdsDBCluster(request *DescribeDrdsDBClusterRequest) (response *DescribeDrdsDBClusterResponse, err error) {
	response = CreateDescribeDrdsDBClusterResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDrdsDBClusterWithChan invokes the drds.DescribeDrdsDBCluster API asynchronously
// api document: https://help.aliyun.com/api/drds/describedrdsdbcluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDrdsDBClusterWithChan(request *DescribeDrdsDBClusterRequest) (<-chan *DescribeDrdsDBClusterResponse, <-chan error) {
	responseChan := make(chan *DescribeDrdsDBClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDrdsDBCluster(request)
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

// DescribeDrdsDBClusterWithCallback invokes the drds.DescribeDrdsDBCluster API asynchronously
// api document: https://help.aliyun.com/api/drds/describedrdsdbcluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDrdsDBClusterWithCallback(request *DescribeDrdsDBClusterRequest, callback func(response *DescribeDrdsDBClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDrdsDBClusterResponse
		var err error
		defer close(result)
		response, err = client.DescribeDrdsDBCluster(request)
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

// DescribeDrdsDBClusterRequest is the request struct for api DescribeDrdsDBCluster
type DescribeDrdsDBClusterRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
	DbInstanceId   string `position:"Query" name:"DbInstanceId"`
}

// DescribeDrdsDBClusterResponse is the response struct for api DescribeDrdsDBCluster
type DescribeDrdsDBClusterResponse struct {
	*responses.BaseResponse
	RequestId  string                            `json:"RequestId" xml:"RequestId"`
	Success    bool                              `json:"Success" xml:"Success"`
	DbInstance DbInstanceInDescribeDrdsDBCluster `json:"DbInstance" xml:"DbInstance"`
}

// CreateDescribeDrdsDBClusterRequest creates a request to invoke DescribeDrdsDBCluster API
func CreateDescribeDrdsDBClusterRequest() (request *DescribeDrdsDBClusterRequest) {
	request = &DescribeDrdsDBClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeDrdsDBCluster", "Drds", "openAPI")
	return
}

// CreateDescribeDrdsDBClusterResponse creates a response to parse from DescribeDrdsDBCluster response
func CreateDescribeDrdsDBClusterResponse() (response *DescribeDrdsDBClusterResponse) {
	response = &DescribeDrdsDBClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
