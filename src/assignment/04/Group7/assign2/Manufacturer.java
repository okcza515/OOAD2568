public class Manufacturer {
    public static Car CreateSUVcar(){
        CarBuilder SUVbuilder = new SUVcar();
        return SUVbuilder.buildCar();
    }
}
//Chitsanucha 65070501016