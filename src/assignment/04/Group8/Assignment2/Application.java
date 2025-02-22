
public class Application {
	// public static void main(String[] args) {
	// 	Car sportCar = new Car(Type.SPORTS_CAR, 2, new Engine(3.0,0), Transmission.SEMI_AUTOMATIC, new TripComputer(), new GPSNavigator());
	// 	System.out.println("Car built:\n" + sportCar.getType());
		
	// 	Manual carManual = new Manual(Type.SPORTS_CAR, 2, new Engine(3.0,0), Transmission.SEMI_AUTOMATIC, new TripComputer(), new GPSNavigator());
	// 	System.out.println("\nCar manual built:\n" + carManual.print());
	// }
	public static void main(String[] args) {
		Manufacturer manufacturer = new Manufacturer();

        // Build Sports Car
        AutoConcreteBuilder autoConcreteBuilder = new AutoConcreteBuilder();
        manufacturer.constructSportsCar(autoConcreteBuilder);
        Car sportsCar = autoConcreteBuilder.getResult();
        System.out.println("Car built:\n" + sportsCar.getType());

	}
}
