public class DeviceFactory {
    public static void main(String[] args) {
        SmartphoneManufacturingProcess smartphone = new SmartphoneManufacturingProcess();
        smartphone.process();
        LaptopManufacturingProcess laptop = new LaptopManufacturingProcess();
        laptop.process();
    }
}