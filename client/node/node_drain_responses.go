// Code generated by go-swagger; DO NOT EDIT.

package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// NodeDrainReader is a Reader for the NodeDrain structure.
type NodeDrainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeDrainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeDrainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNodeDrainDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodeDrainOK creates a NodeDrainOK with default headers values
func NewNodeDrainOK() *NodeDrainOK {
	return &NodeDrainOK{}
}

/*
	NodeDrainOK describes a response with status code 200, with default header values.

Operation completed successfully
*/
type NodeDrainOK struct {
	Payload *models.RegularResponse
}

func (o *NodeDrainOK) Error() string {
	return fmt.Sprintf("[POST /nodes/{node}/drain][%d] nodeDrainOK  %+v", 200, o.Payload)
}
func (o *NodeDrainOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *NodeDrainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeDrainDefault creates a NodeDrainDefault with default headers values
func NewNodeDrainDefault(code int) *NodeDrainDefault {
	return &NodeDrainDefault{
		_statusCode: code,
	}
}

/*
	NodeDrainDefault describes a response with status code -1, with default header values.

Error draining node
*/
type NodeDrainDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the node drain default response
func (o *NodeDrainDefault) Code() int {
	return o._statusCode
}

func (o *NodeDrainDefault) Error() string {
	return fmt.Sprintf("[POST /nodes/{node}/drain][%d] NodeDrain default  %+v", o._statusCode, o.Payload)
}
func (o *NodeDrainDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeDrainDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
