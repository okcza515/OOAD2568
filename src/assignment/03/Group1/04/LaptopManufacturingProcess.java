public class LaptopManufacturingProcess implements GeneralManufacturingProcess {
    @Override
    public void assembly() {
        System.out.println("Laptop assembling");
    }

    @Override
    public void testing() {
        System.out.println("Laptop testing");
    }
}