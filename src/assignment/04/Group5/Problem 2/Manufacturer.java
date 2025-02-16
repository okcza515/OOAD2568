public class Manufacturer {
    public static Car createSportsCar() {
        CarBuilder carBuilder = new ConcreteSportsCar();
        return carBuilder.buildCar();
    }

    public static Car createCityCar() {
        CarBuilder carBuilder = new ConcreteCityCar();
        return carBuilder.buildCar();
    }

    public static Car createSUVCar() {
        CarBuilder carBuilder = new ConcreteSUVCar();
        return carBuilder.buildCar();
    }

    public static Manual buildManual(Car car) {
        return new ManualBuilder().buildManual(car);
    }

    // 65070501048
}