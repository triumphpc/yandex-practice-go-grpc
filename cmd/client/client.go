package main

import (
	"context"
	"fmt"
	// ...
	"google.golang.org/grpc"
	"log"

	pb "github.com/triumphpc/yandex-practice-go-grpc/pkg/api"
)

func main() {
	// устанавливаем соединение с сервером
	conn, err := grpc.Dial(`:3200`, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// получаем переменную интерфейсного типа UsersClient,
	// через которую будем отправлять сообщения
	c := pb.NewUsersClient(conn)

	// функция, в которой будем отправлять сообщения
	TestUsers(c)
}

func TestUsers(c pb.UsersClient) {
	// набор тестовых данных
	users := []*pb.User{
		{Name: `Сергей`, Email: `serge@example.com`, Sex: pb.User_MAN},
		{Name: `Света`, Email: `sveta@example.com`, Sex: pb.User_WOMAN},
		{Name: `Денис`, Email: `den@example.com`, Sex: pb.User_MAN},
		// при добавлении этой записи должна вернуться ошибка:
		// пользователь с email sveta@example.com уже существует
		{Name: `Sveta`, Email: `sveta@example.com`, Sex: pb.User_WOMAN},
	}
	for _, user := range users {
		// добавляем пользователей
		resp, err := c.AddUser(context.Background(), &pb.AddUserRequest{
			User: user,
		})
		if err != nil {
			log.Fatal(err)
		}
		if resp.Error != "" {
			fmt.Println(resp.Error)
		}
	}
	// удаляем одного из пользователей
	if resp, err := c.DelUser(context.Background(), &pb.DelUserRequest{
		Email: `serge@example.com`,
	}); err != nil {
		log.Fatal(err)
	} else if resp.Error != "" {
		fmt.Println(resp.Error)
	}

	// получаем информацию о пользователях
	// во втором случае должна вернуться ошибка:
	// пользователь с email serge@example.com не найден
	for _, userEmail := range []string{`sveta@example.com`, `serge@example.com`} {
		if resp, err := c.GetUser(context.Background(), &pb.GetUserRequest{
			Email: userEmail,
		}); err != nil {
			log.Fatal(err)
		} else if resp.Error == "" {
			fmt.Println(resp.User)
		} else {
			fmt.Println(resp.Error)
		}
	}

	// получаем список email пользователей
	emails, err := c.ListUsers(context.Background(), &pb.ListUsersRequest{
		Offset: 0,
		Limit:  100,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(emails.Count, emails.Emails)
}
