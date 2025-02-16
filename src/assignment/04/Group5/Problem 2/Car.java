//65070501018 Natchanon Phattamanuruk
public class Car {
    public Type type;
    public int seats;
    public Engine engine;
    public Transmission transmission;
    public TripComputer tripComputer;
    public GPSNavigator gpsNavigator;
    public double fuel = 0;

    public Type getType() {
        return type;
    }
    public void setType(Type type) {
        this.type = type;
    }

    public double getFuel() {
        return fuel;
    }
    public void setFuel(double fuel) {
        this.fuel = fuel;
    }

    public int getSeats() {
        return seats;
    }
    public void setSeats(int seats) {
        this.seats = seats;
    }

    public Engine getEngine() {
        return engine;
    }
    public void setEngine(Engine engine) {
        this.engine = engine;
    }

    public Transmission getTransmission() {
        return transmission;
    }
    public void setTransmission(Transmission transmission) {
        this.transmission = transmission;
    }

    public TripComputer getTripComputer() {
        return tripComputer;
    }
    public void setTripComputer(TripComputer tripComputer) {
        this.tripComputer = tripComputer;
    }

    public GPSNavigator getGpsNavigator() {
        return gpsNavigator;
    }
    public void setGpsNavigator(GPSNavigator gpsNavigator) {
        this.gpsNavigator = gpsNavigator;
    }

    public String toString(){
        return "Car [type = " + type + ", Number of seats = " + seats + ", Engine: volume - " + engine.getVolume() + "; mileage - " + engine.getMileage() + ", Transmission = "
                + transmission + "]";
    }
}
