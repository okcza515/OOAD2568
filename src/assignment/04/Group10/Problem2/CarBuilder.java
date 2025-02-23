// 65070501011 Chayapol Wongpuwarak
public abstract  class CarBuilder {
    protected Type type;
    protected int seats;
    protected Engine engine;
    protected Transmission transmission;
    protected TripComputer tripComputer;
    protected GPSNavigator gpsNavigator;

    public CarBuilder setType(Type type) {
        this.type = type;
        return this;
    }

    public CarBuilder setSeats(int seats){
        this.seats = seats;
        return this;
    }

    public CarBuilder setEngine(Engine engine) {
        this.engine = engine;
        return this;
    }

    public CarBuilder setTransmission(Transmission transmission) {
        this.transmission = transmission;
        return this;
    }

    public CarBuilder setTripComputer(TripComputer tripComputer) {
        this.tripComputer = tripComputer;
        return this;
    }

    public CarBuilder setGPSNavigator(GPSNavigator gpsNavigator) {
        this.gpsNavigator = gpsNavigator;
        return this;
    }

    public Car buildCar() {
        return new Car(this.type, this.seats, this.engine, this.transmission, this.tripComputer, this.gpsNavigator);
    }
}
