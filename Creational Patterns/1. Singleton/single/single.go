package single

import (
	"sync"

	"github.com/rs/zerolog/log"
)

type Singleton struct{}

var (
	singletonInstance *Singleton
	lock              sync.Mutex
)

func GetInstance() *Singleton {
	if singletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonInstance == nil {
			singletonInstance = &Singleton{}
			log.Info().Msg("Created Singleton instance.")
			return singletonInstance
		} else {
			log.Info().Msg("Singleton instance already created.")
			return singletonInstance
		}
	} else {
		log.Info().Msg("Singleton instance already created.")
		return singletonInstance
	}
}
