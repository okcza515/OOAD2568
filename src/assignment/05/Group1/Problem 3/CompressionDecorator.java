import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.util.Base64;
import java.util.zip.Deflater;
import java.util.zip.DeflaterOutputStream;
import java.util.zip.InflaterInputStream;

public class CompressionDecorator extends BaseDecorator {

    public CompressionDecorator(FileData source) {
        super(source);
    }

    public void writeData(String data) {
        super.writeData(compress(data));
    }

    public String readData() {
        return decompress(super.readData());
    }

    private String compress(String data) {
        byte[] result = data.getBytes();
        try (ByteArrayOutputStream bos = new ByteArrayOutputStream(result.length);
             DeflaterOutputStream dos = new DeflaterOutputStream(bos, new Deflater())) {
            dos.write(result);
            dos.close();
            return Base64.getEncoder().encodeToString(bos.toByteArray());
        } catch (IOException e) {
            e.printStackTrace();
            return null;
        }
    }

    private String decompress(String data) {
        byte[] result = Base64.getDecoder().decode(data);
        try (InputStream is = new ByteArrayInputStream(result);
             InflaterInputStream iis = new InflaterInputStream(is)) {
            byte[] buffer = new byte[result.length];
            int length = iis.read(buffer);
            return new String(buffer, 0, length);
        } catch (IOException e) {
            e.printStackTrace();
            return null;
        }
    }
}
