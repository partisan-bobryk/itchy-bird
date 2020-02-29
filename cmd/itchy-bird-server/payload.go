package main

type ResponsePayload struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
