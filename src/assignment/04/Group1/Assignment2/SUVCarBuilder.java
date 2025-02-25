public class SUVCarBuilder implements ICarBuilder {
    private Car car;

    public void reset() {
        car = new Car();
    }

    @Override
    public void buildSeats() {
        car.setSeats(7);
    }

    @Override
    public void buildEngine() {
        car.setEngine(2.5f, 0);
    }

    @Override
    public void buildTransmission() {
        car.setTransmission("manual");
    }

    public Car getCar() {
        return car;
    }
}
