//65070503408 Jarukit Jintanasathirakul
public class LaptopManufacturingProcess extends GeneralManufacturingProcess {
    @Override
    protected void assembly() {
        System.out.println("Assembling laptop");
    }
    protected void testing() {
        System.out.println("Testing laptop");
    }
    protected void packaging() {
        System.out.println("Packaging laptop");
    }
    protected void storage() {
        System.out.println("Storing laptop");
    }
}
