import java.util.Base64;

public class EncryptionDecorator extends FileDataSource {
    private FileDataSource wrappee;

    public EncryptionDecorator(FileDataSource source) {
        super(source.name);
        this.wrappee = source;
    }

    @Override
    public void writeData(String data) {
        wrappee.writeData(encode(data));
    }

    @Override
    public String readData() {
        return decode(wrappee.readData());
    }

    private String encode(String data) {
        byte[] result = data.getBytes();
        for (int i = 0; i < result.length; i++) {
            result[i] += 1;
        }
        return Base64.getEncoder().encodeToString(result);
    }

    private String decode(String data) {
        byte[] result = Base64.getDecoder().decode(data);
        for (int i = 0; i < result.length; i++) {
            result[i] -= 1;
        }
        return new String(result);
    }
}
