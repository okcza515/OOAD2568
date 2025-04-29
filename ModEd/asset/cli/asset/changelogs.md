# ✨ Asset Update 18 March

MEP1012 - Asset

## Progress
1. Write description/requirement
2. Design data model
2. Implement Models
3. Try out migration
4. Try to make controller works
5. Try using CLI
5. Discuss with dependency friends

## Issues
Fighting with friends (need more discussion)

## Next
1. Reorganized repository

___

# ✨ Asset Update 22 April

MEP1012 - Asset

## Done
1. Reorganized repository to flat package (again)
2. Somehow, convinced everyone to use `Facade` design pattern to simplify usage between module
3. Add `BaseController` to every controller in Asset (Reusability cool!)
4. Add `InsertMany` method to `BaseController` (which is heavily questionable and would like to open a discussion in class)
5. Create interface for every controller (Interface Segregation: I in SOLID)
6. Create `BaseModel` for everyone to be able to use `RecordInterface` (combined `gorm.Model`, `SerializableRecord`, and some default method)
7. Create CLI layer according to the functional requirement
8. Use `Dependency Injection` design pattern in CLI layer (not in refactoring.guru for some reason)

## Ongoing
1. Add more functionality to CLI and Controller


## Issues
1. We're not sure if my CLI is correct because we saw other group have different solution than our group
2. We don't know how to implement `Record` `RecordFactory` functionality in our model/controller. Thus, we waste a lot of time trying to understand it
3. We might find a "bug" in the `RecordFactory` e.g. pointer to interface, and uninitialized model mapper
4. In CLI package, we can't organize the repository to be within single module due to duplicated `main()` in different module
5. `RetrieveByID` method in `BaseController` need to perform a type assertion after retrieved. However, this issue can be evaded with another approach. (would like to discuss)
6. Due to we reorganized all the package together. The private/public method/variable is not working anymore. Everyone can access to everything.

## Next
1. Keep the #Ongoing
2. idk solve the current problem is hard enough :sad_face:

___

# ✨ Asset Update 29 April
MEP-1012 Asset

### Done
1. Refactor CLI layer using `State` DP
2. Refactor handler using `Strategy` DP
3. Implemented a core migration with `Singleton`, `Builder`, `Strategy`
4. Refactor `AssetControllerFacade` into a `AssetControllerManager`. Instead of using `Facade` DP, we simply encapsulate it with `Singleton`.
5. Implemented cross-controller updated using `Observer` DP
6. Add preloads into `HandlerStrategy`

### Actual New Feature
1. Instrument `Update`
2. Instrument Log `List`, `RetrieveByID`
3. Supply `Insert`,`List`,`Update`,`Delete`

# Implemented Design Pattern Or Principles

### Encapsulation (OOP Principles)
- **AssetControllerManager**: encapsulate entire asset submodule controller into one giant object to be easily accessed by caller. 

### Singleton
- AssetControllerManager
- Migration (core)

### Builder
- Migration (core)

### Strategy
- CLI Handler (core)
- Migration (core)

### State
- CLI Menu (core)


### Observer
- AssetControllerManager

<img src="https://images.ctfassets.net/glsfy1cpffmh/632YZxkYQlLzXz6evBYIHW/d695efc19e73937ac598b046240f69c1/the-singleton-of-dufftown-malt-masters-selection-single-malt-scotch-whisky-70cl-1-transparent.png?fm=webp&w=1920&q=30" alt="My Image" width="300" height="300">
<img src="https://www.pngall.com/wp-content/uploads/15/Bob-The-Builder-PNG.png" alt="My Image" width="160" height="200">
<img src="https://w7.pngwing.com/pngs/773/577/png-transparent-technology-strategy-organization-marketing-strategy-strategic-planning-strategy-love-text-people-thumbnail.png" alt="My Image" width="280" height="200">

.

<img src="https://png.pngtree.com/png-vector/20230904/ourmid/pngtree-united-states-map-america-png-image_9933457.png" alt="My Image" width="300" height="200">
<img src="https://static.wikia.nocookie.net/minecraft_gamepedia/images/3/39/Observer_JE4_BE3.png" alt="My Image" width="200" height="200">

---