package mock

import (
	"context"
	"math/rand"
	"time"

	"github.com/aidarkhanov/nanoid/v2"
	"github.com/genshinsim/gcsim/backend/pkg/services/share"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Server is a mock server for purpose of testing share RPC end points
type Server struct {
	Log  *zap.SugaredLogger
	Rand *rand.Rand
	data map[string]*share.ShareEntry
}

func NewServer(cust ...func(*Server) error) (*Server, error) {
	s := &Server{}
	s.Rand = rand.New(rand.NewSource(time.Now().Unix()))

	for _, f := range cust {
		err := f(s)
		if err != nil {
			return nil, err
		}
	}

	if s.Log == nil {
		logger, err := zap.NewProduction()
		if err != nil {
			return nil, err
		}
		sugar := logger.Sugar()
		sugar.Debugw("logger initiated")

		s.Log = sugar
	}

	return s, nil
}

func (s *Server) Create(ctx context.Context, e *share.ShareEntry) (string, error) {
	key, err := nanoid.New()
	if err != nil {
		s.Log.Infow("err creating nanoid", "err", err)
		return "", status.Error(codes.Internal, "error creating nanoid")
	}
	if _, ok := s.data[key]; ok {
		return "", status.Error(codes.Internal, "error creating nanoid")
	}
	return key, nil
}

func (s *Server) Read(ctx context.Context, key string) (*share.ShareEntry, error) {
	val, ok := s.data[key]
	if !ok {
		return nil, status.Error(codes.NotFound, "key not found")
	}
	n := proto.Clone(val)
	return n.(*share.ShareEntry), nil
}

func (s *Server) Update(ctx context.Context, entry *share.ShareEntry) (string, error) {
	key := entry.GetId()
	_, ok := s.data[key]
	if !ok {
		return "", status.Error(codes.NotFound, "key not found")
	}
	s.data[key] = entry
	return key, nil
}

func (s *Server) SetTTL(ctx context.Context, key string, until uint64) (string, error) {
	_, ok := s.data[key]
	if !ok {
		return "", status.Error(codes.NotFound, "key not found")
	}
	return key, nil
}

func (s *Server) Delete(ctx context.Context, key string) error {
	_, ok := s.data[key]
	if !ok {
		return status.Error(codes.NotFound, "key not found")
	}
	delete(s.data, key)
	return nil
}

func (s *Server) Random(context.Context) (*share.ShareEntry, error) {
	max := len(s.data)
	if max == 0 {
		return nil, status.Error(codes.NotFound, "not found")
	}
	n := s.Rand.Intn(max)
	for _, v := range s.data {
		if n == 0 {
			return v, nil
		}
		n--
	}

	return nil, status.Error(codes.NotFound, "not found")
}
