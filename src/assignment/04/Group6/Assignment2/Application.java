public class Application {
    public static void main(String[] args) {
        Director director = new Director();

        // Sports Car
        CarBuilder carBuilder = new CarBuilder();
        director.constructSportsCar(carBuilder);
        Car sportsCar = carBuilder.build();
        System.out.println("Sports car built:\n" + sportsCar.getType());

        // Sports Car Manual
        ManualBuilder manualBuilder = new ManualBuilder();
        director.constructSportsCar(manualBuilder);
        Manual sportsCarManual = manualBuilder.build();
        System.out.println("\nSports car manual built:\n" + sportsCarManual.print());

        // City Car
        carBuilder = new CarBuilder();
        director.constructCityCar(carBuilder);
        Car cityCar = carBuilder.build();
        System.out.println("\nCity car built:\n" + cityCar.getType());

        // City Car Manual
        manualBuilder = new ManualBuilder();
        director.constructCityCar(manualBuilder);
        Manual cityCarManual = manualBuilder.build();
        System.out.println("\nCity car manual built:\n" + cityCarManual.print());

        // SUV Car
        carBuilder = new CarBuilder();
        director.constructSUVCar(carBuilder);
        Car suvCar = carBuilder.build();
        System.out.println("\nSUV car built:\n" + suvCar.getType());

        // SUV Car Manual
        manualBuilder = new ManualBuilder();
        director.constructSUVCar(manualBuilder);
        Manual suvCarManual = manualBuilder.build();
        System.out.println("\nSUV car manual built:\n" + suvCarManual.print());
    }
}
