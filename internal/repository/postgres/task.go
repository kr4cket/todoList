package postgres

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"time"
	"todoList/internal/common/models"
)

const taskTable = "todoList.tasks"

type TaskRepository struct {
	db *pgx.Conn
	qb sq.StatementBuilderType
}

func NewTaskRepository(db *pgx.Conn) *TaskRepository {
	return &TaskRepository{
		db: db,
		qb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *TaskRepository) Get(ctx *fiber.Ctx) ([]models.Task, error) {
	query, args, err := r.qb.Select("id, title, description, status, created_at, updated_at").
		From(taskTable).ToSql()

	if err != nil {
		logrus.Fatalf("error while getting tasks data from db: %s", err.Error())
		return []models.Task{}, err
	}

	row, err := r.db.Query(ctx.Context(), query, args...)

	tasks, err := pgx.CollectRows(row, func(row pgx.CollectableRow) (models.Task, error) {
		var task models.Task
		err = row.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			logrus.Fatalf("error while scanning row")
			return models.Task{}, fmt.Errorf("error while getting entities from db: %w", err)
		}

		return task, nil
	})

	if err != nil {
		logrus.Fatalf("error while getting task entities from db: %s", err.Error())
		return nil, err
	}

	return tasks, nil
}
func (r *TaskRepository) Insert(ctx *fiber.Ctx, task models.Task) error {
	res := r.qb.Insert(taskTable).
		Columns("title", "description").
		Values(task.Title, task.Description)

	query, args, err := res.ToSql()
	if err != nil {
		logrus.Fatalf("error while creating insert query db: %s", err.Error())
		return fmt.Errorf("build insert query error: %w", err)
	}

	_, err = r.db.Exec(ctx.Context(), query, args...)
	if err != nil {
		logrus.Fatalf("error while inserting in db: %s", err.Error())
		return fmt.Errorf("insert task error: %w", err)
	}

	return nil
}

func (r *TaskRepository) Update(ctx *fiber.Ctx, id int, task models.Task) error {
	data := r.getUpdateData(task)
	updateQuery := r.qb.Update(taskTable).
		SetMap(data).
		Where(sq.Eq{"id": id})

	sqlQuery, arguments, err := updateQuery.ToSql()
	if err != nil {
		logrus.Fatalf("error while creating update entity: %s", err.Error())
		return fmt.Errorf("build query error: %w", err)
	}

	_, err = r.db.Exec(ctx.Context(), sqlQuery, arguments...)
	if err != nil {
		logrus.Fatalf("error while update task entity: %s", err.Error())
		return fmt.Errorf("update task error: %w", err)
	}
	return nil
}

func (r *TaskRepository) Delete(ctx *fiber.Ctx, id int) error {
	query, args, err := r.qb.Delete(taskTable).Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		logrus.Fatalf("error while creatinsg delete query: %s", err.Error())
		return fmt.Errorf("build delete query error: %w", err)
	}

	_, err = r.db.Exec(ctx.Context(), query, args...)
	if err != nil {
		logrus.Fatalf("error while deleting task entity: %s", err.Error())
		return fmt.Errorf("delete task error: %w", err)
	}

	return nil
}

func (r *TaskRepository) getUpdateData(task models.Task) map[string]any {
	updateData := map[string]any{}

	if task.Title != "" {
		updateData["title"] = task.Title
	}
	if task.Description != "" {
		updateData["description"] = task.Description
	}
	if task.Status != "" {
		updateData["status"] = task.Status
	}
	if !time.Time.IsZero(task.CreatedAt) {
		updateData["created_at"] = task.CreatedAt
	}

	updateData["updated_at"] = time.Now()

	return updateData
}
