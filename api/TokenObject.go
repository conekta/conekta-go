package api

import (
)

type TokenObject struct {
    Token  string  `json:"token,omitempty"`
    Id  string  `json:"id,omitempty"`
    Active  bool  `json:"active,omitempty"`
    
}
