package grpc

import (
	"adapter-service/constants"
	"adapter-service/models"
	"adapter-service/proto/proto_models"
	"adapter-service/services/mail"
	"context"
)

type grpcAdapterHandler struct {
	proto_models.UnimplementedAdapterServer
	mailUs mail.MailUsecase
}

func NewGrpcAdapterHandlerImpl(mailUs mail.MailUsecase) proto_models.AdapterServer {
	return &grpcAdapterHandler{
		mailUs: mailUs,
	}
}

func (g *grpcAdapterHandler) SendMail(ctx context.Context, request *proto_models.SendMailRequest) (*proto_models.SendMailResponse, error) {
	form := &models.MailForm{
		SenderName: constants.MAIL_SENDER_NAME,
		To:         request.GetTo(),
		ToName:     request.GetToName(),
		Subject:    request.GetSubject(),
		Body:       request.GetBody(),
	}

	if err := g.mailUs.SendMail(ctx, form); err != nil {
		return nil, err
	}

	return nil, nil
}
