
public class ManualBuilder {

    public Manual buildManual(Car car) {
        return new Manual(
                car.getType(),
                car.getSeats(),
                car.getEngine(),
                car.getTransmission(),
                car.getTripComputer(),
                car.getGpsNavigator()
        );
    }

//    public Manual(Type type,
//    int seats, Engine engine, Transmission transmission,
//                  TripComputer tripComputer, GPSNavigator gpsNavigator) {
}



