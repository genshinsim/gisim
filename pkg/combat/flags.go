package combat

import "github.com/genshinsim/gsim/pkg/def"

func (s *Sim) Flags() def.Flags { return s.flags }

func (s *Sim) SetCustomFlag(key string, val int) {
	s.flags.Custom[key] = val
}

func (s *Sim) GetCustomFlag(key string) (int, bool) {
	val, ok := s.flags.Custom[key]
	return val, ok
}
