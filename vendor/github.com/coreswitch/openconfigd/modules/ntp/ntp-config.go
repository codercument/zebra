// Generated by yang2go from coreswitch-ntp.yang.
//
// Copyright 2017 OpenConfigd Project.
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

package ntp

import (
	"fmt"
)

func mapkey(index int, name string) string {
	if name != "" {
		return name
	}
	return fmt.Sprintf("%v", index)
}

//struct for container coreswitch-ntp:server
type Server struct {
	// original -> coreswitch-ntp:name
	Name string `mapstructure:"name" json:"name,omitempty"`
	// original -> coreswitch-ntp:dynamic
	//coreswitch-ntp:dynamic's original type is empty
	Dynamic bool `mapstructure:"dynamic" json:"dynamic,omitempty"`
	// original -> coreswitch-ntp:noselect
	//coreswitch-ntp:noselect's original type is empty
	Noselect bool `mapstructure:"noselect" json:"noselect,omitempty"`
	// original -> coreswitch-ntp:preempt
	//coreswitch-ntp:preempt's original type is empty
	Preempt bool `mapstructure:"preempt" json:"preempt,omitempty"`
	// original -> coreswitch-ntp:prefer
	//coreswitch-ntp:prefer's original type is empty
	Prefer bool `mapstructure:"prefer" json:"prefer,omitempty"`
}

func (lhs *Server) Equal(rhs *Server) bool {
	if lhs == nil || rhs == nil {
		return false
	}
	if lhs.Name != rhs.Name {
		return false
	}
	if lhs.Dynamic != rhs.Dynamic {
		return false
	}
	if lhs.Noselect != rhs.Noselect {
		return false
	}
	if lhs.Preempt != rhs.Preempt {
		return false
	}
	if lhs.Prefer != rhs.Prefer {
		return false
	}
	return true
}

//struct for container coreswitch-ntp:pool
type Pool struct {
	// original -> coreswitch-ntp:name
	Name string `mapstructure:"name" json:"name,omitempty"`
	// original -> coreswitch-ntp:dynamic
	//coreswitch-ntp:dynamic's original type is empty
	Dynamic bool `mapstructure:"dynamic" json:"dynamic,omitempty"`
	// original -> coreswitch-ntp:noselect
	//coreswitch-ntp:noselect's original type is empty
	Noselect bool `mapstructure:"noselect" json:"noselect,omitempty"`
	// original -> coreswitch-ntp:preempt
	//coreswitch-ntp:preempt's original type is empty
	Preempt bool `mapstructure:"preempt" json:"preempt,omitempty"`
	// original -> coreswitch-ntp:prefer
	//coreswitch-ntp:prefer's original type is empty
	Prefer bool `mapstructure:"prefer" json:"prefer,omitempty"`
}

func (lhs *Pool) Equal(rhs *Pool) bool {
	if lhs == nil || rhs == nil {
		return false
	}
	if lhs.Name != rhs.Name {
		return false
	}
	if lhs.Dynamic != rhs.Dynamic {
		return false
	}
	if lhs.Noselect != rhs.Noselect {
		return false
	}
	if lhs.Preempt != rhs.Preempt {
		return false
	}
	if lhs.Prefer != rhs.Prefer {
		return false
	}
	return true
}

//struct for container coreswitch-ntp:ntp
type Ntp struct {
	// original -> coreswitch-ntp:pool
	PoolList []Pool `mapstructure:"pool" json:"pool,omitempty"`
	// original -> coreswitch-ntp:server
	ServerList []Server `mapstructure:"server" json:"server,omitempty"`
}

func (lhs *Ntp) Equal(rhs *Ntp) bool {
	if lhs == nil || rhs == nil {
		return false
	}
	if len(lhs.PoolList) != len(rhs.PoolList) {
		return false
	}
	{
		lmap := make(map[string]*Pool)
		for i, l := range lhs.PoolList {
			lmap[mapkey(i, string(l.Name))] = &lhs.PoolList[i]
		}
		for i, r := range rhs.PoolList {
			if l, y := lmap[mapkey(i, string(r.Name))]; !y {
				return false
			} else if !r.Equal(l) {
				return false
			}
		}
	}
	if len(lhs.ServerList) != len(rhs.ServerList) {
		return false
	}
	{
		lmap := make(map[string]*Server)
		for i, l := range lhs.ServerList {
			lmap[mapkey(i, string(l.Name))] = &lhs.ServerList[i]
		}
		for i, r := range rhs.ServerList {
			if l, y := lmap[mapkey(i, string(r.Name))]; !y {
				return false
			} else if !r.Equal(l) {
				return false
			}
		}
	}
	return true
}