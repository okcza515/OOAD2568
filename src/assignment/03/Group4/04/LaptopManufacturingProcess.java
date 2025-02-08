public class LaptopManufacturingProcess extends GeneralManufacturingProcess{
    @Override
    protected void assembly() {
        System.out.println("Assembling Labtop...");
    }

    @Override
    protected void testing() {
        System.out.println("Testing Labtop...");
    }

    @Override
    protected void packaging() {
        System.out.println("Packaging Labtop...");
    }

    @Override
    protected void storage() {
        System.out.println("Labtop Storage");
    }
}
