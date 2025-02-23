class ConcreteCarBuilder implements CarBuilder {
    protected Type carType;
    protected int seats;
    protected Engine engine;
    protected Transmission transmission;
    protected TripComputer tripComputer;
    protected GPSNavigator gpsNavigator;

    @Override
    public void setCarType(Type carType) {
        this.carType = carType;
    }   

    @Override
    public void setSeats(int seats) {
        this.seats = seats;
    }

    @Override
    public void setEngine(Engine engine) {
        this.engine = engine;
    }

    @Override
    public void setTransmission(Transmission transmission) {
        this.transmission = transmission;
    }

    @Override
    public void setTripComputer(TripComputer tripComputer) {
        this.tripComputer = tripComputer;
    }

    @Override
    public void setGPSNavigator(GPSNavigator gpsNavigator) {
        this.gpsNavigator = gpsNavigator;
    }

    public Car getCar() {
        return new Car(carType, seats, engine, transmission, tripComputer, gpsNavigator);
    }

}
// 65070503412 Chitsanupong