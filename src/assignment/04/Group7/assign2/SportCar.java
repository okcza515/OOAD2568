public class SportCar extends CarBuilder {
    @Override
    public void createType() {
        getCar().setType(Type.SPORT);
        System.out.println("Set Type.");
    }

    @Override
    public void createFuel() {
        getCar().setFuel(0);
        System.out.println("Set Fuel.");
    }

    @Override
    public void createSeats() {
        getCar().setSeats(2);
        System.out.println("Set Seat.");
    }

    @Override
    public void createEngine() {
        getCar().setEngine(
                new Engine(2.0, 0));
        System.out.println("Set Engine.");
    }

    @Override
    public void createTransmission() {
        getCar().setTransmission(Transmission.SEMI_AUTOMATIC);
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