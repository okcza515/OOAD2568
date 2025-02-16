public class ConcreteCityCar extends CarBuilder{
    @Override
    public void createType() {
        car.setType(Type.CITY_CAR);
    }

    @Override
    public void createSeats() {
        car.setSeats(5);
    }

    @Override
    public void createEngine() {
        car.setEngine(new Engine(1.2, 0));
    }

    @Override
    public void createTransmission() {
        car.setTransmission(Transmission.AUTOMATIC);
    }
}

