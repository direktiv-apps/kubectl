// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/direktiv/apps/go/pkg/apps"
)

// PostHandlerFunc turns a function with the right signature into a post handler
type PostHandlerFunc func(PostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostHandlerFunc) Handle(params PostParams) middleware.Responder {
	return fn(params)
}

// PostHandler interface for that can handle valid post params
type PostHandler interface {
	Handle(PostParams) middleware.Responder
}

// NewPost creates a new http.Handler for the post operation
func NewPost(ctx *middleware.Context, handler PostHandler) *Post {
	return &Post{Context: ctx, Handler: handler}
}

/* Post swagger:route POST / post

Post post API

*/
type Post struct {
	Context *middleware.Context
	Handler PostHandler
}

func (o *Post) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostBody post body
//
// swagger:model PostBody
type PostBody struct {

	// Array of commands.
	Commands []*PostParamsBodyCommandsItems0 `json:"commands"`

	// File to create before running commands. This can include a `kubectl.yaml` file from secrets.
	Files []apps.DirektivFile `json:"files"`

	// Base64 kubectl.yaml file. If not set `kubectl.yaml` will be used. This can be provided via Direktiv files.
	// Example: kubeconfig.yaml
	Kubeconfig string `json:"kubeconfig,omitempty"`
}

// Validate validates this post body
func (o *PostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommands(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostBody) validateCommands(formats strfmt.Registry) error {
	if swag.IsZero(o.Commands) { // not required
		return nil
	}

	for i := 0; i < len(o.Commands); i++ {
		if swag.IsZero(o.Commands[i]) { // not required
			continue
		}

		if o.Commands[i] != nil {
			if err := o.Commands[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "commands" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "commands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PostBody) validateFiles(formats strfmt.Registry) error {
	if swag.IsZero(o.Files) { // not required
		return nil
	}

	for i := 0; i < len(o.Files); i++ {

		if err := o.Files[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "files" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "files" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this post body based on the context it is used
func (o *PostBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCommands(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostBody) contextValidateCommands(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Commands); i++ {

		if o.Commands[i] != nil {
			if err := o.Commands[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "commands" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "commands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PostBody) contextValidateFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Files); i++ {

		if err := o.Files[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "files" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "files" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBody) UnmarshalBinary(b []byte) error {
	var res PostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostOKBody post o k body
//
// swagger:model PostOKBody
type PostOKBody struct {

	// kubectl
	Kubectl []*PostOKBodyKubectlItems0 `json:"kubectl"`
}

// Validate validates this post o k body
func (o *PostOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKubectl(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOKBody) validateKubectl(formats strfmt.Registry) error {
	if swag.IsZero(o.Kubectl) { // not required
		return nil
	}

	for i := 0; i < len(o.Kubectl); i++ {
		if swag.IsZero(o.Kubectl[i]) { // not required
			continue
		}

		if o.Kubectl[i] != nil {
			if err := o.Kubectl[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postOK" + "." + "kubectl" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postOK" + "." + "kubectl" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post o k body based on the context it is used
func (o *PostOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateKubectl(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOKBody) contextValidateKubectl(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Kubectl); i++ {

		if o.Kubectl[i] != nil {
			if err := o.Kubectl[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postOK" + "." + "kubectl" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postOK" + "." + "kubectl" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOKBody) UnmarshalBinary(b []byte) error {
	var res PostOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostOKBodyKubectlItems0 post o k body kubectl items0
//
// swagger:model PostOKBodyKubectlItems0
type PostOKBodyKubectlItems0 struct {

	// result
	// Required: true
	Result interface{} `json:"result"`

	// success
	// Required: true
	Success *bool `json:"success"`
}

// Validate validates this post o k body kubectl items0
func (o *PostOKBodyKubectlItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSuccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOKBodyKubectlItems0) validateResult(formats strfmt.Registry) error {

	if o.Result == nil {
		return errors.Required("result", "body", nil)
	}

	return nil
}

func (o *PostOKBodyKubectlItems0) validateSuccess(formats strfmt.Registry) error {

	if err := validate.Required("success", "body", o.Success); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post o k body kubectl items0 based on context it is used
func (o *PostOKBodyKubectlItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOKBodyKubectlItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOKBodyKubectlItems0) UnmarshalBinary(b []byte) error {
	var res PostOKBodyKubectlItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostParamsBodyCommandsItems0 post params body commands items0
//
// swagger:model PostParamsBodyCommandsItems0
type PostParamsBodyCommandsItems0 struct {

	// Command to run
	// Example: kubectl version --client=true -o json
	Command string `json:"command,omitempty"`

	// Stops excecution if command fails, otherwise proceeds with next command
	Continue bool `json:"continue,omitempty"`

	// If set to false the command will not print the full command with arguments to logs.
	Print *bool `json:"print,omitempty"`

	// If set to false the command will not print output to logs.
	Silent *bool `json:"silent,omitempty"`
}

// Validate validates this post params body commands items0
func (o *PostParamsBodyCommandsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post params body commands items0 based on context it is used
func (o *PostParamsBodyCommandsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostParamsBodyCommandsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostParamsBodyCommandsItems0) UnmarshalBinary(b []byte) error {
	var res PostParamsBodyCommandsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
