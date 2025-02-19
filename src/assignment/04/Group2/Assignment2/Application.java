
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

		Manual sportCarManual = new Manual(sportCar.getCar().getType(), sportCar.getCar().getSeats(),
				sportCar.getCar().getEngine(), sportCar.getCar().getTransmission(), new TripComputer(),
				new GPSNavigator());
		System.out.println("\nSport Car Manual built:\n" + sportCarManual);

		Manual cityCarManual = new Manual(cityCar.getCar().getType(), cityCar.getCar().getSeats(),
				cityCar.getCar().getEngine(), cityCar.getCar().getTransmission(), new TripComputer(),
				new GPSNavigator());
		System.out.println("\nCity Car Manual built:\n" + cityCarManual);

		Manual suvCarManual = new Manual(suvCar.getCar().getType(), suvCar.getCar().getSeats(),
				suvCar.getCar().getEngine(), suvCar.getCar().getTransmission(), new TripComputer(), new GPSNavigator());
		System.out.println("\nSUV Car Manual built:\n" + suvCarManual);
	}
}
// Ratchanon Tarawan 65070503464