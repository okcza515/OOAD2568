public class CityCarBuilder extends CarBuilder {
    public CityCarBuilder() {
        this.carType = Type.CITY_CAR;
        this.tripComputer = new TripComputer();
    }
}

// 65070501039 Pongpon Butseemart