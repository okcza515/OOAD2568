public class SportCarBuilder extends ConcreteCarBuilder {
    public SportCarBuilder() {
        this.carType = Type.SPORTS_CAR;
        this.tripComputer = new TripComputer();
        this.gpsNavigator = new GPSNavigator();
    }
}
// 65070503412 Chitsanupong