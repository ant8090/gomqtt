// Copyright (c) 2014 The gomqtt Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package broker

import (
	"github.com/gomqtt/transport"
)

type Logger func(msg string)

type Broker struct {
	QueueBackend    QueueBackend
	RetainedBackend RetainedBackend
	WillBackend     WillBackend

	Logger Logger
}

// New returns a new Broker with a basic MemoryBackend.
func New() *Broker {
	backend := NewMemoryBackend()

	return &Broker{
		QueueBackend:    backend,
		RetainedBackend: backend,
		WillBackend:     backend,
	}
}

// Handle handles a transport.Conn.
func (b *Broker) Handle(conn transport.Conn) {
	NewClient(b, conn)
}
