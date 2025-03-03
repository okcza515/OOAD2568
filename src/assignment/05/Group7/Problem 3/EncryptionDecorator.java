import java.util.Base64;

public class EncryptionDecorator extends DataSourceDecorator {
    public EncryptionDecorator(DataSource source) {
        super(source);
    }

    @Override
    public void writeData(String data) {
        Encryption encryption = new Encryption();
        super.writeData(encryption.encode(data));
    }

    @Override
    public String readData() {
        Encryption encryption = new Encryption();
        return encryption.decode(super.readData());
    }

}

//65070501016 Chitsanucha
