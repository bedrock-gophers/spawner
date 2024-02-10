# Spawner
The Spawner library simplifies the creation and management of mob spawners within a Minecraft server. It offers an intuitive interface for placing spawners and controlling mob spawning behavior. Notably, the library includes support for entity stacking, which provides significant performance benefits for both the server and client.

## Placing a Spawner
To place a spawner in your Minecraft world using the Spawner library, follow these steps:

Create Spawner: Utilize the spawner.New function to create a new spawner. Specify parameters such as the entity type, position, world, spawn delay, spawn range, and stacking status.

```go
pos := cube.Pos{x, y, z}
s := spawner.New(newEnderman, pos.Vec3Centre(), p.World(), time.Second, 64, true)
```
Register Spawner Block: Ensure the spawner block is registered with the world to enable proper functioning.

```go
world.RegisterBlock(s)
```
Set Spawner Block in World: Set the spawner block at the desired position within the world.

```go
p.World().SetBlock(pos, s, nil)
```
## Entity Stacking
The Spawner library introduces entity stacking, a feature that allows multiple entities of the same type to spawn to a single entity. This capability significantly enhances server and client performance by reducing the number of active entities and optimizing resource utilization.

## Benefits of Entity Stacking
Improved Server Performance: Entity stacking reduces the overall number of entities spawned in the world, leading to improved server performance and decreased resource consumption.

Enhanced Client Performance: With fewer entities to render and process, clients experience smoother gameplay and reduced lag, resulting in a more enjoyable gaming experience.

## Requirements
To utilize Spawner, you must use our Living library for creating your spawnable entities. Entity stacking functionality seamlessly integrates into the Spawner library, requiring no additional configuration.
