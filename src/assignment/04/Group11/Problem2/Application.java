public class Application {
	public static void main(String[] args) {
        //SportCar
        SportCarBuilder sportCarBuilder = new SportCarBuilder();
        sportCarBuilder.setSeats(2);
        sportCarBuilder.setEngine(new Engine(2.0, 0));
        sportCarBuilder.setTransmission(Transmission.SEMI_AUTOMATIC);
		sportCarBuilder.setTripComputer(new TripComputer());
		sportCarBuilder.setGPSNavigator(new GPSNavigator());
        
        Car sportCar = sportCarBuilder.getCar();

        //SUVCar
        SUVCarBuilder suvCarBuilder = new SUVCarBuilder();
        suvCarBuilder.setSeats(7);
        suvCarBuilder.setEngine(new Engine(2.5, 0));
        suvCarBuilder.setTransmission(Transmission.MANUAL);
        suvCarBuilder.setTripComputer(new TripComputer());
        suvCarBuilder.setGPSNavigator(new GPSNavigator());

        Car suvCar = suvCarBuilder.getCar();

        //CityCar
        CityCarBuilder cityCarBuilder = new CityCarBuilder();
        cityCarBuilder.setSeats(5);
        cityCarBuilder.setEngine(new Engine(1.2, 0));
        cityCarBuilder.setTransmission(Transmission.AUTOMATIC);
        cityCarBuilder.setTripComputer(new TripComputer());
        cityCarBuilder.setGPSNavigator(new GPSNavigator());

        Car cityCar = cityCarBuilder.getCar();

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
//65070503408 Jaruikit Jintanasathirakul
