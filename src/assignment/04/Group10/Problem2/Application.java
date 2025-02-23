public class Application {
	public static void main(String[] args) {
		CarBuilder sportCarBuilder = new SportCarBuilder();
		CarBuilder cityCarBuilder = new CityCarBuilder();
		CarBuilder suvCarBuilder = new SUVCarBuilder();

		Car sportCar = sportCarBuilder
						.setSeats(2)
						.setEngine(new Engine(2.0, 0))
						.setTransmission(Transmission.SEMI_AUTOMATIC)
						.buildCar();
		Car cityCar = cityCarBuilder
						.setSeats(5)
						.setEngine(new Engine(1.2, 0))
						.setTransmission(Transmission.AUTOMATIC)
						.buildCar();
		Car suvCar = suvCarBuilder
						.setSeats(7)
						.setEngine(new Engine(2.5, 0))
						.setTransmission(Transmission.MANUAL)
						.buildCar();

		Manual sportCarManual = new Manual(sportCar);
		Manual cityCarManual = new Manual(cityCar);
		Manual suvCarManual = new Manual(suvCar);

		System.out.println("Sport Car: \n");
		System.out.println(sportCarManual.print());
		System.out.println("City Car: \n");
		System.out.println(cityCarManual.print());
		System.out.println("SUV Car: \n");
		System.out.println(suvCarManual.print());
	}
}
