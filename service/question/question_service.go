package question

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	questionpb "grpc-questions/gen/go/proto"
	"grpc-questions/store"
)

type Service struct {
	questionpb.UnimplementedQuestionServiceServer
	db *store.QuestionDB
}

func (s *Service) CreateQuestion(ctx context.Context, request *questionpb.CreateQuestionRequest) (*questionpb.CreateQuestionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ReadQuestion(ctx context.Context, request *questionpb.ReadQuestionRequest) (*questionpb.ReadQuestionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ListQuestions(ctx context.Context, request *questionpb.ListQuestionsRequest) (*questionpb.ListQuestionsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateQuestion(ctx context.Context, request *questionpb.UpdateQuestionRequest) (*questionpb.UpdateQuestionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func RegisterService(svr *grpc.Server, svrEndpoint string, httpGwMux *runtime.ServeMux) {

	questionpb.RegisterQuestionServiceServer(svr, &Service{
		db: store.NewQuestionDB(),
	})

	ctx := context.Background()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if httpGwMux != nil {
		if err := questionpb.RegisterQuestionServiceHandlerFromEndpoint(ctx, httpGwMux, svrEndpoint, opts); err != nil {
			panic(err)
		}
	}
}
