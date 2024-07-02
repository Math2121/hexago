package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonErr(t *testing.T){
	msg := "Hey"
	result := jsonError(msg)
	require.Equal(t, string([]byte(`{"message":"Hey"}`)), string(result))
	
}