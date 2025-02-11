public class LaptopManufacturingProcess implements GeneralManufacturingProcess {
    @Override
    public void assembly() {
        System.out.println("Laptop assembling");
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
        System.out.println("Laptop storing");
    }

    public void launchProcess() {
        assembly();
        testing();
        packaging();
        storage();
    }
}
