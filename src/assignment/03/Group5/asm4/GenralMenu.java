abstract class GenralManu{
    abstract void assembly();
    abstract void testing();
    abstract void packaging();
    abstract void storage();

    public void LaunchProcess() {
        assembly();
        testing();
        packaging();
        storage();
    }
}