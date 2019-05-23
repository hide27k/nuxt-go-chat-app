package db

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/hideUW/nuxt_go_template/server/domain/model"
	"github.com/hideUW/nuxt_go_template/server/domain/repository"
	"github.com/hideUW/nuxt_go_template/server/testutil"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func Test_userRepository_GetUserByID(t *testing.T) {
	// Set sql mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
		return
	}

	// If db has error, close db.
	defer db.Close()

	// Set fake time.
	testutil.SetFakeTime(time.Now())

	type fields struct {
		ctx context.Context
	}

	type args struct {
		m  repository.DBManager
		id uint32
	}

	var validID uint32 = 1
	var inValidID uint32 = 2

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr *model.NoSuchDataError
	}{
		{
			name: "When the specific user exists, return the user.",
			fields: fields{
				ctx: context.Background(),
			},
			args: args{
				m:  db,
				id: validID,
			},
			want: &model.User{
				ID:        validID,
				Name:      "test",
				SessionID: "test12345678",
				Password:  "test",
				CreatedAt: testutil.TimeNow(),
				UpdatedAt: testutil.TimeNow(),
			},
			wantErr: nil,
		},
		{
			name: "When the specific user does not exist, return NoSuchDataError",
			fields: fields{
				ctx: context.Background(),
			},
			args: args{
				m:  db,
				id: inValidID,
			},
			want: nil,
			wantErr: &model.NoSuchDataError{
				PropertyNameForDeveloper:    model.IDPropertyForDeveloper,
				PropertyNameForUser:         model.IDPropertyForUser,
				PropertyValue:               inValidID,
				DomainModelNameForDeveloper: model.DomainModelNameUserForDeveloper,
				DomainModelNameForUser:      model.DomainModelNameUserForUser,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := "SELECT id, name, session_id, password, created_at, updated_at FROM users WHERE id=?"
			prep := mock.ExpectPrepare(q)

			if tt.wantErr != nil {
				prep.ExpectQuery().WillReturnError(tt.wantErr)
			} else {
				rows := sqlmock.NewRows([]string{"id", "name", "session_id", "password", "created_at", "updated_at"}).
					AddRow(tt.want.ID, tt.want.Name, tt.want.SessionID, tt.want.Password, tt.want.CreatedAt, tt.want.UpdatedAt)
				prep.ExpectQuery().WithArgs(tt.want.ID).WillReturnRows(rows)
			}

			repo := &userRepository{
				ctx: tt.fields.ctx,
			}
			got, err := repo.GetUserByID(tt.args.m, tt.args.id)

			if tt.wantErr != nil {
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("userRepository.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}

}
