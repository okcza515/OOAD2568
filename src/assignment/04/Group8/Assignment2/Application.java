
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
        CarConcreteBuilder carConcreteBuilder = new CarConcreteBuilder();
        manufacturer.constructSportsCar(carConcreteBuilder);
        Car sportsCar = carConcreteBuilder.getResult();
        System.out.println("Car built:\n" + sportsCar.print());

		// Build City Car
		CarConcreteBuilder CitycarConcreteBuilder = new CarConcreteBuilder();
        manufacturer.constructCityCar(CitycarConcreteBuilder);
        Car CityCar = CitycarConcreteBuilder.getResult();
        System.out.println("Car built:\n" + CityCar.print());

		// Build SUV Car
		ManualConcreteBuilder manualConcreteBuilder = new ManualConcreteBuilder();
		manufacturer.constructSUVCar(manualConcreteBuilder);
		Manual SUVCar = manualConcreteBuilder.getResult();
		System.out.println("Car built:\n" + SUVCar.print());
	}
}
