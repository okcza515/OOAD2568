public class DeviceFactory {
    public static void main(String[] args) {
        GeneralManufacturingProcess smartphone = new SmartphoneManufacturingProcess();
        GeneralManufacturingProcess laptop = new LaptopManufacturingProcess();

        smartphone.manufacture();
        laptop.manufacture();
    }
}
