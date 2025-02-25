// CarDirector.java //Paratthakon Suksukhon 65070503457
public class CarDirector {
    public void constructSportsCar(CarBuilder builder) {
        builder.setType(Type.SPORTS_CAR)
               .setSeats(2)
               .setEngine(new Engine(2.0, 0))
               .setTransmission(Transmission.SEMI_AUTOMATIC)
               .setTripComputer(new TripComputer())
               .setGPSNavigator(new GPSNavigator());
    }

    public void constructCityCar(CarBuilder builder) {
        builder.setType(Type.CITY_CAR)
               .setSeats(5)
               .setEngine(new Engine(1.2, 0))
               .setTransmission(Transmission.AUTOMATIC)
               .setTripComputer(new TripComputer())
               .setGPSNavigator(new GPSNavigator());
    }

    public void constructSUV(CarBuilder builder) {
        builder.setType(Type.SUV)
               .setSeats(7)
               .setEngine(new Engine(2.5, 0))
               .setTransmission(Transmission.MANUAL)
               .setTripComputer(new TripComputer())
               .setGPSNavigator(new GPSNavigator());
    }
}