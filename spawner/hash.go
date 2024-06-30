package spawner

import "math/rand"

var (
	hashStart   = rand.Uint64()
	hashSpawner = hashStart + 1
)

func (s Spawner) Hash() uint64 {
	return hashSpawner
}
