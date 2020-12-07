// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package lambdaiface provides an interface to enable mocking the AWS Lambda service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package lambdaiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// LambdaAPI provides an interface to enable mocking the
// lambda.Lambda service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Lambda.
//    func myFunc(svc lambdaiface.LambdaAPI) bool {
//        // Make svc.AddPermission request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := lambda.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockLambdaClient struct {
//        lambdaiface.LambdaAPI
//    }
//    func (m *mockLambdaClient) AddPermission(input *lambda.AddPermissionInput) (*lambda.AddPermissionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockLambdaClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type LambdaAPI interface {
	AddPermission(*lambda.AddPermissionInput) (*lambda.AddPermissionOutput, error)
	AddPermissionWithContext(aws.Context, *lambda.AddPermissionInput, ...request.Option) (*lambda.AddPermissionOutput, error)
	AddPermissionRequest(*lambda.AddPermissionInput) (*request.Request, *lambda.AddPermissionOutput)

	CreateAlias(*lambda.CreateAliasInput) (*lambda.AliasConfiguration, error)
	CreateAliasWithContext(aws.Context, *lambda.CreateAliasInput, ...request.Option) (*lambda.AliasConfiguration, error)
	CreateAliasRequest(*lambda.CreateAliasInput) (*request.Request, *lambda.AliasConfiguration)

	CreateEventSourceMapping(*lambda.CreateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
	CreateEventSourceMappingWithContext(aws.Context, *lambda.CreateEventSourceMappingInput, ...request.Option) (*lambda.EventSourceMappingConfiguration, error)
	CreateEventSourceMappingRequest(*lambda.CreateEventSourceMappingInput) (*request.Request, *lambda.EventSourceMappingConfiguration)

	CreateFunction(*lambda.CreateFunctionInput) (*lambda.FunctionConfiguration, error)
	CreateFunctionWithContext(aws.Context, *lambda.CreateFunctionInput, ...request.Option) (*lambda.FunctionConfiguration, error)
	CreateFunctionRequest(*lambda.CreateFunctionInput) (*request.Request, *lambda.FunctionConfiguration)

	DeleteAlias(*lambda.DeleteAliasInput) (*lambda.DeleteAliasOutput, error)
	DeleteAliasWithContext(aws.Context, *lambda.DeleteAliasInput, ...request.Option) (*lambda.DeleteAliasOutput, error)
	DeleteAliasRequest(*lambda.DeleteAliasInput) (*request.Request, *lambda.DeleteAliasOutput)

	DeleteEventSourceMapping(*lambda.DeleteEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
	DeleteEventSourceMappingWithContext(aws.Context, *lambda.DeleteEventSourceMappingInput, ...request.Option) (*lambda.EventSourceMappingConfiguration, error)
	DeleteEventSourceMappingRequest(*lambda.DeleteEventSourceMappingInput) (*request.Request, *lambda.EventSourceMappingConfiguration)

	DeleteFunction(*lambda.DeleteFunctionInput) (*lambda.DeleteFunctionOutput, error)
	DeleteFunctionWithContext(aws.Context, *lambda.DeleteFunctionInput, ...request.Option) (*lambda.DeleteFunctionOutput, error)
	DeleteFunctionRequest(*lambda.DeleteFunctionInput) (*request.Request, *lambda.DeleteFunctionOutput)

	DeleteFunctionConcurrency(*lambda.DeleteFunctionConcurrencyInput) (*lambda.DeleteFunctionConcurrencyOutput, error)
	DeleteFunctionConcurrencyWithContext(aws.Context, *lambda.DeleteFunctionConcurrencyInput, ...request.Option) (*lambda.DeleteFunctionConcurrencyOutput, error)
	DeleteFunctionConcurrencyRequest(*lambda.DeleteFunctionConcurrencyInput) (*request.Request, *lambda.DeleteFunctionConcurrencyOutput)

	GetAccountSettings(*lambda.GetAccountSettingsInput) (*lambda.GetAccountSettingsOutput, error)
	GetAccountSettingsWithContext(aws.Context, *lambda.GetAccountSettingsInput, ...request.Option) (*lambda.GetAccountSettingsOutput, error)
	GetAccountSettingsRequest(*lambda.GetAccountSettingsInput) (*request.Request, *lambda.GetAccountSettingsOutput)

	GetAlias(*lambda.GetAliasInput) (*lambda.AliasConfiguration, error)
	GetAliasWithContext(aws.Context, *lambda.GetAliasInput, ...request.Option) (*lambda.AliasConfiguration, error)
	GetAliasRequest(*lambda.GetAliasInput) (*request.Request, *lambda.AliasConfiguration)

	GetEventSourceMapping(*lambda.GetEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
	GetEventSourceMappingWithContext(aws.Context, *lambda.GetEventSourceMappingInput, ...request.Option) (*lambda.EventSourceMappingConfiguration, error)
	GetEventSourceMappingRequest(*lambda.GetEventSourceMappingInput) (*request.Request, *lambda.EventSourceMappingConfiguration)

	GetFunction(*lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error)
	GetFunctionWithContext(aws.Context, *lambda.GetFunctionInput, ...request.Option) (*lambda.GetFunctionOutput, error)
	GetFunctionRequest(*lambda.GetFunctionInput) (*request.Request, *lambda.GetFunctionOutput)

	GetFunctionConfiguration(*lambda.GetFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)
	GetFunctionConfigurationWithContext(aws.Context, *lambda.GetFunctionConfigurationInput, ...request.Option) (*lambda.FunctionConfiguration, error)
	GetFunctionConfigurationRequest(*lambda.GetFunctionConfigurationInput) (*request.Request, *lambda.FunctionConfiguration)

	GetPolicy(*lambda.GetPolicyInput) (*lambda.GetPolicyOutput, error)
	GetPolicyWithContext(aws.Context, *lambda.GetPolicyInput, ...request.Option) (*lambda.GetPolicyOutput, error)
	GetPolicyRequest(*lambda.GetPolicyInput) (*request.Request, *lambda.GetPolicyOutput)

	Invoke(*lambda.InvokeInput) (*lambda.InvokeOutput, error)
	InvokeWithContext(aws.Context, *lambda.InvokeInput, ...request.Option) (*lambda.InvokeOutput, error)
	InvokeRequest(*lambda.InvokeInput) (*request.Request, *lambda.InvokeOutput)

	InvokeAsync(*lambda.InvokeAsyncInput) (*lambda.InvokeAsyncOutput, error)
	InvokeAsyncWithContext(aws.Context, *lambda.InvokeAsyncInput, ...request.Option) (*lambda.InvokeAsyncOutput, error)
	InvokeAsyncRequest(*lambda.InvokeAsyncInput) (*request.Request, *lambda.InvokeAsyncOutput)

	ListAliases(*lambda.ListAliasesInput) (*lambda.ListAliasesOutput, error)
	ListAliasesWithContext(aws.Context, *lambda.ListAliasesInput, ...request.Option) (*lambda.ListAliasesOutput, error)
	ListAliasesRequest(*lambda.ListAliasesInput) (*request.Request, *lambda.ListAliasesOutput)

	ListEventSourceMappings(*lambda.ListEventSourceMappingsInput) (*lambda.ListEventSourceMappingsOutput, error)
	ListEventSourceMappingsWithContext(aws.Context, *lambda.ListEventSourceMappingsInput, ...request.Option) (*lambda.ListEventSourceMappingsOutput, error)
	ListEventSourceMappingsRequest(*lambda.ListEventSourceMappingsInput) (*request.Request, *lambda.ListEventSourceMappingsOutput)

	ListEventSourceMappingsPages(*lambda.ListEventSourceMappingsInput, func(*lambda.ListEventSourceMappingsOutput, bool) bool) error
	ListEventSourceMappingsPagesWithContext(aws.Context, *lambda.ListEventSourceMappingsInput, func(*lambda.ListEventSourceMappingsOutput, bool) bool, ...request.Option) error

	ListFunctions(*lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, error)
	ListFunctionsWithContext(aws.Context, *lambda.ListFunctionsInput, ...request.Option) (*lambda.ListFunctionsOutput, error)
	ListFunctionsRequest(*lambda.ListFunctionsInput) (*request.Request, *lambda.ListFunctionsOutput)

	ListFunctionsPages(*lambda.ListFunctionsInput, func(*lambda.ListFunctionsOutput, bool) bool) error
	ListFunctionsPagesWithContext(aws.Context, *lambda.ListFunctionsInput, func(*lambda.ListFunctionsOutput, bool) bool, ...request.Option) error

	ListTags(*lambda.ListTagsInput) (*lambda.ListTagsOutput, error)
	ListTagsWithContext(aws.Context, *lambda.ListTagsInput, ...request.Option) (*lambda.ListTagsOutput, error)
	ListTagsRequest(*lambda.ListTagsInput) (*request.Request, *lambda.ListTagsOutput)

	ListVersionsByFunction(*lambda.ListVersionsByFunctionInput) (*lambda.ListVersionsByFunctionOutput, error)
	ListVersionsByFunctionWithContext(aws.Context, *lambda.ListVersionsByFunctionInput, ...request.Option) (*lambda.ListVersionsByFunctionOutput, error)
	ListVersionsByFunctionRequest(*lambda.ListVersionsByFunctionInput) (*request.Request, *lambda.ListVersionsByFunctionOutput)

	PublishVersion(*lambda.PublishVersionInput) (*lambda.FunctionConfiguration, error)
	PublishVersionWithContext(aws.Context, *lambda.PublishVersionInput, ...request.Option) (*lambda.FunctionConfiguration, error)
	PublishVersionRequest(*lambda.PublishVersionInput) (*request.Request, *lambda.FunctionConfiguration)

	PutFunctionConcurrency(*lambda.PutFunctionConcurrencyInput) (*lambda.PutFunctionConcurrencyOutput, error)
	PutFunctionConcurrencyWithContext(aws.Context, *lambda.PutFunctionConcurrencyInput, ...request.Option) (*lambda.PutFunctionConcurrencyOutput, error)
	PutFunctionConcurrencyRequest(*lambda.PutFunctionConcurrencyInput) (*request.Request, *lambda.PutFunctionConcurrencyOutput)

	RemovePermission(*lambda.RemovePermissionInput) (*lambda.RemovePermissionOutput, error)
	RemovePermissionWithContext(aws.Context, *lambda.RemovePermissionInput, ...request.Option) (*lambda.RemovePermissionOutput, error)
	RemovePermissionRequest(*lambda.RemovePermissionInput) (*request.Request, *lambda.RemovePermissionOutput)

	TagResource(*lambda.TagResourceInput) (*lambda.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *lambda.TagResourceInput, ...request.Option) (*lambda.TagResourceOutput, error)
	TagResourceRequest(*lambda.TagResourceInput) (*request.Request, *lambda.TagResourceOutput)

	UntagResource(*lambda.UntagResourceInput) (*lambda.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *lambda.UntagResourceInput, ...request.Option) (*lambda.UntagResourceOutput, error)
	UntagResourceRequest(*lambda.UntagResourceInput) (*request.Request, *lambda.UntagResourceOutput)

	UpdateAlias(*lambda.UpdateAliasInput) (*lambda.AliasConfiguration, error)
	UpdateAliasWithContext(aws.Context, *lambda.UpdateAliasInput, ...request.Option) (*lambda.AliasConfiguration, error)
	UpdateAliasRequest(*lambda.UpdateAliasInput) (*request.Request, *lambda.AliasConfiguration)

	UpdateEventSourceMapping(*lambda.UpdateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
	UpdateEventSourceMappingWithContext(aws.Context, *lambda.UpdateEventSourceMappingInput, ...request.Option) (*lambda.EventSourceMappingConfiguration, error)
	UpdateEventSourceMappingRequest(*lambda.UpdateEventSourceMappingInput) (*request.Request, *lambda.EventSourceMappingConfiguration)

	UpdateFunctionCode(*lambda.UpdateFunctionCodeInput) (*lambda.FunctionConfiguration, error)
	UpdateFunctionCodeWithContext(aws.Context, *lambda.UpdateFunctionCodeInput, ...request.Option) (*lambda.FunctionConfiguration, error)
	UpdateFunctionCodeRequest(*lambda.UpdateFunctionCodeInput) (*request.Request, *lambda.FunctionConfiguration)

	UpdateFunctionConfiguration(*lambda.UpdateFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)
	UpdateFunctionConfigurationWithContext(aws.Context, *lambda.UpdateFunctionConfigurationInput, ...request.Option) (*lambda.FunctionConfiguration, error)
	UpdateFunctionConfigurationRequest(*lambda.UpdateFunctionConfigurationInput) (*request.Request, *lambda.FunctionConfiguration)
}

var _ LambdaAPI = (*lambda.Lambda)(nil)
