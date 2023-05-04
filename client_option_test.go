package gokhttp_http3

import (
	"context"
	"fmt"
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	gokhttp_requests "github.com/BRUHItsABunny/gOkHttp/requests"
	gokhttp_responses "github.com/BRUHItsABunny/gOkHttp/responses"
	"github.com/quic-go/quic-go/http3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClientOption(t *testing.T) {
	hClient, err := gokhttp.NewHTTPClient(NewHTTP3Option(&http3.RoundTripper{}))
	require.NoError(t, err, "NewHTTPClient: errored unexpectedly.")

	req, err := gokhttp_requests.MakeGETRequest(context.Background(), "https://cloudflare-quic.com/")
	require.NoError(t, err, "requests.MakeGETRequest: errored unexpectedly.")

	resp, err := hClient.Do(req)
	require.NoError(t, err, "hClient.Do: errored unexpectedly.")

	respData, err := gokhttp_responses.ResponseText(resp)
	require.NoError(t, err, "gokhttp_responses.ResponseText: errored unexpectedly.")

	fmt.Println(respData)
}
