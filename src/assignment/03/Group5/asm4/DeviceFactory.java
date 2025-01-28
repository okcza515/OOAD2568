public class DeviceFactory {
    public static void main(String[] args) {
        GenralManu laptop = new LaptopManufacturingProcess();
        laptop.assembly();
        laptop.testing();
        laptop.packaging();
        laptop.storage();
    }
}
