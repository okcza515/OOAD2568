
public abstract class GeneralManufacturingProcess {

    protected abstract void assembly();

    protected abstract void testing();

    protected abstract void packaging();

    protected abstract void storage();

    public final void manufacture() {
        System.out.println("Start building...");
        assembly();
        testing();
        packaging();
        storage();
        System.out.println("Done!");
    }
}
// 65070501001 Kantapong Vongpanich
