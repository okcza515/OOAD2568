public class SmartphoneManufacturingProcess extends GeneralManufacturingProcess{
    @Override
    public void assembly() {
        System.out.println("Smartphone assembly");
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
        System.out.println("Smartphone storage");
    }
}
