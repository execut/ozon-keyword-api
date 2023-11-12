package main

import (
	"context"
	"errors"
	"github.com/execut/ozon-keyword-api/internal/api"
	"github.com/execut/ozon-keyword-api/internal/model"
	"github.com/execut/ozon-keyword-api/internal/repo"
	pb "github.com/execut/ozon-keyword-api/pkg/ozon-keyword-api"
	"github.com/rs/zerolog"
	"gotest.tools/v3/assert"
	"log"
	"net"
	"strings"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func testServer() (pb.OzonKeywordApiServiceClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	pb.RegisterOzonKeywordApiServiceServer(baseServer, api.NewKeywordAPI(newStubKeywordRepo()))
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(context.Background(), "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := pb.NewOzonKeywordApiServiceClient(conn)

	zerolog.SetGlobalLevel(zerolog.Disabled)

	return client, closer
}

func newStubKeywordRepo() repo.Repo {
	return &StubRepo{}
}

type StubRepo struct {
}

func (r StubRepo) GetKeyword(ctx context.Context, ozonID uint64) (*model.Keyword, error) {
	return nil, errors.New("GetKeyword unimplemented")
}

func TestOzonKeywordApiServiceServer_CreateKeywordV1(t *testing.T) {
	client, closer := testServer()
	defer closer()

	type expectation struct {
		out *pb.CreateKeywordV1Response
		err string
	}

	errBadNameString := "rpc error: code = InvalidArgument desc = invalid CreateKeywordV1Request.Name: value length must be between 1 and 255 runes, inclusive"
	tests := map[string]struct {
		in       *pb.CreateKeywordV1Request
		expected expectation
	}{
		"Success_Unimplemented": {
			in: &pb.CreateKeywordV1Request{
				Name: "test",
			},
			expected: expectation{
				out: &pb.CreateKeywordV1Response{},
				err: "rpc error: code = Unimplemented desc = CreateKeywordV1 not implemented",
			},
		},
		"WhenNameNil_Error": {
			in: &pb.CreateKeywordV1Request{},
			expected: expectation{
				out: &pb.CreateKeywordV1Response{},
				err: errBadNameString,
			},
		},
		"WhenNameEmpty_Error": {
			in: &pb.CreateKeywordV1Request{
				Name: "",
			},
			expected: expectation{
				out: &pb.CreateKeywordV1Response{},
				err: errBadNameString,
			},
		},
		"WhenNameLenGreaterWhen255_Error": {
			in: &pb.CreateKeywordV1Request{
				Name: strings.Repeat("t", 256),
			},
			expected: expectation{
				out: &pb.CreateKeywordV1Response{},
				err: errBadNameString,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			_, err := client.CreateKeywordV1(context.Background(), tt.in)
			assert.Equal(t, tt.expected.err, err.Error())
		})
	}
}

func TestOzonKeywordApiServiceServer_DescribeKeywordV1(t *testing.T) {
	client, closer := testServer()
	defer closer()

	type expectation struct {
		out *pb.DescribeKeywordV1Response
		err error
	}

	tests := map[string]struct {
		in       *pb.DescribeKeywordV1Request
		expected expectation
	}{
		"Success_Unimplemented": {
			in: &pb.DescribeKeywordV1Request{
				KeywordId: 123,
			},
			expected: expectation{
				out: nil,
				err: errors.New("rpc error: code = Internal desc = GetKeyword unimplemented"),
			},
		},
		"WhenKeywordIdIsZero_Error": {
			in: &pb.DescribeKeywordV1Request{
				KeywordId: 0,
			},
			expected: expectation{
				out: nil,
				err: errors.New("rpc error: code = InvalidArgument desc = invalid DescribeKeywordV1Request.KeywordId: value must be greater than 0"),
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			_, err := client.DescribeKeywordV1(context.Background(), tt.in)
			assert.Equal(t, tt.expected.err.Error(), err.Error())
		})
	}
}

func TestOzonKeywordApiServiceServer_RemoveKeywordV1(t *testing.T) {
	client, closer := testServer()
	defer closer()

	type expectation struct {
		out *pb.RemoveKeywordV1Request
		err error
	}

	tests := map[string]struct {
		in       *pb.RemoveKeywordV1Request
		expected expectation
	}{
		"Success_Unimplemented": {
			in: &pb.RemoveKeywordV1Request{
				KeywordId: 123,
			},
			expected: expectation{
				out: nil,
				err: errors.New("rpc error: code = Unimplemented desc = RemoveKeywordV1 not implemented"),
			},
		},
		"WhenKeywordIdIsZero_Error": {
			in: &pb.RemoveKeywordV1Request{
				KeywordId: 0,
			},
			expected: expectation{
				out: nil,
				err: errors.New("rpc error: code = InvalidArgument desc = invalid RemoveKeywordV1Request.KeywordId: value must be greater than 0"),
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			_, err := client.RemoveKeywordV1(context.Background(), tt.in)
			assert.Equal(t, tt.expected.err.Error(), err.Error())
		})
	}
}

func TestOzonKeywordApiServiceServer_ListKeywordV1(t *testing.T) {
	client, closer := testServer()
	defer closer()

	type expectation struct {
		out *pb.ListKeywordV1Request
		err error
	}

	tests := map[string]struct {
		in       *pb.ListKeywordV1Request
		expected expectation
	}{
		"Success_Unimplemented": {
			in: &pb.ListKeywordV1Request{},
			expected: expectation{
				out: nil,
				err: errors.New("rpc error: code = Unimplemented desc = ListKeywordV1 not implemented"),
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			_, err := client.ListKeywordV1(context.Background(), tt.in)
			assert.Equal(t, tt.expected.err.Error(), err.Error())
		})
	}
}
