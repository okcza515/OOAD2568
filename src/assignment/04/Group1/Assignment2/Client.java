public class Client {
    public static void main(String[] args) {
        SportCarBuilder sportCarBuilder = new SportCarBuilder();
        CityCarBuilder cityCarBuilder = new CityCarBuilder();
        SUVCarBuilder suvCarBuilder = new SUVCarBuilder();

        CarManufacturer carManufacturer = new CarManufacturer();

        carManufacturer.setCarBuilder(sportCarBuilder);
        carManufacturer.constructCar();    
        System.out.println("Sport Car: " + sportCarBuilder.getCar());

        carManufacturer.setCarBuilder(cityCarBuilder);
        carManufacturer.constructCar();  
        System.out.println("City Car: " + cityCarBuilder.getCar());

        carManufacturer.setCarBuilder(suvCarBuilder);
        carManufacturer.constructCar();  
        System.out.println("SUV Car: " + suvCarBuilder.getCar());
    }
    
}
