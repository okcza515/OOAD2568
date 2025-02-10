//65070501078 Nawaphon Promnan
public class DeviceFactory{

    public static void main(String[] args) {
        GeneralManufacturingProcess smartphoneBuilding = new SmartphoneManufacturingProcess();
        GeneralManufacturingProcess laptopBuilding = new LaptopManufacturingProcess();

        smartphoneBuilding.createDevice();
        laptopBuilding.createDevice();
    }
}