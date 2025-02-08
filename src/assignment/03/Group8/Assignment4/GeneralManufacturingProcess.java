abstract public class GeneralManufacturingProcess{
    private String device;

    //encapsulation
    public GeneralManufacturingProcess(String device) {
        this.device = device;
    }
    
    public void startProcess(){
        assemble();
        test();
        packageDevice();
        store();
    }

    public void assemble(){
        System.out.println("Assembling the"+ device);
    }
    public void test(){
        System.out.println("Testing the" + device);
    }
    public void packageDevice(){
        System.out.println("Packaging the"+ device);
    }
    public void store(){
        System.out.println("Storing the"+ device);
    }
}

