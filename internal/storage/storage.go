package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mrbelka12000/mock_server/internal"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetServiceByName(ctx context.Context, serviceName string) (internal.Service, error) {
	var srv internal.Service

	err := s.db.QueryRowContext(ctx, `
SELECT id, name 
FROM services 
WHERE name = $1`, serviceName).Scan(&srv.ID, &srv.Name)
	if err != nil {
		return srv, fmt.Errorf("get service by name %s: %w", serviceName, err)
	}

	srv.Handlers, err = s.GetHandlersByServiceID(ctx, srv.ID)
	if err != nil {
		return srv, fmt.Errorf("get handlers: %w", err)
	}

	return srv, nil
}

func (s *Store) GetHandlersByServiceID(ctx context.Context, serviceID int64) ([]internal.Handler, error) {
	rows, err := s.db.QueryContext(ctx, `
	SELECT id, service_id, route 
	FROM handlers
	WHERE service_id = $1`, serviceID)
	if err != nil {
		return nil, fmt.Errorf("get handlers by service id %d: %w", serviceID, err)
	}
	defer rows.Close()

	var handlers []internal.Handler
	for rows.Next() {
		var handler internal.Handler

		if err := rows.Scan(&handler.ID, &handler.ServiceID, &handler.Route); err != nil {
			return nil, fmt.Errorf("scan handler: %w", err)
		}

		handler.Cases, err = s.GetCasesByHandlerID(ctx, handler.ID)
		if err != nil {
			return nil, fmt.Errorf("get cases by handler: %w", err)
		}

		handlers = append(handlers, handler)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	return handlers, nil
}

func (s *Store) GetServiceByID(ctx context.Context, id int64) (internal.Service, error) {
	var srv internal.Service
	err := s.db.QueryRowContext(ctx, `
SELECT id, name
FROM services
WHERE id = $1`, id).Scan(&srv.ID, &srv.Name)
	if err != nil {
		return srv, fmt.Errorf("get service by id %d: %w", id, err)
	}

	srv.Handlers, err = s.GetHandlersByServiceID(ctx, srv.ID)
	if err != nil {
		return srv, fmt.Errorf("get handlers by service name %d: %w", id, err)
	}

	return srv, nil
}

func (s *Store) AddService(ctx context.Context, service internal.Service) error {
	_, err := s.db.Exec(`
INSERT INTO services
(name) 
VALUES
($1)`, service.Name)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}

func (s *Store) ListServices(ctx context.Context) ([]internal.Service, error) {
	rows, err := s.db.QueryContext(ctx, `
SELECT id, name
FROM services
`)
	if err != nil {
		return nil, fmt.Errorf("list services: %w", err)
	}
	defer rows.Close()

	var services []internal.Service
	for rows.Next() {
		var service internal.Service
		if err := rows.Scan(&service.ID, &service.Name); err != nil {
			return nil, fmt.Errorf("scan service: %w", err)
		}

		service.Handlers, err = s.GetHandlersByServiceID(ctx, service.ID)
		if err != nil {
			return nil, fmt.Errorf("get handlers by service id %d: %w", service.ID, err)
		}

		services = append(services, service)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	return services, nil
}

func (s *Store) AssignHandlerToService(ctx context.Context, serviceID int64, obj internal.HandlerCU) (int64, error) {
	var id int64

	err := s.db.QueryRowContext(ctx, `
INSERT INTO handlers
(service_id, route)
VALUES 
($1, $2)
RETURNING id
`, serviceID, obj.Route).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("insert handler: %w", err)
	}

	return id, nil
}

func (s *Store) AssignCasesToHandler(ctx context.Context, handlerID int64, cases []internal.HandlerCasesCU) error {

	for _, cs := range cases {
		_, err := s.db.ExecContext(ctx, `
INSERT INTO handler_cases
(handler_id, tag_case, request_body, response_body, request_headers, response_headers) 
VALUES 
($1, $2, $3, $4, $5, $6)
`, handlerID, cs.Tag, cs.RequestBody, cs.ResponseBody, cs.RequestHeaders, cs.ResponseHeaders)
		if err != nil {
			return fmt.Errorf("insert handler_cases: %w", err)
		}
	}

	return nil
}

func (s *Store) GetCasesByHandlerID(ctx context.Context, handlerID int64) ([]internal.HandlerCases, error) {
	rows, err := s.db.QueryContext(ctx, `
SELECT id, handler_id, tag_case, request_body, response_body, request_headers, response_headers
FROM handler_cases
WHERE handler_id = $1`, handlerID)
	if err != nil {
		return nil, fmt.Errorf("get cases by handler_id %d: %w", handlerID, err)
	}

	defer rows.Close()

	var cases []internal.HandlerCases

	for rows.Next() {
		var cs internal.HandlerCases
		if err = rows.Scan(
			&cs.ID,
			&cs.HandlerID,
			&cs.Tag,
			&cs.RequestBody,
			&cs.ResponseBody,
			&cs.RequestHeaders,
			&cs.ResponseHeaders,
		); err != nil {
			return nil, fmt.Errorf("scan handler: %w", err)
		}

		cases = append(cases, cs)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	return cases, nil
}

func (s *Store) DeleteCase(ctx context.Context, id int64) error {
	_, err := s.db.Exec(`
DELETE FROM handler_cases
WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete handler_cases: %w", err)
	}

	return nil
}
