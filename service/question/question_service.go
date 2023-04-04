package question

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	questionpb "grpc-questions/gen/go/proto"
	"grpc-questions/model"
	"grpc-questions/store"
)

type Service struct {
	questionpb.UnimplementedQuestionServiceServer
	db *store.QuestionDB
}

func (s *Service) CreateQuestion(ctx context.Context, request *questionpb.CreateQuestionRequest) (*questionpb.CreateQuestionResponse, error) {
	newQn := model.Question{
		Title:       request.Question.GetTitle(),
		Body:        request.GetQuestion().GetBody(),
		Answer:      request.GetQuestion().GetAnswer(),
		Explanation: request.GetQuestion().GetExplanation(),
		Params:      request.GetQuestion().GetParams(),
	}

	out, err := s.db.Create(newQn)
	if err != nil {
		return nil, err
	}

	return &questionpb.CreateQuestionResponse{
		Question: &questionpb.Question{
			Id:          out.ID,
			Title:       out.Title,
			Body:        out.Body,
			Answer:      out.Answer,
			Explanation: out.Explanation,
			Params:      out.Params,
			CreatedAt:   timestamppb.New(out.CreatedAt),
			UpdatedAt:   timestamppb.New(out.UpdatedAt),
		},
	}, nil
}

func (s *Service) ReadQuestion(ctx context.Context, request *questionpb.ReadQuestionRequest) (*questionpb.ReadQuestionResponse, error) {
	out, err := s.db.Read(request.GetId())
	if err != nil {
		return nil, err
	}

	return &questionpb.ReadQuestionResponse{
		Question: &questionpb.Question{
			Id:          out.ID,
			Title:       out.Title,
			Body:        out.Body,
			Answer:      out.Answer,
			Explanation: out.Explanation,
			Params:      out.Params,
			CreatedAt:   timestamppb.New(out.CreatedAt),
			UpdatedAt:   timestamppb.New(out.UpdatedAt),
		},
	}, nil
}

func (s *Service) ListQuestions(ctx context.Context, request *questionpb.ListQuestionsRequest) (*questionpb.ListQuestionsResponse, error) {
	out, err := s.db.List()
	if err != nil {
		return nil, err
	}

	var questions []*questionpb.Question
	for _, q := range out {
		questions = append(questions, &questionpb.Question{
			Id:          q.ID,
			Title:       q.Title,
			Body:        q.Body,
			Answer:      q.Answer,
			Explanation: q.Explanation,
			Params:      q.Params,
			CreatedAt:   timestamppb.New(q.CreatedAt),
			UpdatedAt:   timestamppb.New(q.UpdatedAt),
		})
	}

	return &questionpb.ListQuestionsResponse{
		Questions: questions,
	}, nil
}

func (s *Service) UpdateQuestion(ctx context.Context, request *questionpb.UpdateQuestionRequest) (*questionpb.UpdateQuestionResponse, error) {
	newQn := model.Question{
		ID:          request.GetQuestion().GetId(),
		Title:       request.GetQuestion().GetTitle(),
		Body:        request.GetQuestion().GetBody(),
		Answer:      request.GetQuestion().GetAnswer(),
		Explanation: request.GetQuestion().GetExplanation(),
		Params:      request.GetQuestion().GetParams(),
	}

	out, err := s.db.Update(newQn)
	if err != nil {
		return nil, err
	}

	return &questionpb.UpdateQuestionResponse{
		Question: &questionpb.Question{
			Id:          out.ID,
			Title:       out.Title,
			Body:        out.Body,
			Answer:      out.Answer,
			Explanation: out.Explanation,
			Params:      out.Params,
			CreatedAt:   timestamppb.New(out.CreatedAt),
			UpdatedAt:   timestamppb.New(out.UpdatedAt),
		},
	}, nil
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
