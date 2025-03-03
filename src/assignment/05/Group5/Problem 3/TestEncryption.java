//Nutchanon Boonyato
public class TestEncryption {
    public static void main(String[] args) {
        Encryption encryption = new Encryption();

        String original = "Hello World!";
        String encrypted = encryption.encode(original);
        String decrypted = encryption.decode(encrypted);

        System.out.println("Original: " + original);
        System.out.println("Encrypted: " + encrypted);
        System.out.println("Decrypted: " + decrypted);
    }
}
