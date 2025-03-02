import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.util.Base64;
import java.util.zip.DeflaterOutputStream;
import java.util.zip.InflaterInputStream;

public class CompressionDecorator extends FileDataSource {
    private FileDataSource wrappee;

    public CompressionDecorator(FileDataSource source) {
        super(source.name);
        this.wrappee = source;
    }

    @Override
    public void writeData(String data) {
        wrappee.writeData(compress(data));
    }

    @Override
    public String readData() {
        return decompress(wrappee.readData());
    }

    private String compress(String data) {
        try (ByteArrayOutputStream bout = new ByteArrayOutputStream();
             DeflaterOutputStream dos = new DeflaterOutputStream(bout)) {
            dos.write(data.getBytes());
            dos.close();
            return Base64.getEncoder().encodeToString(bout.toByteArray());
        } catch (IOException e) {
            e.printStackTrace();
            return null;
        }
    }

    private String decompress(String data) {
        try (ByteArrayInputStream bin = new ByteArrayInputStream(Base64.getDecoder().decode(data));
             InflaterInputStream iin = new InflaterInputStream(bin);
             ByteArrayOutputStream bout = new ByteArrayOutputStream()) {
            int b;
            while ((b = iin.read()) != -1) {
                bout.write(b);
            }
            return new String(bout.toByteArray());
        } catch (IOException e) {
            e.printStackTrace();
            return null;
        }
    }
}
