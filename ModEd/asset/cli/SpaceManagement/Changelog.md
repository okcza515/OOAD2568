# MEP-1013

## Progress

- Refactored to use a Base Controller.
- Applied Core Handler Strategy.
- Implemented Space Management Handler Strategy (separate from Core).
- Defined interfaces for methods in each controller.
- Applied Singleton pattern to controller instances.
- Migrated Facade-based menu selector to a State Manager.

## Design Patterns

### Facade

- `SpaceManagementControllerManager` acts as a facade, providing a simplified interface to complex subsystems.

### Singleton

- `SpaceManagementControllerManager` is implemented as a singleton using `GetSpaceManagementInstance()` to ensure a single shared instance.

### State

- Implemented in the CLI menu system to manage transitions between various menu states.
- `CLIMenuStateManager` handles state transitions cleanly and modularly.

### Strategy

- Used in handler classes to encapsulate behavior without modifying client code.
- Ensures consistency with the core service structure.

<!-- ### Command

- `ToString()` methods are used to standardize and encapsulate UI output logic, following the Command pattern principles. -->


Last update (before 13th May)

- add seed data in migration

- Permanent Schedule
    - All function can run.
    - Cannot link data with Faculty and Department Model yet.
