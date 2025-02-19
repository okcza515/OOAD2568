public class CityCarMaker implements CarMaker{
    private Type type;
    private int seats;
    private Engine engine;
    private Transmission transmission;
    private TripComputer tripComputer;
    private GPSNavigator gpsNavigator;

    @Override
    public CarMaker setCarType(Type type) {
        this.type = type;
        return this;
    }

    @Override
    public CarMaker setSeats(int seats){
        this.seats = seats;
        return this;
    }

    @Override
    public CarMaker setEngine(Engine engine){
        this.engine = engine;
        return this;
    }

    @Override
    public CarMaker setTransmission(Transmission transmission){
        this.transmission = transmission;
        return this;
    }

    @Override
    public CarMaker setTripComputer(TripComputer tripComputer){
        this.tripComputer = tripComputer;
        return this;
    }

    @Override
    public CarMaker setGPSNavigator(GPSNavigator gpsNavigator){
        this.gpsNavigator = gpsNavigator;
        return this;
    }

    @Override
    public Car getCar(){
        return new Car(type, seats, engine, transmission, tripComputer, gpsNavigator);
    }
}

//Korawit Sritotum 65070503402