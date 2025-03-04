public class EncryptionDecorator implements DataSource {
    private DataSource wrappee;
    private Encryption encryption = new Encryption();

    public EncryptionDecorator(DataSource source) {
        this.wrappee = source;
    }

    @Override
    public void writeData(String data) {
        wrappee.writeData(encryption.encode(data));
    }

    @Override
    public String readData() {
        return encryption.decode(wrappee.readData());
    }
}