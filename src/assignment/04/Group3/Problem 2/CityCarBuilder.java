public class CityCarBuilder extends CarBuilder {
    public CityCarBuilder() {
        this.carType = Type.CITY_CAR;
        this.tripComputer = new TripComputer();
        this.gpsNavigator = new GPSNavigator();
    }
}

// 65070501039 Pongpon Butseemart