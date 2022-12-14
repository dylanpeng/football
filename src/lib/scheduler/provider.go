package scheduler

import "fmt"

type IProvider interface {
	GetName() string
	GetCronExpression() string
	Run()
	String() string
}

type Provider struct {
	Name           string `json:"name"`
	CronExpression string `json:"cron_expression"`
}

func (p *Provider) GetName() string {
	return p.Name
}

func (p *Provider) GetCronExpression() string {
	return p.CronExpression
}

func (p *Provider) String() string {
	return fmt.Sprintf("%+v", *p)
}
