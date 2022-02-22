package main

import "testing"

func TestHello(t *testing.T) {
	if Hello("Openshift") != "Hello, Openshift" {
		t.Error("Expected Hello(\"Openshift\") to equal \"Hello, Openshift\"")
	}
}
