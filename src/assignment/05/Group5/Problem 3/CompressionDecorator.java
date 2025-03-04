// Step 5: Implement Compression Decorator

import java.util.zip.Deflater;
import java.util.zip.Inflater;
import java.util.Base64;
import java.util.Arrays;
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