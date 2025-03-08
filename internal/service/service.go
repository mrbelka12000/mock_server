package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/mrbelka12000/mock_server/internal"
	"github.com/mrbelka12000/mock_server/pkg/pointer"
)

type (
	Service struct {
		store store
	}

	store interface {
		GetHandlersByServiceID(ctx context.Context, serviceID int64) ([]internal.Handler, error)
		GetServiceByName(ctx context.Context, serviceName string) (internal.Service, error)
		GetServiceByID(ctx context.Context, id int64) (internal.Service, error)
		AddService(ctx context.Context, service internal.Service) error
		ListServices(ctx context.Context) ([]internal.Service, error)
		AssignHandlerToService(ctx context.Context, serviceID int64, handler internal.HandlerCU) (int64, error)
		AssignCasesToHandler(ctx context.Context, handlerID int64, cases []internal.HandlerCasesCU) error
	}
)

func New(
	store store,
) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) HandleRoute(ctx context.Context, serviceName, route string, reqBody []byte, reqHeaders http.Header) ([]byte, internal.Header, error) {
	srv, err := s.store.GetServiceByName(ctx, serviceName)
	if err != nil {
		return nil, nil, fmt.Errorf("get srv: %w", err)
	}

	var (
		h     internal.Handler
		found bool
	)

	for i := 0; i < len(srv.Handlers); i++ {
		if srv.Handlers[i].Route == route {
			found = true
			h = srv.Handlers[i]
			break
		}
	}

	if !found {
		return nil, nil, fmt.Errorf("handler not found: %s", route)
	}

	for _, cs := range h.Cases {
		if cs.Tag == internal.TagEqual {
			if string(reqBody) == pointer.Value(cs.RequestBody) {
				return []byte(pointer.Value(cs.ResponseBody)), pointer.Value(cs.ResponseHeaders), nil
			}
		}
	}

	for _, cs := range h.Cases {
		if cs.Tag == internal.TagDefault {
			return []byte(pointer.Value(cs.ResponseBody)), pointer.Value(cs.ResponseHeaders), nil
		}
	}

	return nil, nil, fmt.Errorf("no suitable case for %s", route)
}

func (s *Service) AddService(ctx context.Context, srv internal.Service) error {
	if srv.Name == "" {
		return errors.New("service name is required")
	}

	return s.store.AddService(ctx, srv)
}

func (s *Service) GetServiceByID(ctx context.Context, serviceID int64) (internal.Service, error) {
	return s.store.GetServiceByID(ctx, serviceID)
}

func (s *Service) ListServices(ctx context.Context) ([]internal.Service, error) {
	return s.store.ListServices(ctx)
}

func (s *Service) AssignHandlerToService(ctx context.Context, obj internal.HandlerCU) error {
	if len(obj.Cases) == 0 {
		return errors.New("handler must have at least one case")
	}
	if obj.ServiceID == nil {
		return errors.New("service id is required")
	}
	if obj.Route == nil {
		return errors.New("route is required")
	}

	for i, cs := range obj.Cases {
		if cs.Tag == nil {
			obj.Cases[i].Tag = pointer.Of(internal.TagDefault)
		}

		if pointer.Value(obj.Cases[i].Tag) != internal.TagEqual && pointer.Value(obj.Cases[i].Tag) != internal.TagDefault {
			return errors.New("handler must have tag equal or default")
		}
	}

	handlerID, err := s.store.AssignHandlerToService(ctx, pointer.Value(obj.ServiceID), obj)
	if err != nil {
		return fmt.Errorf("assign handler to service: %w", err)
	}

	err = s.store.AssignCasesToHandler(ctx, handlerID, obj.Cases)
	if err != nil {
		return fmt.Errorf("assign cases to handlers: %w", err)
	}

	return nil
}
