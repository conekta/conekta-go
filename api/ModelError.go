package api

import (
)

type ModelError struct {
    Object  string  `json:"object,omitempty"`
    Type_  string  `json:"type,omitempty"`
    Message  string  `json:"message,omitempty"`
    MessageToPurchaser  string  `json:"message_to_purchaser,omitempty"`
    Code  string  `json:"code,omitempty"`
    Param  string  `json:"param,omitempty"`
    
}
