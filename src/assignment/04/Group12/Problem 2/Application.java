public class Application {
    public static void main(String[] args) {
        Director director = new Director();
        CarBuilder carBuilder = new CarBuilder();
        director.constructSportCar(carBuilder);

        // Construct sport car
        Car sportCar = carBuilder.getResult();
        System.out.println("Car built:\n" + sportCar.getType());

        // Construct manual for sport car
        ManualBuilder manualBuilder = new ManualBuilder();
        director.constructSportCar(manualBuilder);
        Car manualCar = manualBuilder.getResult();
        System.out.println("Manual built:\n" + manualCar.getType());

        // Construct SUV car
        carBuilder = new CarBuilder();
        director.constructSUVCar(carBuilder);
        Car suvCar = carBuilder.getResult();
        System.out.println("Car built:\n" + suvCar.getType());

        // Construct manual for SUV car
        manualBuilder = new ManualBuilder();
        director.constructSUVCar(manualBuilder);
        Car manualSUVCar = manualBuilder.getResult();
        System.out.println("Manual built:\n" + manualSUVCar.getType());

        // Construct city car
        carBuilder = new CarBuilder();
        director.constructCityCar(carBuilder);
        Car cityCar = carBuilder.getResult();
        System.out.println("Car built:\n" + cityCar.getType());

        // Construct manual for city car
        manualBuilder = new ManualBuilder();
        director.constructCityCar(manualBuilder);
        Car manualCityCar = manualBuilder.getResult();
        System.out.println("Manual built:\n" + manualCityCar.getType());

    }
}
