public class ConcreteSportsCar extends CarBuilder{
    @Override
    public void createType() {
        getCar().setType(Type.SPORTS_CAR);
    }

    @Override
    public void createSeats() {
        getCar().setSeats(2);
    }

    @Override
    public void createEngine() {
        getCar().setEngine(new Engine(2.0, 0));
    }

    @Override
    public void createTransmission() {
        getCar().setTransmission(Transmission.SEMI_AUTOMATIC);
    }
}

