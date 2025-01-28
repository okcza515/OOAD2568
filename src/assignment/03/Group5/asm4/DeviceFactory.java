public class DeviceFactory {
    public static void main(String[] args) {
        GenralManu smartphone = new SmartphoneManufacturingProcess();
        // smartphone.assembly();
        // smartphone.testing();
        // smartphone.packaging();
        // smartphone.storage();
        smartphone.LaunchProcess();

        System.out.println("\n----------------------\n");

        GenralManu laptop = new LaptopManufacturingProcess();
        // laptop.assembly();
        // laptop.testing();
        // laptop.packaging();
        // laptop.storage();
        laptop.LaunchProcess();
    }
}
