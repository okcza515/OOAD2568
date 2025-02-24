public class SportCarBuilder implements ICarBuilder {
    private Car car;

    public void reset() {
        car = new Car();
    }

    @Override
    public void buildSeats() {
        car.setSeats(2);
    }

    @Override
    public void buildEngine() {
        car.setEngine(2.0f, 0);
    }

    @Override
    public void buildTransmission() {
        car.setTransmission("semi automatic");
    }

    public Car getCar() {
        return car;
    }

}
