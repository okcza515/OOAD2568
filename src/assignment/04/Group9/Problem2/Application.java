// Updated Application.java
public class Application {
    public static void main(String[] args) {
        CarDirector director = new CarDirector();
        
        // Building a sports car
        DefaultCarBuilder carBuilder = new DefaultCarBuilder();
        director.constructSportsCar(carBuilder);
        Car sportsCar = carBuilder.build();
        
        // Building a sports car manual
        CarManualBuilder manualBuilder = new CarManualBuilder();
        director.constructSportsCar(manualBuilder);
        Manual sportsCarManual = manualBuilder.build();
        
        System.out.println("Car built:\n" + sportsCar.getType());
        System.out.println("\nCar manual built:\n" + sportsCarManual.print());
        
        // Similarly for city car and SUV...
    }
}