public interface CarMaker {
    CarMaker setCarType(Type type);

    CarMaker setSeats(int seats);

    CarMaker setEngine(Engine engine);

    CarMaker setTransmission(Transmission transmission);

    CarMaker setTripComputer(TripComputer tripComputer);

    CarMaker setGPSNavigator(GPSNavigator gpsNavigator);

    Car getCar();

    Type getCarType();

    int getSeats();

    Engine getEngine();

    Transmission getTransmission();

    TripComputer getTripComputer();

    GPSNavigator getGPSNavigator();

}

// Korawit Sritotum 65070503402
