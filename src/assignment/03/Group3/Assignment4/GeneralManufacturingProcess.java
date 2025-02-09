
public abstract class GeneralManufacturingProcess {

    abstract void assembly();

    abstract void testing();

    abstract void packaging();

    abstract void storage();

    public void manufacture() {
        System.out.println("Start building...");
        assembly();
        testing();
        packaging();
        storage();
        System.out.println("Done!");
    }
}
// 65070501001 Kantapong Vongpanich
