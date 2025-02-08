public class DeviceFactory {
    public static void main(String[] args){

        ManufacturingProcess smartphoneFactory = new SmartphoneManufacturingProcess();
        ManufacturingProcess laptopFactory = new LaptopManufacturingProcess();

        smartphoneFactory.manufacturingProcess();
        System.out.println();
        laptopFactory.manufacturingProcess();
    }
}