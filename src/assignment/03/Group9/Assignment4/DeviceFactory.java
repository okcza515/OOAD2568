// Abstract class defining the template for manufacturing process
//Chayaphon Chaisangkha 65070503409
abstract class GeneralManufacturingProcess {
    // Template method - final so it cannot be overridden
    public final void startManufacturing() {
        assembly();
        testing();
        packaging();
        storage();
    }
    
    // Abstract methods to be implemented by specific processes
    protected abstract void assembly();
    protected abstract void testing();
    protected abstract void packaging();
    protected abstract void storage();
}

// Concrete implementation for smartphone manufacturing
//Top
class SmartphoneManufacturingProcess extends GeneralManufacturingProcess {
    @Override
    protected void assembly() {
        System.out.println("Assembling smartphone components: display, motherboard, battery...");
    }
    
    @Override
    protected void testing() {
        System.out.println("Testing smartphone: touch response, camera, battery life...");
    }
    
    @Override
    protected void packaging() {
        System.out.println("Packaging smartphone: box, charger, manual...");
    }
    
    @Override
    protected void storage() {
        System.out.println("Storing smartphone in temperature-controlled warehouse...");
    }
}

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

        //Chanawat Limpanatewin 65070503445

        System.out.println("\n------------------------\n");
        
        // Manufacturing a laptop
        System.out.println("Starting Laptop Manufacturing Process...");
        GeneralManufacturingProcess laptopProcess = new LaptopManufacturingProcess();
        laptopProcess.startManufacturing();
    }
}