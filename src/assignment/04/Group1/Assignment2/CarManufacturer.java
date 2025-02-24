public class CarManufacturer {
    private ICarBuilder carBuilder;

    public void setCarBuilder(ICarBuilder carBuilder) {
        this.carBuilder = carBuilder;
    }

    public void constructCar() {
        carBuilder.reset();
        carBuilder.buildSeats();
        carBuilder.buildEngine();
        carBuilder.buildTransmission();
    }
}
