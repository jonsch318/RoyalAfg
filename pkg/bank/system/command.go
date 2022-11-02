package system

type ICommand interface {
	GetAggregateId() string
	GetType() string
	GetData() interface{}
}
type Command struct {
	AggregateId string
	Type        string
	Data        interface{}
}

func (c *Command) GetAggregateId() string {
	return c.AggregateId
}

func (c *Command) GetType() string {
	return c.Type
}

func (c *Command) GetData() interface{} {
	return c.Data
}
