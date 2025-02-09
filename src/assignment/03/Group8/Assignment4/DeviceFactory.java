public class DeviceFactory {
    public static void main(String[] args) {
        System.out.println("Starting the device manufacturing process:\n");

        System.out.println("Manufacturing Smartphone:");
        GeneralManufacturingProcess smartphoneProcess = new SmartphoneManufacturingProcess();
        smartphoneProcess.startProcess();startProcess();

        System.out.println("\nManufacturing Laptop:");
        GeneralManufacturingProcess laptopProcess = new LaptopManufacturingProcess();
        laptopProcess.startProcess();startProcess();
    }
}