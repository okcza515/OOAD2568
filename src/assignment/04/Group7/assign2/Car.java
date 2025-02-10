
public class Car {
    public Type type;
    public int seats;
    public Engine engine;
    public Transmission transmission;
    public TripComputer tripComputer;
    public GPSNavigator gpsNavigator;
    public double fuel = 0;

    // public Car(Type type, int seats, Engine engine, Transmission transmission,
    //            TripComputer tripComputer, GPSNavigator gpsNavigator) {
    //     this.type = type;
    //     this.seats = seats;
    //     this.engine = engine;
    //     this.transmission = transmission;
    //     this.tripComputer = tripComputer;
    //     this.tripComputer.setCar(this);
    //     this.gpsNavigator = gpsNavigator;
    // }

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

    public void getSeats(int seats) {
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

    public TripComputer getTripComputer() {
        return tripComputer;
    }

    public GPSNavigator getGpsNavigator() {
        return gpsNavigator;
    }
}

//Jaatupoj 65070501070