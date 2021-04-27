package discov

import (
	"github.com/tal-tech/go-zero/core/discov/internal"
	"github.com/tal-tech/go-zero/core/lang"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/proc"
	"github.com/tal-tech/go-zero/core/syncx"
	"github.com/tal-tech/go-zero/core/threading"
	"go.etcd.io/etcd/clientv3"
)

type (
	// PublisherOption defines the method to customize a Publisher.
	PublisherOption func(client *Publisher)

	// A Publisher can be used to publish the value to an etcd cluster on the given key.
	Publisher struct {
		endpoints  []string
		key        string
		fullKey    string
		id         int64
		value      string
		lease      clientv3.LeaseID
		quit       *syncx.DoneChan
		pauseChan  chan lang.PlaceholderType
		resumeChan chan lang.PlaceholderType
	}
)

// NewPublisher returns a Publisher.
// endpoints is the hosts of the etcd cluster.
// key:value are a pair to be published.
// opts are used to customize the Publisher.
func NewPublisher(endpoints []string, key, value string, opts ...PublisherOption) *Publisher {
	publisher := &Publisher{
		endpoints:  endpoints,
		key:        key,
		value:      value,
		quit:       syncx.NewDoneChan(),
		pauseChan:  make(chan lang.PlaceholderType),
		resumeChan: make(chan lang.PlaceholderType),
	}

	for _, opt := range opts {
		opt(publisher)
	}

	return publisher
}

// KeepAlive keeps key:value alive.
func (p *Publisher) KeepAlive() error {
	cli, err := internal.GetRegistry().GetConn(p.endpoints)
	if err != nil {
		return err
	}

	p.lease, err = p.register(cli)
	if err != nil {
		return err
	}

	proc.AddWrapUpListener(func() {
		p.Stop()
	})

	return p.keepAliveAsync(cli)
}

// Pause pauses the renewing of key:value.
func (p *Publisher) Pause() {
	p.pauseChan <- lang.Placeholder
}

// Resume resumes the renewing of key:value.
func (p *Publisher) Resume() {
	p.resumeChan <- lang.Placeholder
}

// Stop stops the renewing and revokes the registration.
func (p *Publisher) Stop() {
	p.quit.Close()
}

func (p *Publisher) keepAliveAsync(cli internal.EtcdClient) error {
	ch, err := cli.KeepAlive(cli.Ctx(), p.lease)
	if err != nil {
		return err
	}

	threading.GoSafe(func() {
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					p.revoke(cli)
					if err := p.KeepAlive(); err != nil {
						logx.Errorf("KeepAlive: %s", err.Error())
					}
					return
				}
			case <-p.pauseChan:
				logx.Infof("paused etcd renew, key: %s, value: %s", p.key, p.value)
				p.revoke(cli)
				select {
				case <-p.resumeChan:
					if err := p.KeepAlive(); err != nil {
						logx.Errorf("KeepAlive: %s", err.Error())
					}
					return
				case <-p.quit.Done():
					return
				}
			case <-p.quit.Done():
				p.revoke(cli)
				return
			}
		}
	})

	return nil
}

func (p *Publisher) register(client internal.EtcdClient) (clientv3.LeaseID, error) {
	resp, err := client.Grant(client.Ctx(), TimeToLive)
	if err != nil {
		return clientv3.NoLease, err
	}

	lease := resp.ID
	if p.id > 0 {
		p.fullKey = makeEtcdKey(p.key, p.id)
	} else {
		p.fullKey = makeEtcdKey(p.key, int64(lease))
	}
	_, err = client.Put(client.Ctx(), p.fullKey, p.value, clientv3.WithLease(lease))

	return lease, err
}

func (p *Publisher) revoke(cli internal.EtcdClient) {
	if _, err := cli.Revoke(cli.Ctx(), p.lease); err != nil {
		logx.Error(err)
	}
}

// WithId customizes a Publisher with the id.
func WithId(id int64) PublisherOption {
	return func(publisher *Publisher) {
		publisher.id = id
	}
}
