abstract class GeneralManufacturingProcess {
    private String deviceName;

    public GeneralManufacturingProcess(String deviceName) {
        this.deviceName = deviceName;
    }

    public abstract void assembly();
    public abstract void testing();
    public abstract void packaging();
    public abstract void storage();

    public void process() {
        System.out.println("Process of " + this.deviceName + " is starting ...");
        this.assembly();
        this.testing();
        this.packaging();
        this.storage();
        System.out.println("Process finished !");
    }
}