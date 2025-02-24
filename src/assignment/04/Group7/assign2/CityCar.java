public class Citycar extends CarBuilder {
    @Override
    public void createType() {
        getCar().setType(Type.CITY);
        System.out.println("Set Type.");
    }

    @Override
    public void createFuel() {
        getCar().setFuel(0);
        System.out.println("Set Fuel.");
    }

    @Override
    public void createSeats() {
        getCar().setSeats(5);
        System.out.println("Set Seat.");
    }

    @Override
    public void createEngine() {
        getCar().setEngine(
                new Engine(1.2, 0));
        System.out.println("Set Engine.");
    }

    @Override
    public void createTransmission() {
        getCar().setTransmission(Transmission.AUTOMATIC);
        System.out.println("Set Transmission.");
    }

    @Override
    public void createTripComputer() {
        getCar().setTripComputer(new TripComputer());
        System.out.println("Set TripComputer.");
    }

    @Override
    public void createGpsNavigator() {
        getCar().setGpsNavigator(new GPSNavigator());
        System.out.println("Set GPSNavigator.");
    }
}

//Pakaporn Kanteng 65070501042