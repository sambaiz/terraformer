// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the input of an update repository description operation.
type UpdateRepositoryDescriptionInput struct {
	_ struct{} `type:"structure"`

	// The new comment or description for the specified repository. Repository descriptions
	// are limited to 1,000 characters.
	RepositoryDescription *string `locationName:"repositoryDescription" type:"string"`

	// The name of the repository to set or change the comment or description for.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateRepositoryDescriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRepositoryDescriptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateRepositoryDescriptionInput"}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateRepositoryDescriptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateRepositoryDescriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateRepositoryDescription = "UpdateRepositoryDescription"

// UpdateRepositoryDescriptionRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Sets or changes the comment or description for a repository.
//
// The description field for a repository accepts all HTML characters and all
// valid Unicode characters. Applications that do not HTML-encode the description
// and display it in a webpage can expose users to potentially malicious code.
// Make sure that you HTML-encode the description field in any application that
// uses this API to display the repository description on a webpage.
//
//    // Example sending a request using UpdateRepositoryDescriptionRequest.
//    req := client.UpdateRepositoryDescriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/UpdateRepositoryDescription
func (c *Client) UpdateRepositoryDescriptionRequest(input *UpdateRepositoryDescriptionInput) UpdateRepositoryDescriptionRequest {
	op := &aws.Operation{
		Name:       opUpdateRepositoryDescription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateRepositoryDescriptionInput{}
	}

	req := c.newRequest(op, input, &UpdateRepositoryDescriptionOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateRepositoryDescriptionRequest{Request: req, Input: input, Copy: c.UpdateRepositoryDescriptionRequest}
}

// UpdateRepositoryDescriptionRequest is the request type for the
// UpdateRepositoryDescription API operation.
type UpdateRepositoryDescriptionRequest struct {
	*aws.Request
	Input *UpdateRepositoryDescriptionInput
	Copy  func(*UpdateRepositoryDescriptionInput) UpdateRepositoryDescriptionRequest
}

// Send marshals and sends the UpdateRepositoryDescription API request.
func (r UpdateRepositoryDescriptionRequest) Send(ctx context.Context) (*UpdateRepositoryDescriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRepositoryDescriptionResponse{
		UpdateRepositoryDescriptionOutput: r.Request.Data.(*UpdateRepositoryDescriptionOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRepositoryDescriptionResponse is the response type for the
// UpdateRepositoryDescription API operation.
type UpdateRepositoryDescriptionResponse struct {
	*UpdateRepositoryDescriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRepositoryDescription request.
func (r *UpdateRepositoryDescriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
