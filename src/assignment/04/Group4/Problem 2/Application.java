
public class Application {
	public static void main(String[] args) {
		Car sportCar = new SportCarBuilder()
				.setSeats(2)
				.setEngine(new Engine(2.0, 0))
				.setTransmission(Transmission.SEMI_AUTOMATIC)
				.setTripComputer(new TripComputer())
				.buildCar();

		Car cityCar = new CityCarBuilder()
				.setSeats(5)
				.setEngine(new Engine(1.2, 0))
				.setTransmission(Transmission.AUTOMATIC)
				.setGPSNavigator(new GPSNavigator())
				.setTripComputer(new TripComputer())
				.buildCar();

		Car suvCar = new SuvCarBuilder()
				.setSeats(7)
				.setEngine(new Engine(2.5, 0))
				.setTransmission(Transmission.MANUAL)
				.setGPSNavigator(new GPSNavigator())
				.setTripComputer(new TripComputer())
				.buildCar();

		Car brokenIeetan = new CityCarBuilder()
				.setSeats(3)
				.setEngine(new Engine(1.0, 0))
				.setTransmission(Transmission.AUTOMATIC)
				.buildCar();


		Manual sportCarManual = new Manual(sportCar);
		Manual cityCarManual = new Manual(cityCar);
		Manual suvCarManual = new Manual(suvCar);
		Manual ieetanManual = new Manual(brokenIeetan);

		System.out.println("------------------");
		System.out.println("Sport car Manual");
		System.out.println(sportCarManual.print());
		System.out.println("------------------");
		System.out.println("City car Manual");
		System.out.println(cityCarManual.print());
		System.out.println("------------------");
		System.out.println("Suv car Manual");
		System.out.println(suvCarManual.print());
		System.out.println("------------------");
		System.out.println("ieetan Manual");
		System.out.println(ieetanManual.print());
		System.out.println("------------------");
	}
}