public class Application {
	public static void main(String[] args) {
        SportCarBuilder sportCarBuilder = new SportCarBuilder();
        sportCarBuilder.setSeats(2);
        sportCarBuilder.setEngine(new Engine(2.0, 0));
        sportCarBuilder.setTransmission(Transmission.SEMI_AUTOMATIC);
		sportCarBuilder.setTripComputer(new TripComputer());
		sportCarBuilder.setGPSNavigator(new GPSNavigator());
        
        Car sportCar = sportCarBuilder.getCar();
        System.out.println("Car Details:");
        System.out.println("Type: " + sportCar.getType());
        System.out.println("Number of seats: " + sportCar.getSeats());
        System.out.println("Engine: " + sportCar.getEngine().getVolume() + "L (Fuel level: " + sportCar.getEngine().getMileage() + ")");
        System.out.println("Transmission: " + sportCar.getTransmission());
        System.out.println("Trip Computer: " + (sportCar.getTripComputer() != null ? "Installed" : "Not installed"));
        System.out.println("GPS Navigator: " + (sportCar.getGpsNavigator() != null ? "Installed" : "Not installed"));
        System.out.println("Fuel Level: " + sportCar.getFuel() + "L");

        System.out.println("\n--------------------------------\n");
		
    }
}
