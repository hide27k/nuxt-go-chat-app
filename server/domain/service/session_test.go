package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hideUW/nuxt-go-chat-app/server/domain/model"
	"github.com/hideUW/nuxt-go-chat-app/server/domain/repository"
	mock_repository "github.com/hideUW/nuxt-go-chat-app/server/domain/repository/mock"
	"github.com/hideUW/nuxt-go-chat-app/server/infra/db"
	"github.com/hideUW/nuxt-go-chat-app/server/testutil"
	"github.com/pkg/errors"
)

func Test_sessionService_IsAlreadyExistID(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_repository.NewMockSessionRepository(ctrl)

	type fields struct {
		repo repository.SessionRepository
		m    repository.SQLManager
	}

	type args struct {
		ctx context.Context
		id  string
	}

	type returnArgs struct {
		session *model.Session
		err     error
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		returnArgs
		want    bool
		wantErr error
	}{
		{
			name: "When the specific session already exists, return true and nil.",
			fields: fields{
				repo: mock,
				m:    db.NewSQLManager(),
			},
			args: args{
				ctx: context.Background(),
				id:  model.SessionValidIDForTest,
			},
			returnArgs: returnArgs{
				session: &model.Session{
					ID:        model.SessionValidIDForTest,
					UserID:    model.UserValidIDForTest,
					CreatedAt: testutil.TimeNow(),
					UpdatedAt: testutil.TimeNow(),
				},
				err: nil,
			},
			want:    true,
			wantErr: nil,
		},
		{
			name: "When the specific session doesn't exit, return false and nil.",
			fields: fields{
				repo: mock,
				m:    db.NewSQLManager(),
			},
			args: args{
				ctx: context.Background(),
				id:  model.SessionInValidIDForTest,
			},
			returnArgs: returnArgs{
				session: nil,
				err:     nil,
			},
			want:    false,
			wantErr: nil,
		},
		{
			name: "When some errors have ocurred, return false and error",
			fields: fields{
				repo: mock,
				m:    db.NewSQLManager(),
			},
			args: args{
				ctx: context.Background(),
				id:  model.SessionInValidIDForTest,
			},
			returnArgs: returnArgs{
				session: nil,
				err:     errors.New(model.ErrorMessageForTest),
			},
			want:    false,
			wantErr: errors.New(model.ErrorMessageForTest),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sessionService{
				repo: mock,
				m:    tt.fields.m,
			}

			mock.EXPECT().GetSessionByID(tt.fields.m, tt.args.id).Return(tt.returnArgs.session, tt.returnArgs.err)

			got, err := s.IsAlreadyExistID(tt.args.ctx, tt.args.id)
			if tt.wantErr != nil {
				if errors.Cause(err).Error() != tt.wantErr.Error() {
					t.Errorf("sessionService.IsAlreadyExistID() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("sessionService.IsAlreadyExistID() = %v, want %v", got, tt.want)
			}
		})
	}

}
