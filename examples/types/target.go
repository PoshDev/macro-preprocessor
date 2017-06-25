// .----------------------------------------.
// |      WARNING: AUTO-GENERATED CODE      |
// .----------------------------------------.
// This code was automatically generated using
// https://github.com/PoshDev/macro-preprocessor .
// The original source file was source.go.
// XXX: DO NOT CHANGE THIS FILE!!!
package main


type Service interface {
	
	func String() string

}

type service struct {
	
	name string `json:"n"`

}

var _ Service = (*service)(nil)

func NewService() Service {
	
	return &service{"Untitled"}

}


