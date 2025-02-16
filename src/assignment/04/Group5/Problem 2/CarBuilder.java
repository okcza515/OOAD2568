public abstract class CarBuilder {
    Car car;

    public void setCar(Car car) {
        this.car = car;
    }

    public Car getCar() {
        return car;
    }

    public abstract void createType();
    public abstract void createSeats();
    public abstract void createEngine();
    public abstract void createTransmission();

    public final Car buildCar(){
        Car c = new Car();
        setCar(c);

        createType();;
        createSeats();
        createEngine();
        createTransmission();

        return c;
    }

}
