package main

import "context"

type Publisher struct {
	ctx         context.Context
	subscribeCh chan chan<- string //chan <- string 은 string 타입을 넣기만 할 수 있는 write only채널을 의미한다. 이 write only 채널을 받는 채널을 subscribe 채널이라 한다. <- chan 하면 read only임. chan 하면 그냥 양방향 채널임
	publishCh   chan string
	subscribers []chan<- string
}

func NewPublisher(ctx context.Context) *Publisher {
	return &Publisher{
		ctx:         ctx,
		subscribeCh: make(chan chan<- string),
		publishCh:   make(chan string),
		subscribers: make([]chan<- string, 0), //subcriber에서 subscribe 하면 최종적으로 여기에 추가됨.
	}
}

func (p *Publisher) Subscribe(sub chan<- string) {
	p.subscribeCh <- sub
}

func (p *Publisher) Publish(msg string) {
	p.publishCh <- msg
}

func (p *Publisher) Update() {
	for {
		select {
		case sub := <-p.subscribeCh:
			p.subscribers = append(p.subscribers, sub)
		case msg := <-p.publishCh:
			for _, subscriber := range p.subscribers {
				subscriber <- msg
			}
		case <-p.ctx.Done():
			wg.Done()
			return
		}
	}
}
