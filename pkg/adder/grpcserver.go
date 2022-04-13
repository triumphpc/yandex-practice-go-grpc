package adder

import (
	"context"
	"fmt"
	"sort"
	"sync"

	// импортируем пакет со сгенерированными protobuf-файлами
	pb "github.com/triumphpc/yandex-practice-go-grpc/pkg/api"
)

// UsersServer поддерживает все необходимые методы.
type UsersServer struct {
	// нужно встраивать тип pb.Unimplemented<TypeName>
	// для совместимости с будущими версиями
	pb.UnimplementedUsersServer
	// используем sync.Map для хранения пользователей
	users sync.Map
}

// AddUser реализует интерфейс добавления пользователя.
func (s *UsersServer) AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	var response pb.AddUserResponse

	if _, ok := s.users.Load(in.User.Email); ok {
		response.Error = fmt.Sprintf(`Пользователь с email %s уже существует`, in.User.Email)
	} else {
		s.users.Store(in.User.Email, in.User)
	}
	return &response, nil
}

// ListUsers реализует интерфейс получения списка пользователей.
func (s *UsersServer) ListUsers(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	var list []string

	s.users.Range(func(key, _ interface{}) bool {
		list = append(list, key.(string))
		return true
	})
	// сортируем слайс из email
	sort.Strings(list)

	offset := int(in.Offset)
	end := int(in.Offset + in.Limit)
	if end > len(list) {
		end = len(list)
	}
	if offset >= end {
		offset = 0
		end = 0
	}
	response := pb.ListUsersResponse{
		Count:  int32(len(list)),
		Emails: list[offset:end],
	}
	return &response, nil
}

// GetUser реализация интерфейса получения пользователя
func (s *UsersServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var response pb.GetUserResponse

	if user, ok := s.users.Load(in.Email); ok {
		response.User = user.(*pb.User)
	} else {
		response.Error = fmt.Sprintf(`Пользователь с email %s не найден`, in.Email)
	}

	return &response, nil
}

// DelUser реализует интерфейс удаления информации о пользователе.
func (s *UsersServer) DelUser(ctx context.Context, in *pb.DelUserRequest) (*pb.DelUserResponse, error) {
	var response pb.DelUserResponse

	if _, ok := s.users.LoadAndDelete(in.Email); !ok {
		response.Error = fmt.Sprintf(`Пользователь с email %s не найден`, in.Email)
	}
	return &response, nil
}
