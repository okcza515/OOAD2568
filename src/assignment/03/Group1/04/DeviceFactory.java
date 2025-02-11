public class DeviceFactory {
    public static void main(String[] args) {
        SmartphoneManufacturingProcess smartphone = new SmartphoneManufacturingProcess();
        LaptopManufacturingProcess laptop = new LaptopManufacturingProcess();

        smartphone.launchProcess();
        laptop.launchProcess();
    }
}
