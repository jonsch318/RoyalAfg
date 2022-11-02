package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/bank/system"
	"github.com/gofrs/uuid"
)

type EventStoreRepository[T system.ICreatableAggregate] struct {
	client      *esdb.Client
	eventParser system.IEventParser[esdb.RecordedEvent]
}

func NewEventStoreRepository[T system.ICreatableAggregate](client *esdb.Client, eventParser system.IEventParser[esdb.RecordedEvent]) *EventStoreRepository[T] {
	esr := &EventStoreRepository[T]{
		client:      client,
		eventParser: eventParser,
	}
	return esr
}

func (esr *EventStoreRepository[T]) Load(ctx context.Context, id string) (*T, error) {
	ropt := esdb.ReadStreamOptions{
		From:      esdb.Start{},
		Direction: esdb.Forwards,
	}

	stream, err := esr.client.ReadStream(ctx, id, ropt, 100)

	if err != nil {
		return nil, fmt.Errorf("error reading stream: %w", err)
	}

	defer stream.Close()

	var aggregate T

	events := make([]system.IEvent[any], 0)

	for {
		rawEvent, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("error reading stream: %w", err)
		}

		event, err := esr.eventParser.Parse(rawEvent.Event)
		events = append(events, event)
	}
	aggregate.CreateFromEvents(events)

	return &aggregate, nil
}

func (esr *EventStoreRepository[T]) Save(ctx context.Context, aggregate T) error {
	id := aggregate.GetId()

	aopt := esdb.AppendToStreamOptions{
		ExpectedRevision: esdb.NoStream{},
	}

	events := make([]esdb.EventData, len(aggregate.GetChanges()))
	for i, v := range aggregate.GetChanges() {
		contentType := esdb.JsonContentType
		if v.GetContentType() != "application/json" {
			contentType = esdb.BinaryContentType
		}

		events[i] = esdb.EventData{
			EventID:     uuid.Must(uuid.NewV4()),
			EventType:   v.GetType(),
			Data:        v.GetDataRaw(),
			ContentType: contentType,
			Metadata:    generateMetadata(aggregate),
		}
	}

	_, err := esr.client.AppendToStream(ctx, id, aopt, events...)
	return err
}

func (esr *EventStoreRepository[T]) Delete(ctx context.Context, id string) error {

	_, err := esr.client.DeleteStream(ctx, id, esdb.DeleteStreamOptions{})

	return err
}

func generateMetadata[T system.ICreatableAggregate](aggregate T) []byte {
	metadata := make(map[string]interface{})
	metadata["aggregateType"] = aggregate.GetType()
	metadata["aggregateId"] = aggregate.GetId()
	metadata["aggregateVersion"] = aggregate.GetVersion()
	if b, err := json.Marshal(metadata); err != nil {
		return nil
	} else {
		return b
	}
}
