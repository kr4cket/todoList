package connection

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"todoList/internal/common/configs"
)

func GetPostgresConnection(cfg configs.DBConfig) (*pgx.Conn, error) {
	connection, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSslMode))

	if err != nil {
		logrus.Fatalf("unable to create connection: %v\n", err)
		return nil, err
	}

	err = connection.Ping(context.Background())
	if err != nil {
		logrus.Fatalf("unable to ping DB: %v\n", err)
		return nil, err
	}

	return connection, nil
}
