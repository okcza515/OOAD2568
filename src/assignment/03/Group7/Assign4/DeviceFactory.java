public class DeviceFactory {
    public static void initializeManufacturing(GeneralManufacturingProcess manufacturingProcess) {
        manufacturingProcess.manufactureDevice();
    }

    public static void main(String[] args) {
        // Done
        GeneralManufacturingProcess smartphone = new SmartphoneManufacturingProcess();
        GeneralManufacturingProcess laptop = new LaptopManufacturingProcess();

        initializeManufacturing(smartphone);
        initializeManufacturing(laptop);
    }
}

// 65070501085