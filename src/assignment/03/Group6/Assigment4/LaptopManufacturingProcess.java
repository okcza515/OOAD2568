public class LaptopManufacturingProcess extends GeneralManufacturingProcess{
    @Override
    public void assembly() {
        System.out.println("Laptop assembly");
    }

    @Override
    public void testing() {
        System.out.println("Laptop testing");
    }
    
    @Override
    public void packaging() {
        System.out.println("Laptop packaging");
    }

    @Override
    public void storage() {
        System.out.println("Laptop storage");
    }
}
