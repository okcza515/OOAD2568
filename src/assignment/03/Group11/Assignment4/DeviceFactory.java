public class DeviceFactory {
    public static void main(String[] args) {
        GeneralManufacturingProcess LaptopProduction = new LaptopManufacturingProcess();

        LaptopProduction.manufacture();
    }
}
