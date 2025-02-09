// Concrete implementation for laptop manufacturing
// Sawitt Ngamvilaisiriwong 65070503469
class LaptopManufacturingProcess extends GeneralManufacturingProcess {
    @Override
    protected void assembly() {
        System.out.println("Assembling laptop components: screen, keyboard, motherboard...");
    }
    
    @Override
    protected void testing() {
        System.out.println("Testing laptop: keyboard, display, hardware diagnostics...");
    }
    
    @Override
    protected void packaging() {
        System.out.println("Packaging laptop: protective foam, charger, manual...");
    }
    
    @Override
    protected void storage() {
        System.out.println("Storing laptop in secure warehouse section...");
    }
}

public class DeviceFactory {
    public static void main(String[] args) {
        // Manufacturing a smartphone
        //Bew
        System.out.println("Starting Smartphone Manufacturing Process...");
        GeneralManufacturingProcess smartphoneProcess = new SmartphoneManufacturingProcess();
        smartphoneProcess.startManufacturing();
    }
}