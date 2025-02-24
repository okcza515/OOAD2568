public class CityCarBuilder implements ICarBuilder {
    private Car car;

    public void reset() {
        car = new Car();
    }

    @Override
    public void buildSeats() {
        car.setSeats(5);
    }

    @Override
    public void buildEngine() {
        car.setEngine(1.2f, 0);
    }

    @Override
    public void buildTransmission() {
        car.setTransmission("automatic");
    }

    public Car getCar() {
        return car;
    }
}
