import java.io.File;
import java.io.FileOutputStream;
import java.io.FileReader;
import java.io.IOException;
import java.io.OutputStream;
import java.util.zip.Deflater;
import java.util.zip.Inflater;
import java.util.Base64;
import java.util.Arrays;

//Nutchanon Boonyato
// Step 1: Create a common interface
interface DataSource {
    void writeData(String data);
    String readData();
}

// Step 2: Implement the base FileDataSource
class FileDataSource implements DataSource {
    private String name;

    public FileDataSource(String name) {
        this.name = name;
    }

    @Override
    public void writeData(String data) {
        File file = new File(name);
        try (OutputStream fos = new FileOutputStream(file)) {
            fos.write(data.getBytes(), 0, data.length());
        } catch (IOException ex) {
            System.out.println(ex.getMessage());
        }
    }

    @Override
    public String readData() {
        char[] buffer = null;
        File file = new File(name);
        try (FileReader reader = new FileReader(file)) {
            buffer = new char[(int) file.length()];
            reader.read(buffer);
        } catch (IOException ex) {
            System.out.println(ex.getMessage());
        }
        return new String(buffer);
    }
}

// Step 3: Create an abstract decorator class
abstract class DataSourceDecorator implements DataSource {
    protected DataSource wrappee;

    public DataSourceDecorator(DataSource source) {
        this.wrappee = source;
    }

    @Override
    public void writeData(String data) {
        wrappee.writeData(data);
    }

    @Override
    public String readData() {
        return wrappee.readData();
    }
}

// Step 4: Implement Encryption Decorator
class EncryptionDecorator extends DataSourceDecorator {
    public EncryptionDecorator(DataSource source) {
        super(source);
    }

    @Override
    public void writeData(String data) {
        super.writeData(encode(data));
    }

    @Override
    public String readData() {
        return decode(super.readData());
    }

    private String encode(String data) {
        return Base64.getEncoder().encodeToString(data.getBytes());
    }

    private String decode(String data) {
        return new String(Base64.getDecoder().decode(data));
    }
}

// Step 5: Implement Compression Decorator
class CompressionDecorator extends DataSourceDecorator {
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
        try {
            byte[] input = data.getBytes();
            Deflater deflater = new Deflater();
            deflater.setInput(input);
            deflater.finish();
            byte[] output = new byte[100];
            int compressedDataLength = deflater.deflate(output);
            return Base64.getEncoder().encodeToString(Arrays.copyOf(output, compressedDataLength));
        } catch (Exception e) {
            return data;
        }
    }

    private String decompress(String data) {
        try {
            byte[] input = Base64.getDecoder().decode(data);
            Inflater inflater = new Inflater();
            inflater.setInput(input);
            byte[] output = new byte[100];
            int decompressedDataLength = inflater.inflate(output);
            return new String(output, 0, decompressedDataLength);
        } catch (Exception e) {
            return data;
        }
    }
}

// Step 6: Usage Example
class Application {
    public static void main(String[] args) {
        DataSource source = new FileDataSource("OutputDemo.txt");
        
        // Writing with encryption and compression
        DataSource encryptedCompressed = new EncryptionDecorator(new CompressionDecorator(source));
        encryptedCompressed.writeData("Hello, Secure World!");

        // Reading the data
        String result = encryptedCompressed.readData();
        System.out.println("Decrypted & Decompressed Data: " + result);
    }
}
