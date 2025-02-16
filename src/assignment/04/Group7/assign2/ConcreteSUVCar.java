public class ConcreteSUVCar extends CarBuilder {
    @Override
    public void createType() {
        car.setType(Type.SUV);
    }

    @Override
    public void createSeats() {
        car.setSeats(7);
    }

    @Override
    public void createEngine() {
        car.setEngine(new Engine(2.5, 0));
    }

    @Override
    public void createTransmission() {
        car.setTransmission(Transmission.MANUAL);
    }
}