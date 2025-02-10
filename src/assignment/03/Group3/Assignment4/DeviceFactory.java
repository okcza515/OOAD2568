
public class DeviceFactory {

    public static void main(String[] args) {
        GeneralManufacturingProcess smartphoneFactory = new SmartphoneManufacturingProcess();
        GeneralManufacturingProcess laptopFactory = new LaptopManufacturingProcess();

        smartphoneFactory.manufacture();
        System.out.println();
        laptopFactory.manufacture();
    }
}

// 65070501023 Thanaphol Thangthaweesuk
// 65070501088 Sopida Keawjongkool
