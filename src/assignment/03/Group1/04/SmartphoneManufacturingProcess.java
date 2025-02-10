public class SmartphoneManufacturingProcess implements GeneralManufacturingProcess {

    @Override
    public void assembly() {
        System.out.println("Smartphone assembling");
    }

    @Override
    public void testing() {
        System.out.println("Smartphone testing");
    }

    @Override
    public void packaging() {
        System.out.println("Smartphone packaging");
    }

    @Override
    public void storage() {
        System.out.println("Smartphone storing");
    }
    
    public void launchProcess() {
        assembly();
        testing();
        packaging();
        storage();
    }
}
