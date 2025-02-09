public class DeviceFactory {
    public static void main(String[] args) {
        SmartphoneManufacturingProcess smartphone = new SmartphoneManufacturingProcess();
        LaptopManufacturingProcess laptop = new LaptopManufacturingProcess();
        
        System.out.println("Starting smartphone manufacturing process...");
        smartphone.assembly();
        smartphone.testing();
        smartphone.packaging();
        smartphone.storage();
        
        System.out.println("Starting laptop manufacturing process...");
        laptop.assembly();
        laptop.testing();
        laptop.packaging();
        laptop.storage();
    }
}
//Ratchanon Tarawan 65070503464