
public class Application {
	public static void main(String[] args) {
		// Car sportCar = new Car(Type.SPORTS_CAR, 2, new Engine(3.0,0), Transmission.SEMI_AUTOMATIC, new TripComputer(), new GPSNavigator());
		// System.out.println("Car built:\n" + sportCar.getType());
		
		// Manual carManual = new Manual(Type.SPORTS_CAR, 2, new Engine(3.0,0), Transmission.SEMI_AUTOMATIC, new TripComputer(), new GPSNavigator());
		// System.out.println("\nCar manual built:\n" + carManual.print());

		Car sports = Manufacturer.createSportsCar();
		Manual sportsCarManual = Manufacturer.buildManual(sports);
		System.out.println("\nCar built:\n" + sports.getType()+"\n");
		System.out.println(sportsCarManual.print());
		System.out.println("\n----------------------\n");

		Car city = Manufacturer.createCityCar();
		Manual cityCarManual = Manufacturer.buildManual(city);
		System.out.println("Car built:\n" + city.getType()+"\n");
		System.out.println(cityCarManual.print());
		System.out.println("\n----------------------\n");

		Car suv = Manufacturer.createSUVCar();
		Manual suvCarManual = Manufacturer.buildManual(suv);
		System.out.println("Car built:\n" + suv.getType()+"\n");
		System.out.println(suvCarManual.print());
		System.out.println("\n----------------------\n");
		
		// 65070501048
	}
}
