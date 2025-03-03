import java.util.Base64;
//Nutchanon Boonyato
public class Encryption {
    
    // Encrypts the data by shifting bytes and encoding to Base64
    public String encode(String data) {
        if (data == null || data.isEmpty()) {
            return data; // Avoid processing null/empty strings
        }
        
        byte[] result = data.getBytes();
        for (int i = 0; i < result.length; i++) {
            result[i] += (byte) 1; // Shift bytes for simple encryption
        }
        
        return Base64.getEncoder().encodeToString(result); // Encode to Base64
    }

    // Decrypts the Base64-encoded data by reversing the shift
    public String decode(String data) {
        if (data == null || data.isEmpty()) {
            return data; // Avoid errors on null/empty input
        }
        
        byte[] result = Base64.getDecoder().decode(data); // Decode from Base64
        for (int i = 0; i < result.length; i++) {
            result[i] -= (byte) 1; // Reverse byte shift for decryption
        }
        
        return new String(result);
    }
}

