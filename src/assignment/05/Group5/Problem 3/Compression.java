import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.util.Base64;
import java.util.zip.Deflater;
import java.util.zip.DeflaterOutputStream;
import java.util.zip.InflaterInputStream;

public class Compression {
    // 65070501048 Rattipong Sakunjeen
    private int compLevel = 6; // Default compression level

    // Getter and Setter for compression level
    public int getCompressionLevel() {
        return compLevel;
    }

    public void setCompressionLevel(int value) {
        if (value >= 0 && value <= 9) { // Valid Deflater levels (0-9)
            compLevel = value;
        }
    }

    // Compress the input string and return a Base64-encoded result
    public String compress(String stringData) {
        if (stringData == null || stringData.isEmpty()) {
            return stringData; // Return as is if empty/null
        }

        byte[] data = stringData.getBytes();
        try (ByteArrayOutputStream bout = new ByteArrayOutputStream();
             DeflaterOutputStream dos = new DeflaterOutputStream(bout, new Deflater(compLevel))) {
            
            dos.write(data);
            dos.finish(); // Ensure all data is written
            return Base64.getEncoder().encodeToString(bout.toByteArray());
        } catch (IOException ex) {
            return stringData; // Return original data if compression fails
        }
    }

    // Decompress a Base64-encoded compressed string
    public String decompress(String stringData) {
        if (stringData == null || stringData.isEmpty()) {
            return stringData; // Return as is if empty/null
        }

        byte[] data = Base64.getDecoder().decode(stringData);
        try (InputStream in = new ByteArrayInputStream(data);
             InflaterInputStream iin = new InflaterInputStream(in);
             ByteArrayOutputStream bout = new ByteArrayOutputStream()) {
            
            int b;
            while ((b = iin.read()) != -1) {
                bout.write(b);
            }
            return new String(bout.toByteArray());
        } catch (IOException ex) {
            return stringData; // Return original data if decompression fails
        }
    }
}
