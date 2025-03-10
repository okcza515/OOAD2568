import java.io.*;
import java.util.Base64;
import java.util.zip.Deflater;
import java.util.zip.DeflaterOutputStream;
import java.util.zip.InflaterInputStream;

public class CompressionDecorator extends DataSourceDecorator {
    private int compLevel = 6;

    public CompressionDecorator(DataSource source) {
        super(source);
    }

    @Override
    public void writeData(String data) {
        super.writeData(compress(data));
    }

    @Override
    public String readData() {
        return decompress(super.readData());
    }

    private String compress(String data) {
        byte[] input = data.getBytes();
        try (ByteArrayOutputStream bout = new ByteArrayOutputStream(512);
             DeflaterOutputStream dos = new DeflaterOutputStream(bout, new Deflater(compLevel))) {
            dos.write(input);
            dos.close();
            return Base64.getEncoder().encodeToString(bout.toByteArray());
        } catch (IOException ex) {
            return null;
        }
    }

    private String decompress(String data) {
        byte[] input = Base64.getDecoder().decode(data);
        try (ByteArrayInputStream in = new ByteArrayInputStream(input);
             InflaterInputStream iin = new InflaterInputStream(in);
             ByteArrayOutputStream bout = new ByteArrayOutputStream(512)) {
            int b;
            while ((b = iin.read()) != -1) {
                bout.write(b);
            }
            return new String(bout.toByteArray());
        } catch (IOException ex) {
            return null;
        }
    }
}
