//65070503408 Jarukit Jintanasathirakul
public abstract class GeneralManufacturingProcess {
    protected abstract void assembly();
    protected abstract void testing();
    protected abstract void packaging();
    protected abstract void storage();

    public final void manufacture() {
        assembly();
        testing();
        packaging();
        storage();
    }
}
