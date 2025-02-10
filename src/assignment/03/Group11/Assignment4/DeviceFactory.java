public class DeviceFactory {
    public static void main(String[] args) {
        GeneralManufacturingProcess LaptopProduction = new LaptopManufacturingProcess();
        GeneralManufacturingProcess SmartphoneProduction = new SmartphoneManufacturingProcess();
        LaptopProduction.manufacture();
        SmartphoneProduction.manufacture();
    }
}
