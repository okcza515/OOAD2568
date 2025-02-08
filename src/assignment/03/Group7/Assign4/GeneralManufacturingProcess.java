public abstract class GeneralManufacturingProcess {

    protected abstract void assembly();

    protected abstract void testing();

    protected abstract void packaging();

    protected abstract void storage();

    public final void manufactureDevice() {
        assembly();
        testing();
        packaging();
        storage();
    }
}

// 65070501085