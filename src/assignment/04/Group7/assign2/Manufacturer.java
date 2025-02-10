public class Manufacturer {
    public static Car Sportscar(){
        CarBuilder Sportsbuilder = new Sportscar();
        return Sportsbuilder.buildCar();
    }

    public static Car Citycar(){
        CarBuilder Citybuilder = new Citycar();
        return Citybuilder.buildCar();
    }

    public static Car SUVcar(){
        CarBuilder SUVbuilder = new SUVcar();
        return SUVbuilder.buildCar();
    }
}
//Chitsanucha 65070501016