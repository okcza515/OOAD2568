
public class Application {
	public static void main(String[] args) {
		SportCarMaker sportCar = new SportCarMaker();
		sportCar.setSeats(6);
		sportCar.setCarType(Type.SPORTS_CAR);
		sportCar.setEngine(new Engine(2.0, 0));
		sportCar.setTransmission(Transmission.SEMI_AUTOMATIC);
		System.out.println("Sport Car built:\n" + sportCar.getCar().getType());

		CityCarMaker cityCar = new CityCarMaker();
		cityCar.setSeats(5);
		cityCar.setCarType(Type.SPORTS_CAR);
		cityCar.setEngine(new Engine(1.2, 10));
		cityCar.setTransmission(Transmission.AUTOMATIC);
		System.out.println("City Car built:\n" + cityCar.getCar().getType());

		SUVCarMaker suvCar = new SUVCarMaker();
		suvCar.setSeats(5);
		suvCar.setCarType(Type.SPORTS_CAR);
		suvCar.setEngine(new Engine(1.2, 10));
		suvCar.setTransmission(Transmission.MANUAL);
		System.out.println("SUV Car built:\n" + suvCar.getCar().getType());

		// Manual sportCarManual = new Manual(Type.SPORTS_CAR, 2, new Engine(2.0, 0),
		// Transmission.SEMI_AUTOMATIC,
		// new TripComputer(), new GPSNavigator());

		// System.out.println("\nCar manual built:\n" + carManual.print());
	}
}
