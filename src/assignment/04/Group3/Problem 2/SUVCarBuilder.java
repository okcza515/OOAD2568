public class SUVCarBuilder extends CarBuilder {

    public SUVCarBuilder() {
        this.carType = Type.SUV;
        this.tripComputer = new TripComputer();
        this.gpsNavigator = new GPSNavigator();
    }
    
}

//65070501088 Sopida Keawjongkool