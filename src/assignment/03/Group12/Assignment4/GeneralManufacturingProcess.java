abstract class GeneralManufacturingProcess {
    private String deviceName;

    public GeneralManufacturingProcess(String deviceName) {
        this.deviceName = deviceName;
    }

    protected abstract void assembly();
    protected abstract void testing();
    protected abstract void packaging();
    protected abstract void storage();

    public void process() {
        System.out.println("Process of " + this.deviceName + " is starting ...");
        this.assembly();
        this.testing();
        this.packaging();
        this.storage();
        System.out.println("Process finished !");
    }
}