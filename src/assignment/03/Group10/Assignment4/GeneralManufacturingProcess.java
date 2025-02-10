public abstract class GeneralManufacturingProcess {
    // 65070501049 Roodfan Maimahad
    public abstract void assembly();
    public abstract void testing();
    public abstract void packaging();
    public abstract void storage();

    // 65070501076 Danai Saengbuamad
    public final void createDevice(){
        System.out.println("Start Building Device Process");
        assembly();
        testing();
        packaging();
        storage();
        System.out.println("Device Builded");
    }
}