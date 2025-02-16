
public class ManualBuilder {
    Manual manual;

    public ManualBuilder() {
        this.manual = new Manual();
    }

    public Manual buildManual(Car car) {
        manual.setType(car.getType());
        manual.setFuel(car.getFuel());
        manual.setSeats(car.getSeats());
        manual.setEngine(car.getEngine());
        manual.setTransmission(car.getTransmission());
        manual.setTripComputer(car.getTripComputer());
        manual.setGpsNavigator(car.getGpsNavigator());
        return manual;
    }
}



