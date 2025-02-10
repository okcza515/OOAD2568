// 65070503412 Chitsanupong
public class SmartphoneManufacturingProcess extends GeneralManufacturingProcess {
    @Override
    protected void assembly() {
        System.out.println("Assembling smartphone");
    }
    protected void testing() {
        System.out.println("Testing smartphone");
    }
    protected void packaging() {
        System.out.println("Packaging smartphone");
    }
    protected void storage() {
        System.out.println("Storing smartphone");
    }
}
