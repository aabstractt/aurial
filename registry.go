package aurial

import (
	"sync"
	"sync/atomic"
)

type Registry[V any] struct {
	values map[int32]V
	mu     sync.RWMutex

	number *atomic.Int32
}

// Store stores a value in the registry and returns the index at which it was stored.
func (r *Registry[V]) Store(v V) int32 {
	r.mu.Lock()

	id := r.number.Add(1)
	r.values[id] = v

	r.mu.Unlock()

	return id
}

// Remove removes a value from the registry by its index.
func (r *Registry[V]) Remove(index int32) {
	r.mu.Lock()
	delete(r.values, index)
	r.mu.Unlock()
}

func (r *Registry[V]) all() []V {
	r.mu.RLock()
	defer r.mu.RUnlock()

	values := make([]V, 0, len(r.values))
	for _, v := range r.values {
		values = append(values, v)
	}

	return values
}

func newRegistry[V any]() *Registry[V] {
	return &Registry[V]{
		values: make(map[int32]V),
	}
}

func BlockBreakRegistry() *Registry[BlockBreakHandler] {
	return blockBreakRegistry
}

func BlockPlaceRegistry() *Registry[BlockPlaceHandler] {
	return blockPlaceRegistry
}

func AttackEntityRegistry() *Registry[AttackEntityHandler] {
	return attackEntityRegistry
}

func HurtRegistry() *Registry[HurtHandler] {
	return hurtRegistry
}

func DeathRegistry() *Registry[DeathHandler] {
	return deathRegistry
}

func FoodLossRegistry() *Registry[FoodLossHandler] {
	return foodLossRegistry
}

func HealRegistry() *Registry[HealHandler] {
	return healRegistry
}

func ChatRegistry() *Registry[ChatHandler] {
	return chatRegistry
}

func ChangeWorldRegistry() *Registry[ChangeWorldHandler] {
	return changeWorldRegistry
}

func TeleportRegistry() *Registry[TeleportHandler] {
	return teleportRegistry
}

func MoveRegistry() *Registry[MoveHandler] {
	return moveRegistry
}
