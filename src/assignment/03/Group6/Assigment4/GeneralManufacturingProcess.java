public abstract class GeneralManufacturingProcess {
    public abstract void assembly();
    public abstract void testing();
    public abstract void packaging();
    public abstract void storage();

    public void manufacture() {
        assembly();
        testing();
        packaging();
        storage();
    }
}
